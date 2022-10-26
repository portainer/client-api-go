// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new registries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for registries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RegistryConfigure(params *RegistryConfigureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryConfigureNoContent, error)

	RegistryCreate(params *RegistryCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryCreateOK, error)

	RegistryDelete(params *RegistryDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryDeleteNoContent, error)

	RegistryInspect(params *RegistryInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryInspectOK, error)

	RegistryList(params *RegistryListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryListOK, error)

	RegistryUpdate(params *RegistryUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryUpdateOK, error)

	EcrDeleteRepository(params *EcrDeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EcrDeleteRepositoryOK, error)

	EcrDeleteTags(params *EcrDeleteTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EcrDeleteTagsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	RegistryConfigure configures a registry

	Configures a registry.

**Access policy**: restricted
*/
func (a *Client) RegistryConfigure(params *RegistryConfigureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryConfigureNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryConfigureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryConfigure",
		Method:             "POST",
		PathPattern:        "/registries/{id}/configure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryConfigureReader{formats: a.formats},
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
	success, ok := result.(*RegistryConfigureNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryConfigure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RegistryCreate creates a new registry

	Create a new registry.

**Access policy**: restricted
*/
func (a *Client) RegistryCreate(params *RegistryCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryCreate",
		Method:             "POST",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryCreateReader{formats: a.formats},
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
	success, ok := result.(*RegistryCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RegistryDelete removes a registry

	Remove a registry

**Access policy**: restricted
*/
func (a *Client) RegistryDelete(params *RegistryDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryDelete",
		Method:             "DELETE",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryDeleteReader{formats: a.formats},
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
	success, ok := result.(*RegistryDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RegistryInspect inspects a registry

	Retrieve details about a registry.

**Access policy**: restricted
*/
func (a *Client) RegistryInspect(params *RegistryInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryInspect",
		Method:             "GET",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryInspectReader{formats: a.formats},
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
	success, ok := result.(*RegistryInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RegistryList lists registries

	List all registries based on the current user authorizations.

Will return all registries if using an administrator account otherwise it
will only return authorized registries.
**Access policy**: restricted
*/
func (a *Client) RegistryList(params *RegistryListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryList",
		Method:             "GET",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryListReader{formats: a.formats},
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
	success, ok := result.(*RegistryListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RegistryUpdate updates a registry

	Update a registry

**Access policy**: restricted
*/
func (a *Client) RegistryUpdate(params *RegistryUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegistryUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegistryUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegistryUpdate",
		Method:             "PUT",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegistryUpdateReader{formats: a.formats},
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
	success, ok := result.(*RegistryUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RegistryUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EcrDeleteRepository deletes e c r repository

	Delete ECR repository.

**Access policy**: restricted
*/
func (a *Client) EcrDeleteRepository(params *EcrDeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EcrDeleteRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEcrDeleteRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ecrDeleteRepository",
		Method:             "DELETE",
		PathPattern:        "/registries/{id}/ecr/repositories/{repositoryName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EcrDeleteRepositoryReader{formats: a.formats},
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
	success, ok := result.(*EcrDeleteRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ecrDeleteRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EcrDeleteTags deletes tags

	Delete tags for a given ECR repository

**Access policy**: restricted
*/
func (a *Client) EcrDeleteTags(params *EcrDeleteTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EcrDeleteTagsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEcrDeleteTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ecrDeleteTags",
		Method:             "DELETE",
		PathPattern:        "/registries/{id}/ecr/repositories/{repositoryName}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EcrDeleteTagsReader{formats: a.formats},
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
	success, ok := result.(*EcrDeleteTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ecrDeleteTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
