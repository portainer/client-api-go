// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerOAuthSettings portainer o auth settings
//
// swagger:model portainer.OAuthSettings
type PortainerOAuthSettings struct {

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

	// kube secret key
	KubeSecretKey []int64 `json:"KubeSecretKey"`

	// logout URI
	LogoutURI string `json:"LogoutURI,omitempty"`

	// o auth auto create users
	OAuthAutoCreateUsers bool `json:"OAuthAutoCreateUsers,omitempty"`

	// redirect URI
	RedirectURI string `json:"RedirectURI,omitempty"`

	// resource URI
	ResourceURI string `json:"ResourceURI,omitempty"`

	// s s o
	SSO bool `json:"SSO,omitempty"`

	// scopes
	Scopes string `json:"Scopes,omitempty"`

	// user identifier
	UserIdentifier string `json:"UserIdentifier,omitempty"`
}

// Validate validates this portainer o auth settings
func (m *PortainerOAuthSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer o auth settings based on context it is used
func (m *PortainerOAuthSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerOAuthSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerOAuthSettings) UnmarshalBinary(b []byte) error {
	var res PortainerOAuthSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
