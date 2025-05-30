// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Duration v1 duration
//
// swagger:model v1.Duration
type V1Duration struct {

	// time duration
	TimeDuration int64 `json:"time.Duration,omitempty"`
}

// Validate validates this v1 duration
func (m *V1Duration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 duration based on context it is used
func (m *V1Duration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Duration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Duration) UnmarshalBinary(b []byte) error {
	var res V1Duration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
