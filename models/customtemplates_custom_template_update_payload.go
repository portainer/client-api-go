// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomtemplatesCustomTemplateUpdatePayload customtemplates custom template update payload
//
// swagger:model customtemplates.customTemplateUpdatePayload
type CustomtemplatesCustomTemplateUpdatePayload struct {

	// Description of the template
	// Example: High performance web server
	// Required: true
	Description *string `json:"description"`

	// Content of stack file
	// Required: true
	FileContent *string `json:"fileContent"`

	// URL of the template's logo
	// Example: https://cloudinovasi.id/assets/img/logos/nginx.png
	Logo string `json:"logo,omitempty"`

	// A note that will be displayed in the UI. Supports HTML content
	// Example: This is my \u003cb\u003ecustom\u003c/b\u003e template
	Note string `json:"note,omitempty"`

	// Platform associated to the template.
	// Valid values are: 1 - 'linux', 2 - 'windows'
	// Required for Docker stacks
	// Example: 1
	// Enum: [1 2]
	Platform int64 `json:"platform,omitempty"`

	// Title of the template
	// Example: Nginx
	// Required: true
	Title *string `json:"title"`

	// Type of created stack (1 - swarm, 2 - compose, 3 - kubernetes)
	// Example: 1
	// Required: true
	// Enum: [1 2 3]
	Type *int64 `json:"type"`
}

// Validate validates this customtemplates custom template update payload
func (m *CustomtemplatesCustomTemplateUpdatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateUpdatePayload) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateUpdatePayload) validateFileContent(formats strfmt.Registry) error {

	if err := validate.Required("fileContent", "body", m.FileContent); err != nil {
		return err
	}

	return nil
}

var customtemplatesCustomTemplateUpdatePayloadTypePlatformPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customtemplatesCustomTemplateUpdatePayloadTypePlatformPropEnum = append(customtemplatesCustomTemplateUpdatePayloadTypePlatformPropEnum, v)
	}
}

// prop value enum
func (m *CustomtemplatesCustomTemplateUpdatePayload) validatePlatformEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, customtemplatesCustomTemplateUpdatePayloadTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateUpdatePayload) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateUpdatePayload) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

var customtemplatesCustomTemplateUpdatePayloadTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customtemplatesCustomTemplateUpdatePayloadTypeTypePropEnum = append(customtemplatesCustomTemplateUpdatePayloadTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *CustomtemplatesCustomTemplateUpdatePayload) validateTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, customtemplatesCustomTemplateUpdatePayloadTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateUpdatePayload) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this customtemplates custom template update payload based on context it is used
func (m *CustomtemplatesCustomTemplateUpdatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomtemplatesCustomTemplateUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomtemplatesCustomTemplateUpdatePayload) UnmarshalBinary(b []byte) error {
	var res CustomtemplatesCustomTemplateUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
