// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapTestLoginResponse ldap test login response
//
// swagger:model ldap.testLoginResponse
type LdapTestLoginResponse struct {

	// valid
	Valid *bool `json:"valid,omitempty"`
}

// Validate validates this ldap test login response
func (m *LdapTestLoginResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ldap test login response based on context it is used
func (m *LdapTestLoginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapTestLoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapTestLoginResponse) UnmarshalBinary(b []byte) error {
	var res LdapTestLoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
