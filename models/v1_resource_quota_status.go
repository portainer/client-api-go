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

// V1ResourceQuotaStatus v1 resource quota status
//
// swagger:model v1.ResourceQuotaStatus
type V1ResourceQuotaStatus struct {

	// Hard is the set of enforced hard limits for each named resource.
	// More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	// +optional
	Hard V1ResourceList `json:"hard,omitempty"`

	// Used is the current observed total usage of the resource in the namespace.
	// +optional
	Used V1ResourceList `json:"used,omitempty"`
}

// Validate validates this v1 resource quota status
func (m *V1ResourceQuotaStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ResourceQuotaStatus) validateHard(formats strfmt.Registry) error {
	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if m.Hard != nil {
		if err := m.Hard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hard")
			}
			return err
		}
	}

	return nil
}

func (m *V1ResourceQuotaStatus) validateUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if m.Used != nil {
		if err := m.Used.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 resource quota status based on the context it is used
func (m *V1ResourceQuotaStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ResourceQuotaStatus) contextValidateHard(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if err := m.Hard.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hard")
		}
		return err
	}

	return nil
}

func (m *V1ResourceQuotaStatus) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if err := m.Used.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("used")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("used")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceQuotaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceQuotaStatus) UnmarshalBinary(b []byte) error {
	var res V1ResourceQuotaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
