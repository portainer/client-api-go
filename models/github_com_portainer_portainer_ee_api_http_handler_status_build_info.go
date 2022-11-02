// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo github com portainer portainer ee api http handler status build info
//
// swagger:model github.com_portainer_portainer-ee_api_http_handler_status.BuildInfo
type GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo struct {

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

// Validate validates this github com portainer portainer ee api http handler status build info
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com portainer portainer ee api http handler status build info based on context it is used
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo) UnmarshalBinary(b []byte) error {
	var res GithubComPortainerPortainerEeAPIHTTPHandlerStatusBuildInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
