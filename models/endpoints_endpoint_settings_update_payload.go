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

// EndpointsEndpointSettingsUpdatePayload endpoints endpoint settings update payload
//
// swagger:model endpoints.endpointSettingsUpdatePayload
type EndpointsEndpointSettingsUpdatePayload struct {

	// Whether GitOps update time restrictions are enabled
	ChangeWindow *PortainereeEndpointChangeWindow `json:"changeWindow,omitempty"`

	// Hide manual deployment forms for an environment
	DeploymentOptions *PortainereeDeploymentOptions `json:"deploymentOptions,omitempty"`

	// enable g p u management
	// Example: false
	EnableGPUManagement bool `json:"enableGPUManagement,omitempty"`

	// enable image notification
	// Example: true
	EnableImageNotification bool `json:"enableImageNotification,omitempty"`

	// gpus
	Gpus []*PortainereePair `json:"gpus"`

	// security settings
	SecuritySettings *EndpointsEndpointSettingsUpdatePayloadSecuritySettings `json:"securitySettings,omitempty"`
}

// Validate validates this endpoints endpoint settings update payload
func (m *EndpointsEndpointSettingsUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecuritySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) validateChangeWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangeWindow) { // not required
		return nil
	}

	if m.ChangeWindow != nil {
		if err := m.ChangeWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeWindow")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) validateDeploymentOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentOptions) { // not required
		return nil
	}

	if m.DeploymentOptions != nil {
		if err := m.DeploymentOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) validateGpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Gpus) { // not required
		return nil
	}

	for i := 0; i < len(m.Gpus); i++ {
		if swag.IsZero(m.Gpus[i]) { // not required
			continue
		}

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) validateSecuritySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SecuritySettings) { // not required
		return nil
	}

	if m.SecuritySettings != nil {
		if err := m.SecuritySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securitySettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securitySettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoints endpoint settings update payload based on the context it is used
func (m *EndpointsEndpointSettingsUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChangeWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecuritySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) contextValidateChangeWindow(ctx context.Context, formats strfmt.Registry) error {

	if m.ChangeWindow != nil {

		if swag.IsZero(m.ChangeWindow) { // not required
			return nil
		}

		if err := m.ChangeWindow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeWindow")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) contextValidateDeploymentOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentOptions != nil {

		if swag.IsZero(m.DeploymentOptions) { // not required
			return nil
		}

		if err := m.DeploymentOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) contextValidateGpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gpus); i++ {

		if m.Gpus[i] != nil {

			if swag.IsZero(m.Gpus[i]) { // not required
				return nil
			}

			if err := m.Gpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EndpointsEndpointSettingsUpdatePayload) contextValidateSecuritySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SecuritySettings != nil {

		if swag.IsZero(m.SecuritySettings) { // not required
			return nil
		}

		if err := m.SecuritySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securitySettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securitySettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsEndpointSettingsUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsEndpointSettingsUpdatePayload) UnmarshalBinary(b []byte) error {
	var res EndpointsEndpointSettingsUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EndpointsEndpointSettingsUpdatePayloadSecuritySettings Security settings updates
//
// swagger:model EndpointsEndpointSettingsUpdatePayloadSecuritySettings
type EndpointsEndpointSettingsUpdatePayloadSecuritySettings struct {

	// Whether non-administrator should be able to use bind mounts when creating containers
	// Example: false
	AllowBindMountsForRegularUsers bool `json:"allowBindMountsForRegularUsers,omitempty"`

	// Whether non-administrator should be able to use container capabilities
	// Example: true
	AllowContainerCapabilitiesForRegularUsers bool `json:"allowContainerCapabilitiesForRegularUsers,omitempty"`

	// Whether non-administrator should be able to use device mapping
	// Example: true
	AllowDeviceMappingForRegularUsers bool `json:"allowDeviceMappingForRegularUsers,omitempty"`

	// Whether non-administrator should be able to use the host pid
	// Example: true
	AllowHostNamespaceForRegularUsers bool `json:"allowHostNamespaceForRegularUsers,omitempty"`

	// Whether non-administrator should be able to use privileged mode when creating containers
	// Example: false
	AllowPrivilegedModeForRegularUsers bool `json:"allowPrivilegedModeForRegularUsers,omitempty"`

	// Whether non-administrator should be able to manage stacks
	// Example: true
	AllowStackManagementForRegularUsers bool `json:"allowStackManagementForRegularUsers,omitempty"`

	// Whether non-administrator should be able to use sysctl settings
	// Example: true
	AllowSysctlSettingForRegularUsers bool `json:"allowSysctlSettingForRegularUsers,omitempty"`

	// Whether non-administrator should be able to browse volumes
	// Example: true
	AllowVolumeBrowserForRegularUsers bool `json:"allowVolumeBrowserForRegularUsers,omitempty"`

	// Whether host management features are enabled
	// Example: true
	EnableHostManagementFeatures bool `json:"enableHostManagementFeatures,omitempty"`
}

// Validate validates this endpoints endpoint settings update payload security settings
func (m *EndpointsEndpointSettingsUpdatePayloadSecuritySettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this endpoints endpoint settings update payload security settings based on context it is used
func (m *EndpointsEndpointSettingsUpdatePayloadSecuritySettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsEndpointSettingsUpdatePayloadSecuritySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsEndpointSettingsUpdatePayloadSecuritySettings) UnmarshalBinary(b []byte) error {
	var res EndpointsEndpointSettingsUpdatePayloadSecuritySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
