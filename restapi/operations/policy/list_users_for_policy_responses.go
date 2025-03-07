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

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListUsersForPolicyOKCode is the HTTP code returned for type ListUsersForPolicyOK
const ListUsersForPolicyOKCode int = 200

/*
ListUsersForPolicyOK A successful response.

swagger:response listUsersForPolicyOK
*/
type ListUsersForPolicyOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListUsersForPolicyOK creates ListUsersForPolicyOK with default headers values
func NewListUsersForPolicyOK() *ListUsersForPolicyOK {

	return &ListUsersForPolicyOK{}
}

// WithPayload adds the payload to the list users for policy o k response
func (o *ListUsersForPolicyOK) WithPayload(payload []string) *ListUsersForPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users for policy o k response
func (o *ListUsersForPolicyOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersForPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
ListUsersForPolicyDefault Generic error response.

swagger:response listUsersForPolicyDefault
*/
type ListUsersForPolicyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListUsersForPolicyDefault creates ListUsersForPolicyDefault with default headers values
func NewListUsersForPolicyDefault(code int) *ListUsersForPolicyDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUsersForPolicyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list users for policy default response
func (o *ListUsersForPolicyDefault) WithStatusCode(code int) *ListUsersForPolicyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list users for policy default response
func (o *ListUsersForPolicyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list users for policy default response
func (o *ListUsersForPolicyDefault) WithPayload(payload *models.Error) *ListUsersForPolicyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users for policy default response
func (o *ListUsersForPolicyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersForPolicyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
