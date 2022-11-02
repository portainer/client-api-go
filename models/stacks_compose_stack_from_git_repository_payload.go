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
	"github.com/go-openapi/validate"
)

// StacksComposeStackFromGitRepositoryPayload stacks compose stack from git repository payload
//
// swagger:model stacks.composeStackFromGitRepositoryPayload
type StacksComposeStackFromGitRepositoryPayload struct {

	// Applicable when deploying with multiple stack files
	// Example: ["[nz.compose.yml"," uat.compose.yml]"]
	AdditionalFiles []string `json:"additionalFiles"`

	// Optional auto update configuration
	AutoUpdate *PortainereeStackAutoUpdate `json:"autoUpdate,omitempty"`

	// Path to the Stack file inside the Git repository
	// Example: docker-compose.yml
	ComposeFile *string `json:"composeFile,omitempty"`

	// A list of environment(endpoint) variables used during stack deployment
	Env []*PortainereePair `json:"env"`

	// Whether the stack is from a app template
	// Example: false
	FromAppTemplate bool `json:"fromAppTemplate,omitempty"`

	// Name of the stack
	// Example: myStack
	// Required: true
	Name *string `json:"name"`

	// Use basic authentication to clone the Git repository
	// Example: true
	RepositoryAuthentication bool `json:"repositoryAuthentication,omitempty"`

	// GitCredentialID used to identify the bound git credential. Required when RepositoryAuthentication
	// is true and RepositoryUsername/RepositoryPassword are not provided
	// Example: 0
	RepositoryGitCredentialID int64 `json:"repositoryGitCredentialID,omitempty"`

	// Password used in basic authentication. Required when RepositoryAuthentication is true
	// and RepositoryGitCredentialID is 0
	// Example: myGitPassword
	RepositoryPassword string `json:"repositoryPassword,omitempty"`

	// Reference name of a Git repository hosting the Stack file
	// Example: refs/heads/master
	RepositoryReferenceName string `json:"repositoryReferenceName,omitempty"`

	// URL of a Git repository hosting the Stack file
	// Example: https://github.com/openfaas/faas
	// Required: true
	RepositoryURL *string `json:"repositoryURL"`

	// Username used in basic authentication. Required when RepositoryAuthentication is true
	// and RepositoryGitCredentialID is 0
	// Example: myGitUsername
	RepositoryUsername string `json:"repositoryUsername,omitempty"`
}

// Validate validates this stacks compose stack from git repository payload
func (m *StacksComposeStackFromGitRepositoryPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksComposeStackFromGitRepositoryPayload) validateAutoUpdate(formats strfmt.Registry) error {
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

func (m *StacksComposeStackFromGitRepositoryPayload) validateEnv(formats strfmt.Registry) error {
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

func (m *StacksComposeStackFromGitRepositoryPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StacksComposeStackFromGitRepositoryPayload) validateRepositoryURL(formats strfmt.Registry) error {

	if err := validate.Required("repositoryURL", "body", m.RepositoryURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stacks compose stack from git repository payload based on the context it is used
func (m *StacksComposeStackFromGitRepositoryPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *StacksComposeStackFromGitRepositoryPayload) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpdate != nil {
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

func (m *StacksComposeStackFromGitRepositoryPayload) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {
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
func (m *StacksComposeStackFromGitRepositoryPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksComposeStackFromGitRepositoryPayload) UnmarshalBinary(b []byte) error {
	var res StacksComposeStackFromGitRepositoryPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
