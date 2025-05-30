// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImagesStatusResponse images status response
//
// swagger:model images.StatusResponse
type ImagesStatusResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this images status response
func (m *ImagesStatusResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images status response based on context it is used
func (m *ImagesStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImagesStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImagesStatusResponse) UnmarshalBinary(b []byte) error {
	var res ImagesStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
