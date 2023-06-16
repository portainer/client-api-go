// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainersContainerGpusResponse containers container gpus response
//
// swagger:model containers.containerGpusResponse
type ContainersContainerGpusResponse struct {

	// gpus
	Gpus string `json:"gpus,omitempty"`
}

// Validate validates this containers container gpus response
func (m *ContainersContainerGpusResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this containers container gpus response based on context it is used
func (m *ContainersContainerGpusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainersContainerGpusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainersContainerGpusResponse) UnmarshalBinary(b []byte) error {
	var res ContainersContainerGpusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
