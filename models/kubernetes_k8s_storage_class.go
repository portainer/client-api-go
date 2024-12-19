// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesK8sStorageClass kubernetes k8s storage class
//
// swagger:model kubernetes.K8sStorageClass
type KubernetesK8sStorageClass struct {

	// allow volume expansion
	AllowVolumeExpansion bool `json:"allowVolumeExpansion,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provisioner
	Provisioner string `json:"provisioner,omitempty"`

	// reclaim policy
	ReclaimPolicy string `json:"reclaimPolicy,omitempty"`
}

// Validate validates this kubernetes k8s storage class
func (m *KubernetesK8sStorageClass) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes k8s storage class based on context it is used
func (m *KubernetesK8sStorageClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesK8sStorageClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesK8sStorageClass) UnmarshalBinary(b []byte) error {
	var res KubernetesK8sStorageClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
