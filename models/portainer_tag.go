// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerTag portainer tag
//
// swagger:model portainer.Tag
type PortainerTag struct {

	// A set of endpoint group ids that have this tag
	EndpointGroups map[string]bool `json:"EndpointGroups,omitempty"`

	// A set of endpoint ids that have this tag
	Endpoints map[string]bool `json:"Endpoints,omitempty"`

	// Tag name
	// Example: org/acme
	Name string `json:"Name,omitempty"`

	// Tag identifier
	// Example: 1
	ID int64 `json:"id,omitempty"`
}

// Validate validates this portainer tag
func (m *PortainerTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer tag based on context it is used
func (m *PortainerTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerTag) UnmarshalBinary(b []byte) error {
	var res PortainerTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
