// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgestacksUpdateStatusPayload edgestacks update status payload
//
// swagger:model edgestacks.updateStatusPayload
type EdgestacksUpdateStatusPayload struct {

	// endpoint ID
	EndpointID int64 `json:"endpointID,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// RollbackTo specifies the stack file version to rollback to (only support to rollback to the last version currently)
	RollbackTo int64 `json:"rollbackTo,omitempty"`

	// Deprecated
	Status int64 `json:"status,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`
}

// Validate validates this edgestacks update status payload
func (m *EdgestacksUpdateStatusPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edgestacks update status payload based on context it is used
func (m *EdgestacksUpdateStatusPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgestacksUpdateStatusPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgestacksUpdateStatusPayload) UnmarshalBinary(b []byte) error {
	var res EdgestacksUpdateStatusPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
