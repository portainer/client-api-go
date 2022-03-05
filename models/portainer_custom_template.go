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

// PortainerCustomTemplate portainer custom template
//
// swagger:model portainer.CustomTemplate
type PortainerCustomTemplate struct {

	// User identifier who created this template
	// Example: 3
	CreatedByUserID int64 `json:"CreatedByUserId,omitempty"`

	// Description of the template
	// Example: High performance web server
	Description string `json:"Description,omitempty"`

	// Path to the Stack file
	// Example: docker-compose.yml
	EntryPoint string `json:"EntryPoint,omitempty"`

	// CustomTemplate Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// URL of the template's logo
	// Example: https://cloudinovasi.id/assets/img/logos/nginx.png
	Logo string `json:"Logo,omitempty"`

	// A note that will be displayed in the UI. Supports HTML content
	// Example: This is my \u003cb\u003ecustom\u003c/b\u003e template
	Note string `json:"Note,omitempty"`

	// Platform associated to the template.
	// Valid values are: 1 - 'linux', 2 - 'windows'
	// Example: 1
	// Enum: [1 2]
	Platform int64 `json:"Platform,omitempty"`

	// Path on disk to the repository hosting the Stack file
	// Example: /data/custom_template/3
	ProjectPath string `json:"ProjectPath,omitempty"`

	// resource control
	ResourceControl *PortainerResourceControl `json:"ResourceControl,omitempty"`

	// Title of the template
	// Example: Nginx
	Title string `json:"Title,omitempty"`

	// Type of created stack (1 - swarm, 2 - compose)
	// Example: 1
	Type int64 `json:"Type,omitempty"`
}

// Validate validates this portainer custom template
func (m *PortainerCustomTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var portainerCustomTemplateTypePlatformPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portainerCustomTemplateTypePlatformPropEnum = append(portainerCustomTemplateTypePlatformPropEnum, v)
	}
}

// prop value enum
func (m *PortainerCustomTemplate) validatePlatformEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, portainerCustomTemplateTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortainerCustomTemplate) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("Platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *PortainerCustomTemplate) validateResourceControl(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceControl) { // not required
		return nil
	}

	if m.ResourceControl != nil {
		if err := m.ResourceControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceControl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceControl")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portainer custom template based on the context it is used
func (m *PortainerCustomTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerCustomTemplate) contextValidateResourceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceControl != nil {
		if err := m.ResourceControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceControl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceControl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerCustomTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerCustomTemplate) UnmarshalBinary(b []byte) error {
	var res PortainerCustomTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
