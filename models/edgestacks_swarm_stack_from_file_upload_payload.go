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

// EdgestacksSwarmStackFromFileUploadPayload edgestacks swarm stack from file upload payload
//
// swagger:model edgestacks.swarmStackFromFileUploadPayload
type EdgestacksSwarmStackFromFileUploadPayload struct {

	// Deployment type to deploy this stack
	// Valid values are: 0 - 'compose', 1 - 'kubernetes', 2 - 'nomad'
	// for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints)
	// kubernetes deploytype is enabled only for kubernetes environments(endpoints)
	// nomad deploytype is enabled only for nomad environments(endpoints)
	// Example: 0
	// Enum: [0 1 2]
	DeploymentType int64 `json:"deploymentType,omitempty"`

	// edge groups
	EdgeGroups []int64 `json:"edgeGroups"`

	// name
	Name string `json:"name,omitempty"`

	// Pre Pull image
	// Example: false
	PrePullImage bool `json:"prePullImage,omitempty"`

	// registries
	Registries []int64 `json:"registries"`

	// stack file content
	StackFileContent []int64 `json:"stackFileContent"`

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces bool `json:"useManifestNamespaces,omitempty"`
}

// Validate validates this edgestacks swarm stack from file upload payload
func (m *EdgestacksSwarmStackFromFileUploadPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var edgestacksSwarmStackFromFileUploadPayloadTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksSwarmStackFromFileUploadPayloadTypeDeploymentTypePropEnum = append(edgestacksSwarmStackFromFileUploadPayloadTypeDeploymentTypePropEnum, v)
	}
}

// prop value enum
func (m *EdgestacksSwarmStackFromFileUploadPayload) validateDeploymentTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, edgestacksSwarmStackFromFileUploadPayloadTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksSwarmStackFromFileUploadPayload) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edgestacks swarm stack from file upload payload based on context it is used
func (m *EdgestacksSwarmStackFromFileUploadPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromFileUploadPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksSwarmStackFromFileUploadPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksSwarmStackFromFileUploadPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
