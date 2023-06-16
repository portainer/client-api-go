// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerSSLSettings portainer s s l settings
//
// swagger:model portainer.SSLSettings
type PortainerSSLSettings struct {

	// cert path
	CertPath string `json:"certPath,omitempty"`

	// http enabled
	HTTPEnabled *bool `json:"httpEnabled,omitempty"`

	// key path
	KeyPath string `json:"keyPath,omitempty"`

	// self signed
	SelfSigned *bool `json:"selfSigned,omitempty"`
}

// Validate validates this portainer s s l settings
func (m *PortainerSSLSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer s s l settings based on context it is used
func (m *PortainerSSLSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerSSLSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerSSLSettings) UnmarshalBinary(b []byte) error {
	var res PortainerSSLSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
