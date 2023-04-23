// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerEdgeStackStatusDetails portainer edge stack status details
//
// swagger:model portainer.EdgeStackStatusDetails
type PortainerEdgeStackStatusDetails struct {

	// acknowledged
	Acknowledged *bool `json:"acknowledged,omitempty"`

	// error
	Error *bool `json:"error,omitempty"`

	// images pulled
	ImagesPulled *bool `json:"imagesPulled,omitempty"`

	// ok
	Ok *bool `json:"ok,omitempty"`

	// pending
	Pending *bool `json:"pending,omitempty"`

	// remote update success
	RemoteUpdateSuccess *bool `json:"remoteUpdateSuccess,omitempty"`

	// remove
	Remove *bool `json:"remove,omitempty"`
}

// Validate validates this portainer edge stack status details
func (m *PortainerEdgeStackStatusDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer edge stack status details based on context it is used
func (m *PortainerEdgeStackStatusDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEdgeStackStatusDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEdgeStackStatusDetails) UnmarshalBinary(b []byte) error {
	var res PortainerEdgeStackStatusDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
