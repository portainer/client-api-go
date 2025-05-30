// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesPod kubernetes pod
//
// swagger:model kubernetes.Pod
type KubernetesPod struct {

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this kubernetes pod
func (m *KubernetesPod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes pod based on context it is used
func (m *KubernetesPod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesPod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesPod) UnmarshalBinary(b []byte) error {
	var res KubernetesPod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
