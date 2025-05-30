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

// BackupRestorePayload backup restore payload
//
// swagger:model backup.restorePayload
type BackupRestorePayload struct {

	// Content of the backup
	// Required: true
	FileContent []int64 `json:"fileContent"`

	// File name
	// Required: true
	FileName *string `json:"fileName"`

	// Password to decrypt the backup with
	Password string `json:"password,omitempty"`
}

// Validate validates this backup restore payload
func (m *BackupRestorePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestorePayload) validateFileContent(formats strfmt.Registry) error {

	if err := validate.Required("fileContent", "body", m.FileContent); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestorePayload) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("fileName", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup restore payload based on context it is used
func (m *BackupRestorePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRestorePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestorePayload) UnmarshalBinary(b []byte) error {
	var res BackupRestorePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
