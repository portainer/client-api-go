// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsersUserUpdatePayload users user update payload
//
// swagger:model users.userUpdatePayload
type UsersUserUpdatePayload struct {

	// new password
	// Example: asfj2emv
	// Required: true
	NewPassword *string `json:"newPassword"`

	// password
	// Example: cg9Wgky3
	// Required: true
	Password *string `json:"password"`

	// User role
	// 1 = administrator account
	// 2 = regular account
	// 3 = edge administrator account
	// Example: 2
	// Required: true
	// Enum: [1,2,3]
	Role *int64 `json:"role"`

	// theme
	Theme *UsersThemePayload `json:"theme,omitempty"`

	// use cache
	// Example: true
	// Required: true
	UseCache *bool `json:"useCache"`

	// username
	// Example: bob
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this users user update payload
func (m *UsersUserUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersUserUpdatePayload) validateNewPassword(formats strfmt.Registry) error {

	if err := validate.Required("newPassword", "body", m.NewPassword); err != nil {
		return err
	}

	return nil
}

func (m *UsersUserUpdatePayload) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var usersUserUpdatePayloadTypeRolePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usersUserUpdatePayloadTypeRolePropEnum = append(usersUserUpdatePayloadTypeRolePropEnum, v)
	}
}

// prop value enum
func (m *UsersUserUpdatePayload) validateRoleEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, usersUserUpdatePayloadTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UsersUserUpdatePayload) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

func (m *UsersUserUpdatePayload) validateTheme(formats strfmt.Registry) error {
	if swag.IsZero(m.Theme) { // not required
		return nil
	}

	if m.Theme != nil {
		if err := m.Theme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("theme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("theme")
			}
			return err
		}
	}

	return nil
}

func (m *UsersUserUpdatePayload) validateUseCache(formats strfmt.Registry) error {

	if err := validate.Required("useCache", "body", m.UseCache); err != nil {
		return err
	}

	return nil
}

func (m *UsersUserUpdatePayload) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this users user update payload based on the context it is used
func (m *UsersUserUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersUserUpdatePayload) contextValidateTheme(ctx context.Context, formats strfmt.Registry) error {

	if m.Theme != nil {

		if swag.IsZero(m.Theme) { // not required
			return nil
		}

		if err := m.Theme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("theme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("theme")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsersUserUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersUserUpdatePayload) UnmarshalBinary(b []byte) error {
	var res UsersUserUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
