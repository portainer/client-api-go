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

// UseractivityLogsListResponse useractivity logs list response
//
// swagger:model useractivity.logsListResponse
type UseractivityLogsListResponse struct {

	// logs
	Logs []*PortainereeUserActivityLog `json:"logs"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this useractivity logs list response
func (m *UseractivityLogsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UseractivityLogsListResponse) validateLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this useractivity logs list response based on the context it is used
func (m *UseractivityLogsListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UseractivityLogsListResponse) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logs); i++ {

		if m.Logs[i] != nil {

			if swag.IsZero(m.Logs[i]) { // not required
				return nil
			}

			if err := m.Logs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UseractivityLogsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UseractivityLogsListResponse) UnmarshalBinary(b []byte) error {
	var res UseractivityLogsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
