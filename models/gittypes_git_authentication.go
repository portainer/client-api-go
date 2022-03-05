// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GittypesGitAuthentication gittypes git authentication
//
// swagger:model gittypes.GitAuthentication
type GittypesGitAuthentication struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this gittypes git authentication
func (m *GittypesGitAuthentication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gittypes git authentication based on context it is used
func (m *GittypesGitAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GittypesGitAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GittypesGitAuthentication) UnmarshalBinary(b []byte) error {
	var res GittypesGitAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
