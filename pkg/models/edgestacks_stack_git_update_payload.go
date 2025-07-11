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

// EdgestacksStackGitUpdatePayload edgestacks stack git update payload
//
// swagger:model edgestacks.stackGitUpdatePayload
type EdgestacksStackGitUpdatePayload struct {

	// authentication
	Authentication *GittypesGitAuthentication `json:"authentication,omitempty"`

	// auto update
	AutoUpdate *PortainerAutoUpdateSettings `json:"autoUpdate,omitempty"`

	// Options to control the deployer behaviour
	DeployerOptions *EdgestacksDeployerOptionsPayload `json:"deployerOptions,omitempty"`

	// Deployment type to deploy this stack
	// Valid values are: 0 - 'compose', 1 - 'kubernetes'
	// compose is enabled only for docker environments
	// kubernetes is enabled only for kubernetes environments
	// Example: 0
	// Enum: [0,1]
	DeploymentType int64 `json:"deploymentType,omitempty"`

	// env vars
	EnvVars []*PortainerPair `json:"envVars"`

	// group ids
	GroupIds []int64 `json:"groupIds"`

	// pre pull image
	PrePullImage bool `json:"prePullImage,omitempty"`

	// re pull image
	RePullImage bool `json:"rePullImage,omitempty"`

	// ref name
	RefName string `json:"refName,omitempty"`

	// List of Registries to use for this stack
	Registries []int64 `json:"registries"`

	// retry deploy
	RetryDeploy bool `json:"retryDeploy,omitempty"`

	// RetryPeriod specifies the duration, in seconds, for which the agent should continue attempting to deploy the stack after a failure
	RetryPeriod int64 `json:"retryPeriod,omitempty"`

	// Configuration for stagger updates
	StaggerConfig *PortainereeEdgeStaggerConfig `json:"staggerConfig,omitempty"`

	// Update the stack file content from the git repository
	// If this is set to true, it indicates that the stack is being redeployed,
	// if it is false, it indicates that the stack is being updated
	UpdateVersion bool `json:"updateVersion,omitempty"`
}

// Validate validates this edgestacks stack git update payload
func (m *EdgestacksStackGitUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployerOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
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

func (m *EdgestacksStackGitUpdatePayload) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *EdgestacksStackGitUpdatePayload) validateAutoUpdate(formats strfmt.Registry) error {
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

func (m *EdgestacksStackGitUpdatePayload) validateDeployerOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DeployerOptions) { // not required
		return nil
	}

	if m.DeployerOptions != nil {
		if err := m.DeployerOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployerOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployerOptions")
			}
			return err
		}
	}

	return nil
}

var edgestacksStackGitUpdatePayloadTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgestacksStackGitUpdatePayloadTypeDeploymentTypePropEnum = append(edgestacksStackGitUpdatePayloadTypeDeploymentTypePropEnum, v)
	}
}

// prop value enum
func (m *EdgestacksStackGitUpdatePayload) validateDeploymentTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, edgestacksStackGitUpdatePayloadTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgestacksStackGitUpdatePayload) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deploymentType", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

func (m *EdgestacksStackGitUpdatePayload) validateEnvVars(formats strfmt.Registry) error {
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

func (m *EdgestacksStackGitUpdatePayload) validateStaggerConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this edgestacks stack git update payload based on the context it is used
func (m *EdgestacksStackGitUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployerOptions(ctx, formats); err != nil {
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

func (m *EdgestacksStackGitUpdatePayload) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {

		if swag.IsZero(m.Authentication) { // not required
			return nil
		}

		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *EdgestacksStackGitUpdatePayload) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgestacksStackGitUpdatePayload) contextValidateDeployerOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DeployerOptions != nil {

		if swag.IsZero(m.DeployerOptions) { // not required
			return nil
		}

		if err := m.DeployerOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployerOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployerOptions")
			}
			return err
		}
	}

	return nil
}

func (m *EdgestacksStackGitUpdatePayload) contextValidateEnvVars(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnvVars); i++ {

		if m.EnvVars[i] != nil {

			if swag.IsZero(m.EnvVars[i]) { // not required
				return nil
			}

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

func (m *EdgestacksStackGitUpdatePayload) contextValidateStaggerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StaggerConfig != nil {

		if swag.IsZero(m.StaggerConfig) { // not required
			return nil
		}

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
func (m *EdgestacksStackGitUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksStackGitUpdatePayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksStackGitUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
