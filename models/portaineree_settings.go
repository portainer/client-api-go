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

// PortainereeSettings portaineree settings
//
// swagger:model portaineree.Settings
type PortainereeSettings struct {

	// Container environment parameter AGENT_SECRET
	AgentSecret string `json:"AgentSecret,omitempty"`

	// allow bind mounts for regular users
	AllowBindMountsForRegularUsers bool `json:"AllowBindMountsForRegularUsers,omitempty"`

	// allow container capabilities for regular users
	AllowContainerCapabilitiesForRegularUsers bool `json:"AllowContainerCapabilitiesForRegularUsers,omitempty"`

	// allow device mapping for regular users
	AllowDeviceMappingForRegularUsers bool `json:"AllowDeviceMappingForRegularUsers,omitempty"`

	// allow host namespace for regular users
	AllowHostNamespaceForRegularUsers bool `json:"AllowHostNamespaceForRegularUsers,omitempty"`

	// allow privileged mode for regular users
	AllowPrivilegedModeForRegularUsers bool `json:"AllowPrivilegedModeForRegularUsers,omitempty"`

	// allow stack management for regular users
	AllowStackManagementForRegularUsers bool `json:"AllowStackManagementForRegularUsers,omitempty"`

	// allow volume browser for regular users
	AllowVolumeBrowserForRegularUsers bool `json:"AllowVolumeBrowserForRegularUsers,omitempty"`

	// Active authentication method for the Portainer instance. Valid values are: 1 for internal, 2 for LDAP, or 3 for oauth
	// Example: 1
	AuthenticationMethod int64 `json:"AuthenticationMethod,omitempty"`

	// A list of label name & value that will be used to hide containers when querying containers
	BlackListedLabels []*PortainereePair `json:"BlackListedLabels"`

	// CloudAPIKeys
	CloudAPIKeys *PortainereeCloudAPIKeys `json:"CloudApiKeys,omitempty"`

	// The default check in interval for edge agent (in seconds)
	// Example: 5
	EdgeAgentCheckinInterval int64 `json:"EdgeAgentCheckinInterval,omitempty"`

	// EdgePortainerURL is the URL that is exposed to edge agents
	EdgePortainerURL string `json:"EdgePortainerUrl,omitempty"`

	// Whether edge compute features are enabled
	EnableEdgeComputeFeatures bool `json:"EnableEdgeComputeFeatures,omitempty"`

	// Deprecated fields v26
	EnableHostManagementFeatures bool `json:"EnableHostManagementFeatures,omitempty"`

	// Whether telemetry is enabled
	// Example: false
	EnableTelemetry bool `json:"EnableTelemetry,omitempty"`

	// EnforceEdgeID makes Portainer store the Edge ID instead of accepting anyone
	// Example: false
	EnforceEdgeID bool `json:"EnforceEdgeID,omitempty"`

	// feature flag settings
	FeatureFlagSettings map[string]bool `json:"FeatureFlagSettings,omitempty"`

	// Helm repository URL, defaults to "https://charts.bitnami.com/bitnami"
	// Example: https://charts.bitnami.com/bitnami
	HelmRepositoryURL string `json:"HelmRepositoryURL,omitempty"`

	// internal auth settings
	InternalAuthSettings *PortainereeInternalAuthSettings `json:"InternalAuthSettings,omitempty"`

	// The expiry of a Kubeconfig
	// Example: 24h
	KubeconfigExpiry string `json:"KubeconfigExpiry,omitempty"`

	// KubectlImage, defaults to portainer/kubectl-shell
	// Example: portainer/kubectl-shell
	KubectlShellImage string `json:"KubectlShellImage,omitempty"`

	// l d a p settings
	LDAPSettings *PortainereeLDAPSettings `json:"LDAPSettings,omitempty"`

	// URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string
	// Example: https://mycompany.mydomain.tld/logo.png
	LogoURL string `json:"LogoURL,omitempty"`

	// o auth settings
	OAuthSettings *PortainereeOAuthSettings `json:"OAuthSettings,omitempty"`

	// The interval in which environment(endpoint) snapshots are created
	// Example: 5m
	SnapshotInterval string `json:"SnapshotInterval,omitempty"`

	// URL to the templates that will be displayed in the UI when navigating to App Templates
	// Example: https://raw.githubusercontent.com/portainer/templates/master/templates.json
	TemplatesURL string `json:"TemplatesURL,omitempty"`

	// TrustOnFirstConnect makes Portainer accepting edge agent connection by default
	// Example: false
	TrustOnFirstConnect bool `json:"TrustOnFirstConnect,omitempty"`

	// The duration of a user session
	// Example: 5m
	UserSessionTimeout string `json:"UserSessionTimeout,omitempty"`

	// default registry
	DefaultRegistry *PortainereeSettingsDefaultRegistry `json:"defaultRegistry,omitempty"`

	// Deprecated fields
	DisplayDonationHeader bool `json:"displayDonationHeader,omitempty"`

	// display external contributors
	DisplayExternalContributors bool `json:"displayExternalContributors,omitempty"`

	// edge
	Edge *PortainereeSettingsEdge `json:"edge,omitempty"`

	// fdo configuration
	FdoConfiguration *PortainereeFDOConfiguration `json:"fdoConfiguration,omitempty"`

	// open a m t configuration
	OpenAMTConfiguration *PortainerOpenAMTConfiguration `json:"openAMTConfiguration,omitempty"`
}

// Validate validates this portaineree settings
func (m *PortainereeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlackListedLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAPIKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalAuthSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLDAPSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOAuthSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFdoConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenAMTConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeSettings) validateBlackListedLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.BlackListedLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.BlackListedLabels); i++ {
		if swag.IsZero(m.BlackListedLabels[i]) { // not required
			continue
		}

		if m.BlackListedLabels[i] != nil {
			if err := m.BlackListedLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlackListedLabels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlackListedLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeSettings) validateCloudAPIKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudAPIKeys) { // not required
		return nil
	}

	if m.CloudAPIKeys != nil {
		if err := m.CloudAPIKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CloudApiKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CloudApiKeys")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateInternalAuthSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.InternalAuthSettings) { // not required
		return nil
	}

	if m.InternalAuthSettings != nil {
		if err := m.InternalAuthSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InternalAuthSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InternalAuthSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateLDAPSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.LDAPSettings) { // not required
		return nil
	}

	if m.LDAPSettings != nil {
		if err := m.LDAPSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LDAPSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LDAPSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateOAuthSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.OAuthSettings) { // not required
		return nil
	}

	if m.OAuthSettings != nil {
		if err := m.OAuthSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OAuthSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OAuthSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateDefaultRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultRegistry) { // not required
		return nil
	}

	if m.DefaultRegistry != nil {
		if err := m.DefaultRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRegistry")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateEdge(formats strfmt.Registry) error {
	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateFdoConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.FdoConfiguration) { // not required
		return nil
	}

	if m.FdoConfiguration != nil {
		if err := m.FdoConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fdoConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fdoConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) validateOpenAMTConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenAMTConfiguration) { // not required
		return nil
	}

	if m.OpenAMTConfiguration != nil {
		if err := m.OpenAMTConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openAMTConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openAMTConfiguration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portaineree settings based on the context it is used
func (m *PortainereeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlackListedLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudAPIKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInternalAuthSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLDAPSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOAuthSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFdoConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenAMTConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeSettings) contextValidateBlackListedLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlackListedLabels); i++ {

		if m.BlackListedLabels[i] != nil {
			if err := m.BlackListedLabels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlackListedLabels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlackListedLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeSettings) contextValidateCloudAPIKeys(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudAPIKeys != nil {
		if err := m.CloudAPIKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CloudApiKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CloudApiKeys")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateInternalAuthSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.InternalAuthSettings != nil {
		if err := m.InternalAuthSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InternalAuthSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InternalAuthSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateLDAPSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.LDAPSettings != nil {
		if err := m.LDAPSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LDAPSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LDAPSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateOAuthSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.OAuthSettings != nil {
		if err := m.OAuthSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OAuthSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OAuthSettings")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateDefaultRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultRegistry != nil {
		if err := m.DefaultRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRegistry")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateEdge(ctx context.Context, formats strfmt.Registry) error {

	if m.Edge != nil {
		if err := m.Edge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateFdoConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.FdoConfiguration != nil {
		if err := m.FdoConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fdoConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fdoConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeSettings) contextValidateOpenAMTConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenAMTConfiguration != nil {
		if err := m.OpenAMTConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openAMTConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openAMTConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeSettings) UnmarshalBinary(b []byte) error {
	var res PortainereeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortainereeSettingsDefaultRegistry the default builtin registry now is anonymous docker hub registry
//
// swagger:model PortainereeSettingsDefaultRegistry
type PortainereeSettingsDefaultRegistry struct {

	// hide
	// Example: false
	Hide bool `json:"Hide,omitempty"`
}

// Validate validates this portaineree settings default registry
func (m *PortainereeSettingsDefaultRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree settings default registry based on context it is used
func (m *PortainereeSettingsDefaultRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeSettingsDefaultRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeSettingsDefaultRegistry) UnmarshalBinary(b []byte) error {
	var res PortainereeSettingsDefaultRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortainereeSettingsEdge portaineree settings edge
//
// swagger:model PortainereeSettingsEdge
type PortainereeSettingsEdge struct {

	// The command list interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	CommandInterval int64 `json:"CommandInterval,omitempty"`

	// The ping interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	PingInterval int64 `json:"PingInterval,omitempty"`

	// The snapshot interval for edge agent - used in edge async mode (in seconds)
	// Example: 5
	SnapshotInterval int64 `json:"SnapshotInterval,omitempty"`

	// EdgeAsyncMode enables edge async mode by default
	AsyncMode bool `json:"asyncMode,omitempty"`
}

// Validate validates this portaineree settings edge
func (m *PortainereeSettingsEdge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree settings edge based on context it is used
func (m *PortainereeSettingsEdge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeSettingsEdge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeSettingsEdge) UnmarshalBinary(b []byte) error {
	var res PortainereeSettingsEdge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
