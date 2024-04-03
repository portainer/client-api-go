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

// UsersAccessTokenResponse users access token response
//
// swagger:model users.accessTokenResponse
type UsersAccessTokenResponse struct {

	// api key
	APIKey *PortainereeAPIKey `json:"apiKey,omitempty"`

	// raw API key
	RawAPIKey string `json:"rawAPIKey,omitempty"`
}

// Validate validates this users access token response
func (m *UsersAccessTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersAccessTokenResponse) validateAPIKey(formats strfmt.Registry) error {
	if swag.IsZero(m.APIKey) { // not required
		return nil
	}

	if m.APIKey != nil {
		if err := m.APIKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this users access token response based on the context it is used
func (m *UsersAccessTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersAccessTokenResponse) contextValidateAPIKey(ctx context.Context, formats strfmt.Registry) error {

	if m.APIKey != nil {

		if swag.IsZero(m.APIKey) { // not required
			return nil
		}

		if err := m.APIKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsersAccessTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersAccessTokenResponse) UnmarshalBinary(b []byte) error {
	var res UsersAccessTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
