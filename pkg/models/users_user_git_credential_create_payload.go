// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsersUserGitCredentialCreatePayload users user git credential create payload
//
// swagger:model users.userGitCredentialCreatePayload
type UsersUserGitCredentialCreatePayload struct {

	// name
	// Example: my-credential
	// Required: true
	Name *string `json:"name"`

	// password
	// Required: true
	Password *string `json:"password"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this users user git credential create payload
func (m *UsersUserGitCredentialCreatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
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

func (m *UsersUserGitCredentialCreatePayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UsersUserGitCredentialCreatePayload) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UsersUserGitCredentialCreatePayload) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this users user git credential create payload based on context it is used
func (m *UsersUserGitCredentialCreatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsersUserGitCredentialCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersUserGitCredentialCreatePayload) UnmarshalBinary(b []byte) error {
	var res UsersUserGitCredentialCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
