// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EdgestacksSwarmStackFromGitRepositoryPayload edgestacks swarm stack from git repository payload
//
// swagger:model edgestacks.swarmStackFromGitRepositoryPayload
type EdgestacksSwarmStackFromGitRepositoryPayload struct {

	// Deployment type to deploy this stack
	// Valid values are: 0 - 'compose', 1 - 'kubernetes', 2 - 'nomad'
	// for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints)
	// kubernetes deploy type is enabled only for kubernetes environments(endpoints)
	// nomad deploy type is enabled only for nomad environments(endpoints)
	// Example: 0
	// Enum: [0 1 2]
	DeploymentType int64 `json:"deploymentType,omitempty"`

	// List of identifiers of EdgeGroups
	// Example: [1]
	EdgeGroups []int64 `json:"edgeGroups"`

	// Path to the Stack file inside the Git repository
	// Example: docker-compose.yml
	FilePathInRepository *string `json:"filePathInRepository,omitempty"`

	// Name of the stack
	// Example: myStack
	// Required: true
	Name *string `json:"name"`

	// Pre Pull image
	// Example: false
	PrePullImage *bool `json:"prePullImage,omitempty"`

	// List of Registries to use for this stack
	Registries []int64 `json:"registries"`

	// Use basic authentication to clone the Git repository
	// Example: true
	RepositoryAuthentication *bool `json:"repositoryAuthentication,omitempty"`

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

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces *bool `json:"useManifestNamespaces,omitempty"`
}

// Validate validates this edgestacks swarm stack from git repository payload
func (m *EdgestacksSwarmStackFromGitRepositoryPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentType(formats); err != nil {
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

var edgestacksSwarmStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksSwarmStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum = append(edgestacksSwarmStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum, v)
	}
}

// prop value enum
func (m *EdgestacksSwarmStackFromGitRepositoryPayload) validateDeploymentTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, edgestacksSwarmStackFromGitRepositoryPayloadTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksSwarmStackFromGitRepositoryPayload) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksSwarmStackFromGitRepositoryPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksSwarmStackFromGitRepositoryPayload) validateRepositoryURL(formats strfmt.Registry) error {

	if err := validate.Required("repositoryURL", "body", m.RepositoryURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edgestacks swarm stack from git repository payload based on context it is used
func (m *EdgestacksSwarmStackFromGitRepositoryPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromGitRepositoryPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromGitRepositoryPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksSwarmStackFromGitRepositoryPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
