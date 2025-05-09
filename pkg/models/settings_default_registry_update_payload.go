// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingsDefaultRegistryUpdatePayload settings default registry update payload
//
// swagger:model settings.defaultRegistryUpdatePayload
type SettingsDefaultRegistryUpdatePayload struct {

	// hide
	// Example: false
	Hide bool `json:"Hide,omitempty"`
}

// Validate validates this settings default registry update payload
func (m *SettingsDefaultRegistryUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this settings default registry update payload based on context it is used
func (m *SettingsDefaultRegistryUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsDefaultRegistryUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsDefaultRegistryUpdatePayload) UnmarshalBinary(b []byte) error {
	var res SettingsDefaultRegistryUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
