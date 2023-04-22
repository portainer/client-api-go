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

// PortainereeUserThemeSettings portaineree user theme settings
//
// swagger:model portaineree.UserThemeSettings
type PortainereeUserThemeSettings struct {

	// Color represents the color theme of the UI
	// Example: dark
	// Enum: [dark light highcontrast auto]
	Color string `json:"color,omitempty"`

	// SubtleUpgradeButton indicates if the upgrade banner should be displayed in a subtle way
	SubtleUpgradeButton bool `json:"subtleUpgradeButton,omitempty"`
}

// Validate validates this portaineree user theme settings
func (m *PortainereeUserThemeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var portainereeUserThemeSettingsTypeColorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dark","light","highcontrast","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portainereeUserThemeSettingsTypeColorPropEnum = append(portainereeUserThemeSettingsTypeColorPropEnum, v)
	}
}

const (

	// PortainereeUserThemeSettingsColorDark captures enum value "dark"
	PortainereeUserThemeSettingsColorDark string = "dark"

	// PortainereeUserThemeSettingsColorLight captures enum value "light"
	PortainereeUserThemeSettingsColorLight string = "light"

	// PortainereeUserThemeSettingsColorHighcontrast captures enum value "highcontrast"
	PortainereeUserThemeSettingsColorHighcontrast string = "highcontrast"

	// PortainereeUserThemeSettingsColorAuto captures enum value "auto"
	PortainereeUserThemeSettingsColorAuto string = "auto"
)

// prop value enum
func (m *PortainereeUserThemeSettings) validateColorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portainereeUserThemeSettingsTypeColorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortainereeUserThemeSettings) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	// value enum
	if err := m.validateColorEnum("color", "body", m.Color); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this portaineree user theme settings based on context it is used
func (m *PortainereeUserThemeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeUserThemeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeUserThemeSettings) UnmarshalBinary(b []byte) error {
	var res PortainereeUserThemeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
