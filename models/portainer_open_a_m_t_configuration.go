// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerOpenAMTConfiguration portainer open a m t configuration
//
// swagger:model portainer.OpenAMTConfiguration
type PortainerOpenAMTConfiguration struct {

	// cert file content
	CertFileContent string `json:"certFileContent,omitempty"`

	// cert file name
	CertFileName string `json:"certFileName,omitempty"`

	// cert file password
	CertFilePassword string `json:"certFilePassword,omitempty"`

	// domain name
	DomainName string `json:"domainName,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// mps password
	MpsPassword string `json:"mpsPassword,omitempty"`

	// mps server
	MpsServer string `json:"mpsServer,omitempty"`

	// retrieved from API
	MpsToken string `json:"mpsToken,omitempty"`

	// mps user
	MpsUser string `json:"mpsUser,omitempty"`
}

// Validate validates this portainer open a m t configuration
func (m *PortainerOpenAMTConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer open a m t configuration based on context it is used
func (m *PortainerOpenAMTConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerOpenAMTConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerOpenAMTConfiguration) UnmarshalBinary(b []byte) error {
	var res PortainerOpenAMTConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
