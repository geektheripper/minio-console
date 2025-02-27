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

package k_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// KMSAPIsOKCode is the HTTP code returned for type KMSAPIsOK
const KMSAPIsOKCode int = 200

/*
KMSAPIsOK A successful response.

swagger:response kMSAPIsOK
*/
type KMSAPIsOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsAPIsResponse `json:"body,omitempty"`
}

// NewKMSAPIsOK creates KMSAPIsOK with default headers values
func NewKMSAPIsOK() *KMSAPIsOK {

	return &KMSAPIsOK{}
}

// WithPayload adds the payload to the k m s a p is o k response
func (o *KMSAPIsOK) WithPayload(payload *models.KmsAPIsResponse) *KMSAPIsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s a p is o k response
func (o *KMSAPIsOK) SetPayload(payload *models.KmsAPIsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSAPIsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSAPIsDefault Generic error response.

swagger:response kMSAPIsDefault
*/
type KMSAPIsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewKMSAPIsDefault creates KMSAPIsDefault with default headers values
func NewKMSAPIsDefault(code int) *KMSAPIsDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSAPIsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s a p is default response
func (o *KMSAPIsDefault) WithStatusCode(code int) *KMSAPIsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s a p is default response
func (o *KMSAPIsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s a p is default response
func (o *KMSAPIsDefault) WithPayload(payload *models.Error) *KMSAPIsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s a p is default response
func (o *KMSAPIsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSAPIsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
