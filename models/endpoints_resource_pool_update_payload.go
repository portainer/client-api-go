// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointsResourcePoolUpdatePayload endpoints resource pool update payload
//
// swagger:model endpoints.resourcePoolUpdatePayload
type EndpointsResourcePoolUpdatePayload struct {

	// teams to add
	TeamsToAdd []int64 `json:"teamsToAdd"`

	// teams to remove
	TeamsToRemove []int64 `json:"teamsToRemove"`

	// users to add
	UsersToAdd []int64 `json:"usersToAdd"`

	// users to remove
	UsersToRemove []int64 `json:"usersToRemove"`
}

// Validate validates this endpoints resource pool update payload
func (m *EndpointsResourcePoolUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this endpoints resource pool update payload based on context it is used
func (m *EndpointsResourcePoolUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsResourcePoolUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsResourcePoolUpdatePayload) UnmarshalBinary(b []byte) error {
	var res EndpointsResourcePoolUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
