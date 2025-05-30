// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhooksWebhookCreatePayload webhooks webhook create payload
//
// swagger:model webhooks.webhookCreatePayload
type WebhooksWebhookCreatePayload struct {

	// Environment(Endpoint) identifier. Reference the environment(endpoint) that will be used for deployment
	// Example: 1
	EndpointID int64 `json:"endpointID,omitempty"`

	// Registry Identifier
	// Example: 1
	RegistryID int64 `json:"registryID,omitempty"`

	// resource ID
	ResourceID string `json:"resourceID,omitempty"`

	// Type of webhook (1 - service, 2 - container)
	WebhookType int64 `json:"webhookType,omitempty"`
}

// Validate validates this webhooks webhook create payload
func (m *WebhooksWebhookCreatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhooks webhook create payload based on context it is used
func (m *WebhooksWebhookCreatePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhooksWebhookCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhooksWebhookCreatePayload) UnmarshalBinary(b []byte) error {
	var res WebhooksWebhookCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
