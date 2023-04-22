// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PortainereeTemplateType portaineree template type
//
// swagger:model portaineree.TemplateType
type PortainereeTemplateType int64

// for schema
var portainereeTemplateTypeEnum []interface{}

func init() {
	var res []PortainereeTemplateType
	if err := json.Unmarshal([]byte(`[0,1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portainereeTemplateTypeEnum = append(portainereeTemplateTypeEnum, v)
	}
}

func (m PortainereeTemplateType) validatePortainereeTemplateTypeEnum(path, location string, value PortainereeTemplateType) error {
	if err := validate.EnumCase(path, location, value, portainereeTemplateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this portaineree template type
func (m PortainereeTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePortainereeTemplateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this portaineree template type based on context it is used
func (m PortainereeTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
