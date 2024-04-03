// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingsPublicSettingsResponse settings public settings response
//
// swagger:model settings.publicSettingsResponse
type SettingsPublicSettingsResponse struct {

	// Active authentication method for the Portainer instance. Valid values are: 1 for internal, 2 for LDAP, or 3 for oauth
	// Example: 1
	AuthenticationMethod int64 `json:"AuthenticationMethod,omitempty"`

	// The content in plaintext used to display in the login page. Will hide when value is empty string
	// Example: notice or agreement
	CustomLoginBanner string `json:"CustomLoginBanner,omitempty"`

	// Whether edge compute features are enabled
	// Example: true
	EnableEdgeComputeFeatures bool `json:"EnableEdgeComputeFeatures,omitempty"`

	// Whether telemetry is enabled
	// Example: true
	EnableTelemetry bool `json:"EnableTelemetry,omitempty"`

	// Supported feature flags
	Features map[string]bool `json:"Features,omitempty"`

	// Deployment options for encouraging deployment as code
	GlobalDeploymentOptions *PortainereeGlobalDeploymentOptions `json:"GlobalDeploymentOptions,omitempty"`

	// URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string
	// Example: https://mycompany.mydomain.tld/logo.png
	LogoURL string `json:"LogoURL,omitempty"`

	// Whether portainer internal auth view will be hidden
	// Example: true
	OAuthHideInternalAuth bool `json:"OAuthHideInternalAuth,omitempty"`

	// The URL used for oauth login
	// Example: https://gitlab.com/oauth
	OAuthLoginURI string `json:"OAuthLoginURI,omitempty"`

	// The URL used for oauth logout
	// Example: https://gitlab.com/oauth/logout
	OAuthLogoutURI string `json:"OAuthLogoutURI,omitempty"`

	// The minimum required length for a password of any user when using internal auth mode
	// Example: 1
	RequiredPasswordLength int64 `json:"RequiredPasswordLength,omitempty"`

	// Show the Kompose build option (discontinued in 2.18)
	// Example: false
	ShowKomposeBuildOption bool `json:"ShowKomposeBuildOption,omitempty"`

	// Whether team sync is enabled
	// Example: true
	TeamSync bool `json:"TeamSync,omitempty"`

	// default registry
	DefaultRegistry *SettingsPublicSettingsResponseDefaultRegistry `json:"defaultRegistry,omitempty"`

	// edge
	Edge *SettingsPublicSettingsResponseEdge `json:"edge,omitempty"`

	// Whether AMT is enabled
	IsAMTEnabled bool `json:"isAMTEnabled,omitempty"`

	// Whether FDO is enabled
	IsFDOEnabled bool `json:"isFDOEnabled,omitempty"`

	// The expiry of a Kubeconfig
	// Example: 24h
	KubeconfigExpiry *string `json:"kubeconfigExpiry,omitempty"`
}

// Validate validates this settings public settings response
func (m *SettingsPublicSettingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalDeploymentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsPublicSettingsResponse) validateGlobalDeploymentOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalDeploymentOptions) { // not required
		return nil
	}

	if m.GlobalDeploymentOptions != nil {
		if err := m.GlobalDeploymentOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GlobalDeploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GlobalDeploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsPublicSettingsResponse) validateDefaultRegistry(formats strfmt.Registry) error {
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

func (m *SettingsPublicSettingsResponse) validateEdge(formats strfmt.Registry) error {
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

// ContextValidate validate this settings public settings response based on the context it is used
func (m *SettingsPublicSettingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobalDeploymentOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsPublicSettingsResponse) contextValidateGlobalDeploymentOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalDeploymentOptions != nil {

		if swag.IsZero(m.GlobalDeploymentOptions) { // not required
			return nil
		}

		if err := m.GlobalDeploymentOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GlobalDeploymentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GlobalDeploymentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsPublicSettingsResponse) contextValidateDefaultRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultRegistry != nil {

		if swag.IsZero(m.DefaultRegistry) { // not required
			return nil
		}

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

func (m *SettingsPublicSettingsResponse) contextValidateEdge(ctx context.Context, formats strfmt.Registry) error {

	if m.Edge != nil {

		if swag.IsZero(m.Edge) { // not required
			return nil
		}

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

// MarshalBinary interface implementation
func (m *SettingsPublicSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsPublicSettingsResponse) UnmarshalBinary(b []byte) error {
	var res SettingsPublicSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsPublicSettingsResponseDefaultRegistry settings public settings response default registry
//
// swagger:model SettingsPublicSettingsResponseDefaultRegistry
type SettingsPublicSettingsResponseDefaultRegistry struct {

	// hide
	// Example: false
	Hide bool `json:"Hide,omitempty"`
}

// Validate validates this settings public settings response default registry
func (m *SettingsPublicSettingsResponseDefaultRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this settings public settings response default registry based on context it is used
func (m *SettingsPublicSettingsResponseDefaultRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsPublicSettingsResponseDefaultRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsPublicSettingsResponseDefaultRegistry) UnmarshalBinary(b []byte) error {
	var res SettingsPublicSettingsResponseDefaultRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsPublicSettingsResponseEdge settings public settings response edge
//
// swagger:model SettingsPublicSettingsResponseEdge
type SettingsPublicSettingsResponseEdge struct {

	// The command list interval for edge agent - used in edge async mode [seconds]
	// Example: 60
	CommandInterval int64 `json:"CommandInterval,omitempty"`

	// The ping interval for edge agent - used in edge async mode [seconds]
	// Example: 60
	PingInterval int64 `json:"PingInterval,omitempty"`

	// The snapshot interval for edge agent - used in edge async mode [seconds]
	// Example: 60
	SnapshotInterval int64 `json:"SnapshotInterval,omitempty"`

	// The check in interval for edge agent (in seconds) - used in non async mode [seconds]
	// Example: 60
	CheckinInterval int64 `json:"checkinInterval,omitempty"`

	// mtls
	Mtls *PortainereeMTLSSettings `json:"mtls,omitempty"`
}

// Validate validates this settings public settings response edge
func (m *SettingsPublicSettingsResponseEdge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMtls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsPublicSettingsResponseEdge) validateMtls(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtls) { // not required
		return nil
	}

	if m.Mtls != nil {
		if err := m.Mtls.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge" + "." + "mtls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge" + "." + "mtls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this settings public settings response edge based on the context it is used
func (m *SettingsPublicSettingsResponseEdge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMtls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsPublicSettingsResponseEdge) contextValidateMtls(ctx context.Context, formats strfmt.Registry) error {

	if m.Mtls != nil {

		if swag.IsZero(m.Mtls) { // not required
			return nil
		}

		if err := m.Mtls.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge" + "." + "mtls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge" + "." + "mtls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsPublicSettingsResponseEdge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsPublicSettingsResponseEdge) UnmarshalBinary(b []byte) error {
	var res SettingsPublicSettingsResponseEdge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
