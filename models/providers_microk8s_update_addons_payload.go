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

// ProvidersMicrok8sUpdateAddonsPayload providers microk8s update addons payload
//
// swagger:model providers.Microk8sUpdateAddonsPayload
type ProvidersMicrok8sUpdateAddonsPayload struct {

	// addons
	Addons []*PortainereeMicroK8sAddon `json:"addons"`
}

// Validate validates this providers microk8s update addons payload
func (m *ProvidersMicrok8sUpdateAddonsPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersMicrok8sUpdateAddonsPayload) validateAddons(formats strfmt.Registry) error {
	if swag.IsZero(m.Addons) { // not required
		return nil
	}

	for i := 0; i < len(m.Addons); i++ {
		if swag.IsZero(m.Addons[i]) { // not required
			continue
		}

		if m.Addons[i] != nil {
			if err := m.Addons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this providers microk8s update addons payload based on the context it is used
func (m *ProvidersMicrok8sUpdateAddonsPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersMicrok8sUpdateAddonsPayload) contextValidateAddons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addons); i++ {

		if m.Addons[i] != nil {
			if err := m.Addons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersMicrok8sUpdateAddonsPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersMicrok8sUpdateAddonsPayload) UnmarshalBinary(b []byte) error {
	var res ProvidersMicrok8sUpdateAddonsPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
