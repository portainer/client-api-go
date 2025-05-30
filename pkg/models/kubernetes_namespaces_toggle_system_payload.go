// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesNamespacesToggleSystemPayload kubernetes namespaces toggle system payload
//
// swagger:model kubernetes.namespacesToggleSystemPayload
type KubernetesNamespacesToggleSystemPayload struct {

	// Toggle the system state of this namespace to true or false
	// Example: true
	System bool `json:"system,omitempty"`
}

// Validate validates this kubernetes namespaces toggle system payload
func (m *KubernetesNamespacesToggleSystemPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes namespaces toggle system payload based on context it is used
func (m *KubernetesNamespacesToggleSystemPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesNamespacesToggleSystemPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesNamespacesToggleSystemPayload) UnmarshalBinary(b []byte) error {
	var res KubernetesNamespacesToggleSystemPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
