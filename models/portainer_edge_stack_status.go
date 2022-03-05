// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerEdgeStackStatus portainer edge stack status
//
// swagger:model portainer.EdgeStackStatus
type PortainerEdgeStackStatus struct {

	// endpoint ID
	EndpointID int64 `json:"EndpointID,omitempty"`

	// error
	Error string `json:"Error,omitempty"`

	// type
	Type int64 `json:"Type,omitempty"`
}

// Validate validates this portainer edge stack status
func (m *PortainerEdgeStackStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer edge stack status based on context it is used
func (m *PortainerEdgeStackStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEdgeStackStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEdgeStackStatus) UnmarshalBinary(b []byte) error {
	var res PortainerEdgeStackStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
