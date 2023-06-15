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

// HelmHelmUserRepositoryResponse helm helm user repository response
//
// swagger:model helm.helmUserRepositoryResponse
type HelmHelmUserRepositoryResponse struct {

	// global repository
	GlobalRepository string `json:"GlobalRepository,omitempty"`

	// user repositories
	UserRepositories []*PortainerHelmUserRepository `json:"UserRepositories"`
}

// Validate validates this helm helm user repository response
func (m *HelmHelmUserRepositoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserRepositories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmHelmUserRepositoryResponse) validateUserRepositories(formats strfmt.Registry) error {
	if swag.IsZero(m.UserRepositories) { // not required
		return nil
	}

	for i := 0; i < len(m.UserRepositories); i++ {
		if swag.IsZero(m.UserRepositories[i]) { // not required
			continue
		}

		if m.UserRepositories[i] != nil {
			if err := m.UserRepositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UserRepositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UserRepositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this helm helm user repository response based on the context it is used
func (m *HelmHelmUserRepositoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserRepositories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmHelmUserRepositoryResponse) contextValidateUserRepositories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserRepositories); i++ {

		if m.UserRepositories[i] != nil {
			if err := m.UserRepositories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UserRepositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UserRepositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HelmHelmUserRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelmHelmUserRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res HelmHelmUserRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
