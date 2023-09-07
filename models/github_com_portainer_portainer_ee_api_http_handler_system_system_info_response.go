// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse github com portainer portainer ee api http handler system system info response
//
// swagger:model github.com_portainer_portainer-ee_api_http_handler_system.systemInfoResponse
type GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse struct {

	// agents
	Agents int64 `json:"agents,omitempty"`

	// edge agents
	EdgeAgents int64 `json:"edgeAgents,omitempty"`

	// edge devices
	EdgeDevices int64 `json:"edgeDevices,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`
}

// Validate validates this github com portainer portainer ee api http handler system system info response
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com portainer portainer ee api http handler system system info response based on context it is used
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse) UnmarshalBinary(b []byte) error {
	var res GithubComPortainerPortainerEeAPIHTTPHandlerSystemSystemInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
