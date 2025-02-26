// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ResourceClaim v1 resource claim
//
// swagger:model v1.ResourceClaim
type V1ResourceClaim struct {

	// Name must match the name of one entry in pod.spec.resourceClaims of
	// the Pod where this field is used. It makes that resource available
	// inside a container.
	Name string `json:"name,omitempty"`

	// Request is the name chosen for a request in the referenced claim.
	// If empty, everything from the claim is made available, otherwise
	// only the result of this request.
	//
	// +optional
	Request string `json:"request,omitempty"`
}

// Validate validates this v1 resource claim
func (m *V1ResourceClaim) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 resource claim based on context it is used
func (m *V1ResourceClaim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceClaim) UnmarshalBinary(b []byte) error {
	var res V1ResourceClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
