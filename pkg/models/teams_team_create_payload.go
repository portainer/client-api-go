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

// TeamsTeamCreatePayload teams team create payload
//
// swagger:model teams.teamCreatePayload
type TeamsTeamCreatePayload struct {

	// Name
	// Example: developers
	// Required: true
	Name *string `json:"name"`

	// TeamLeaders
	// Example: [3,5]
	TeamLeaders []int64 `json:"teamLeaders"`
}

// Validate validates this teams team create payload
func (m *TeamsTeamCreatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamsTeamCreatePayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this teams team create payload based on context it is used
func (m *TeamsTeamCreatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamsTeamCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamsTeamCreatePayload) UnmarshalBinary(b []byte) error {
	var res TeamsTeamCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
