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

// PortainereeEdgeAsyncCommand portaineree edge async command
//
// swagger:model portaineree.EdgeAsyncCommand
type PortainereeEdgeAsyncCommand struct {

	// endpoint ID
	EndpointID int64 `json:"endpointID,omitempty"`

	// executed
	Executed bool `json:"executed,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// op
	Op PortainereeEdgeAsyncCommandOperation `json:"op,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// scheduled time
	ScheduledTime string `json:"scheduledTime,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// type
	Type PortainereeEdgeAsyncCommandType `json:"type,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this portaineree edge async command
func (m *PortainereeEdgeAsyncCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
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

func (m *PortainereeEdgeAsyncCommand) validateOp(formats strfmt.Registry) error {
	if swag.IsZero(m.Op) { // not required
		return nil
	}

	if err := m.Op.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("op")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("op")
		}
		return err
	}

	return nil
}

func (m *PortainereeEdgeAsyncCommand) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this portaineree edge async command based on the context it is used
func (m *PortainereeEdgeAsyncCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOp(ctx, formats); err != nil {
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

func (m *PortainereeEdgeAsyncCommand) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Op.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("op")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("op")
		}
		return err
	}

	return nil
}

func (m *PortainereeEdgeAsyncCommand) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortainereeEdgeAsyncCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEdgeAsyncCommand) UnmarshalBinary(b []byte) error {
	var res PortainereeEdgeAsyncCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
