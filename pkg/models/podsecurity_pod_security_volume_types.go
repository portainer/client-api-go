// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodsecurityPodSecurityVolumeTypes podsecurity pod security volume types
//
// swagger:model podsecurity.PodSecurityVolumeTypes
type PodsecurityPodSecurityVolumeTypes struct {

	// allowed types
	AllowedTypes []string `json:"allowedTypes"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this podsecurity pod security volume types
func (m *PodsecurityPodSecurityVolumeTypes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this podsecurity pod security volume types based on context it is used
func (m *PodsecurityPodSecurityVolumeTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodsecurityPodSecurityVolumeTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodsecurityPodSecurityVolumeTypes) UnmarshalBinary(b []byte) error {
	var res PodsecurityPodSecurityVolumeTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
