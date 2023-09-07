// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CustomTemplateCreateFile(params *CustomTemplateCreateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateFileOK, error)

	CustomTemplateCreateRepository(params *CustomTemplateCreateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateRepositoryOK, error)

	CustomTemplateCreateString(params *CustomTemplateCreateStringParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateStringOK, error)

	CustomTemplateDelete(params *CustomTemplateDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateDeleteNoContent, error)

	CustomTemplateFile(params *CustomTemplateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateFileOK, error)

	CustomTemplateGitFetch(params *CustomTemplateGitFetchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateGitFetchOK, error)

	CustomTemplateInspect(params *CustomTemplateInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateInspectOK, error)

	CustomTemplateList(params *CustomTemplateListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateListOK, error)

	CustomTemplateUpdate(params *CustomTemplateUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CustomTemplateCreateFile creates a custom template

	Create a custom template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateCreateFile(params *CustomTemplateCreateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateCreateFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateCreateFile",
		Method:             "POST",
		PathPattern:        "/custom_templates/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateCreateFileReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateCreateFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateCreateFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateCreateRepository creates a custom template

	Create a custom template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateCreateRepository(params *CustomTemplateCreateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateCreateRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateCreateRepository",
		Method:             "POST",
		PathPattern:        "/custom_templates/repository",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateCreateRepositoryReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateCreateRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateCreateRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateCreateString creates a custom template

	Create a custom template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateCreateString(params *CustomTemplateCreateStringParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateCreateStringOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateCreateStringParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateCreateString",
		Method:             "POST",
		PathPattern:        "/custom_templates/string",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateCreateStringReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateCreateStringOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateCreateString: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateDelete removes a template

	Remove a template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateDelete(params *CustomTemplateDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateDelete",
		Method:             "DELETE",
		PathPattern:        "/custom_templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateDeleteReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateFile gets template stack file content

	Retrieve the content of the Stack file for the specified custom template

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateFile(params *CustomTemplateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateFile",
		Method:             "GET",
		PathPattern:        "/custom_templates/{id}/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateFileReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateGitFetch fetches the latest config file content based on custom template s git repository configuration

	Retrieve details about a template created from git repository method.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateGitFetch(params *CustomTemplateGitFetchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateGitFetchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateGitFetchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateGitFetch",
		Method:             "PUT",
		PathPattern:        "/custom_templates/{id}/git_fetch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateGitFetchReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateGitFetchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateGitFetch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateInspect inspects a custom template

	Retrieve details about a template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateInspect(params *CustomTemplateInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateInspect",
		Method:             "GET",
		PathPattern:        "/custom_templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateInspectReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateList lists available custom templates

	List available custom templates.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateList(params *CustomTemplateListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateList",
		Method:             "GET",
		PathPattern:        "/custom_templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateListReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CustomTemplateUpdate updates a template

	Update a template.

**Access policy**: authenticated
*/
func (a *Client) CustomTemplateUpdate(params *CustomTemplateUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CustomTemplateUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomTemplateUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomTemplateUpdate",
		Method:             "PUT",
		PathPattern:        "/custom_templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CustomTemplateUpdateReader{formats: a.formats},
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
	success, ok := result.(*CustomTemplateUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CustomTemplateUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
