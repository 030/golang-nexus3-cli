// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProxyAttributes proxy attributes
//
// swagger:model ProxyAttributes
type ProxyAttributes struct {

	// How long to cache artifacts before rechecking the remote repository (in minutes)
	// Example: 1440
	// Required: true
	ContentMaxAge *int32 `json:"contentMaxAge"`

	// How long to cache metadata before rechecking the remote repository (in minutes)
	// Example: 1440
	// Required: true
	MetadataMaxAge *int32 `json:"metadataMaxAge"`

	// Location of the remote repository being proxied
	// Example: https://remote.repository.com
	RemoteURL string `json:"remoteUrl,omitempty"`
}

// Validate validates this proxy attributes
func (m *ProxyAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadataMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyAttributes) validateContentMaxAge(formats strfmt.Registry) error {

	if err := validate.Required("contentMaxAge", "body", m.ContentMaxAge); err != nil {
		return err
	}

	return nil
}

func (m *ProxyAttributes) validateMetadataMaxAge(formats strfmt.Registry) error {

	if err := validate.Required("metadataMaxAge", "body", m.MetadataMaxAge); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this proxy attributes based on context it is used
func (m *ProxyAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProxyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyAttributes) UnmarshalBinary(b []byte) error {
	var res ProxyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}