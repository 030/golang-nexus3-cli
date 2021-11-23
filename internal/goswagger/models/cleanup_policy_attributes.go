// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CleanupPolicyAttributes cleanup policy attributes
//
// swagger:model CleanupPolicyAttributes
type CleanupPolicyAttributes struct {

	// Components that match any of the applied policies will be deleted
	PolicyNames []string `json:"policyNames"`
}

// Validate validates this cleanup policy attributes
func (m *CleanupPolicyAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cleanup policy attributes based on context it is used
func (m *CleanupPolicyAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CleanupPolicyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CleanupPolicyAttributes) UnmarshalBinary(b []byte) error {
	var res CleanupPolicyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
