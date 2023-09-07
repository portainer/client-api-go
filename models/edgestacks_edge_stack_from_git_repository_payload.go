// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EdgestacksEdgeStackFromGitRepositoryPayload edgestacks edge stack from git repository payload
//
// swagger:model edgestacks.edgeStackFromGitRepositoryPayload
type EdgestacksEdgeStackFromGitRepositoryPayload struct {

	// Optional GitOps update configuration
	AutoUpdate *PortainereeAutoUpdateSettings `json:"autoUpdate,omitempty"`

	// Deployment type to deploy this stack
	// Valid values are: 0 - 'compose', 1 - 'kubernetes', 2 - 'nomad'
	// compose is enabled only for docker environments
	// kubernetes is enabled only for kubernetes environments
	// nomad is enabled only for nomad environments
	// Example: 0
	// Enum: [0 1 2]
	DeploymentType int64 `json:"deploymentType,omitempty"`

	// List of identifiers of EdgeGroups
	// Example: [1]
	EdgeGroups []int64 `json:"edgeGroups"`

	// List of environment variables
	EnvVars []*PortainerPair `json:"envVars"`

	// Path to the Stack file inside the Git repository
	// Example: docker-compose.yml
	FilePathInRepository *string `json:"filePathInRepository,omitempty"`

	// Local filesystem path
	// Example: /mnt
	FilesystemPath string `json:"filesystemPath,omitempty"`

	// Name of the stack
	// Example: myStack
	// Required: true
	Name *string `json:"name"`

	// Per device configs match type
	// Example: file
	// Enum: [file  dir]
	PerDeviceConfigsMatchType string `json:"perDeviceConfigsMatchType,omitempty"`

	// Per device configs path
	// Example: configs
	PerDeviceConfigsPath string `json:"perDeviceConfigsPath,omitempty"`

	// Pre Pull image
	// Example: false
	PrePullImage bool `json:"prePullImage,omitempty"`

	// List of Registries to use for this stack
	Registries []int64 `json:"registries"`

	// Use basic authentication to clone the Git repository
	// Example: true
	RepositoryAuthentication bool `json:"repositoryAuthentication,omitempty"`

	// GitCredentialID used to identify the binded git credential
	// Example: 0
	RepositoryGitCredentialID int64 `json:"repositoryGitCredentialID,omitempty"`

	// Password used in basic authentication. Required when RepositoryAuthentication is true.
	// Example: myGitPassword
	RepositoryPassword string `json:"repositoryPassword,omitempty"`

	// Reference name of a Git repository hosting the Stack file
	// Example: refs/heads/master
	RepositoryReferenceName string `json:"repositoryReferenceName,omitempty"`

	// URL of a Git repository hosting the Stack file
	// Example: https://github.com/openfaas/faas
	// Required: true
	RepositoryURL *string `json:"repositoryURL"`

	// Username used in basic authentication. Required when RepositoryAuthentication is true.
	// Example: myGitUsername
	RepositoryUsername string `json:"repositoryUsername,omitempty"`

	// Retry deploy
	// Example: false
	RetryDeploy bool `json:"retryDeploy,omitempty"`

	// Configuration for stagger updates
	StaggerConfig *PortainereeEdgeStaggerConfig `json:"staggerConfig,omitempty"`

	// Whether the edge stack supports per device configs
	// Example: false
	SupportPerDeviceConfigs bool `json:"supportPerDeviceConfigs,omitempty"`

	// Whether the stack supports relative path volume
	// Example: false
	SupportRelativePath bool `json:"supportRelativePath,omitempty"`

	// TLSSkipVerify skips SSL verification when cloning the Git repository
	// Example: false
	TlsskipVerify bool `json:"tlsskipVerify,omitempty"`

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces bool `json:"useManifestNamespaces,omitempty"`
}

// Validate validates this edgestacks edge stack from git repository payload
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerDeviceConfigsMatchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaggerConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateAutoUpdate(formats strfmt.Registry) error {
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

var edgestacksEdgeStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksEdgeStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum = append(edgestacksEdgeStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum, v)
	}
}

// prop value enum
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateDeploymentTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, edgestacksEdgeStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateEnvVars(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvVars) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvVars); i++ {
		if swag.IsZero(m.EnvVars[i]) { // not required
			continue
		}

		if m.EnvVars[i] != nil {
			if err := m.EnvVars[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envVars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envVars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var edgestacksEdgeStackFromGitRepositoryPayloadTypePerDeviceConfigsMatchTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file"," dir"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksEdgeStackFromGitRepositoryPayloadTypePerDeviceConfigsMatchTypePropEnum = append(edgestacksEdgeStackFromGitRepositoryPayloadTypePerDeviceConfigsMatchTypePropEnum, v)
	}
}

const (

	// EdgestacksEdgeStackFromGitRepositoryPayloadPerDeviceConfigsMatchTypeFile captures enum value "file"
	EdgestacksEdgeStackFromGitRepositoryPayloadPerDeviceConfigsMatchTypeFile string = "file"

	// EdgestacksEdgeStackFromGitRepositoryPayloadPerDeviceConfigsMatchTypeDir captures enum value " dir"
	EdgestacksEdgeStackFromGitRepositoryPayloadPerDeviceConfigsMatchTypeDir string = " dir"
)

// prop value enum
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validatePerDeviceConfigsMatchTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgestacksEdgeStackFromGitRepositoryPayloadTypePerDeviceConfigsMatchTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validatePerDeviceConfigsMatchType(formats strfmt.Registry) error {
	if swag.IsZero(m.PerDeviceConfigsMatchType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePerDeviceConfigsMatchTypeEnum("perDeviceConfigsMatchType", "body", m.PerDeviceConfigsMatchType); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateRepositoryURL(formats strfmt.Registry) error {

	if err := validate.Required("repositoryURL", "body", m.RepositoryURL); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) validateStaggerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StaggerConfig) { // not required
		return nil
	}

	if m.StaggerConfig != nil {
		if err := m.StaggerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("staggerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("staggerConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edgestacks edge stack from git repository payload based on the context it is used
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvVars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaggerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) contextValidateEnvVars(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnvVars); i++ {

		if m.EnvVars[i] != nil {
			if err := m.EnvVars[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envVars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envVars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgestacksEdgeStackFromGitRepositoryPayload) contextValidateStaggerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StaggerConfig != nil {
		if err := m.StaggerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("staggerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("staggerConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksEdgeStackFromGitRepositoryPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksEdgeStackFromGitRepositoryPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
