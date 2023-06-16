// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerEdgeJobEndpointMeta portainer edge job endpoint meta
//
// swagger:model portainer.EdgeJobEndpointMeta
type PortainerEdgeJobEndpointMeta struct {

	// collect logs
	CollectLogs *bool `json:"collectLogs,omitempty"`

	// logs status
	LogsStatus int64 `json:"logsStatus,omitempty"`
}

// Validate validates this portainer edge job endpoint meta
func (m *PortainerEdgeJobEndpointMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer edge job endpoint meta based on context it is used
func (m *PortainerEdgeJobEndpointMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEdgeJobEndpointMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEdgeJobEndpointMeta) UnmarshalBinary(b []byte) error {
	var res PortainerEdgeJobEndpointMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
