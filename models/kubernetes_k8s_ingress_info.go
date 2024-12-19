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

// KubernetesK8sIngressInfo kubernetes k8s ingress info
//
// swagger:model kubernetes.K8sIngressInfo
type KubernetesK8sIngressInfo struct {

	// annotations
	Annotations map[string]string `json:"Annotations,omitempty"`

	// class name
	ClassName string `json:"ClassName,omitempty"`

	// creation date
	CreationDate string `json:"CreationDate,omitempty"`

	// hosts
	Hosts []string `json:"Hosts"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// namespace
	Namespace string `json:"Namespace,omitempty"`

	// paths
	Paths []*KubernetesK8sIngressPath `json:"Paths"`

	// TLS
	TLS []*KubernetesK8sIngressTLS `json:"TLS"`

	// type
	Type string `json:"Type,omitempty"`

	// UID
	UID string `json:"UID,omitempty"`
}

// Validate validates this kubernetes k8s ingress info
func (m *KubernetesK8sIngressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sIngressInfo) validatePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.Paths) { // not required
		return nil
	}

	for i := 0; i < len(m.Paths); i++ {
		if swag.IsZero(m.Paths[i]) { // not required
			continue
		}

		if m.Paths[i] != nil {
			if err := m.Paths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesK8sIngressInfo) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	for i := 0; i < len(m.TLS); i++ {
		if swag.IsZero(m.TLS[i]) { // not required
			continue
		}

		if m.TLS[i] != nil {
			if err := m.TLS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TLS" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TLS" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this kubernetes k8s ingress info based on the context it is used
func (m *KubernetesK8sIngressInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sIngressInfo) contextValidatePaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Paths); i++ {

		if m.Paths[i] != nil {

			if swag.IsZero(m.Paths[i]) { // not required
				return nil
			}

			if err := m.Paths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesK8sIngressInfo) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TLS); i++ {

		if m.TLS[i] != nil {

			if swag.IsZero(m.TLS[i]) { // not required
				return nil
			}

			if err := m.TLS[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TLS" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TLS" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesK8sIngressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesK8sIngressInfo) UnmarshalBinary(b []byte) error {
	var res KubernetesK8sIngressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
