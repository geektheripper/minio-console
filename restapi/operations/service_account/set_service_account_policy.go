// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
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
//

package service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/minio/console/models"
)

// SetServiceAccountPolicyHandlerFunc turns a function with the right signature into a set service account policy handler
type SetServiceAccountPolicyHandlerFunc func(SetServiceAccountPolicyParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SetServiceAccountPolicyHandlerFunc) Handle(params SetServiceAccountPolicyParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SetServiceAccountPolicyHandler interface for that can handle valid set service account policy params
type SetServiceAccountPolicyHandler interface {
	Handle(SetServiceAccountPolicyParams, *models.Principal) middleware.Responder
}

// NewSetServiceAccountPolicy creates a new http.Handler for the set service account policy operation
func NewSetServiceAccountPolicy(ctx *middleware.Context, handler SetServiceAccountPolicyHandler) *SetServiceAccountPolicy {
	return &SetServiceAccountPolicy{Context: ctx, Handler: handler}
}

/*
	SetServiceAccountPolicy swagger:route PUT /service-accounts/{access_key}/policy ServiceAccount setServiceAccountPolicy

Set Service Account Policy
*/
type SetServiceAccountPolicy struct {
	Context *middleware.Context
	Handler SetServiceAccountPolicyHandler
}

func (o *SetServiceAccountPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSetServiceAccountPolicyParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
