// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/minio/console/pkg/logger"
	"github.com/minio/console/pkg/utils"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/klauspost/compress/gzhttp"

	portal_ui "github.com/minio/console/portal-ui"
	"github.com/minio/pkg/env"
	"github.com/minio/pkg/mimedb"
	xnet "github.com/minio/pkg/net"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/minio/console/models"
	"github.com/minio/console/pkg/auth"
	"github.com/minio/console/restapi/operations"
	"github.com/unrolled/secure"
)

//go:generate swagger generate server --target ../../console --name Console --spec ../swagger.yml

var additionalServerFlags = struct {
	CertsDir string `long:"certs-dir" description:"path to certs directory" env:"CONSOLE_CERTS_DIR"`
}{}

const (
	SubPath = "CONSOLE_SUBPATH"
)

var (
	subPath     = "/"
	subPathOnce sync.Once
)

func configureFlags(api *operations.ConsoleAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "additional server flags",
			Options:          &additionalServerFlags,
		},
	}
}

func configureAPI(api *operations.ConsoleAPI) http.Handler {
	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string, scopes []string) (*models.Principal, error) {
		// we are validating the session token by decrypting the claims inside, if the operation succeed that means the jwt
		// was generated and signed by us in the first place
		if token == "Anonymous" {
			return &models.Principal{}, nil
		}
		claims, err := auth.ParseClaimsFromToken(token)
		if err != nil {
			api.Logger("Unable to validate the session token %s: %v", token, err)
			return nil, errors.New(401, "incorrect api key auth")
		}
		return &models.Principal{
			STSAccessKeyID:     claims.STSAccessKeyID,
			STSSecretAccessKey: claims.STSSecretAccessKey,
			STSSessionToken:    claims.STSSessionToken,
			AccountAccessKey:   claims.AccountAccessKey,
			Hm:                 claims.HideMenu,
			Ob:                 claims.ObjectBrowser,
			CustomStyleOb:      claims.CustomStyleOB,
		}, nil
	}
	api.AnonymousAuth = func(s string) (*models.Principal, error) {
		return &models.Principal{}, nil
	}

	// Register login handlers
	registerLoginHandlers(api)
	// Register logout handlers
	registerLogoutHandlers(api)
	// Register bucket handlers
	registerBucketsHandlers(api)
	// Register all users handlers
	registerUsersHandlers(api)
	// Register groups handlers
	registerGroupsHandlers(api)
	// Register policies handlers
	registersPoliciesHandler(api)
	// Register configurations handlers
	registerConfigHandlers(api)
	// Register bucket events handlers
	registerBucketEventsHandlers(api)
	// Register bucket lifecycle handlers
	registerBucketsLifecycleHandlers(api)
	// Register service handlers
	registerServiceHandlers(api)
	// Register session handlers
	registerSessionHandlers(api)
	// Register version handlers
	registerVersionHandlers(api)
	// Register admin info handlers
	registerAdminInfoHandlers(api)
	// Register admin arns handlers
	registerAdminArnsHandlers(api)
	// Register admin notification endpoints handlers
	registerAdminNotificationEndpointsHandlers(api)
	// Register admin Service Account Handlers
	registerServiceAccountsHandlers(api)
	// Register admin remote buckets
	registerAdminBucketRemoteHandlers(api)
	// Register admin log search
	registerLogSearchHandlers(api)
	// Register admin subnet handlers
	registerSubnetHandlers(api)
	// Register admin KMS handlers
	registerKMSHandlers(api)
	// Register admin IDP handlers
	registerIDPHandlers(api)
	// Register Account handlers
	registerAdminTiersHandlers(api)
	// Register Inspect Handler
	registerInspectHandler(api)
	// Register nodes handlers
	registerNodesHandler(api)

	registerSiteReplicationHandler(api)
	registerSiteReplicationStatusHandler(api)
	// Register Support Handler
	registerSupportHandlers(api)

	// Operator Console

	// Register Object's Handlers
	registerObjectsHandlers(api)
	// Register Bucket Quota's Handlers
	registerBucketQuotaHandlers(api)
	// Register Account handlers
	registerAccountHandlers(api)

	registerReleasesHandlers(api)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	// do an initial subnet plan caching
	fetchLicensePlan()

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	tlsConfig.RootCAs = GlobalRootCAs
	tlsConfig.GetCertificate = GlobalTLSCertsManager.GetCertificate
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, err := utils.NewUUID()
		if err != nil && err != auth.ErrNoAuthToken {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), utils.ContextRequestID, requestID)
		ctx = context.WithValue(ctx, utils.ContextRequestUserAgent, r.UserAgent())
		ctx = context.WithValue(ctx, utils.ContextRequestHost, r.Host)
		ctx = context.WithValue(ctx, utils.ContextRequestRemoteAddr, r.RemoteAddr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuditLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := logger.NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		if strings.HasPrefix(r.URL.Path, "/ws") || strings.HasPrefix(r.URL.Path, "/api") {
			logger.AuditLog(r.Context(), rw, r, map[string]interface{}{}, "Authorization", "Cookie", "Set-Cookie")
		}
	})
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logger and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	gnext := gzhttp.GzipHandler(handler)
	// if audit-log is enabled console will log all incoming request
	next := AuditLogMiddleware(gnext)
	// serve static files
	next = FileServerMiddleware(next)
	// add information to request context
	next = ContextMiddleware(next)
	// handle cookie or authorization header for session
	next = AuthenticationMiddleware(next)

	sslHostFn := secure.SSLHostFunc(func(host string) string {
		xhost, err := xnet.ParseHost(host)
		if err != nil {
			return host
		}
		return net.JoinHostPort(xhost.Name, TLSPort)
	})

	// Secure middleware, this middleware wrap all the previous handlers and add
	// HTTP security headers
	secureOptions := secure.Options{
		AllowedHosts:                    GetSecureAllowedHosts(),
		AllowedHostsAreRegex:            GetSecureAllowedHostsAreRegex(),
		HostsProxyHeaders:               GetSecureHostsProxyHeaders(),
		SSLRedirect:                     GetTLSRedirect() == "on" && len(GlobalPublicCerts) > 0,
		SSLHostFunc:                     &sslHostFn,
		SSLHost:                         GetSecureTLSHost(),
		STSSeconds:                      GetSecureSTSSeconds(),
		STSIncludeSubdomains:            GetSecureSTSIncludeSubdomains(),
		STSPreload:                      GetSecureSTSPreload(),
		SSLTemporaryRedirect:            false,
		ForceSTSHeader:                  GetSecureForceSTSHeader(),
		FrameDeny:                       GetSecureFrameDeny(),
		ContentTypeNosniff:              GetSecureContentTypeNonSniff(),
		BrowserXssFilter:                GetSecureBrowserXSSFilter(),
		ContentSecurityPolicy:           GetSecureContentSecurityPolicy(),
		ContentSecurityPolicyReportOnly: GetSecureContentSecurityPolicyReportOnly(),
		PublicKey:                       GetSecurePublicKey(),
		ReferrerPolicy:                  GetSecureReferrerPolicy(),
		FeaturePolicy:                   GetSecureFeaturePolicy(),
		ExpectCTHeader:                  GetSecureExpectCTHeader(),
		IsDevelopment:                   false,
	}
	secureMiddleware := secure.New(secureOptions)
	next = secureMiddleware.Handler(next)
	return RejectS3Middleware(next)
}

const apiRequestErr = `<?xml version="1.0" encoding="UTF-8"?><Error><Code>InvalidArgument</Code><Message>S3 API Requests must be made to API port.</Message><RequestId>0</RequestId></Error>`

// RejectS3Middleware will reject requests that have AWS S3 specific headers.
func RejectS3Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("X-Amz-Content-Sha256")) > 0 ||
			len(r.Header.Get("X-Amz-Date")) > 0 ||
			strings.HasPrefix(r.Header.Get("Authorization"), "AWS4-HMAC-SHA256") ||
			r.URL.Query().Get("AWSAccessKeyId") != "" {

			w.Header().Set("Location", getMinIOServer())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(apiRequestErr))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetTokenFromRequest(r)
		if err != nil && err != auth.ErrNoAuthToken {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		sessionToken, _ := auth.DecryptToken(token)
		// All handlers handle appropriately to return errors
		// based on their swagger rules, we do not need to
		// additionally return error here, let the next ServeHTTPs
		// handle it appropriately.
		if len(sessionToken) > 0 {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer  %s", string(sessionToken)))
		} else {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "Anonymous"))
		}
		ctx := r.Context()
		claims, _ := auth.ParseClaimsFromToken(string(sessionToken))
		if claims != nil {
			// save user session id context
			ctx = context.WithValue(r.Context(), utils.ContextRequestUserID, claims.STSSessionToken)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FileServerMiddleware serves files from the static folder
func FileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", globalAppName) // do not add version information
		switch {
		case strings.HasPrefix(r.URL.Path, "/ws"):
			serveWS(w, r)
		case strings.HasPrefix(r.URL.Path, "/api"):
			next.ServeHTTP(w, r)
		default:
			buildFs, err := fs.Sub(portal_ui.GetStaticAssets(), "build")
			if err != nil {
				panic(err)
			}
			wrapHandlerSinglePageApplication(requestBounce(http.FileServer(http.FS(buildFs)))).ServeHTTP(w, r)
		}
	})
}

type notFoundRedirectRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func (w *notFoundRedirectRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w *notFoundRedirectRespWr) Write(p []byte) (int, error) {
	if w.status != http.StatusNotFound {
		return w.ResponseWriter.Write(p)
	}
	return len(p), nil // Lie that we successfully wrote it
}

// handleSPA handles the serving of the React Single Page Application
func handleSPA(w http.ResponseWriter, r *http.Request) {
	basePath := "/"
	// For SPA mode we will replace root base with a sub path if configured unless we received cp=y and cpb=/NEW/BASE
	if v := r.URL.Query().Get("cp"); v == "y" {
		if base := r.URL.Query().Get("cpb"); base != "" {
			// make sure the subpath has a trailing slash
			if !strings.HasSuffix(base, "/") {
				base = fmt.Sprintf("%s/", base)
			}
			basePath = base
		}
	}

	indexPage, err := portal_ui.GetStaticAssets().Open("build/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sts := r.URL.Query().Get("sts")
	stsAccessKey := r.URL.Query().Get("sts_a")
	stsSecretKey := r.URL.Query().Get("sts_s")
	overridenStyles := r.URL.Query().Get("ov_st")

	// if these three parameters are present we are being asked to issue a session with these values
	if sts != "" && stsAccessKey != "" && stsSecretKey != "" {
		creds := credentials.NewStaticV4(stsAccessKey, stsSecretKey, sts)
		consoleCreds := &ConsoleCredentials{
			ConsoleCredentials: creds,
			AccountAccessKey:   stsAccessKey,
		}
		sf := &auth.SessionFeatures{}
		sf.HideMenu = true
		sf.ObjectBrowser = true

		if overridenStyles != "" {
			err := ValidateEncodedStyles(overridenStyles)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			sf.CustomStyleOB = overridenStyles
		}

		sessionID, err := login(consoleCreds, sf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie := NewSessionCookieForConsole(*sessionID)

		http.SetCookie(w, &cookie)

		// Allow us to be iframed
		w.Header().Del("X-Frame-Options")
	}

	indexPageBytes, err := io.ReadAll(indexPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if we have a seeded basePath. This should override CONSOLE_SUBPATH every time, thus the `if else`
	if basePath != "/" {
		indexPageBytes = replaceBaseInIndex(indexPageBytes, basePath)
		// if we have a custom subpath replace it in
	} else if getSubPath() != "/" {
		indexPageBytes = replaceBaseInIndex(indexPageBytes, getSubPath())
	}
	indexPageBytes = replaceLicense(indexPageBytes)

	mimeType := mimedb.TypeByExtension(filepath.Ext(r.URL.Path))

	if mimeType == "application/octet-stream" {
		mimeType = "text/html"
	}

	w.Header().Set("Content-Type", mimeType)
	http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(indexPageBytes))
}

// wrapHandlerSinglePageApplication handles a http.FileServer returning a 404 and overrides it with index.html
func wrapHandlerSinglePageApplication(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handleSPA(w, r)
			return
		}

		w.Header().Set("Content-Type", mimedb.TypeByExtension(filepath.Ext(r.URL.Path)))
		nfw := &notFoundRedirectRespWr{ResponseWriter: w}
		h.ServeHTTP(nfw, r)
		if nfw.status == http.StatusNotFound {
			handleSPA(w, r)
		}
	}
}

type nullWriter struct{}

func (lw nullWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, _, _ string) {
	// Turn-off random logger by Go net/http
	s.ErrorLog = log.New(&nullWriter{}, "", 0)
}

func getSubPath() string {
	subPathOnce.Do(func() {
		subPath = parseSubPath(env.Get(SubPath, ""))
	})
	return subPath
}

func parseSubPath(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return SlashSeparator
	}
	// Replace all unnecessary `\` to `/`
	// also add pro-actively at the end.
	subPath = path.Clean(filepath.ToSlash(v))
	if !strings.HasPrefix(subPath, SlashSeparator) {
		subPath = SlashSeparator + subPath
	}
	if !strings.HasSuffix(subPath, SlashSeparator) {
		subPath += SlashSeparator
	}
	return subPath
}

func replaceBaseInIndex(indexPageBytes []byte, basePath string) []byte {
	if basePath != "" {
		validBasePath := regexp.MustCompile(`^[0-9a-zA-Z\/-]+$`)
		if !validBasePath.MatchString(basePath) {
			return indexPageBytes
		}
		indexPageStr := string(indexPageBytes)
		newBase := fmt.Sprintf("<base href=\"%s\"/>", basePath)
		indexPageStr = strings.Replace(indexPageStr, "<base href=\"/\"/>", newBase, 1)
		indexPageBytes = []byte(indexPageStr)

	}
	return indexPageBytes
}

func replaceLicense(indexPageBytes []byte) []byte {
	indexPageStr := string(indexPageBytes)
	newPlan := fmt.Sprintf("<meta name=\"minio-license\" content=\"%s\" />", InstanceLicensePlan.String())
	indexPageStr = strings.Replace(indexPageStr, "<meta name=\"minio-license\" content=\"apgl\"/>", newPlan, 1)
	indexPageBytes = []byte(indexPageStr)
	return indexPageBytes
}

func requestBounce(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
