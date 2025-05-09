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

// UsersThemePayload users theme payload
//
// swagger:model users.themePayload
type UsersThemePayload struct {

	// Color represents the color theme of the UI
	// Example: dark
	// Enum: ["dark","light","highcontrast","auto"]
	Color string `json:"color,omitempty"`

	// SubtleUpgradeButton indicates if the upgrade banner should be displayed in a subtle way
	// Example: false
	SubtleUpgradeButton bool `json:"subtleUpgradeButton,omitempty"`
}

// Validate validates this users theme payload
func (m *UsersThemePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var usersThemePayloadTypeColorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dark","light","highcontrast","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usersThemePayloadTypeColorPropEnum = append(usersThemePayloadTypeColorPropEnum, v)
	}
}

const (

	// UsersThemePayloadColorDark captures enum value "dark"
	UsersThemePayloadColorDark string = "dark"

	// UsersThemePayloadColorLight captures enum value "light"
	UsersThemePayloadColorLight string = "light"

	// UsersThemePayloadColorHighcontrast captures enum value "highcontrast"
	UsersThemePayloadColorHighcontrast string = "highcontrast"

	// UsersThemePayloadColorAuto captures enum value "auto"
	UsersThemePayloadColorAuto string = "auto"
)

// prop value enum
func (m *UsersThemePayload) validateColorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, usersThemePayloadTypeColorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UsersThemePayload) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	// value enum
	if err := m.validateColorEnum("color", "body", m.Color); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this users theme payload based on context it is used
func (m *UsersThemePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsersThemePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersThemePayload) UnmarshalBinary(b []byte) error {
	var res UsersThemePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
