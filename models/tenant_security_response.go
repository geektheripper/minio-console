// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantSecurityResponse tenant security response
//
// swagger:model tenantSecurityResponse
type TenantSecurityResponse struct {

	// auto cert
	AutoCert bool `json:"autoCert,omitempty"`

	// custom certificates
	CustomCertificates *TenantSecurityResponseCustomCertificates `json:"customCertificates,omitempty"`

	// security context
	SecurityContext *SecurityContext `json:"securityContext,omitempty"`
}

// Validate validates this tenant security response
func (m *TenantSecurityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomCertificates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantSecurityResponse) validateCustomCertificates(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomCertificates) { // not required
		return nil
	}

	if m.CustomCertificates != nil {
		if err := m.CustomCertificates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customCertificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customCertificates")
			}
			return err
		}
	}

	return nil
}

func (m *TenantSecurityResponse) validateSecurityContext(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityContext) { // not required
		return nil
	}

	if m.SecurityContext != nil {
		if err := m.SecurityContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tenant security response based on the context it is used
func (m *TenantSecurityResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantSecurityResponse) contextValidateCustomCertificates(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomCertificates != nil {
		if err := m.CustomCertificates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customCertificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customCertificates")
			}
			return err
		}
	}

	return nil
}

func (m *TenantSecurityResponse) contextValidateSecurityContext(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityContext != nil {
		if err := m.SecurityContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantSecurityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantSecurityResponse) UnmarshalBinary(b []byte) error {
	var res TenantSecurityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TenantSecurityResponseCustomCertificates tenant security response custom certificates
//
// swagger:model TenantSecurityResponseCustomCertificates
type TenantSecurityResponseCustomCertificates struct {

	// client
	Client []*CertificateInfo `json:"client"`

	// minio
	Minio []*CertificateInfo `json:"minio"`

	// minio c as
	MinioCAs []*CertificateInfo `json:"minioCAs"`
}

// Validate validates this tenant security response custom certificates
func (m *TenantSecurityResponseCustomCertificates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinioCAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantSecurityResponseCustomCertificates) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	for i := 0; i < len(m.Client); i++ {
		if swag.IsZero(m.Client[i]) { // not required
			continue
		}

		if m.Client[i] != nil {
			if err := m.Client[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "client" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "client" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantSecurityResponseCustomCertificates) validateMinio(formats strfmt.Registry) error {
	if swag.IsZero(m.Minio) { // not required
		return nil
	}

	for i := 0; i < len(m.Minio); i++ {
		if swag.IsZero(m.Minio[i]) { // not required
			continue
		}

		if m.Minio[i] != nil {
			if err := m.Minio[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantSecurityResponseCustomCertificates) validateMinioCAs(formats strfmt.Registry) error {
	if swag.IsZero(m.MinioCAs) { // not required
		return nil
	}

	for i := 0; i < len(m.MinioCAs); i++ {
		if swag.IsZero(m.MinioCAs[i]) { // not required
			continue
		}

		if m.MinioCAs[i] != nil {
			if err := m.MinioCAs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minioCAs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minioCAs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenant security response custom certificates based on the context it is used
func (m *TenantSecurityResponseCustomCertificates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinioCAs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantSecurityResponseCustomCertificates) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Client); i++ {

		if m.Client[i] != nil {
			if err := m.Client[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "client" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "client" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantSecurityResponseCustomCertificates) contextValidateMinio(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Minio); i++ {

		if m.Minio[i] != nil {
			if err := m.Minio[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minio" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantSecurityResponseCustomCertificates) contextValidateMinioCAs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MinioCAs); i++ {

		if m.MinioCAs[i] != nil {
			if err := m.MinioCAs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customCertificates" + "." + "minioCAs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customCertificates" + "." + "minioCAs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantSecurityResponseCustomCertificates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantSecurityResponseCustomCertificates) UnmarshalBinary(b []byte) error {
	var res TenantSecurityResponseCustomCertificates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
