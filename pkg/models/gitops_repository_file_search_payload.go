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

// GitopsRepositoryFileSearchPayload gitops repository file search payload
//
// swagger:model gitops.repositoryFileSearchPayload
type GitopsRepositoryFileSearchPayload struct {

	// created from custom template ID
	CreatedFromCustomTemplateID int64 `json:"createdFromCustomTemplateID,omitempty"`

	// DirOnly List directories only
	// Example: false
	DirOnly bool `json:"dirOnly,omitempty"`

	// git credential ID
	GitCredentialID int64 `json:"gitCredentialID,omitempty"`

	// Allow to provide specific file extension as the search result. If empty, the file extensions yml,yaml,hcl,json will be set by default
	// Example: json,yml
	Include string `json:"include,omitempty"`

	// Partial or completed file name. If empty, all filenames with included extensions will be returned
	// Example: docker-compose
	Keyword string `json:"keyword,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// Specific Git repository reference. If empty, the reference ref/heads/main will be set by default
	// Example: refs/heads/develop
	Reference string `json:"reference,omitempty"`

	// repository
	// Required: true
	Repository *string `json:"repository"`

	// TLSSkipVerify skips SSL verification when cloning the Git repository
	// Example: false
	TlsskipVerify bool `json:"tlsskipVerify,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this gitops repository file search payload
func (m *GitopsRepositoryFileSearchPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitopsRepositoryFileSearchPayload) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gitops repository file search payload based on context it is used
func (m *GitopsRepositoryFileSearchPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitopsRepositoryFileSearchPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitopsRepositoryFileSearchPayload) UnmarshalBinary(b []byte) error {
	var res GitopsRepositoryFileSearchPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
