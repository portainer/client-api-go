// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerStack portainer stack
//
// swagger:model portainer.Stack
type PortainerStack struct {

	// Only applies when deploying stack with multiple files
	AdditionalFiles []string `json:"AdditionalFiles"`

	// The auto update settings of a git stack
	AutoUpdate *PortainerStackAutoUpdate `json:"AutoUpdate,omitempty"`

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

	// The date in unix time when stack was created
	// Example: 1587399600
	CreationDate int64 `json:"creationDate,omitempty"`

	// Whether the stack is from a app template
	// Example: false
	FromAppTemplate *bool `json:"fromAppTemplate,omitempty"`

	// The git config of this stack
	GitConfig *GittypesRepoConfig `json:"gitConfig,omitempty"`

	// IsComposeFormat indicates if the Kubernetes stack is created from a Docker Compose file
	// Example: false
	IsComposeFormat *bool `json:"isComposeFormat,omitempty"`

	// Kubernetes namespace if stack is a kube application
	// Example: default
	Namespace string `json:"namespace,omitempty"`

	// Path on disk to the repository hosting the Stack file
	// Example: /data/compose/myStack_jpofkc0i9uo9wtx1zesuk649w
	ProjectPath string `json:"projectPath,omitempty"`

	// The date in unix time when stack was last updated
	// Example: 1587399600
	UpdateDate int64 `json:"updateDate,omitempty"`

	// The username which last updated this stack
	// Example: bob
	UpdatedBy string `json:"updatedBy,omitempty"`
}

// Validate validates this portainer stack
func (m *PortainerStack) Validate(formats strfmt.Registry) error {
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

func (m *PortainerStack) validateAutoUpdate(formats strfmt.Registry) error {
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

func (m *PortainerStack) validateEnv(formats strfmt.Registry) error {
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

func (m *PortainerStack) validateOption(formats strfmt.Registry) error {
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

func (m *PortainerStack) validateResourceControl(formats strfmt.Registry) error {
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

func (m *PortainerStack) validateGitConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this portainer stack based on the context it is used
func (m *PortainerStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PortainerStack) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoUpdate != nil {
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

func (m *PortainerStack) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {
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

func (m *PortainerStack) contextValidateOption(ctx context.Context, formats strfmt.Registry) error {

	if m.Option != nil {
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

func (m *PortainerStack) contextValidateResourceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceControl != nil {
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

func (m *PortainerStack) contextValidateGitConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.GitConfig != nil {
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
func (m *PortainerStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerStack) UnmarshalBinary(b []byte) error {
	var res PortainerStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
