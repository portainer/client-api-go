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

// StacksStackGitUpdatePayload stacks stack git update payload
//
// swagger:model stacks.stackGitUpdatePayload
type StacksStackGitUpdatePayload struct {

	// auto update
	AutoUpdate *PortainerAutoUpdateSettings `json:"autoUpdate,omitempty"`

	// env
	Env []*PortainerPair `json:"env"`

	// prune
	Prune bool `json:"prune,omitempty"`

	// repository authentication
	RepositoryAuthentication bool `json:"repositoryAuthentication,omitempty"`

	// repository git credential ID
	RepositoryGitCredentialID int64 `json:"repositoryGitCredentialID,omitempty"`

	// repository password
	RepositoryPassword string `json:"repositoryPassword,omitempty"`

	// repository reference name
	RepositoryReferenceName string `json:"repositoryReferenceName,omitempty"`

	// repository username
	RepositoryUsername string `json:"repositoryUsername,omitempty"`

	// tlsskip verify
	TlsskipVerify bool `json:"tlsskipVerify,omitempty"`
}

// Validate validates this stacks stack git update payload
func (m *StacksStackGitUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksStackGitUpdatePayload) validateAutoUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoUpdate) { // not required
		return nil
	}

	if m.AutoUpdate != nil {
		if err := m.AutoUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *StacksStackGitUpdatePayload) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stacks stack git update payload based on the context it is used
func (m *StacksStackGitUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksStackGitUpdatePayload) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpdate != nil {

		if swag.IsZero(m.AutoUpdate) { // not required
			return nil
		}

		if err := m.AutoUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *StacksStackGitUpdatePayload) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {

			if swag.IsZero(m.Env[i]) { // not required
				return nil
			}

			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksStackGitUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksStackGitUpdatePayload) UnmarshalBinary(b []byte) error {
	var res StacksStackGitUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
