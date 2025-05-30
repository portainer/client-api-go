// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerAccessPolicy portainer access policy
//
// swagger:model portainer.AccessPolicy
type PortainerAccessPolicy struct {

	// Role identifier. Reference the role that will be associated to this access policy
	// Example: 1
	RoleID int64 `json:"RoleId,omitempty"`
}

// Validate validates this portainer access policy
func (m *PortainerAccessPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer access policy based on context it is used
func (m *PortainerAccessPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerAccessPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerAccessPolicy) UnmarshalBinary(b []byte) error {
	var res PortainerAccessPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
