// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseChart release chart
//
// swagger:model release.Chart
type ReleaseChart struct {

	// Files are miscellaneous files in a chart archive,
	// e.g. README, LICENSE, etc.
	Files []*ReleaseFile `json:"files"`

	// Lock is the contents of Chart.lock.
	Lock *ReleaseLock `json:"lock,omitempty"`

	// Metadata is the contents of the Chartfile.
	Metadata *ReleaseMetadata `json:"metadata,omitempty"`

	// Schema is an optional JSON schema for imposing structure on Values
	Schema []int64 `json:"schema"`

	// Templates for this chart.
	Templates []*ReleaseFile `json:"templates"`

	// Values are default config for this chart.
	Values map[string]interface{} `json:"values,omitempty"`
}

// Validate validates this release chart
func (m *ReleaseChart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseChart) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseChart) validateLock(formats strfmt.Registry) error {
	if swag.IsZero(m.Lock) { // not required
		return nil
	}

	if m.Lock != nil {
		if err := m.Lock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lock")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseChart) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseChart) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this release chart based on the context it is used
func (m *ReleaseChart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseChart) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {

			if swag.IsZero(m.Files[i]) { // not required
				return nil
			}

			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseChart) contextValidateLock(ctx context.Context, formats strfmt.Registry) error {

	if m.Lock != nil {

		if swag.IsZero(m.Lock) { // not required
			return nil
		}

		if err := m.Lock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lock")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseChart) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseChart) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Templates); i++ {

		if m.Templates[i] != nil {

			if swag.IsZero(m.Templates[i]) { // not required
				return nil
			}

			if err := m.Templates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseChart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseChart) UnmarshalBinary(b []byte) error {
	var res ReleaseChart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
