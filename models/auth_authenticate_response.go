// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthAuthenticateResponse auth authenticate response
//
// swagger:model auth.authenticateResponse
type AuthAuthenticateResponse struct {

	// JWT token used to authenticate against the API
	// Example: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzAB
	Jwt string `json:"jwt,omitempty"`
}

// Validate validates this auth authenticate response
func (m *AuthAuthenticateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth authenticate response based on context it is used
func (m *AuthAuthenticateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthAuthenticateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthAuthenticateResponse) UnmarshalBinary(b []byte) error {
	var res AuthAuthenticateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
