// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// PortainereeEndpointAuthorizations portaineree endpoint authorizations
//
// swagger:model portaineree.EndpointAuthorizations
type PortainereeEndpointAuthorizations map[string]PortainereeAuthorizations

// Validate validates this portaineree endpoint authorizations
func (m PortainereeEndpointAuthorizations) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this portaineree endpoint authorizations based on the context it is used
func (m PortainereeEndpointAuthorizations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
