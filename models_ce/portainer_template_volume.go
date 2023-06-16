// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerTemplateVolume portainer template volume
//
// swagger:model portainer.TemplateVolume
type PortainerTemplateVolume struct {

	// Path on the host
	// Example: /tmp
	Bind string `json:"bind,omitempty"`

	// Path inside the container
	// Example: /data
	Container string `json:"container,omitempty"`

	// Whether the volume used should be readonly
	// Example: true
	Readonly *bool `json:"readonly,omitempty"`
}

// Validate validates this portainer template volume
func (m *PortainerTemplateVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer template volume based on context it is used
func (m *PortainerTemplateVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerTemplateVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerTemplateVolume) UnmarshalBinary(b []byte) error {
	var res PortainerTemplateVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
