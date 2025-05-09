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

// PortainerDockerSnapshot portainer docker snapshot
//
// swagger:model portainer.DockerSnapshot
type PortainerDockerSnapshot struct {

	// container count
	ContainerCount int64 `json:"ContainerCount,omitempty"`

	// diagnostics data
	DiagnosticsData *PortainerDiagnosticsData `json:"DiagnosticsData,omitempty"`

	// docker snapshot raw
	DockerSnapshotRaw PortainerDockerSnapshotRaw `json:"DockerSnapshotRaw,omitempty"`

	// docker version
	DockerVersion string `json:"DockerVersion,omitempty"`

	// gpu use all
	GpuUseAll bool `json:"GpuUseAll,omitempty"`

	// gpu use list
	GpuUseList []string `json:"GpuUseList"`

	// healthy container count
	HealthyContainerCount int64 `json:"HealthyContainerCount,omitempty"`

	// image count
	ImageCount int64 `json:"ImageCount,omitempty"`

	// is podman
	IsPodman bool `json:"IsPodman,omitempty"`

	// node count
	NodeCount int64 `json:"NodeCount,omitempty"`

	// running container count
	RunningContainerCount int64 `json:"RunningContainerCount,omitempty"`

	// service count
	ServiceCount int64 `json:"ServiceCount,omitempty"`

	// stack count
	StackCount int64 `json:"StackCount,omitempty"`

	// stopped container count
	StoppedContainerCount int64 `json:"StoppedContainerCount,omitempty"`

	// swarm
	Swarm bool `json:"Swarm,omitempty"`

	// time
	Time int64 `json:"Time,omitempty"`

	// total CPU
	TotalCPU int64 `json:"TotalCPU,omitempty"`

	// total memory
	TotalMemory int64 `json:"TotalMemory,omitempty"`

	// unhealthy container count
	UnhealthyContainerCount int64 `json:"UnhealthyContainerCount,omitempty"`

	// volume count
	VolumeCount int64 `json:"VolumeCount,omitempty"`
}

// Validate validates this portainer docker snapshot
func (m *PortainerDockerSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiagnosticsData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerDockerSnapshot) validateDiagnosticsData(formats strfmt.Registry) error {
	if swag.IsZero(m.DiagnosticsData) { // not required
		return nil
	}

	if m.DiagnosticsData != nil {
		if err := m.DiagnosticsData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DiagnosticsData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DiagnosticsData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portainer docker snapshot based on the context it is used
func (m *PortainerDockerSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiagnosticsData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerDockerSnapshot) contextValidateDiagnosticsData(ctx context.Context, formats strfmt.Registry) error {

	if m.DiagnosticsData != nil {

		if swag.IsZero(m.DiagnosticsData) { // not required
			return nil
		}

		if err := m.DiagnosticsData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DiagnosticsData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DiagnosticsData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerDockerSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerDockerSnapshot) UnmarshalBinary(b []byte) error {
	var res PortainerDockerSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
