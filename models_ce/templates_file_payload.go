// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TemplatesFilePayload templates file payload
//
// swagger:model templates.filePayload
type TemplatesFilePayload struct {

	// Path to the file inside the git repository
	// Example: ./subfolder/docker-compose.yml
	// Required: true
	ComposeFilePathInRepository *string `json:"composeFilePathInRepository"`

	// URL of a git repository where the file is stored
	// Example: https://github.com/portainer/portainer-compose
	// Required: true
	RepositoryURL *string `json:"repositoryURL"`
}

// Validate validates this templates file payload
func (m *TemplatesFilePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComposeFilePathInRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplatesFilePayload) validateComposeFilePathInRepository(formats strfmt.Registry) error {

	if err := validate.Required("composeFilePathInRepository", "body", m.ComposeFilePathInRepository); err != nil {
		return err
	}

	return nil
}

func (m *TemplatesFilePayload) validateRepositoryURL(formats strfmt.Registry) error {

	if err := validate.Required("repositoryURL", "body", m.RepositoryURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this templates file payload based on context it is used
func (m *TemplatesFilePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplatesFilePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplatesFilePayload) UnmarshalBinary(b []byte) error {
	var res TemplatesFilePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
