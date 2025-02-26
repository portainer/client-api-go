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

// GitopsRepositoryReferenceListPayload gitops repository reference list payload
//
// swagger:model gitops.repositoryReferenceListPayload
type GitopsRepositoryReferenceListPayload struct {

	// created from custom template ID
	CreatedFromCustomTemplateID int64 `json:"createdFromCustomTemplateID,omitempty"`

	// git credential ID
	GitCredentialID int64 `json:"gitCredentialID,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// repository
	// Required: true
	Repository *string `json:"repository"`

	// stack ID
	StackID int64 `json:"stackID,omitempty"`

	// TLSSkipVerify skips SSL verification when cloning the Git repository
	// Example: false
	TlsskipVerify bool `json:"tlsskipVerify,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this gitops repository reference list payload
func (m *GitopsRepositoryReferenceListPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitopsRepositoryReferenceListPayload) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gitops repository reference list payload based on context it is used
func (m *GitopsRepositoryReferenceListPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitopsRepositoryReferenceListPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitopsRepositoryReferenceListPayload) UnmarshalBinary(b []byte) error {
	var res GitopsRepositoryReferenceListPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
