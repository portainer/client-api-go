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

// TypesContainer types container
//
// swagger:model types.Container
type TypesContainer struct {

	// Id
	ID string `json:"Id,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// host config
	HostConfig *TypesContainerHostConfig `json:"hostConfig,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// image ID
	ImageID string `json:"imageID,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// mounts
	Mounts []*TypesMountPoint `json:"mounts"`

	// names
	Names []string `json:"names"`

	// network settings
	NetworkSettings *TypesSummaryNetworkSettings `json:"networkSettings,omitempty"`

	// ports
	Ports []*TypesPort `json:"ports"`

	// size root fs
	SizeRootFs int64 `json:"sizeRootFs,omitempty"`

	// size rw
	SizeRw int64 `json:"sizeRw,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this types container
func (m *TypesContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesContainer) validateHostConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {
		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *TypesContainer) validateMounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Mounts); i++ {
		if swag.IsZero(m.Mounts[i]) { // not required
			continue
		}

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesContainer) validateNetworkSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *TypesContainer) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this types container based on the context it is used
func (m *TypesContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesContainer) contextValidateHostConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HostConfig != nil {
		if err := m.HostConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *TypesContainer) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mounts); i++ {

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesContainer) contextValidateNetworkSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *TypesContainer) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesContainer) UnmarshalBinary(b []byte) error {
	var res TypesContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TypesContainerHostConfig types container host config
//
// swagger:model TypesContainerHostConfig
type TypesContainerHostConfig struct {

	// network mode
	NetworkMode string `json:"networkMode,omitempty"`
}

// Validate validates this types container host config
func (m *TypesContainerHostConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types container host config based on context it is used
func (m *TypesContainerHostConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesContainerHostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesContainerHostConfig) UnmarshalBinary(b []byte) error {
	var res TypesContainerHostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
