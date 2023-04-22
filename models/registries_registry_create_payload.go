// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistriesRegistryCreatePayload registries registry create payload
//
// swagger:model registries.registryCreatePayload
type RegistriesRegistryCreatePayload struct {

	// Is authentication against this registry enabled
	// Example: false
	// Required: true
	Authentication *bool `json:"authentication"`

	// BaseURL required for ProGet registry
	// Example: registry.mydomain.tld:2375
	BaseURL string `json:"baseURL,omitempty"`

	// ECR specific details, required when type = 7
	Ecr struct {
		PortainereeEcrData
	} `json:"ecr,omitempty"`

	// Github specific details, required when type = 8
	Github struct {
		PortainereeGithubRegistryData
	} `json:"github,omitempty"`

	// Gitlab specific details, required when type = 4
	Gitlab struct {
		PortainereeGitlabRegistryData
	} `json:"gitlab,omitempty"`

	// Name that will be used to identify this registry
	// Example: my-registry
	// Required: true
	Name *string `json:"name"`

	// Password used to authenticate against this registry. required when Authentication is true
	// Example: registry_password
	Password string `json:"password,omitempty"`

	// Quay specific details, required when type = 1
	Quay struct {
		PortainereeQuayRegistryData
	} `json:"quay,omitempty"`

	// Registry Type. Valid values are:
	// 	1 (Quay.io),
	// 	2 (Azure container registry),
	// 	3 (custom registry),
	// 	4 (Gitlab registry),
	// 	5 (ProGet registry),
	// 	6 (DockerHub)
	// 	7 (ECR)
	// 	8 (Github registry)
	// Example: 1
	// Required: true
	// Enum: [1 2 3 4 5 6 7 8]
	Type struct {
		PortainereeRegistryType
	} `json:"type"`

	// URL or IP address of the Docker registry
	// Example: registry.mydomain.tld:2375/feed
	// Required: true
	URL *string `json:"url"`

	// Username used to authenticate against this registry. Required when Authentication is true
	// Example: registry_user
	Username string `json:"username,omitempty"`
}

// Validate validates this registries registry create payload
func (m *RegistriesRegistryCreatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEcr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitlab(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistriesRegistryCreatePayload) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateEcr(formats strfmt.Registry) error {
	if swag.IsZero(m.Ecr) { // not required
		return nil
	}

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateGithub(formats strfmt.Registry) error {
	if swag.IsZero(m.Github) { // not required
		return nil
	}

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateGitlab(formats strfmt.Registry) error {
	if swag.IsZero(m.Gitlab) { // not required
		return nil
	}

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateQuay(formats strfmt.Registry) error {
	if swag.IsZero(m.Quay) { // not required
		return nil
	}

	return nil
}

var registriesRegistryCreatePayloadTypeTypePropEnum []interface{}

func init() {
	var res []struct {
		PortainereeRegistryType
	}
	if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registriesRegistryCreatePayloadTypeTypePropEnum = append(registriesRegistryCreatePayloadTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *RegistriesRegistryCreatePayload) validateTypeEnum(path, location string, value *struct {
	PortainereeRegistryType
}) error {
	if err := validate.EnumCase(path, location, value, registriesRegistryCreatePayloadTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RegistriesRegistryCreatePayload) validateType(formats strfmt.Registry) error {

	return nil
}

func (m *RegistriesRegistryCreatePayload) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this registries registry create payload based on the context it is used
func (m *RegistriesRegistryCreatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEcr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGithub(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitlab(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistriesRegistryCreatePayload) contextValidateEcr(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RegistriesRegistryCreatePayload) contextValidateGithub(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RegistriesRegistryCreatePayload) contextValidateGitlab(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RegistriesRegistryCreatePayload) contextValidateQuay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RegistriesRegistryCreatePayload) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *RegistriesRegistryCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistriesRegistryCreatePayload) UnmarshalBinary(b []byte) error {
	var res RegistriesRegistryCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
