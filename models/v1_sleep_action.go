// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SleepAction v1 sleep action
//
// swagger:model v1.SleepAction
type V1SleepAction struct {

	// Seconds is the number of seconds to sleep.
	Seconds int64 `json:"seconds,omitempty"`
}

// Validate validates this v1 sleep action
func (m *V1SleepAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 sleep action based on context it is used
func (m *V1SleepAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SleepAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SleepAction) UnmarshalBinary(b []byte) error {
	var res V1SleepAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
