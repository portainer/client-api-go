// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeLDAPGroupSearchSettings portaineree l d a p group search settings
//
// swagger:model portaineree.LDAPGroupSearchSettings
type PortainereeLDAPGroupSearchSettings struct {

	// LDAP attribute which denotes the group membership
	// Example: member
	GroupAttribute string `json:"GroupAttribute,omitempty"`

	// The distinguished name of the element from which the LDAP server will search for groups
	// Example: dc=ldap,dc=domain,dc=tld
	GroupBaseDN string `json:"GroupBaseDN,omitempty"`

	// The LDAP search filter used to select group elements, optional
	// Example: (objectClass=account
	GroupFilter string `json:"GroupFilter,omitempty"`
}

// Validate validates this portaineree l d a p group search settings
func (m *PortainereeLDAPGroupSearchSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree l d a p group search settings based on context it is used
func (m *PortainereeLDAPGroupSearchSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeLDAPGroupSearchSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeLDAPGroupSearchSettings) UnmarshalBinary(b []byte) error {
	var res PortainereeLDAPGroupSearchSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
