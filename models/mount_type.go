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

// MountType mount type
//
// swagger:model mount.Type
type MountType string

func NewMountType(value MountType) *MountType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MountType.
func (m MountType) Pointer() *MountType {
	return &m
}

const (

	// MountTypeBind captures enum value "bind"
	MountTypeBind MountType = "bind"

	// MountTypeVolume captures enum value "volume"
	MountTypeVolume MountType = "volume"

	// MountTypeTmpfs captures enum value "tmpfs"
	MountTypeTmpfs MountType = "tmpfs"

	// MountTypeNpipe captures enum value "npipe"
	MountTypeNpipe MountType = "npipe"
)

// for schema
var mountTypeEnum []interface{}

func init() {
	var res []MountType
	if err := json.Unmarshal([]byte(`["bind","volume","tmpfs","npipe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mountTypeEnum = append(mountTypeEnum, v)
	}
}

func (m MountType) validateMountTypeEnum(path, location string, value MountType) error {
	if err := validate.EnumCase(path, location, value, mountTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this mount type
func (m MountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this mount type based on context it is used
func (m MountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
