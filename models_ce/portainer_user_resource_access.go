// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerUserResourceAccess portainer user resource access
//
// swagger:model portainer.UserResourceAccess
type PortainerUserResourceAccess struct {

	// access level
	AccessLevel int64 `json:"AccessLevel,omitempty"`

	// user Id
	UserID int64 `json:"UserId,omitempty"`
}

// Validate validates this portainer user resource access
func (m *PortainerUserResourceAccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer user resource access based on context it is used
func (m *PortainerUserResourceAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerUserResourceAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerUserResourceAccess) UnmarshalBinary(b []byte) error {
	var res PortainerUserResourceAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
