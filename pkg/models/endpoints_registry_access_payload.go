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

// EndpointsRegistryAccessPayload endpoints registry access payload
//
// swagger:model endpoints.registryAccessPayload
type EndpointsRegistryAccessPayload struct {

	// namespaces
	Namespaces []string `json:"namespaces"`

	// team access policies
	TeamAccessPolicies PortainerTeamAccessPolicies `json:"teamAccessPolicies,omitempty"`

	// user access policies
	UserAccessPolicies PortainerUserAccessPolicies `json:"userAccessPolicies,omitempty"`
}

// Validate validates this endpoints registry access payload
func (m *EndpointsRegistryAccessPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsRegistryAccessPayload) validateTeamAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamAccessPolicies) { // not required
		return nil
	}

	if m.TeamAccessPolicies != nil {
		if err := m.TeamAccessPolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamAccessPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamAccessPolicies")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsRegistryAccessPayload) validateUserAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.UserAccessPolicies) { // not required
		return nil
	}

	if m.UserAccessPolicies != nil {
		if err := m.UserAccessPolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userAccessPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userAccessPolicies")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoints registry access payload based on the context it is used
func (m *EndpointsRegistryAccessPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeamAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsRegistryAccessPayload) contextValidateTeamAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.TeamAccessPolicies) { // not required
		return nil
	}

	if err := m.TeamAccessPolicies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("teamAccessPolicies")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("teamAccessPolicies")
		}
		return err
	}

	return nil
}

func (m *EndpointsRegistryAccessPayload) contextValidateUserAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.UserAccessPolicies) { // not required
		return nil
	}

	if err := m.UserAccessPolicies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userAccessPolicies")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userAccessPolicies")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsRegistryAccessPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsRegistryAccessPayload) UnmarshalBinary(b []byte) error {
	var res EndpointsRegistryAccessPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
