// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProvidersMicrok8sTestSSHPayload providers microk8s test SSH payload
//
// swagger:model providers.Microk8sTestSSHPayload
type ProvidersMicrok8sTestSSHPayload struct {

	// CredentialID holds an ID of the credential used to create the cluster
	// Example: 1
	// Required: true
	CredentialID *int64 `json:"credentialID"`

	// node i ps
	// Required: true
	NodeIPs []string `json:"nodeIPs"`
}

// Validate validates this providers microk8s test SSH payload
func (m *ProvidersMicrok8sTestSSHPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeIPs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersMicrok8sTestSSHPayload) validateCredentialID(formats strfmt.Registry) error {

	if err := validate.Required("credentialID", "body", m.CredentialID); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersMicrok8sTestSSHPayload) validateNodeIPs(formats strfmt.Registry) error {

	if err := validate.Required("nodeIPs", "body", m.NodeIPs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this providers microk8s test SSH payload based on context it is used
func (m *ProvidersMicrok8sTestSSHPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersMicrok8sTestSSHPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersMicrok8sTestSSHPayload) UnmarshalBinary(b []byte) error {
	var res ProvidersMicrok8sTestSSHPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
