// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerKubernetesConfiguration portainer kubernetes configuration
//
// swagger:model portainer.KubernetesConfiguration
type PortainerKubernetesConfiguration struct {

	// allow none ingress class
	AllowNoneIngressClass bool `json:"AllowNoneIngressClass,omitempty"`

	// enable resource over commit
	EnableResourceOverCommit bool `json:"EnableResourceOverCommit,omitempty"`

	// ingress availability per namespace
	IngressAvailabilityPerNamespace bool `json:"IngressAvailabilityPerNamespace,omitempty"`

	// ingress classes
	IngressClasses []*PortainerKubernetesIngressClassConfig `json:"IngressClasses"`

	// resource over commit percentage
	ResourceOverCommitPercentage int64 `json:"ResourceOverCommitPercentage,omitempty"`

	// restrict default namespace
	RestrictDefaultNamespace bool `json:"RestrictDefaultNamespace,omitempty"`

	// storage classes
	StorageClasses []*PortainerKubernetesStorageClassConfig `json:"StorageClasses"`

	// use load balancer
	UseLoadBalancer bool `json:"UseLoadBalancer,omitempty"`

	// use server metrics
	UseServerMetrics bool `json:"UseServerMetrics,omitempty"`
}

// Validate validates this portainer kubernetes configuration
func (m *PortainerKubernetesConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngressClasses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageClasses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerKubernetesConfiguration) validateIngressClasses(formats strfmt.Registry) error {
	if swag.IsZero(m.IngressClasses) { // not required
		return nil
	}

	for i := 0; i < len(m.IngressClasses); i++ {
		if swag.IsZero(m.IngressClasses[i]) { // not required
			continue
		}

		if m.IngressClasses[i] != nil {
			if err := m.IngressClasses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("IngressClasses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("IngressClasses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerKubernetesConfiguration) validateStorageClasses(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageClasses) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageClasses); i++ {
		if swag.IsZero(m.StorageClasses[i]) { // not required
			continue
		}

		if m.StorageClasses[i] != nil {
			if err := m.StorageClasses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageClasses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StorageClasses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this portainer kubernetes configuration based on the context it is used
func (m *PortainerKubernetesConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIngressClasses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageClasses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerKubernetesConfiguration) contextValidateIngressClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngressClasses); i++ {

		if m.IngressClasses[i] != nil {

			if swag.IsZero(m.IngressClasses[i]) { // not required
				return nil
			}

			if err := m.IngressClasses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("IngressClasses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("IngressClasses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerKubernetesConfiguration) contextValidateStorageClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageClasses); i++ {

		if m.StorageClasses[i] != nil {

			if swag.IsZero(m.StorageClasses[i]) { // not required
				return nil
			}

			if err := m.StorageClasses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageClasses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StorageClasses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerKubernetesConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerKubernetesConfiguration) UnmarshalBinary(b []byte) error {
	var res PortainerKubernetesConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
