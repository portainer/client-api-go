// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodsecurityPodSecuritySecComp podsecurity pod security sec comp
//
// swagger:model podsecurity.PodSecuritySecComp
type PodsecurityPodSecuritySecComp struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// sec comp type
	SecCompType []string `json:"secCompType"`
}

// Validate validates this podsecurity pod security sec comp
func (m *PodsecurityPodSecuritySecComp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this podsecurity pod security sec comp based on context it is used
func (m *PodsecurityPodSecuritySecComp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodsecurityPodSecuritySecComp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodsecurityPodSecuritySecComp) UnmarshalBinary(b []byte) error {
	var res PodsecurityPodSecuritySecComp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
