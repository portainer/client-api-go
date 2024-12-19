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

// V1ScopeSelector v1 scope selector
//
// swagger:model v1.ScopeSelector
type V1ScopeSelector struct {

	// A list of scope selector requirements by scope of the resources.
	// +optional
	MatchExpressions []*V1ScopedResourceSelectorRequirement `json:"matchExpressions"`
}

// Validate validates this v1 scope selector
func (m *V1ScopeSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchExpressions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ScopeSelector) validateMatchExpressions(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchExpressions); i++ {
		if swag.IsZero(m.MatchExpressions[i]) { // not required
			continue
		}

		if m.MatchExpressions[i] != nil {
			if err := m.MatchExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 scope selector based on the context it is used
func (m *V1ScopeSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchExpressions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ScopeSelector) contextValidateMatchExpressions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchExpressions); i++ {

		if m.MatchExpressions[i] != nil {

			if swag.IsZero(m.MatchExpressions[i]) { // not required
				return nil
			}

			if err := m.MatchExpressions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ScopeSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ScopeSelector) UnmarshalBinary(b []byte) error {
	var res V1ScopeSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
