// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerStackOption portainer stack option
//
// swagger:model portainer.StackOption
type PortainerStackOption struct {

	// Prune services that are no longer referenced
	// Example: false
	Prune *bool `json:"prune,omitempty"`
}

// Validate validates this portainer stack option
func (m *PortainerStackOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer stack option based on context it is used
func (m *PortainerStackOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerStackOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerStackOption) UnmarshalBinary(b []byte) error {
	var res PortainerStackOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
