// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeFDOConfiguration portaineree f d o configuration
//
// swagger:model portaineree.FDOConfiguration
type PortainereeFDOConfiguration struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// owner password
	OwnerPassword string `json:"ownerPassword,omitempty"`

	// owner URL
	OwnerURL string `json:"ownerURL,omitempty"`

	// owner username
	OwnerUsername string `json:"ownerUsername,omitempty"`
}

// Validate validates this portaineree f d o configuration
func (m *PortainereeFDOConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree f d o configuration based on context it is used
func (m *PortainereeFDOConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeFDOConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeFDOConfiguration) UnmarshalBinary(b []byte) error {
	var res PortainereeFDOConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
