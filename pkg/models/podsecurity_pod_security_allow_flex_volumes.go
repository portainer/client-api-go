// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodsecurityPodSecurityAllowFlexVolumes podsecurity pod security allow flex volumes
//
// swagger:model podsecurity.PodSecurityAllowFlexVolumes
type PodsecurityPodSecurityAllowFlexVolumes struct {

	// allowed volumes
	AllowedVolumes []string `json:"allowedVolumes"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this podsecurity pod security allow flex volumes
func (m *PodsecurityPodSecurityAllowFlexVolumes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this podsecurity pod security allow flex volumes based on context it is used
func (m *PodsecurityPodSecurityAllowFlexVolumes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodsecurityPodSecurityAllowFlexVolumes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodsecurityPodSecurityAllowFlexVolumes) UnmarshalBinary(b []byte) error {
	var res PodsecurityPodSecurityAllowFlexVolumes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
