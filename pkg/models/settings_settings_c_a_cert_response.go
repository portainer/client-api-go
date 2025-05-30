// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SettingsSettingsCACertResponse settings settings c a cert response
//
// swagger:model settings.settingsCACertResponse
type SettingsSettingsCACertResponse struct {

	// MTLSCACertificate is the X.509 Certificate of the MTLS CA Certificate
	MTLSCACertificate *SslCertificate `json:"MTLSCACertificate,omitempty"`
}

// Validate validates this settings settings c a cert response
func (m *SettingsSettingsCACertResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMTLSCACertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsSettingsCACertResponse) validateMTLSCACertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.MTLSCACertificate) { // not required
		return nil
	}

	if m.MTLSCACertificate != nil {
		if err := m.MTLSCACertificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MTLSCACertificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MTLSCACertificate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this settings settings c a cert response based on the context it is used
func (m *SettingsSettingsCACertResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMTLSCACertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsSettingsCACertResponse) contextValidateMTLSCACertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.MTLSCACertificate != nil {

		if swag.IsZero(m.MTLSCACertificate) { // not required
			return nil
		}

		if err := m.MTLSCACertificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MTLSCACertificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MTLSCACertificate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsSettingsCACertResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsSettingsCACertResponse) UnmarshalBinary(b []byte) error {
	var res SettingsSettingsCACertResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
