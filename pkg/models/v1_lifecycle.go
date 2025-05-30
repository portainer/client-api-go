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

// V1Lifecycle v1 lifecycle
//
// swagger:model v1.Lifecycle
type V1Lifecycle struct {

	// PostStart is called immediately after a container is created. If the handler fails,
	// the container is terminated and restarted according to its restart policy.
	// Other management of the container blocks until the hook completes.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PostStart *V1LifecycleHandler `json:"postStart,omitempty"`

	// PreStop is called immediately before a container is terminated due to an
	// API request or management event such as liveness/startup probe failure,
	// preemption, resource contention, etc. The handler is not called if the
	// container crashes or exits. The Pod's termination grace period countdown begins before the
	// PreStop hook is executed. Regardless of the outcome of the handler, the
	// container will eventually terminate within the Pod's termination grace
	// period (unless delayed by finalizers). Other management of the container blocks until the hook completes
	// or until the termination grace period is reached.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PreStop *V1LifecycleHandler `json:"preStop,omitempty"`
}

// Validate validates this v1 lifecycle
func (m *V1Lifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreStop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Lifecycle) validatePostStart(formats strfmt.Registry) error {
	if swag.IsZero(m.PostStart) { // not required
		return nil
	}

	if m.PostStart != nil {
		if err := m.PostStart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *V1Lifecycle) validatePreStop(formats strfmt.Registry) error {
	if swag.IsZero(m.PreStop) { // not required
		return nil
	}

	if m.PreStop != nil {
		if err := m.PreStop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 lifecycle based on the context it is used
func (m *V1Lifecycle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePostStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreStop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Lifecycle) contextValidatePostStart(ctx context.Context, formats strfmt.Registry) error {

	if m.PostStart != nil {

		if swag.IsZero(m.PostStart) { // not required
			return nil
		}

		if err := m.PostStart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postStart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postStart")
			}
			return err
		}
	}

	return nil
}

func (m *V1Lifecycle) contextValidatePreStop(ctx context.Context, formats strfmt.Registry) error {

	if m.PreStop != nil {

		if swag.IsZero(m.PreStop) { // not required
			return nil
		}

		if err := m.PreStop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preStop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preStop")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Lifecycle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Lifecycle) UnmarshalBinary(b []byte) error {
	var res V1Lifecycle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
