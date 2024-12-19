// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2CrossVersionObjectReference v2 cross version object reference
//
// swagger:model v2.CrossVersionObjectReference
type V2CrossVersionObjectReference struct {

	// apiVersion is the API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
}

// Validate validates this v2 cross version object reference
func (m *V2CrossVersionObjectReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2 cross version object reference based on context it is used
func (m *V2CrossVersionObjectReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2CrossVersionObjectReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2CrossVersionObjectReference) UnmarshalBinary(b []byte) error {
	var res V2CrossVersionObjectReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
