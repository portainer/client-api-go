// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeTeamResourceAccess portaineree team resource access
//
// swagger:model portaineree.TeamResourceAccess
type PortainereeTeamResourceAccess struct {

	// access level
	AccessLevel PortainereeResourceAccessLevel `json:"AccessLevel,omitempty"`

	// team Id
	TeamID int64 `json:"TeamId,omitempty"`
}

// Validate validates this portaineree team resource access
func (m *PortainereeTeamResourceAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeTeamResourceAccess) validateAccessLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessLevel) { // not required
		return nil
	}

	if err := m.AccessLevel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccessLevel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccessLevel")
		}
		return err
	}

	return nil
}

// ContextValidate validate this portaineree team resource access based on the context it is used
func (m *PortainereeTeamResourceAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeTeamResourceAccess) contextValidateAccessLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccessLevel.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccessLevel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccessLevel")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeTeamResourceAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeTeamResourceAccess) UnmarshalBinary(b []byte) error {
	var res PortainereeTeamResourceAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
