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

// EndpointedgeEdgeAsyncResponse endpointedge edge async response
//
// swagger:model endpointedge.EdgeAsyncResponse
type EndpointedgeEdgeAsyncResponse struct {

	// command interval
	CommandInterval int64 `json:"commandInterval,omitempty"`

	// commands
	Commands []*PortainereeEdgeAsyncCommand `json:"commands"`

	// endpoint ID
	EndpointID int64 `json:"endpointID,omitempty"`

	// need full snapshot
	NeedFullSnapshot bool `json:"needFullSnapshot,omitempty"`

	// ping interval
	PingInterval int64 `json:"pingInterval,omitempty"`

	// snapshot interval
	SnapshotInterval int64 `json:"snapshotInterval,omitempty"`
}

// Validate validates this endpointedge edge async response
func (m *EndpointedgeEdgeAsyncResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointedgeEdgeAsyncResponse) validateCommands(formats strfmt.Registry) error {
	if swag.IsZero(m.Commands) { // not required
		return nil
	}

	for i := 0; i < len(m.Commands); i++ {
		if swag.IsZero(m.Commands[i]) { // not required
			continue
		}

		if m.Commands[i] != nil {
			if err := m.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this endpointedge edge async response based on the context it is used
func (m *EndpointedgeEdgeAsyncResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointedgeEdgeAsyncResponse) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Commands); i++ {

		if m.Commands[i] != nil {

			if swag.IsZero(m.Commands[i]) { // not required
				return nil
			}

			if err := m.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointedgeEdgeAsyncResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointedgeEdgeAsyncResponse) UnmarshalBinary(b []byte) error {
	var res EndpointedgeEdgeAsyncResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
