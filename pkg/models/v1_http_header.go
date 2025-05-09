// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1HTTPHeader v1 HTTP header
//
// swagger:model v1.HTTPHeader
type V1HTTPHeader struct {

	// The header field name.
	// This will be canonicalized upon output, so case-variant names will be understood as the same header.
	Name string `json:"name,omitempty"`

	// The header field value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 HTTP header
func (m *V1HTTPHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 HTTP header based on context it is used
func (m *V1HTTPHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1HTTPHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HTTPHeader) UnmarshalBinary(b []byte) error {
	var res V1HTTPHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
