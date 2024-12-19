// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// UsersNamespaceMapping users namespace mapping
//
// swagger:model users.namespaceMapping
type UsersNamespaceMapping map[string]map[string]PortainerAuthorizations

// Validate validates this users namespace mapping
func (m UsersNamespaceMapping) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			if val, ok := m[k][kk]; ok {
				if err := val.Validate(formats); err != nil {
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this users namespace mapping based on the context it is used
func (m UsersNamespaceMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			if val, ok := m[k][kk]; ok {
				if err := val.ContextValidate(ctx, formats); err != nil {
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
