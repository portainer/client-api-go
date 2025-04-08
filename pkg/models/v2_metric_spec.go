// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2MetricSpec v2 metric spec
//
// swagger:model v2.MetricSpec
type V2MetricSpec struct {

	// containerResource refers to a resource metric (such as those specified in
	// requests and limits) known to Kubernetes describing a single container in
	// each pod of the current scale target (e.g. CPU or memory). Such metrics are
	// built in to Kubernetes, and have special scaling options on top of those
	// available to normal per-pod metrics using the "pods" source.
	// +optional
	ContainerResource *V2ContainerResourceMetricSource `json:"containerResource,omitempty"`

	// external refers to a global metric that is not associated
	// with any Kubernetes object. It allows autoscaling based on information
	// coming from components running outside of cluster
	// (for example length of queue in cloud messaging service, or
	// QPS from loadbalancer running outside of cluster).
	// +optional
	External *V2ExternalMetricSource `json:"external,omitempty"`

	// object refers to a metric describing a single kubernetes object
	// (for example, hits-per-second on an Ingress object).
	// +optional
	Object *V2ObjectMetricSource `json:"object,omitempty"`

	// pods refers to a metric describing each pod in the current scale target
	// (for example, transactions-processed-per-second).  The values will be
	// averaged together before being compared to the target value.
	// +optional
	Pods *V2PodsMetricSource `json:"pods,omitempty"`

	// resource refers to a resource metric (such as those specified in
	// requests and limits) known to Kubernetes describing each pod in the
	// current scale target (e.g. CPU or memory). Such metrics are built in to
	// Kubernetes, and have special scaling options on top of those available
	// to normal per-pod metrics using the "pods" source.
	// +optional
	Resource *V2ResourceMetricSource `json:"resource,omitempty"`

	// type is the type of metric source.  It should be one of "ContainerResource", "External",
	// "Object", "Pods" or "Resource", each mapping to a matching field in the object.
	Type string `json:"type,omitempty"`
}

// Validate validates this v2 metric spec
func (m *V2MetricSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2MetricSpec) validateContainerResource(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerResource) { // not required
		return nil
	}

	if m.ContainerResource != nil {
		if err := m.ContainerResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerResource")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) validateExternal(formats strfmt.Registry) error {
	if swag.IsZero(m.External) { // not required
		return nil
	}

	if m.External != nil {
		if err := m.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) validateObject(formats strfmt.Registry) error {
	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) validatePods(formats strfmt.Registry) error {
	if swag.IsZero(m.Pods) { // not required
		return nil
	}

	if m.Pods != nil {
		if err := m.Pods.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2 metric spec based on the context it is used
func (m *V2MetricSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2MetricSpec) contextValidateContainerResource(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerResource != nil {

		if swag.IsZero(m.ContainerResource) { // not required
			return nil
		}

		if err := m.ContainerResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerResource")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) contextValidateExternal(ctx context.Context, formats strfmt.Registry) error {

	if m.External != nil {

		if swag.IsZero(m.External) { // not required
			return nil
		}

		if err := m.External.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) contextValidateObject(ctx context.Context, formats strfmt.Registry) error {

	if m.Object != nil {

		if swag.IsZero(m.Object) { // not required
			return nil
		}

		if err := m.Object.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) contextValidatePods(ctx context.Context, formats strfmt.Registry) error {

	if m.Pods != nil {

		if swag.IsZero(m.Pods) { // not required
			return nil
		}

		if err := m.Pods.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *V2MetricSpec) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if swag.IsZero(m.Resource) { // not required
			return nil
		}

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2MetricSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2MetricSpec) UnmarshalBinary(b []byte) error {
	var res V2MetricSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
