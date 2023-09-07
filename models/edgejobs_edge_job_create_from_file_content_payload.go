// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgejobsEdgeJobCreateFromFileContentPayload edgejobs edge job create from file content payload
//
// swagger:model edgejobs.edgeJobCreateFromFileContentPayload
type EdgejobsEdgeJobCreateFromFileContentPayload struct {

	// cron expression
	CronExpression string `json:"cronExpression,omitempty"`

	// edge groups
	EdgeGroups []int64 `json:"edgeGroups"`

	// endpoints
	Endpoints []int64 `json:"endpoints"`

	// file content
	FileContent string `json:"fileContent,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// recurring
	Recurring bool `json:"recurring,omitempty"`
}

// Validate validates this edgejobs edge job create from file content payload
func (m *EdgejobsEdgeJobCreateFromFileContentPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edgejobs edge job create from file content payload based on context it is used
func (m *EdgejobsEdgeJobCreateFromFileContentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgejobsEdgeJobCreateFromFileContentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgejobsEdgeJobCreateFromFileContentPayload) UnmarshalBinary(b []byte) error {
	var res EdgejobsEdgeJobCreateFromFileContentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
