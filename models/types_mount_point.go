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

// TypesMountPoint types mount point
//
// swagger:model types.MountPoint
type TypesMountPoint struct {

	// destination
	Destination string `json:"destination,omitempty"`

	// driver
	Driver string `json:"driver,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// propagation
	Propagation MountPropagation `json:"propagation,omitempty"`

	// rw
	Rw bool `json:"rw,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// type
	Type MountType `json:"type,omitempty"`
}

// Validate validates this types mount point
func (m *TypesMountPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropagation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesMountPoint) validatePropagation(formats strfmt.Registry) error {
	if swag.IsZero(m.Propagation) { // not required
		return nil
	}

	if err := m.Propagation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("propagation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("propagation")
		}
		return err
	}

	return nil
}

func (m *TypesMountPoint) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this types mount point based on the context it is used
func (m *TypesMountPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePropagation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesMountPoint) contextValidatePropagation(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Propagation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("propagation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("propagation")
		}
		return err
	}

	return nil
}

func (m *TypesMountPoint) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesMountPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesMountPoint) UnmarshalBinary(b []byte) error {
	var res TypesMountPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
