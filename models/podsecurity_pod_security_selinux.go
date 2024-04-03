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

// PodsecurityPodSecuritySelinux podsecurity pod security selinux
//
// swagger:model podsecurity.PodSecuritySelinux
type PodsecurityPodSecuritySelinux struct {

	// allowed capabilities
	AllowedCapabilities []*PodsecurityPodSecurityAllowedCapabilities `json:"allowedCapabilities"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this podsecurity pod security selinux
func (m *PodsecurityPodSecuritySelinux) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodsecurityPodSecuritySelinux) validateAllowedCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedCapabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedCapabilities); i++ {
		if swag.IsZero(m.AllowedCapabilities[i]) { // not required
			continue
		}

		if m.AllowedCapabilities[i] != nil {
			if err := m.AllowedCapabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedCapabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedCapabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this podsecurity pod security selinux based on the context it is used
func (m *PodsecurityPodSecuritySelinux) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodsecurityPodSecuritySelinux) contextValidateAllowedCapabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedCapabilities); i++ {

		if m.AllowedCapabilities[i] != nil {

			if swag.IsZero(m.AllowedCapabilities[i]) { // not required
				return nil
			}

			if err := m.AllowedCapabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedCapabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedCapabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodsecurityPodSecuritySelinux) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodsecurityPodSecuritySelinux) UnmarshalBinary(b []byte) error {
	var res PodsecurityPodSecuritySelinux
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
