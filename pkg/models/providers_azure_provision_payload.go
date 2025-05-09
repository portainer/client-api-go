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

// ProvidersAzureProvisionPayload providers azure provision payload
//
// swagger:model providers.AzureProvisionPayload
type ProvidersAzureProvisionPayload struct {

	// availability zones
	AvailabilityZones []string `json:"availabilityZones"`

	// CredentialID holds an ID of the credential used to create the cluster
	// Example: 1
	// Required: true
	CredentialID *int64 `json:"credentialID"`

	// dns prefix
	DNSPrefix string `json:"dnsPrefix,omitempty"`

	// kubernetes version
	// Example: 1.23
	// Required: true
	KubernetesVersion *string `json:"kubernetesVersion"`

	// meta
	Meta *TypesEnvironmentMetadata `json:"meta,omitempty"`

	// name
	// Example: myDevCluster
	// Required: true
	Name *string `json:"name"`

	// network ID
	// Example: 8465fb26-632e-4fa3-bb9b-21c449629026
	NetworkID string `json:"networkID,omitempty"`

	// node count
	// Example: 3
	// Required: true
	NodeCount *int64 `json:"nodeCount"`

	// node size
	// Example: g3.small
	// Required: true
	NodeSize *string `json:"nodeSize"`

	// pool name
	PoolName string `json:"poolName,omitempty"`

	// region
	// Example: NYC1
	// Required: true
	Region *string `json:"region"`

	// Azure specific fields
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// resource group name
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`
}

// Validate validates this providers azure provision payload
func (m *ProvidersAzureProvisionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersAzureProvisionPayload) validateCredentialID(formats strfmt.Registry) error {

	if err := validate.Required("credentialID", "body", m.CredentialID); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateKubernetesVersion(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesVersion", "body", m.KubernetesVersion); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateNodeSize(formats strfmt.Registry) error {

	if err := validate.Required("nodeSize", "body", m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersAzureProvisionPayload) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this providers azure provision payload based on the context it is used
func (m *ProvidersAzureProvisionPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersAzureProvisionPayload) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {

		if swag.IsZero(m.Meta) { // not required
			return nil
		}

		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersAzureProvisionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersAzureProvisionPayload) UnmarshalBinary(b []byte) error {
	var res ProvidersAzureProvisionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
