// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeEdgeAsyncCommand portaineree edge async command
//
// swagger:model portaineree.EdgeAsyncCommand
type PortainereeEdgeAsyncCommand struct {

	// endpoint ID
	EndpointID int64 `json:"endpointID,omitempty"`

	// executed
	Executed bool `json:"executed,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// op
	Op string `json:"op,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// scheduled time
	ScheduledTime string `json:"scheduledTime,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this portaineree edge async command
func (m *PortainereeEdgeAsyncCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree edge async command based on context it is used
func (m *PortainereeEdgeAsyncCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeEdgeAsyncCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEdgeAsyncCommand) UnmarshalBinary(b []byte) error {
	var res PortainereeEdgeAsyncCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
