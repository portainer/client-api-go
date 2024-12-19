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

// V2ObjectMetricSource v2 object metric source
//
// swagger:model v2.ObjectMetricSource
type V2ObjectMetricSource struct {

	// describedObject specifies the descriptions of a object,such as kind,name apiVersion
	DescribedObject *V2CrossVersionObjectReference `json:"describedObject,omitempty"`

	// metric identifies the target metric by name and selector
	Metric *V2MetricIdentifier `json:"metric,omitempty"`

	// target specifies the target value for the given metric
	Target *V2MetricTarget `json:"target,omitempty"`
}

// Validate validates this v2 object metric source
func (m *V2ObjectMetricSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescribedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ObjectMetricSource) validateDescribedObject(formats strfmt.Registry) error {
	if swag.IsZero(m.DescribedObject) { // not required
		return nil
	}

	if m.DescribedObject != nil {
		if err := m.DescribedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describedObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricSource) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricSource) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2 object metric source based on the context it is used
func (m *V2ObjectMetricSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescribedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ObjectMetricSource) contextValidateDescribedObject(ctx context.Context, formats strfmt.Registry) error {

	if m.DescribedObject != nil {

		if swag.IsZero(m.DescribedObject) { // not required
			return nil
		}

		if err := m.DescribedObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describedObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("describedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricSource) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {

		if swag.IsZero(m.Metric) { // not required
			return nil
		}

		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricSource) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2ObjectMetricSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2ObjectMetricSource) UnmarshalBinary(b []byte) error {
	var res V2ObjectMetricSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
