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

// PortainereeEdgeStack portaineree edge stack
//
// swagger:model portaineree.EdgeStack
type PortainereeEdgeStack struct {

	// The GitOps update settings of a git stack
	AutoUpdate *PortainereeAutoUpdateSettings `json:"AutoUpdate,omitempty"`

	// creation date
	CreationDate int64 `json:"CreationDate,omitempty"`

	// deployment type
	DeploymentType int64 `json:"DeploymentType,omitempty"`

	// edge groups
	EdgeGroups []int64 `json:"EdgeGroups"`

	// entry point
	EntryPoint string `json:"EntryPoint,omitempty"`

	// EdgeStack Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// manifest path
	ManifestPath string `json:"ManifestPath,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// num deployments
	NumDeployments int64 `json:"NumDeployments,omitempty"`

	// Pre Pull Image
	PrePullImage bool `json:"PrePullImage,omitempty"`

	// PreviousDeploymentInfo represents the previous deployment info of the stack
	PreviousDeploymentInfo *PortainerStackDeploymentInfo `json:"PreviousDeploymentInfo,omitempty"`

	// project path
	ProjectPath string `json:"ProjectPath,omitempty"`

	// Deprecated
	Prune bool `json:"Prune,omitempty"`

	// Re-Pull Image
	RePullImage bool `json:"RePullImage,omitempty"`

	// registries
	Registries []int64 `json:"Registries"`

	// StackFileVersion represents the version of the stack file, such yaml, hcl, manifest file
	// Example: 1
	StackFileVersion int64 `json:"StackFileVersion,omitempty"`

	// status
	Status map[string]PortainerEdgeStackStatus `json:"Status,omitempty"`

	// version
	Version int64 `json:"Version,omitempty"`

	// EdgeUpdateID represents the parent update ID, will be zero if this stack is not part of an update
	EdgeUpdateID int64 `json:"edgeUpdateID,omitempty"`

	// EnvVars is a list of environment variables to inject into the stack
	EnvVars []*PortainerPair `json:"envVars"`

	// Local filesystem path
	// Example: /tmp
	FilesystemPath string `json:"filesystemPath,omitempty"`

	// The git configuration of a git stack
	GitConfig *GittypesRepoConfig `json:"gitConfig,omitempty"`

	// Per device configs match type
	// Example: file
	// Enum: [file  dir]
	PerDeviceConfigsMatchType string `json:"perDeviceConfigsMatchType,omitempty"`

	// Per device configs path
	// Example: configs
	PerDeviceConfigsPath string `json:"perDeviceConfigsPath,omitempty"`

	// Retry deploy
	// Example: false
	RetryDeploy bool `json:"retryDeploy,omitempty"`

	// Schedule represents the schedule of the Edge stack (optional, format - 'YYYY-MM-DD HH:mm:ss')
	// Example: 2020-11-13 14:53:00
	ScheduledTime string `json:"scheduledTime,omitempty"`

	// StaggerConfig is the configuration for staggered update
	StaggerConfig *PortainereeEdgeStaggerConfig `json:"staggerConfig,omitempty"`

	// Whether the edge stack supports per device configs
	// Example: false
	SupportPerDeviceConfigs bool `json:"supportPerDeviceConfigs,omitempty"`

	// Whether the stack supports relative path volume
	// Example: false
	SupportRelativePath bool `json:"supportRelativePath,omitempty"`

	// Uses the manifest's namespaces instead of the default one
	UseManifestNamespaces bool `json:"useManifestNamespaces,omitempty"`

	// A UUID to identify a webhook. The stack will be force updated and pull the latest image when the webhook was invoked.
	// Example: c11fdf23-183e-428a-9bb6-16db01032174
	Webhook string `json:"webhook,omitempty"`
}

// Validate validates this portaineree edge stack
func (m *PortainereeEdgeStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousDeploymentInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerDeviceConfigsMatchType(formats); err != nil {
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

func (m *PortainereeEdgeStack) validateAutoUpdate(formats strfmt.Registry) error {
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

func (m *PortainereeEdgeStack) validatePreviousDeploymentInfo(formats strfmt.Registry) error {
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

func (m *PortainereeEdgeStack) validateStatus(formats strfmt.Registry) error {
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

func (m *PortainereeEdgeStack) validateEnvVars(formats strfmt.Registry) error {
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

func (m *PortainereeEdgeStack) validateGitConfig(formats strfmt.Registry) error {
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

var portainereeEdgeStackTypePerDeviceConfigsMatchTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file"," dir"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portainereeEdgeStackTypePerDeviceConfigsMatchTypePropEnum = append(portainereeEdgeStackTypePerDeviceConfigsMatchTypePropEnum, v)
	}
}

const (

	// PortainereeEdgeStackPerDeviceConfigsMatchTypeFile captures enum value "file"
	PortainereeEdgeStackPerDeviceConfigsMatchTypeFile string = "file"

	// PortainereeEdgeStackPerDeviceConfigsMatchTypeDir captures enum value " dir"
	PortainereeEdgeStackPerDeviceConfigsMatchTypeDir string = " dir"
)

// prop value enum
func (m *PortainereeEdgeStack) validatePerDeviceConfigsMatchTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portainereeEdgeStackTypePerDeviceConfigsMatchTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortainereeEdgeStack) validatePerDeviceConfigsMatchType(formats strfmt.Registry) error {
	if swag.IsZero(m.PerDeviceConfigsMatchType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePerDeviceConfigsMatchTypeEnum("perDeviceConfigsMatchType", "body", m.PerDeviceConfigsMatchType); err != nil {
		return err
	}

	return nil
}

func (m *PortainereeEdgeStack) validateStaggerConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this portaineree edge stack based on the context it is used
func (m *PortainereeEdgeStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreviousDeploymentInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvVars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitConfig(ctx, formats); err != nil {
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

func (m *PortainereeEdgeStack) contextValidateAutoUpdate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainereeEdgeStack) contextValidatePreviousDeploymentInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PreviousDeploymentInfo != nil {
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

func (m *PortainereeEdgeStack) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Status {

		if val, ok := m.Status[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeEdgeStack) contextValidateEnvVars(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainereeEdgeStack) contextValidateGitConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainereeEdgeStack) contextValidateStaggerConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortainereeEdgeStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEdgeStack) UnmarshalBinary(b []byte) error {
	var res PortainereeEdgeStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
