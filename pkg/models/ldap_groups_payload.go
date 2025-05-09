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

// LdapGroupsPayload ldap groups payload
//
// swagger:model ldap.groupsPayload
type LdapGroupsPayload struct {

	// ldapsettings
	Ldapsettings *PortainereeLDAPSettings `json:"ldapsettings,omitempty"`
}

// Validate validates this ldap groups payload
func (m *LdapGroupsPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapsettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapGroupsPayload) validateLdapsettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Ldapsettings) { // not required
		return nil
	}

	if m.Ldapsettings != nil {
		if err := m.Ldapsettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapsettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldapsettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap groups payload based on the context it is used
func (m *LdapGroupsPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLdapsettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapGroupsPayload) contextValidateLdapsettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Ldapsettings != nil {

		if swag.IsZero(m.Ldapsettings) { // not required
			return nil
		}

		if err := m.Ldapsettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapsettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldapsettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapGroupsPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapGroupsPayload) UnmarshalBinary(b []byte) error {
	var res LdapGroupsPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
