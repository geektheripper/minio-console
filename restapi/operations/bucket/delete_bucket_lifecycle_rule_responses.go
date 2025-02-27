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

package bucket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// DeleteBucketLifecycleRuleNoContentCode is the HTTP code returned for type DeleteBucketLifecycleRuleNoContent
const DeleteBucketLifecycleRuleNoContentCode int = 204

/*
DeleteBucketLifecycleRuleNoContent A successful response.

swagger:response deleteBucketLifecycleRuleNoContent
*/
type DeleteBucketLifecycleRuleNoContent struct {
}

// NewDeleteBucketLifecycleRuleNoContent creates DeleteBucketLifecycleRuleNoContent with default headers values
func NewDeleteBucketLifecycleRuleNoContent() *DeleteBucketLifecycleRuleNoContent {

	return &DeleteBucketLifecycleRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteBucketLifecycleRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteBucketLifecycleRuleDefault Generic error response.

swagger:response deleteBucketLifecycleRuleDefault
*/
type DeleteBucketLifecycleRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBucketLifecycleRuleDefault creates DeleteBucketLifecycleRuleDefault with default headers values
func NewDeleteBucketLifecycleRuleDefault(code int) *DeleteBucketLifecycleRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBucketLifecycleRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete bucket lifecycle rule default response
func (o *DeleteBucketLifecycleRuleDefault) WithStatusCode(code int) *DeleteBucketLifecycleRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete bucket lifecycle rule default response
func (o *DeleteBucketLifecycleRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete bucket lifecycle rule default response
func (o *DeleteBucketLifecycleRuleDefault) WithPayload(payload *models.Error) *DeleteBucketLifecycleRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bucket lifecycle rule default response
func (o *DeleteBucketLifecycleRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBucketLifecycleRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
