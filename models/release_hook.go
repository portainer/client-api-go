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

// ReleaseHook release hook
//
// swagger:model release.Hook
type ReleaseHook struct {

	// DeletePolicies are the policies that indicate when to delete the hook
	DeletePolicies []string `json:"delete_policies"`

	// Events are the events that this hook fires on.
	Events []string `json:"events"`

	// Kind is the Kubernetes kind.
	Kind string `json:"kind,omitempty"`

	// LastRun indicates the date/time this was last run.
	LastRun *ReleaseHookExecution `json:"last_run,omitempty"`

	// Manifest is the manifest contents.
	Manifest string `json:"manifest,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Path is the chart-relative path to the template.
	Path string `json:"path,omitempty"`

	// Weight indicates the sort order for execution among similar Hook type
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this release hook
func (m *ReleaseHook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseHook) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if m.LastRun != nil {
		if err := m.LastRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_run")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this release hook based on the context it is used
func (m *ReleaseHook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseHook) contextValidateLastRun(ctx context.Context, formats strfmt.Registry) error {

	if m.LastRun != nil {

		if swag.IsZero(m.LastRun) { // not required
			return nil
		}

		if err := m.LastRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseHook) UnmarshalBinary(b []byte) error {
	var res ReleaseHook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
