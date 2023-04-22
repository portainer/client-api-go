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

// PortainerEdgeStack portainer edge stack
//
// swagger:model portainer.EdgeStack
type PortainerEdgeStack struct {

	// creation date
	CreationDate int64 `json:"CreationDate,omitempty"`

	// edge groups
	EdgeGroups []int64 `json:"EdgeGroups"`

	// entry point
	EntryPoint string `json:"EntryPoint,omitempty"`

	// EdgeStack Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// num deployments
	NumDeployments int64 `json:"NumDeployments,omitempty"`

	// project path
	ProjectPath string `json:"ProjectPath,omitempty"`

	// Deprecated
	Prune bool `json:"Prune,omitempty"`

	// status
	Status map[string]PortainerEdgeStackStatus `json:"Status,omitempty"`

	// version
	Version int64 `json:"Version,omitempty"`

	// deployment type
	DeploymentType PortainerEdgeStackDeploymentType `json:"deploymentType,omitempty"`

	// manifest path
	ManifestPath string `json:"manifestPath,omitempty"`

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces bool `json:"useManifestNamespaces,omitempty"`
}

// Validate validates this portainer edge stack
func (m *PortainerEdgeStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerEdgeStack) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for k := range m.Status {

		if err := validate.Required("Status"+"."+k, "body", m.Status[k]); err != nil {
			return err
		}
		if val, ok := m.Status[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerEdgeStack) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if err := m.DeploymentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this portainer edge stack based on the context it is used
func (m *PortainerEdgeStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerEdgeStack) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Status {

		if val, ok := m.Status[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PortainerEdgeStack) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploymentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploymentType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEdgeStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEdgeStack) UnmarshalBinary(b []byte) error {
	var res PortainerEdgeStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
