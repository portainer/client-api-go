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

// PortainerKubernetesData portainer kubernetes data
//
// swagger:model portainer.KubernetesData
type PortainerKubernetesData struct {

	// configuration
	Configuration *PortainerKubernetesConfiguration `json:"Configuration,omitempty"`

	// flags
	Flags *PortainerKubernetesFlags `json:"Flags,omitempty"`

	// snapshots
	Snapshots []*PortainerKubernetesSnapshot `json:"Snapshots"`
}

// Validate validates this portainer kubernetes data
func (m *PortainerKubernetesData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerKubernetesData) validateConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Configuration")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerKubernetesData) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	if m.Flags != nil {
		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Flags")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerKubernetesData) validateSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this portainer kubernetes data based on the context it is used
func (m *PortainerKubernetesData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerKubernetesData) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.Configuration != nil {
		if err := m.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Configuration")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerKubernetesData) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	if m.Flags != nil {
		if err := m.Flags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Flags")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerKubernetesData) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerKubernetesData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerKubernetesData) UnmarshalBinary(b []byte) error {
	var res PortainerKubernetesData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
