// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EndpointgroupsEndpointGroupCreatePayload endpointgroups endpoint group create payload
//
// swagger:model endpointgroups.endpointGroupCreatePayload
type EndpointgroupsEndpointGroupCreatePayload struct {

	// List of environment(endpoint) identifiers that will be part of this group
	// Example: [1,3]
	AssociatedEndpoints []int64 `json:"associatedEndpoints"`

	// Environment(Endpoint) group description
	// Example: description
	Description string `json:"description,omitempty"`

	// Environment(Endpoint) group name
	// Example: my-environment-group
	// Required: true
	Name *string `json:"name"`

	// List of tag identifiers to which this environment(endpoint) group is associated
	// Example: [1,2]
	TagIDs []int64 `json:"tagIDs"`
}

// Validate validates this endpointgroups endpoint group create payload
func (m *EndpointgroupsEndpointGroupCreatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointgroupsEndpointGroupCreatePayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this endpointgroups endpoint group create payload based on context it is used
func (m *EndpointgroupsEndpointGroupCreatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointgroupsEndpointGroupCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointgroupsEndpointGroupCreatePayload) UnmarshalBinary(b []byte) error {
	var res EndpointgroupsEndpointGroupCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
