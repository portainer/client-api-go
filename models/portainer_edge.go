// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerEdge portainer edge
//
// swagger:model portainer.Edge
type PortainerEdge struct {

	// Deprecated 2.18
	// Example: false
	AsyncMode bool `json:"AsyncMode,omitempty"`

	// The command list interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	CommandInterval int64 `json:"CommandInterval,omitempty"`

	// The ping interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	PingInterval int64 `json:"PingInterval,omitempty"`

	// The snapshot interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	SnapshotInterval int64 `json:"SnapshotInterval,omitempty"`
}

// Validate validates this portainer edge
func (m *PortainerEdge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer edge based on context it is used
func (m *PortainerEdge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEdge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEdge) UnmarshalBinary(b []byte) error {
	var res PortainerEdge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
