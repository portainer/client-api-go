// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesK8sResourceQuota kubernetes k8s resource quota
//
// swagger:model kubernetes.K8sResourceQuota
type KubernetesK8sResourceQuota struct {

	// cpu
	CPU string `json:"cpu,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// memory
	Memory string `json:"memory,omitempty"`
}

// Validate validates this kubernetes k8s resource quota
func (m *KubernetesK8sResourceQuota) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes k8s resource quota based on context it is used
func (m *KubernetesK8sResourceQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesK8sResourceQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesK8sResourceQuota) UnmarshalBinary(b []byte) error {
	var res KubernetesK8sResourceQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
