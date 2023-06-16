// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerTemplateEnv portainer template env
//
// swagger:model portainer.TemplateEnv
type PortainerTemplateEnv struct {

	// Default value that will be set for the variable
	// Example: default_value
	Default string `json:"default,omitempty"`

	// Content of the tooltip that will be generated in the UI
	// Example: MySQL root account password
	Description string `json:"description,omitempty"`

	// Text for the label that will be generated in the UI
	// Example: Root password
	Label string `json:"label,omitempty"`

	// name of the environment(endpoint) variable
	// Example: MYSQL_ROOT_PASSWORD
	Name string `json:"name,omitempty"`

	// If set to true, will not generate any input for this variable in the UI
	// Example: false
	Preset *bool `json:"preset,omitempty"`

	// A list of name/value that will be used to generate a dropdown in the UI
	Select []*PortainerTemplateEnvSelect `json:"select"`
}

// Validate validates this portainer template env
func (m *PortainerTemplateEnv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelect(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerTemplateEnv) validateSelect(formats strfmt.Registry) error {
	if swag.IsZero(m.Select) { // not required
		return nil
	}

	for i := 0; i < len(m.Select); i++ {
		if swag.IsZero(m.Select[i]) { // not required
			continue
		}

		if m.Select[i] != nil {
			if err := m.Select[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("select" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("select" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this portainer template env based on the context it is used
func (m *PortainerTemplateEnv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerTemplateEnv) contextValidateSelect(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Select); i++ {

		if m.Select[i] != nil {
			if err := m.Select[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("select" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("select" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerTemplateEnv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerTemplateEnv) UnmarshalBinary(b []byte) error {
	var res PortainerTemplateEnv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
