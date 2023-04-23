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

// StacksKubernetesGitDeploymentPayload stacks kubernetes git deployment payload
//
// swagger:model stacks.kubernetesGitDeploymentPayload
type StacksKubernetesGitDeploymentPayload struct {

	// additional files
	AdditionalFiles []string `json:"additionalFiles"`

	// auto update
	AutoUpdate *PortainereeAutoUpdateSettings `json:"autoUpdate,omitempty"`

	// compose format
	ComposeFormat *bool `json:"composeFormat,omitempty"`

	// manifest file
	ManifestFile string `json:"manifestFile,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// repository authentication
	RepositoryAuthentication *bool `json:"repositoryAuthentication,omitempty"`

	// repository git credential ID
	RepositoryGitCredentialID int64 `json:"repositoryGitCredentialID,omitempty"`

	// repository password
	RepositoryPassword string `json:"repositoryPassword,omitempty"`

	// repository reference name
	RepositoryReferenceName string `json:"repositoryReferenceName,omitempty"`

	// repository URL
	RepositoryURL string `json:"repositoryURL,omitempty"`

	// repository username
	RepositoryUsername string `json:"repositoryUsername,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// TLSSkipVerify skips SSL verification when cloning the Git repository
	// Example: false
	TlsskipVerify *bool `json:"tlsskipVerify,omitempty"`
}

// Validate validates this stacks kubernetes git deployment payload
func (m *StacksKubernetesGitDeploymentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksKubernetesGitDeploymentPayload) validateAutoUpdate(formats strfmt.Registry) error {
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

// ContextValidate validate this stacks kubernetes git deployment payload based on the context it is used
func (m *StacksKubernetesGitDeploymentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksKubernetesGitDeploymentPayload) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *StacksKubernetesGitDeploymentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksKubernetesGitDeploymentPayload) UnmarshalBinary(b []byte) error {
	var res StacksKubernetesGitDeploymentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
