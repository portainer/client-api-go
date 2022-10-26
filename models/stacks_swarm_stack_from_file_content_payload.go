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
	"github.com/go-openapi/validate"
)

// StacksSwarmStackFromFileContentPayload stacks swarm stack from file content payload
//
// swagger:model stacks.swarmStackFromFileContentPayload
type StacksSwarmStackFromFileContentPayload struct {

	// A list of environment(endpoint) variables used during stack deployment
	Env []*PortainereePair `json:"env"`

	// Whether the stack is from a app template
	// Example: false
	FromAppTemplate bool `json:"fromAppTemplate,omitempty"`

	// Name of the stack
	// Example: myStack
	// Required: true
	Name *string `json:"name"`

	// Content of the Stack file
	// Example: version: 3\n services:\n web:\n image:nginx
	// Required: true
	StackFileContent *string `json:"stackFileContent"`

	// Swarm cluster identifier
	// Example: jpofkc0i9uo9wtx1zesuk649w
	// Required: true
	SwarmID *string `json:"swarmID"`

	// A UUID to identify a webhook. The stack will be force updated and pull the latest image when the webhook was invoked.
	// Example: c11fdf23-183e-428a-9bb6-16db01032174
	Webhook string `json:"webhook,omitempty"`
}

// Validate validates this stacks swarm stack from file content payload
func (m *StacksSwarmStackFromFileContentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackFileContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwarmID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksSwarmStackFromFileContentPayload) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksSwarmStackFromFileContentPayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StacksSwarmStackFromFileContentPayload) validateStackFileContent(formats strfmt.Registry) error {

	if err := validate.Required("stackFileContent", "body", m.StackFileContent); err != nil {
		return err
	}

	return nil
}

func (m *StacksSwarmStackFromFileContentPayload) validateSwarmID(formats strfmt.Registry) error {

	if err := validate.Required("swarmID", "body", m.SwarmID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stacks swarm stack from file content payload based on the context it is used
func (m *StacksSwarmStackFromFileContentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksSwarmStackFromFileContentPayload) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {
			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksSwarmStackFromFileContentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksSwarmStackFromFileContentPayload) UnmarshalBinary(b []byte) error {
	var res StacksSwarmStackFromFileContentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
