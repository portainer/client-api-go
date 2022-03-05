// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgejobsEdgeJobCreateFromFilePayload edgejobs edge job create from file payload
//
// swagger:model edgejobs.edgeJobCreateFromFilePayload
type EdgejobsEdgeJobCreateFromFilePayload struct {

	// cron expression
	CronExpression string `json:"cronExpression,omitempty"`

	// endpoints
	Endpoints []int64 `json:"endpoints"`

	// file
	File []int64 `json:"file"`

	// name
	Name string `json:"name,omitempty"`

	// recurring
	Recurring bool `json:"recurring,omitempty"`
}

// Validate validates this edgejobs edge job create from file payload
func (m *EdgejobsEdgeJobCreateFromFilePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edgejobs edge job create from file payload based on context it is used
func (m *EdgejobsEdgeJobCreateFromFilePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgejobsEdgeJobCreateFromFilePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgejobsEdgeJobCreateFromFilePayload) UnmarshalBinary(b []byte) error {
	var res EdgejobsEdgeJobCreateFromFilePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
