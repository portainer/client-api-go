// Code generated by go-swagger; DO NOT EDIT.

package models_ce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupRestorePayload backup restore payload
//
// swagger:model backup.restorePayload
type BackupRestorePayload struct {

	// file content
	FileContent []int64 `json:"fileContent"`

	// file name
	FileName string `json:"fileName,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this backup restore payload
func (m *BackupRestorePayload) Validate(formats strfmt.Registry) error {
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
