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

// PortainereeLDAPSettings portaineree l d a p settings
//
// swagger:model portaineree.LDAPSettings
type PortainereeLDAPSettings struct {

	// Whether auto admin population is switched on or not
	// Example: true
	AdminAutoPopulate *bool `json:"AdminAutoPopulate,omitempty"`

	// admin group search settings
	AdminGroupSearchSettings []*PortainereeLDAPGroupSearchSettings `json:"AdminGroupSearchSettings"`

	// Saved admin group list, the user will be populated as an admin role if any user group matches the record in the list
	// Example: ["['manager'","'operator']"]
	AdminGroups []string `json:"AdminGroups"`

	// Enable this option if the server is configured for Anonymous access. When enabled, ReaderDN and Password will not be used
	// Example: true
	AnonymousMode *bool `json:"AnonymousMode,omitempty"`

	// Automatically provision users and assign them to matching LDAP group names
	// Example: true
	AutoCreateUsers *bool `json:"AutoCreateUsers,omitempty"`

	// group search settings
	GroupSearchSettings []*PortainereeLDAPGroupSearchSettings `json:"GroupSearchSettings"`

	// Password of the account that will be used to search users
	// Example: readonly-password
	Password string `json:"Password,omitempty"`

	// Account that will be used to search for users
	// Example: cn=readonly-account,dc=ldap,dc=domain,dc=tld
	ReaderDN string `json:"ReaderDN,omitempty"`

	// search settings
	SearchSettings []*PortainereeLDAPSearchSettings `json:"SearchSettings"`

	// server type
	// Example: 1
	ServerType int64 `json:"ServerType,omitempty"`

	// Whether LDAP connection should use StartTLS
	// Example: true
	StartTLS *bool `json:"StartTLS,omitempty"`

	// TLS config
	TLSConfig *PortainereeTLSConfiguration `json:"TLSConfig,omitempty"`

	// Deprecated
	URL string `json:"URL,omitempty"`

	// URLs or IP addresses of the LDAP server
	URLs []string `json:"URLs"`
}

// Validate validates this portaineree l d a p settings
func (m *PortainereeLDAPSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminGroupSearchSettings(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PortainereeLDAPSettings) validateAdminGroupSearchSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminGroupSearchSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.AdminGroupSearchSettings); i++ {
		if swag.IsZero(m.AdminGroupSearchSettings[i]) { // not required
			continue
		}

		if m.AdminGroupSearchSettings[i] != nil {
			if err := m.AdminGroupSearchSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdminGroupSearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AdminGroupSearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeLDAPSettings) validateGroupSearchSettings(formats strfmt.Registry) error {
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

func (m *PortainereeLDAPSettings) validateSearchSettings(formats strfmt.Registry) error {
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

func (m *PortainereeLDAPSettings) validateTLSConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this portaineree l d a p settings based on the context it is used
func (m *PortainereeLDAPSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminGroupSearchSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *PortainereeLDAPSettings) contextValidateAdminGroupSearchSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdminGroupSearchSettings); i++ {

		if m.AdminGroupSearchSettings[i] != nil {
			if err := m.AdminGroupSearchSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdminGroupSearchSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AdminGroupSearchSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeLDAPSettings) contextValidateGroupSearchSettings(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainereeLDAPSettings) contextValidateSearchSettings(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainereeLDAPSettings) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortainereeLDAPSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeLDAPSettings) UnmarshalBinary(b []byte) error {
	var res PortainereeLDAPSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
