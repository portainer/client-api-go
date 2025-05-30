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

// V2ObjectMetricStatus v2 object metric status
//
// swagger:model v2.ObjectMetricStatus
type V2ObjectMetricStatus struct {

	// current contains the current value for the given metric
	Current *V2MetricValueStatus `json:"current,omitempty"`

	// DescribedObject specifies the descriptions of a object,such as kind,name apiVersion
	DescribedObject *V2CrossVersionObjectReference `json:"describedObject,omitempty"`

	// metric identifies the target metric by name and selector
	Metric *V2MetricIdentifier `json:"metric,omitempty"`
}

// Validate validates this v2 object metric status
func (m *V2ObjectMetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescribedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ObjectMetricStatus) validateCurrent(formats strfmt.Registry) error {
	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricStatus) validateDescribedObject(formats strfmt.Registry) error {
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

func (m *V2ObjectMetricStatus) validateMetric(formats strfmt.Registry) error {
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

// ContextValidate validate this v2 object metric status based on the context it is used
func (m *V2ObjectMetricStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescribedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ObjectMetricStatus) contextValidateCurrent(ctx context.Context, formats strfmt.Registry) error {

	if m.Current != nil {

		if swag.IsZero(m.Current) { // not required
			return nil
		}

		if err := m.Current.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *V2ObjectMetricStatus) contextValidateDescribedObject(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V2ObjectMetricStatus) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V2ObjectMetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2ObjectMetricStatus) UnmarshalBinary(b []byte) error {
	var res V2ObjectMetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
