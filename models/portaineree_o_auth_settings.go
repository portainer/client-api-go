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

// PortainereeOAuthSettings portaineree o auth settings
//
// swagger:model portaineree.OAuthSettings
type PortainereeOAuthSettings struct {

	// access token URI
	AccessTokenURI string `json:"AccessTokenURI,omitempty"`

	// authorization URI
	AuthorizationURI string `json:"AuthorizationURI,omitempty"`

	// client ID
	ClientID string `json:"ClientID,omitempty"`

	// client secret
	ClientSecret string `json:"ClientSecret,omitempty"`

	// default team ID
	DefaultTeamID int64 `json:"DefaultTeamID,omitempty"`

	// hide internal auth
	HideInternalAuth *bool `json:"HideInternalAuth,omitempty"`

	// kube secret key
	KubeSecretKey []int64 `json:"KubeSecretKey"`

	// logout URI
	LogoutURI string `json:"LogoutURI,omitempty"`

	// microsoft tenant ID
	MicrosoftTenantID string `json:"MicrosoftTenantID,omitempty"`

	// o auth auto create users
	OAuthAutoCreateUsers *bool `json:"OAuthAutoCreateUsers,omitempty"`

	// o auth auto map team memberships
	OAuthAutoMapTeamMemberships *bool `json:"OAuthAutoMapTeamMemberships,omitempty"`

	// redirect URI
	RedirectURI string `json:"RedirectURI,omitempty"`

	// resource URI
	ResourceURI string `json:"ResourceURI,omitempty"`

	// s s o
	SSO *bool `json:"SSO,omitempty"`

	// scopes
	Scopes string `json:"Scopes,omitempty"`

	// team memberships
	TeamMemberships *PortainereeTeamMemberships `json:"TeamMemberships,omitempty"`

	// user identifier
	UserIdentifier string `json:"UserIdentifier,omitempty"`
}

// Validate validates this portaineree o auth settings
func (m *PortainereeOAuthSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamMemberships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeOAuthSettings) validateTeamMemberships(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamMemberships) { // not required
		return nil
	}

	if m.TeamMemberships != nil {
		if err := m.TeamMemberships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TeamMemberships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TeamMemberships")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portaineree o auth settings based on the context it is used
func (m *PortainereeOAuthSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeamMemberships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeOAuthSettings) contextValidateTeamMemberships(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamMemberships != nil {
		if err := m.TeamMemberships.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TeamMemberships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TeamMemberships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeOAuthSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeOAuthSettings) UnmarshalBinary(b []byte) error {
	var res PortainereeOAuthSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
