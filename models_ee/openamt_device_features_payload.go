// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenamtDeviceFeaturesPayload openamt device features payload
//
// swagger:model openamt.deviceFeaturesPayload
type OpenamtDeviceFeaturesPayload struct {

	// features
	Features *PortainerOpenAMTDeviceEnabledFeatures `json:"features,omitempty"`
}

// Validate validates this openamt device features payload
func (m *OpenamtDeviceFeaturesPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenamtDeviceFeaturesPayload) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openamt device features payload based on the context it is used
func (m *OpenamtDeviceFeaturesPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenamtDeviceFeaturesPayload) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.Features != nil {
		if err := m.Features.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenamtDeviceFeaturesPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenamtDeviceFeaturesPayload) UnmarshalBinary(b []byte) error {
	var res OpenamtDeviceFeaturesPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
