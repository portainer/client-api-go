// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerTemplateEnvSelect portainer template env select
//
// swagger:model portainer.TemplateEnvSelect
type PortainerTemplateEnvSelect struct {

	// Will set this choice as the default choice
	// Example: false
	Default *bool `json:"default,omitempty"`

	// Some text that will displayed as a choice
	// Example: text value
	Text string `json:"text,omitempty"`

	// A value that will be associated to the choice
	// Example: value
	Value string `json:"value,omitempty"`
}

// Validate validates this portainer template env select
func (m *PortainerTemplateEnvSelect) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer template env select based on context it is used
func (m *PortainerTemplateEnvSelect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerTemplateEnvSelect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerTemplateEnvSelect) UnmarshalBinary(b []byte) error {
	var res PortainerTemplateEnvSelect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
