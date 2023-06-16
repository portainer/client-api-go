// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgegroupsEdgeGroupUpdatePayload edgegroups edge group update payload
//
// swagger:model edgegroups.edgeGroupUpdatePayload
type EdgegroupsEdgeGroupUpdatePayload struct {

	// dynamic
	Dynamic *bool `json:"dynamic,omitempty"`

	// endpoints
	Endpoints []int64 `json:"endpoints"`

	// name
	Name string `json:"name,omitempty"`

	// partial match
	PartialMatch *bool `json:"partialMatch,omitempty"`

	// tag i ds
	TagIDs []int64 `json:"tagIDs"`
}

// Validate validates this edgegroups edge group update payload
func (m *EdgegroupsEdgeGroupUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edgegroups edge group update payload based on context it is used
func (m *EdgegroupsEdgeGroupUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgegroupsEdgeGroupUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgegroupsEdgeGroupUpdatePayload) UnmarshalBinary(b []byte) error {
	var res EdgegroupsEdgeGroupUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
