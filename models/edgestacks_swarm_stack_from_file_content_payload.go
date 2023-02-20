// Code generated by go-swagger; DO NOT EDIT.

package models

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

// EdgestacksSwarmStackFromFileContentPayload edgestacks swarm stack from file content payload
//
// swagger:model edgestacks.swarmStackFromFileContentPayload
type EdgestacksSwarmStackFromFileContentPayload struct {

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

	// Name of the stack
	// Example: myStack
	// Required: true
	Name *string `json:"name"`

	// Pre Pull image
	// Example: false
	PrePullImage bool `json:"prePullImage,omitempty"`

	// List of Registries to use for this stack
	Registries []int64 `json:"registries"`

	// Content of the Stack file
	// Example: version: 3\n services:\n web:\n image:nginx
	// Required: true
	StackFileContent *string `json:"stackFileContent"`

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces bool `json:"useManifestNamespaces,omitempty"`
}

// Validate validates this edgestacks swarm stack from file content payload
func (m *EdgestacksSwarmStackFromFileContentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackFileContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var edgestacksSwarmStackFromFileContentPayloadTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksSwarmStackFromFileContentPayloadTypeDeploymentTypePropEnum = append(edgestacksSwarmStackFromFileContentPayloadTypeDeploymentTypePropEnum, v)
	}
}

// prop value enum
func (m *EdgestacksSwarmStackFromFileContentPayload) validateDeploymentTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, edgestacksSwarmStackFromFileContentPayloadTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksSwarmStackFromFileContentPayload) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksSwarmStackFromFileContentPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksSwarmStackFromFileContentPayload) validateStackFileContent(formats strfmt.Registry) error {

	if err := validate.Required("stackFileContent", "body", m.StackFileContent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edgestacks swarm stack from file content payload based on context it is used
func (m *EdgestacksSwarmStackFromFileContentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromFileContentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromFileContentPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksSwarmStackFromFileContentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
