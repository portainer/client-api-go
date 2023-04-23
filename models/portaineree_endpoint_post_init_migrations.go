// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeEndpointPostInitMigrations portaineree endpoint post init migrations
//
// swagger:model portaineree.EndpointPostInitMigrations
type PortainereeEndpointPostInitMigrations struct {

	// migrate g p us
	MigrateGPUs *bool `json:"MigrateGPUs,omitempty"`

	// migrate gate keeper
	MigrateGateKeeper *bool `json:"MigrateGateKeeper,omitempty"`

	// migrate ingresses
	MigrateIngresses *bool `json:"MigrateIngresses,omitempty"`
}

// Validate validates this portaineree endpoint post init migrations
func (m *PortainereeEndpointPostInitMigrations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree endpoint post init migrations based on context it is used
func (m *PortainereeEndpointPostInitMigrations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeEndpointPostInitMigrations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEndpointPostInitMigrations) UnmarshalBinary(b []byte) error {
	var res PortainereeEndpointPostInitMigrations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
