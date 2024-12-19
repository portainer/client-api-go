// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DockerContainerStats docker container stats
//
// swagger:model docker.ContainerStats
type DockerContainerStats struct {

	// healthy
	Healthy int64 `json:"healthy,omitempty"`

	// running
	Running int64 `json:"running,omitempty"`

	// stopped
	Stopped int64 `json:"stopped,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// unhealthy
	Unhealthy int64 `json:"unhealthy,omitempty"`
}

// Validate validates this docker container stats
func (m *DockerContainerStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this docker container stats based on context it is used
func (m *DockerContainerStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DockerContainerStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerContainerStats) UnmarshalBinary(b []byte) error {
	var res DockerContainerStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
