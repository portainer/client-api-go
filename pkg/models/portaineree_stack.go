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

// PortainereeStack portaineree stack
//
// swagger:model portaineree.Stack
type PortainereeStack struct {

	// Only applies when deploying stack with multiple files
	AdditionalFiles []string `json:"AdditionalFiles"`

	// The GitOps update settings of a git stack
	AutoUpdate *PortainerAutoUpdateSettings `json:"AutoUpdate,omitempty"`

	// Environment(Endpoint) identifier. Reference the environment(endpoint) that will be used for deployment
	// Example: 1
	EndpointID int64 `json:"EndpointId,omitempty"`

	// Path to the Stack file
	// Example: docker-compose.yml
	EntryPoint string `json:"EntryPoint,omitempty"`

	// A list of environment(endpoint) variables used during stack deployment
	Env []*PortainerPair `json:"Env"`

	// Stack Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// Stack name
	// Example: myStack
	Name string `json:"Name,omitempty"`

	// The stack deployment option
	Option *PortainerStackOption `json:"Option,omitempty"`

	// The previous deployment info of the stack
	PreviousDeploymentInfo *PortainerStackDeploymentInfo `json:"PreviousDeploymentInfo,omitempty"`

	// resource control
	ResourceControl *PortainerResourceControl `json:"ResourceControl,omitempty"`

	// Stack status (1 - active, 2 - inactive)
	// Example: 1
	Status int64 `json:"Status,omitempty"`

	// Cluster identifier of the Swarm cluster where the stack is deployed
	// Example: jpofkc0i9uo9wtx1zesuk649w
	SwarmID string `json:"SwarmId,omitempty"`

	// Stack type. 1 for a Swarm stack, 2 for a Compose stack
	// Example: 2
	Type int64 `json:"Type,omitempty"`

	// The username which created this stack
	// Example: admin
	CreatedBy string `json:"createdBy,omitempty"`

	// The username id which created this stack
	// Example: 1
	CreatedByUserID string `json:"createdByUserId,omitempty"`

	// The date in unix time when stack was created
	// Example: 1587399600
	CreationDate int64 `json:"creationDate,omitempty"`

	// Network(Swarm) or local(Standalone) filesystem path
	// Example: /tmp
	FilesystemPath string `json:"filesystemPath,omitempty"`

	// Whether the stack is from a app template
	// Example: false
	FromAppTemplate bool `json:"fromAppTemplate,omitempty"`

	// The git config of this stack
	GitConfig *GittypesRepoConfig `json:"gitConfig,omitempty"`

	// Whether the stack is detached from git
	// Example: false
	IsDetachedFromGit bool `json:"isDetachedFromGit,omitempty"`

	// Kubernetes namespace if stack is a kube application
	// Example: default
	Namespace string `json:"namespace,omitempty"`

	// Path on disk to the repository hosting the Stack file
	// Example: /data/compose/myStack_jpofkc0i9uo9wtx1zesuk649w
	ProjectPath string `json:"projectPath,omitempty"`

	// StackFileVersion indicates the stack file version, such as yaml, hcl, and manifest
	// Example: 1
	StackFileVersion int64 `json:"stackFileVersion,omitempty"`

	// If stack support relative path volume
	// Example: false
	SupportRelativePath bool `json:"supportRelativePath,omitempty"`

	// The date in unix time when stack was last updated
	// Example: 1587399600
	UpdateDate int64 `json:"updateDate,omitempty"`

	// The username which last updated this stack
	// Example: bob
	UpdatedBy string `json:"updatedBy,omitempty"`

	// A UUID to identify a webhook. The stack will be force updated and pull the latest image when the webhook was invoked.
	// Example: c11fdf23-183e-428a-9bb6-16db01032174
	Webhook string `json:"webhook,omitempty"`
}

// Validate validates this portaineree stack
func (m *PortainereeStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousDeploymentInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeStack) validateAutoUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoUpdate) { // not required
		return nil
	}

	if m.AutoUpdate != nil {
		if err := m.AutoUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutoUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AutoUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) validateEnv(formats strfmt.Registry) error {
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
					return ve.ValidateName("Env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeStack) validateOption(formats strfmt.Registry) error {
	if swag.IsZero(m.Option) { // not required
		return nil
	}

	if m.Option != nil {
		if err := m.Option.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Option")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) validatePreviousDeploymentInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviousDeploymentInfo) { // not required
		return nil
	}

	if m.PreviousDeploymentInfo != nil {
		if err := m.PreviousDeploymentInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PreviousDeploymentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PreviousDeploymentInfo")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) validateResourceControl(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceControl) { // not required
		return nil
	}

	if m.ResourceControl != nil {
		if err := m.ResourceControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceControl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceControl")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) validateGitConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.GitConfig) { // not required
		return nil
	}

	if m.GitConfig != nil {
		if err := m.GitConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portaineree stack based on the context it is used
func (m *PortainereeStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreviousDeploymentInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeStack) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpdate != nil {

		if swag.IsZero(m.AutoUpdate) { // not required
			return nil
		}

		if err := m.AutoUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutoUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AutoUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {

			if swag.IsZero(m.Env[i]) { // not required
				return nil
			}

			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeStack) contextValidateOption(ctx context.Context, formats strfmt.Registry) error {

	if m.Option != nil {

		if swag.IsZero(m.Option) { // not required
			return nil
		}

		if err := m.Option.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Option")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) contextValidatePreviousDeploymentInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PreviousDeploymentInfo != nil {

		if swag.IsZero(m.PreviousDeploymentInfo) { // not required
			return nil
		}

		if err := m.PreviousDeploymentInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PreviousDeploymentInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PreviousDeploymentInfo")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) contextValidateResourceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceControl != nil {

		if swag.IsZero(m.ResourceControl) { // not required
			return nil
		}

		if err := m.ResourceControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceControl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceControl")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeStack) contextValidateGitConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.GitConfig != nil {

		if swag.IsZero(m.GitConfig) { // not required
			return nil
		}

		if err := m.GitConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeStack) UnmarshalBinary(b []byte) error {
	var res PortainereeStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
