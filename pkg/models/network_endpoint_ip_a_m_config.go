// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkEndpointIPAMConfig network endpoint IP a m config
//
// swagger:model network.EndpointIPAMConfig
type NetworkEndpointIPAMConfig struct {

	// ipv4 address
	IPV4Address string `json:"ipv4Address,omitempty"`

	// ipv6 address
	IPV6Address string `json:"ipv6Address,omitempty"`

	// link local i ps
	LinkLocalIPs []string `json:"linkLocalIPs"`
}

// Validate validates this network endpoint IP a m config
func (m *NetworkEndpointIPAMConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network endpoint IP a m config based on context it is used
func (m *NetworkEndpointIPAMConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkEndpointIPAMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkEndpointIPAMConfig) UnmarshalBinary(b []byte) error {
	var res NetworkEndpointIPAMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
