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

// V1LifecycleHandler v1 lifecycle handler
//
// swagger:model v1.LifecycleHandler
type V1LifecycleHandler struct {

	// Exec specifies a command to execute in the container.
	// +optional
	Exec *V1ExecAction `json:"exec,omitempty"`

	// HTTPGet specifies an HTTP GET request to perform.
	// +optional
	HTTPGet *V1HTTPGetAction `json:"httpGet,omitempty"`

	// Sleep represents a duration that the container should sleep.
	// +featureGate=PodLifecycleSleepAction
	// +optional
	Sleep *V1SleepAction `json:"sleep,omitempty"`

	// Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
	// for backward compatibility. There is no validation of this field and
	// lifecycle hooks will fail at runtime when it is specified.
	// +optional
	TCPSocket *V1TCPSocketAction `json:"tcpSocket,omitempty"`
}

// Validate validates this v1 lifecycle handler
func (m *V1LifecycleHandler) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPGet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSleep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPSocket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LifecycleHandler) validateExec(formats strfmt.Registry) error {
	if swag.IsZero(m.Exec) { // not required
		return nil
	}

	if m.Exec != nil {
		if err := m.Exec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) validateHTTPGet(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPGet) { // not required
		return nil
	}

	if m.HTTPGet != nil {
		if err := m.HTTPGet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) validateSleep(formats strfmt.Registry) error {
	if swag.IsZero(m.Sleep) { // not required
		return nil
	}

	if m.Sleep != nil {
		if err := m.Sleep.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sleep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sleep")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) validateTCPSocket(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPSocket) { // not required
		return nil
	}

	if m.TCPSocket != nil {
		if err := m.TCPSocket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 lifecycle handler based on the context it is used
func (m *V1LifecycleHandler) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPGet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSleep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPSocket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LifecycleHandler) contextValidateExec(ctx context.Context, formats strfmt.Registry) error {

	if m.Exec != nil {

		if swag.IsZero(m.Exec) { // not required
			return nil
		}

		if err := m.Exec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) contextValidateHTTPGet(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPGet != nil {

		if swag.IsZero(m.HTTPGet) { // not required
			return nil
		}

		if err := m.HTTPGet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) contextValidateSleep(ctx context.Context, formats strfmt.Registry) error {

	if m.Sleep != nil {

		if swag.IsZero(m.Sleep) { // not required
			return nil
		}

		if err := m.Sleep.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sleep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sleep")
			}
			return err
		}
	}

	return nil
}

func (m *V1LifecycleHandler) contextValidateTCPSocket(ctx context.Context, formats strfmt.Registry) error {

	if m.TCPSocket != nil {

		if swag.IsZero(m.TCPSocket) { // not required
			return nil
		}

		if err := m.TCPSocket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LifecycleHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LifecycleHandler) UnmarshalBinary(b []byte) error {
	var res V1LifecycleHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
