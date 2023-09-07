// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SettingsExperimentalInspect(params *SettingsExperimentalInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsExperimentalInspectOK, error)

	SettingsExperimentalUpdate(params *SettingsExperimentalUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsExperimentalUpdateNoContent, error)

	SettingsInspect(params *SettingsInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsInspectOK, error)

	SettingsPublic(params *SettingsPublicParams, opts ...ClientOption) (*SettingsPublicOK, error)

	SettingsUpdate(params *SettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	SettingsExperimentalInspect retrieves portainer experimental settings

	Retrieve Portainer experimental settings.

**Access policy**: authenticated
*/
func (a *Client) SettingsExperimentalInspect(params *SettingsExperimentalInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsExperimentalInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsExperimentalInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingsExperimentalInspect",
		Method:             "GET",
		PathPattern:        "/settings/experimental",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingsExperimentalInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettingsExperimentalInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SettingsExperimentalInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SettingsExperimentalUpdate updates portainer experimental settings

	Update Portainer experimental settings.

**Access policy**: administrator
*/
func (a *Client) SettingsExperimentalUpdate(params *SettingsExperimentalUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsExperimentalUpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsExperimentalUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingsExperimentalUpdate",
		Method:             "PUT",
		PathPattern:        "/settings/experimental",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingsExperimentalUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettingsExperimentalUpdateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SettingsExperimentalUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SettingsInspect retrieves portainer settings

	Retrieve Portainer settings.

**Access policy**: administrator
*/
func (a *Client) SettingsInspect(params *SettingsInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingsInspect",
		Method:             "GET",
		PathPattern:        "/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingsInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettingsInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SettingsInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SettingsPublic retrieves portainer public settings

	Retrieve public settings. Returns a small set of settings that are not reserved to administrators only.

**Access policy**: public
*/
func (a *Client) SettingsPublic(params *SettingsPublicParams, opts ...ClientOption) (*SettingsPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsPublicParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingsPublic",
		Method:             "GET",
		PathPattern:        "/settings/public",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingsPublicReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettingsPublicOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SettingsPublic: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SettingsUpdate updates portainer settings

	Update Portainer settings.

**Access policy**: administrator
*/
func (a *Client) SettingsUpdate(params *SettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SettingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingsUpdate",
		Method:             "PUT",
		PathPattern:        "/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettingsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SettingsUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
