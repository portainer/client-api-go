// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseRelease release release
//
// swagger:model release.Release
type ReleaseRelease struct {

	// Info provides information about a release
	// Info *Info `json:"info,omitempty"`
	// Chart is the chart that was released.
	Chart *ReleaseChart `json:"chart,omitempty"`

	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	Config interface{} `json:"config,omitempty"`

	// Hooks are all of the hooks declared for this release.
	Hooks []*ReleaseHook `json:"hooks"`

	// Manifest is the string representation of the rendered template.
	Manifest string `json:"manifest,omitempty"`

	// Name is the name of the release
	Name string `json:"name,omitempty"`

	// Namespace is the kubernetes namespace of the release.
	Namespace string `json:"namespace,omitempty"`

	// Version is an int which represents the revision of the release.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this release release
func (m *ReleaseRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseRelease) validateChart(formats strfmt.Registry) error {
	if swag.IsZero(m.Chart) { // not required
		return nil
	}

	if m.Chart != nil {
		if err := m.Chart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseRelease) validateHooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Hooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Hooks); i++ {
		if swag.IsZero(m.Hooks[i]) { // not required
			continue
		}

		if m.Hooks[i] != nil {
			if err := m.Hooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this release release based on the context it is used
func (m *ReleaseRelease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseRelease) contextValidateChart(ctx context.Context, formats strfmt.Registry) error {

	if m.Chart != nil {
		if err := m.Chart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseRelease) contextValidateHooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hooks); i++ {

		if m.Hooks[i] != nil {
			if err := m.Hooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseRelease) UnmarshalBinary(b []byte) error {
	var res ReleaseRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
