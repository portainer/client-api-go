// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StacksStackMigratePayload stacks stack migrate payload
//
// swagger:model stacks.stackMigratePayload
type StacksStackMigratePayload struct {

	// endpoint ID
	// Example: 2
	// Required: true
	EndpointID *int64 `json:"endpointID"`

	// name
	// Example: new-stack
	Name string `json:"name,omitempty"`

	// swarm ID
	// Example: jpofkc0i9uo9wtx1zesuk649w
	SwarmID string `json:"swarmID,omitempty"`
}

// Validate validates this stacks stack migrate payload
func (m *StacksStackMigratePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksStackMigratePayload) validateEndpointID(formats strfmt.Registry) error {

	if err := validate.Required("endpointID", "body", m.EndpointID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stacks stack migrate payload based on context it is used
func (m *StacksStackMigratePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StacksStackMigratePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksStackMigratePayload) UnmarshalBinary(b []byte) error {
	var res StacksStackMigratePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
