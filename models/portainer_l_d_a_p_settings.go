// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerLDAPSettings portainer l d a p settings
//
// swagger:model portainer.LDAPSettings
type PortainerLDAPSettings struct {

	// Enable this option if the server is configured for Anonymous access. When enabled, ReaderDN and Password will not be used
	// Example: true
	AnonymousMode bool `json:"AnonymousMode,omitempty"`

	// Automatically provision users and assign them to matching LDAP group names
	// Example: true
	AutoCreateUsers bool `json:"AutoCreateUsers,omitempty"`

	// group search settings
	GroupSearchSettings []*PortainerLDAPGroupSearchSettings `json:"GroupSearchSettings"`

	// Password of the account that will be used to search users
	// Example: readonly-password
	Password string `json:"Password,omitempty"`

	// Account that will be used to search for users
	// Example: cn=readonly-account,dc=ldap,dc=domain,dc=tld
	ReaderDN string `json:"ReaderDN,omitempty"`

	// search settings
	SearchSettings []*PortainerLDAPSearchSettings `json:"SearchSettings"`

	// Whether LDAP connection should use StartTLS
	// Example: true
	StartTLS bool `json:"StartTLS,omitempty"`

	// TLS config
	TLSConfig *PortainerTLSConfiguration `json:"TLSConfig,omitempty"`

	// URL or IP address of the LDAP server
	// Example: myldap.domain.tld:389
	URL string `json:"URL,omitempty"`
}

// Validate validates this portainer l d a p settings
func (m *PortainerLDAPSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSearchSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerLDAPSettings) validateGroupSearchSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupSearchSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupSearchSettings); i++ {
		if swag.IsZero(m.GroupSearchSettings[i]) { // not required
			continue
		}

		if m.GroupSearchSettings[i] != nil {
			if err := m.GroupSearchSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GroupSearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GroupSearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerLDAPSettings) validateSearchSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.SearchSettings); i++ {
		if swag.IsZero(m.SearchSettings[i]) { // not required
			continue
		}

		if m.SearchSettings[i] != nil {
			if err := m.SearchSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerLDAPSettings) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portainer l d a p settings based on the context it is used
func (m *PortainerLDAPSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupSearchSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerLDAPSettings) contextValidateGroupSearchSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupSearchSettings); i++ {

		if m.GroupSearchSettings[i] != nil {
			if err := m.GroupSearchSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GroupSearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GroupSearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerLDAPSettings) contextValidateSearchSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SearchSettings); i++ {

		if m.SearchSettings[i] != nil {
			if err := m.SearchSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerLDAPSettings) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {
		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerLDAPSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerLDAPSettings) UnmarshalBinary(b []byte) error {
	var res PortainerLDAPSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
