// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SeccompProfile v1 seccomp profile
//
// swagger:model v1.SeccompProfile
type V1SeccompProfile struct {

	// localhostProfile indicates a profile defined in a file on the node should be used.
	// The profile must be preconfigured on the node to work.
	// Must be a descending path, relative to the kubelet's configured seccomp profile location.
	// Must be set if type is "Localhost". Must NOT be set for any other type.
	// +optional
	LocalhostProfile string `json:"localhostProfile,omitempty"`

	// type indicates which kind of seccomp profile will be applied.
	// Valid options are:
	//
	// Localhost - a profile defined in a file on the node should be used.
	// RuntimeDefault - the container runtime default profile should be used.
	// Unconfined - no profile should be applied.
	// +unionDiscriminator
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 seccomp profile
func (m *V1SeccompProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 seccomp profile based on context it is used
func (m *V1SeccompProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SeccompProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SeccompProfile) UnmarshalBinary(b []byte) error {
	var res V1SeccompProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
