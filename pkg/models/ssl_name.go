// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SslName ssl name
//
// swagger:model ssl.Name
type SslName struct {

	// country
	Country []string `json:"country"`

	// locality
	Locality []string `json:"locality"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// street address
	StreetAddress []string `json:"streetAddress"`
}

// Validate validates this ssl name
func (m *SslName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ssl name based on context it is used
func (m *SslName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SslName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslName) UnmarshalBinary(b []byte) error {
	var res SslName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
