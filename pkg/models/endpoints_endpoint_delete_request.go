// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointsEndpointDeleteRequest endpoints endpoint delete request
//
// swagger:model endpoints.endpointDeleteRequest
type EndpointsEndpointDeleteRequest struct {

	// delete cluster
	DeleteCluster bool `json:"deleteCluster,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this endpoints endpoint delete request
func (m *EndpointsEndpointDeleteRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this endpoints endpoint delete request based on context it is used
func (m *EndpointsEndpointDeleteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsEndpointDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsEndpointDeleteRequest) UnmarshalBinary(b []byte) error {
	var res EndpointsEndpointDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
