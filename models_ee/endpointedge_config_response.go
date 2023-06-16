// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointedgeConfigResponse endpointedge config response
//
// swagger:model endpointedge.configResponse
type EndpointedgeConfigResponse struct {

	// name
	Name string `json:"name,omitempty"`

	// Namespace to use for Kubernetes manifests, leave empty to use the namespaces defined in the manifest
	Namespace string `json:"namespace,omitempty"`

	// pre pull image
	PrePullImage *bool `json:"prePullImage,omitempty"`

	// re pull image
	RePullImage *bool `json:"rePullImage,omitempty"`

	// registry credentials
	RegistryCredentials []*PortainereeEdgeRegistryCredential `json:"registryCredentials"`

	// stack file content
	StackFileContent string `json:"stackFileContent,omitempty"`
}

// Validate validates this endpointedge config response
func (m *EndpointedgeConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistryCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointedgeConfigResponse) validateRegistryCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistryCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.RegistryCredentials); i++ {
		if swag.IsZero(m.RegistryCredentials[i]) { // not required
			continue
		}

		if m.RegistryCredentials[i] != nil {
			if err := m.RegistryCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registryCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registryCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this endpointedge config response based on the context it is used
func (m *EndpointedgeConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistryCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointedgeConfigResponse) contextValidateRegistryCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RegistryCredentials); i++ {

		if m.RegistryCredentials[i] != nil {
			if err := m.RegistryCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registryCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registryCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointedgeConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointedgeConfigResponse) UnmarshalBinary(b []byte) error {
	var res EndpointedgeConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
