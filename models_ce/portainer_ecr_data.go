// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerEcrData portainer ecr data
//
// swagger:model portainer.EcrData
type PortainerEcrData struct {

	// region
	// Example: ap-southeast-2
	Region string `json:"Region,omitempty"`
}

// Validate validates this portainer ecr data
func (m *PortainerEcrData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer ecr data based on context it is used
func (m *PortainerEcrData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEcrData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEcrData) UnmarshalBinary(b []byte) error {
	var res PortainerEcrData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
