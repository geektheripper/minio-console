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

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CheckMinIOVersionHandlerFunc turns a function with the right signature into a check min i o version handler
type CheckMinIOVersionHandlerFunc func(CheckMinIOVersionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckMinIOVersionHandlerFunc) Handle(params CheckMinIOVersionParams) middleware.Responder {
	return fn(params)
}

// CheckMinIOVersionHandler interface for that can handle valid check min i o version params
type CheckMinIOVersionHandler interface {
	Handle(CheckMinIOVersionParams) middleware.Responder
}

// NewCheckMinIOVersion creates a new http.Handler for the check min i o version operation
func NewCheckMinIOVersion(ctx *middleware.Context, handler CheckMinIOVersionHandler) *CheckMinIOVersion {
	return &CheckMinIOVersion{Context: ctx, Handler: handler}
}

/*
	CheckMinIOVersion swagger:route GET /check-version System checkMinIOVersion

Checks the current MinIO version against the latest
*/
type CheckMinIOVersion struct {
	Context *middleware.Context
	Handler CheckMinIOVersionHandler
}

func (o *CheckMinIOVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCheckMinIOVersionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
