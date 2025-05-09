// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// KubernetesK8sIngressDeleteRequests kubernetes k8s ingress delete requests
//
// swagger:model kubernetes.K8sIngressDeleteRequests
type KubernetesK8sIngressDeleteRequests map[string][]string

// Validate validates this kubernetes k8s ingress delete requests
func (m KubernetesK8sIngressDeleteRequests) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes k8s ingress delete requests based on context it is used
func (m KubernetesK8sIngressDeleteRequests) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
