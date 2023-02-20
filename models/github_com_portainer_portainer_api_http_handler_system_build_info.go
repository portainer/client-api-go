// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo github com portainer portainer api http handler system build info
//
// swagger:model github.com_portainer_portainer_api_http_handler_system.BuildInfo
type GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo struct {

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// go version
	GoVersion string `json:"goVersion,omitempty"`

	// image tag
	ImageTag string `json:"imageTag,omitempty"`

	// nodejs version
	NodejsVersion string `json:"nodejsVersion,omitempty"`

	// webpack version
	WebpackVersion string `json:"webpackVersion,omitempty"`

	// yarn version
	YarnVersion string `json:"yarnVersion,omitempty"`
}

// Validate validates this github com portainer portainer api http handler system build info
func (m *GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com portainer portainer api http handler system build info based on context it is used
func (m *GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo) UnmarshalBinary(b []byte) error {
	var res GithubComPortainerPortainerAPIHTTPHandlerSystemBuildInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
