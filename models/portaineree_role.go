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

// PortainereeRole portaineree role
//
// swagger:model portaineree.Role
type PortainereeRole struct {

	// Authorizations associated to a role
	// Required: true
	Authorizations PortainereeAuthorizations `json:"Authorizations"`

	// Role description
	// Example: Read-only access of all resources in an environment(endpoint)
	// Required: true
	Description *string `json:"Description"`

	// Role Identifier
	// Example: 1
	// Required: true
	ID *int64 `json:"Id"`

	// Role name
	// Example: HelpDesk
	// Required: true
	Name *string `json:"Name"`

	// priority
	// Required: true
	Priority *int64 `json:"Priority"`
}

// Validate validates this portaineree role
func (m *PortainereeRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeRole) validateAuthorizations(formats strfmt.Registry) error {

	if err := validate.Required("Authorizations", "body", m.Authorizations); err != nil {
		return err
	}

	if m.Authorizations != nil {
		if err := m.Authorizations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authorizations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Authorizations")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeRole) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PortainereeRole) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PortainereeRole) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PortainereeRole) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("Priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this portaineree role based on the context it is used
func (m *PortainereeRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeRole) contextValidateAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Authorizations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Authorizations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Authorizations")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeRole) UnmarshalBinary(b []byte) error {
	var res PortainereeRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
