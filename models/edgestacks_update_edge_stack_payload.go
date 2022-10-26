// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgestacksUpdateEdgeStackPayload edgestacks update edge stack payload
//
// swagger:model edgestacks.updateEdgeStackPayload
type EdgestacksUpdateEdgeStackPayload struct {

	// Deployment type to deploy this stack
	// Valid values are: 0 - 'compose', 1 - 'kubernetes', 2 - 'nomad'
	// for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints)
	// kubernetes deploytype is enabled only for kubernetes environments(endpoints)
	// nomad deploytype is enabled only for nomad environments(endpoints)
	DeploymentType int64 `json:"deploymentType,omitempty"`

	// edge groups
	EdgeGroups []int64 `json:"edgeGroups"`

	// registries
	Registries []int64 `json:"registries"`

	// stack file content
	StackFileContent string `json:"stackFileContent,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this edgestacks update edge stack payload
func (m *EdgestacksUpdateEdgeStackPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edgestacks update edge stack payload based on context it is used
func (m *EdgestacksUpdateEdgeStackPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksUpdateEdgeStackPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksUpdateEdgeStackPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksUpdateEdgeStackPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
