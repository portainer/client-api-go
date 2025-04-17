// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse github com portainer portainer ee api http handler kubernetes describe resource response
//
// swagger:model github.com_portainer_portainer-ee_api_http_handler_kubernetes.describeResourceResponse
type GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse struct {

	// describe
	Describe string `json:"describe,omitempty"`
}

// Validate validates this github com portainer portainer ee api http handler kubernetes describe resource response
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com portainer portainer ee api http handler kubernetes describe resource response based on context it is used
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse) UnmarshalBinary(b []byte) error {
	var res GithubComPortainerPortainerEeAPIHTTPHandlerKubernetesDescribeResourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
