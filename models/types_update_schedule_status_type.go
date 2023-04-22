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

// TypesUpdateScheduleStatusType types update schedule status type
//
// swagger:model types.UpdateScheduleStatusType
type TypesUpdateScheduleStatusType int64

// for schema
var typesUpdateScheduleStatusTypeEnum []interface{}

func init() {
	var res []TypesUpdateScheduleStatusType
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		typesUpdateScheduleStatusTypeEnum = append(typesUpdateScheduleStatusTypeEnum, v)
	}
}

func (m TypesUpdateScheduleStatusType) validateTypesUpdateScheduleStatusTypeEnum(path, location string, value TypesUpdateScheduleStatusType) error {
	if err := validate.EnumCase(path, location, value, typesUpdateScheduleStatusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this types update schedule status type
func (m TypesUpdateScheduleStatusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTypesUpdateScheduleStatusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this types update schedule status type based on context it is used
func (m TypesUpdateScheduleStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
