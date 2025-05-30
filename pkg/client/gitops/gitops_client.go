// Code generated by go-swagger; DO NOT EDIT.

package gitops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new gitops API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new gitops API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new gitops API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for gitops API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GitOperationRepoFilePreview(params *GitOperationRepoFilePreviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoFilePreviewOK, error)

	GitOperationRepoFilesSearch(params *GitOperationRepoFilesSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoFilesSearchOK, error)

	GitOperationRepoRefs(params *GitOperationRepoRefsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoRefsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GitOperationRepoFilePreview previews the content of target file in the git repository

	Retrieve the compose file content based on git repository configuration

**Access policy**: authenticated
*/
func (a *Client) GitOperationRepoFilePreview(params *GitOperationRepoFilePreviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoFilePreviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGitOperationRepoFilePreviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GitOperationRepoFilePreview",
		Method:             "POST",
		PathPattern:        "/gitops/repo/file/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GitOperationRepoFilePreviewReader{formats: a.formats},
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
	success, ok := result.(*GitOperationRepoFilePreviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GitOperationRepoFilePreview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GitOperationRepoFilesSearch searches the file path from a git repository files with specified extensions

	Search the file path from the git repository based on partial or completed filename

**Access policy**: authenticated
*/
func (a *Client) GitOperationRepoFilesSearch(params *GitOperationRepoFilesSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoFilesSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGitOperationRepoFilesSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GitOperationRepoFilesSearch",
		Method:             "POST",
		PathPattern:        "/gitops/repo/files/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GitOperationRepoFilesSearchReader{formats: a.formats},
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
	success, ok := result.(*GitOperationRepoFilesSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GitOperationRepoFilesSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GitOperationRepoRefs lists the refs of a git repository

	List all the refs of a git repository

Will return all refs of a git repository
**Access policy**: authenticated
*/
func (a *Client) GitOperationRepoRefs(params *GitOperationRepoRefsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GitOperationRepoRefsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGitOperationRepoRefsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GitOperationRepoRefs",
		Method:             "POST",
		PathPattern:        "/gitops/repo/refs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GitOperationRepoRefsReader{formats: a.formats},
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
	success, ok := result.(*GitOperationRepoRefsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GitOperationRepoRefs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
