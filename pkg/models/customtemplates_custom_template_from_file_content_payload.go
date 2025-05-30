// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomtemplatesCustomTemplateFromFileContentPayload customtemplates custom template from file content payload
//
// swagger:model customtemplates.customTemplateFromFileContentPayload
type CustomtemplatesCustomTemplateFromFileContentPayload struct {

	// Description of the template
	// Example: High performance web server
	// Required: true
	Description *string `json:"description"`

	// edge settings
	EdgeSettings *PortainereeCustomTemplateEdgeSettings `json:"edgeSettings,omitempty"`

	// EdgeTemplate indicates if this template purpose for Edge Stack
	// Example: false
	EdgeTemplate bool `json:"edgeTemplate,omitempty"`

	// Content of stack file
	// Required: true
	FileContent *string `json:"fileContent"`

	// URL of the template's logo
	// Example: https://portainer.io/img/logo.svg
	Logo string `json:"logo,omitempty"`

	// A note that will be displayed in the UI. Supports HTML content
	// Example: This is my \u003cb\u003ecustom\u003c/b\u003e template
	Note string `json:"note,omitempty"`

	// Platform associated to the template.
	// Valid values are: 1 - 'linux', 2 - 'windows'
	// Required for Docker stacks
	// Example: 1
	// Enum: [1,2]
	Platform int64 `json:"platform,omitempty"`

	// Title of the template
	// Example: Nginx
	// Required: true
	Title *string `json:"title"`

	// Type of created stack:
	// * 1 - swarm
	// * 2 - compose
	// * 3 - kubernetes
	// Example: 1
	// Required: true
	// Enum: [1,2,3]
	Type *int64 `json:"type"`

	// Definitions of variables in the stack file
	Variables []*PortainerCustomTemplateVariableDefinition `json:"variables"`
}

// Validate validates this customtemplates custom template from file content payload
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeSettings(formats); err != nil {
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

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateEdgeSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeSettings) { // not required
		return nil
	}

	if m.EdgeSettings != nil {
		if err := m.EdgeSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateFileContent(formats strfmt.Registry) error {

	if err := validate.Required("fileContent", "body", m.FileContent); err != nil {
		return err
	}

	return nil
}

var customtemplatesCustomTemplateFromFileContentPayloadTypePlatformPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customtemplatesCustomTemplateFromFileContentPayloadTypePlatformPropEnum = append(customtemplatesCustomTemplateFromFileContentPayloadTypePlatformPropEnum, v)
	}
}

// prop value enum
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validatePlatformEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, customtemplatesCustomTemplateFromFileContentPayloadTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

var customtemplatesCustomTemplateFromFileContentPayloadTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customtemplatesCustomTemplateFromFileContentPayloadTypeTypePropEnum = append(customtemplatesCustomTemplateFromFileContentPayloadTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, customtemplatesCustomTemplateFromFileContentPayloadTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) validateVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this customtemplates custom template from file content payload based on the context it is used
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) contextValidateEdgeSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeSettings != nil {

		if swag.IsZero(m.EdgeSettings) { // not required
			return nil
		}

		if err := m.EdgeSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CustomtemplatesCustomTemplateFromFileContentPayload) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {

			if swag.IsZero(m.Variables[i]) { // not required
				return nil
			}

			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomtemplatesCustomTemplateFromFileContentPayload) UnmarshalBinary(b []byte) error {
	var res CustomtemplatesCustomTemplateFromFileContentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
