// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeEndpointStatusMessage portaineree endpoint status message
//
// swagger:model portaineree.EndpointStatusMessage
type PortainereeEndpointStatusMessage struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// TODO: in future versions, we should think about removing these fields and
	// create a separate bucket to store cluster operation messages instead or try to find a better way.
	// Operation/OperationStatus blank means, nothing is happening
	Operation string `json:"operation,omitempty"`

	// ,processing,error
	OperationStatus string `json:"operationStatus,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this portaineree endpoint status message
func (m *PortainereeEndpointStatusMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree endpoint status message based on context it is used
func (m *PortainereeEndpointStatusMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeEndpointStatusMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEndpointStatusMessage) UnmarshalBinary(b []byte) error {
	var res PortainereeEndpointStatusMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
