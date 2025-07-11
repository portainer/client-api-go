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

// V1Descriptor v1 descriptor
//
// swagger:model v1.Descriptor
type V1Descriptor struct {

	// Annotations contains arbitrary metadata relating to the targeted content.
	Annotations map[string]string `json:"annotations,omitempty"`

	// ArtifactType is the IANA media type of this artifact.
	ArtifactType string `json:"artifactType,omitempty"`

	// Data is an embedding of the targeted content. This is encoded as a base64
	// string when marshalled to JSON (automatically, by encoding/json). If
	// present, Data can be used directly to avoid fetching the targeted content.
	Data []int64 `json:"data"`

	// Digest is the digest of the targeted content.
	Digest string `json:"digest,omitempty"`

	// MediaType is the media type of the object this schema refers to.
	MediaType string `json:"mediaType,omitempty"`

	// Platform describes the platform which the image in the manifest runs on.
	//
	// This should only be used when referring to a manifest.
	Platform *V1Platform `json:"platform,omitempty"`

	// Size specifies the size in bytes of the blob.
	Size int64 `json:"size,omitempty"`

	// URLs specifies a list of URLs from which this object MAY be downloaded
	Urls []string `json:"urls"`
}

// Validate validates this v1 descriptor
func (m *V1Descriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Descriptor) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 descriptor based on the context it is used
func (m *V1Descriptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Descriptor) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {

		if swag.IsZero(m.Platform) { // not required
			return nil
		}

		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Descriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Descriptor) UnmarshalBinary(b []byte) error {
	var res V1Descriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
