// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeAPIKey portaineree API key
//
// swagger:model portaineree.APIKey
type PortainereeAPIKey struct {

	// Unix timestamp (UTC) when the API key was created
	DateCreated int64 `json:"dateCreated,omitempty"`

	// description
	// Example: portainer-api-key
	Description string `json:"description,omitempty"`

	// Digest represents SHA256 hash of the raw API key
	Digest []int64 `json:"digest"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Unix timestamp (UTC) when the API key was last used
	LastUsed int64 `json:"lastUsed,omitempty"`

	// API key identifier (7 char prefix)
	Prefix string `json:"prefix,omitempty"`

	// User Identifier
	// Example: 1
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this portaineree API key
func (m *PortainereeAPIKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree API key based on context it is used
func (m *PortainereeAPIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeAPIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeAPIKey) UnmarshalBinary(b []byte) error {
	var res PortainereeAPIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
