// Code generated by go-swagger; DO NOT EDIT.

package models_ee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainerHelmUserRepository portainer helm user repository
//
// swagger:model portainer.HelmUserRepository
type PortainerHelmUserRepository struct {

	// Membership Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// Helm repository URL
	// Example: https://charts.bitnami.com/bitnami
	URL string `json:"URL,omitempty"`

	// User identifier
	// Example: 1
	UserID int64 `json:"UserId,omitempty"`
}

// Validate validates this portainer helm user repository
func (m *PortainerHelmUserRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portainer helm user repository based on context it is used
func (m *PortainerHelmUserRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainerHelmUserRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerHelmUserRepository) UnmarshalBinary(b []byte) error {
	var res PortainerHelmUserRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
