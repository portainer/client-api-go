// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FdoDeviceConfigurePayload fdo device configure payload
//
// swagger:model fdo.deviceConfigurePayload
type FdoDeviceConfigurePayload struct {

	// edge ID
	EdgeID string `json:"edgeID,omitempty"`

	// edge key
	EdgeKey string `json:"edgeKey,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// profile
	Profile int64 `json:"profile,omitempty"`
}

// Validate validates this fdo device configure payload
func (m *FdoDeviceConfigurePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fdo device configure payload based on context it is used
func (m *FdoDeviceConfigurePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FdoDeviceConfigurePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FdoDeviceConfigurePayload) UnmarshalBinary(b []byte) error {
	var res FdoDeviceConfigurePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
