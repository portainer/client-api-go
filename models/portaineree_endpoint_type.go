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

// PortainereeEndpointType portaineree endpoint type
//
// swagger:model portaineree.EndpointType
type PortainereeEndpointType int64

// for schema
var portainereeEndpointTypeEnum []interface{}

func init() {
	var res []PortainereeEndpointType
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7,8,9]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portainereeEndpointTypeEnum = append(portainereeEndpointTypeEnum, v)
	}
}

func (m PortainereeEndpointType) validatePortainereeEndpointTypeEnum(path, location string, value PortainereeEndpointType) error {
	if err := validate.EnumCase(path, location, value, portainereeEndpointTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this portaineree endpoint type
func (m PortainereeEndpointType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePortainereeEndpointTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this portaineree endpoint type based on context it is used
func (m PortainereeEndpointType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
