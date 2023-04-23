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

// PortainereeKubernetesConfiguration portaineree kubernetes configuration
//
// swagger:model portaineree.KubernetesConfiguration
type PortainereeKubernetesConfiguration struct {

	// allow none ingress class
	AllowNoneIngressClass bool `json:"AllowNoneIngressClass,omitempty"`

	// enable resource over commit
	EnableResourceOverCommit bool `json:"EnableResourceOverCommit,omitempty"`

	// ingress availability per namespace
	IngressAvailabilityPerNamespace bool `json:"IngressAvailabilityPerNamespace,omitempty"`

	// ingress classes
	IngressClasses []*PortainereeKubernetesIngressClassConfig `json:"IngressClasses"`

	// resource over commit percentage
	ResourceOverCommitPercentage int64 `json:"ResourceOverCommitPercentage,omitempty"`

	// restrict default namespace
	RestrictDefaultNamespace bool `json:"RestrictDefaultNamespace,omitempty"`

	// restrict standard user ingress w
	RestrictStandardUserIngressW bool `json:"RestrictStandardUserIngressW,omitempty"`

	// storage classes
	StorageClasses []*PortainereeKubernetesStorageClassConfig `json:"StorageClasses"`

	// use load balancer
	UseLoadBalancer bool `json:"UseLoadBalancer,omitempty"`

	// use server metrics
	UseServerMetrics bool `json:"UseServerMetrics,omitempty"`
}

// Validate validates this portaineree kubernetes configuration
func (m *PortainereeKubernetesConfiguration) Validate(formats strfmt.Registry) error {
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

func (m *PortainereeKubernetesConfiguration) validateIngressClasses(formats strfmt.Registry) error {
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

func (m *PortainereeKubernetesConfiguration) validateStorageClasses(formats strfmt.Registry) error {
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

// ContextValidate validate this portaineree kubernetes configuration based on the context it is used
func (m *PortainereeKubernetesConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PortainereeKubernetesConfiguration) contextValidateIngressClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngressClasses); i++ {

		if m.IngressClasses[i] != nil {
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

func (m *PortainereeKubernetesConfiguration) contextValidateStorageClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageClasses); i++ {

		if m.StorageClasses[i] != nil {
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
func (m *PortainereeKubernetesConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeKubernetesConfiguration) UnmarshalBinary(b []byte) error {
	var res PortainereeKubernetesConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
