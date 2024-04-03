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

// PortainereeUser portaineree user
//
// swagger:model portaineree.User
type PortainereeUser struct {

	// endpoint authorizations
	EndpointAuthorizations PortainereeEndpointAuthorizations `json:"EndpointAuthorizations,omitempty"`

	// User Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// OpenAI integration parameters
	// Example: sk-1234567890
	OpenAIAPIKey string `json:"OpenAIApiKey,omitempty"`

	// portainer authorizations
	PortainerAuthorizations PortainereeAuthorizations `json:"PortainerAuthorizations,omitempty"`

	// User role (1 for administrator account and 2 for regular account)
	// Example: 1
	Role int64 `json:"Role,omitempty"`

	// token issue at
	// Example: 1639408208
	TokenIssueAt int64 `json:"TokenIssueAt,omitempty"`

	// username
	// Example: bob
	Username string `json:"Username,omitempty"`

	// theme settings
	ThemeSettings *PortainereeUserThemeSettings `json:"themeSettings,omitempty"`

	// Deprecated
	// Example: dark
	UserTheme string `json:"userTheme,omitempty"`
}

// Validate validates this portaineree user
func (m *PortainereeUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortainerAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThemeSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeUser) validateEndpointAuthorizations(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointAuthorizations) { // not required
		return nil
	}

	if m.EndpointAuthorizations != nil {
		if err := m.EndpointAuthorizations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndpointAuthorizations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EndpointAuthorizations")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeUser) validatePortainerAuthorizations(formats strfmt.Registry) error {
	if swag.IsZero(m.PortainerAuthorizations) { // not required
		return nil
	}

	if m.PortainerAuthorizations != nil {
		if err := m.PortainerAuthorizations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PortainerAuthorizations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PortainerAuthorizations")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeUser) validateThemeSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ThemeSettings) { // not required
		return nil
	}

	if m.ThemeSettings != nil {
		if err := m.ThemeSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("themeSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("themeSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portaineree user based on the context it is used
func (m *PortainereeUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpointAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortainerAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThemeSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeUser) contextValidateEndpointAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointAuthorizations) { // not required
		return nil
	}

	if err := m.EndpointAuthorizations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EndpointAuthorizations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EndpointAuthorizations")
		}
		return err
	}

	return nil
}

func (m *PortainereeUser) contextValidatePortainerAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PortainerAuthorizations) { // not required
		return nil
	}

	if err := m.PortainerAuthorizations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PortainerAuthorizations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PortainerAuthorizations")
		}
		return err
	}

	return nil
}

func (m *PortainereeUser) contextValidateThemeSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ThemeSettings != nil {

		if swag.IsZero(m.ThemeSettings) { // not required
			return nil
		}

		if err := m.ThemeSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("themeSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("themeSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeUser) UnmarshalBinary(b []byte) error {
	var res PortainereeUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
