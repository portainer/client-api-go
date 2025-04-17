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

// SslCertificate ssl certificate
//
// swagger:model ssl.Certificate
type SslCertificate struct {

	// authority key Id
	AuthorityKeyID string `json:"AuthorityKeyId,omitempty"`

	// c r l distribution points
	CRLDistributionPoints []string `json:"CRLDistributionPoints"`

	// extended key usages
	ExtendedKeyUsages []string `json:"ExtendedKeyUsages"`

	// is certificate authority
	IsCertificateAuthority bool `json:"IsCertificateAuthority,omitempty"`

	// issuer
	Issuer *SslName `json:"Issuer,omitempty"`

	// issuing certificate u r ls
	IssuingCertificateURLs []string `json:"IssuingCertificateURLs"`

	// key usages
	KeyUsages []string `json:"KeyUsages"`

	// o c s p servers
	OCSPServers []string `json:"OCSPServers"`

	// policies
	Policies []string `json:"Policies"`

	// public key
	PublicKey *SslPublicKey `json:"PublicKey,omitempty"`

	// s h a1 fingerprint
	SHA1Fingerprint string `json:"SHA1Fingerprint,omitempty"`

	// s h a256 fingerprint
	SHA256Fingerprint string `json:"SHA256Fingerprint,omitempty"`

	// serial number
	SerialNumber string `json:"SerialNumber,omitempty"`

	// signature algorithm
	SignatureAlgorithm string `json:"SignatureAlgorithm,omitempty"`

	// subject
	Subject *SslName `json:"Subject,omitempty"`

	// subject alt DNS names
	SubjectAltDNSNames []string `json:"SubjectAltDNSNames"`

	// subject alt email addresses
	SubjectAltEmailAddresses []string `json:"SubjectAltEmailAddresses"`

	// subject alt Ip addresses
	SubjectAltIPAddresses []string `json:"SubjectAltIpAddresses"`

	// subject alt u r is
	SubjectAltURIs []string `json:"SubjectAltURIs"`

	// subject key Id
	SubjectKeyID string `json:"SubjectKeyId,omitempty"`

	// valid not after
	ValidNotAfter string `json:"ValidNotAfter,omitempty"`

	// valid not before
	ValidNotBefore string `json:"ValidNotBefore,omitempty"`

	// version
	Version int64 `json:"Version,omitempty"`
}

// Validate validates this ssl certificate
func (m *SslCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslCertificate) validateIssuer(formats strfmt.Registry) error {
	if swag.IsZero(m.Issuer) { // not required
		return nil
	}

	if m.Issuer != nil {
		if err := m.Issuer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issuer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issuer")
			}
			return err
		}
	}

	return nil
}

func (m *SslCertificate) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PublicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PublicKey")
			}
			return err
		}
	}

	return nil
}

func (m *SslCertificate) validateSubject(formats strfmt.Registry) error {
	if swag.IsZero(m.Subject) { // not required
		return nil
	}

	if m.Subject != nil {
		if err := m.Subject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Subject")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ssl certificate based on the context it is used
func (m *SslCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssuer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslCertificate) contextValidateIssuer(ctx context.Context, formats strfmt.Registry) error {

	if m.Issuer != nil {

		if swag.IsZero(m.Issuer) { // not required
			return nil
		}

		if err := m.Issuer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issuer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issuer")
			}
			return err
		}
	}

	return nil
}

func (m *SslCertificate) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {

		if swag.IsZero(m.PublicKey) { // not required
			return nil
		}

		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PublicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PublicKey")
			}
			return err
		}
	}

	return nil
}

func (m *SslCertificate) contextValidateSubject(ctx context.Context, formats strfmt.Registry) error {

	if m.Subject != nil {

		if swag.IsZero(m.Subject) { // not required
			return nil
		}

		if err := m.Subject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Subject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Subject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SslCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslCertificate) UnmarshalBinary(b []byte) error {
	var res SslCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
