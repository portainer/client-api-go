// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge stacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge stacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EdgeStackCreateFile(params *EdgeStackCreateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateFileOK, error)

	EdgeStackCreateRepository(params *EdgeStackCreateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateRepositoryOK, error)

	EdgeStackCreateString(params *EdgeStackCreateStringParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateStringOK, error)

	EdgeStackDelete(params *EdgeStackDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackDeleteNoContent, error)

	EdgeStackFile(params *EdgeStackFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackFileOK, error)

	EdgeStackInspect(params *EdgeStackInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackInspectOK, error)

	EdgeStackList(params *EdgeStackListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackListOK, error)

	EdgeStackLogsCollect(params *EdgeStackLogsCollectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsCollectNoContent, error)

	EdgeStackLogsDelete(params *EdgeStackLogsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsDeleteNoContent, error)

	EdgeStackLogsStatusGet(params *EdgeStackLogsStatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsStatusGetOK, error)

	EdgeStackStaggerStatusInspect(params *EdgeStackStaggerStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackStaggerStatusInspectOK, error)

	EdgeStackStatusDelete(params *EdgeStackStatusDeleteParams, opts ...ClientOption) (*EdgeStackStatusDeleteOK, error)

	EdgeStackStatusUpdate(params *EdgeStackStatusUpdateParams, opts ...ClientOption) (*EdgeStackStatusUpdateOK, error)

	EdgeStackUpdate(params *EdgeStackUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackUpdateOK, error)

	EdgeStackUpdateFromGit(params *EdgeStackUpdateFromGitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackUpdateFromGitNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
EdgeStackCreateFile creates an edge stack from file

**Access policy**: administrator
*/
func (a *Client) EdgeStackCreateFile(params *EdgeStackCreateFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackCreateFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackCreateFile",
		Method:             "POST",
		PathPattern:        "/edge_stacks/create/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackCreateFileReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackCreateFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackCreateFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackCreateRepository creates an edge stack from a git repository

**Access policy**: administrator
*/
func (a *Client) EdgeStackCreateRepository(params *EdgeStackCreateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackCreateRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackCreateRepository",
		Method:             "POST",
		PathPattern:        "/edge_stacks/create/repository",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackCreateRepositoryReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackCreateRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackCreateRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackCreateString creates an edge stack from a text

**Access policy**: administrator
*/
func (a *Client) EdgeStackCreateString(params *EdgeStackCreateStringParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackCreateStringOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackCreateStringParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackCreateString",
		Method:             "POST",
		PathPattern:        "/edge_stacks/create/string",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackCreateStringReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackCreateStringOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackCreateString: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackDelete deletes an edge stack

**Access policy**: administrator
*/
func (a *Client) EdgeStackDelete(params *EdgeStackDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackDelete",
		Method:             "DELETE",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackDeleteReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackFile fetches the stack file for an edge stack

**Access policy**: administrator
*/
func (a *Client) EdgeStackFile(params *EdgeStackFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackFile",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackFileReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackInspect inspects an edge stack

**Access policy**: administrator
*/
func (a *Client) EdgeStackInspect(params *EdgeStackInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackInspect",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackInspectReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackList fetches the list of edge stacks

**Access policy**: administrator
*/
func (a *Client) EdgeStackList(params *EdgeStackListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackList",
		Method:             "GET",
		PathPattern:        "/edge_stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackListReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackLogsCollect schedules the collection of logs for a given endpoint and edge stack

**Access policy**: administrator
*/
func (a *Client) EdgeStackLogsCollect(params *EdgeStackLogsCollectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsCollectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackLogsCollectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackLogsCollect",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}/logs/{endpoint_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackLogsCollectReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackLogsCollectNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackLogsCollect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackLogsDelete deletes the available logs for a given edge stack and endpoint

**Access policy**: administrator
*/
func (a *Client) EdgeStackLogsDelete(params *EdgeStackLogsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackLogsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackLogsDelete",
		Method:             "DELETE",
		PathPattern:        "/edge_stacks/{id}/logs/{endpoint_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackLogsDeleteReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackLogsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackLogsDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackLogsStatusGet gets the status of the log collection for a given edgestack and environment

**Access policy**: administrator
*/
func (a *Client) EdgeStackLogsStatusGet(params *EdgeStackLogsStatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackLogsStatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackLogsStatusGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackLogsStatusGet",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}/logs/{endpoint_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackLogsStatusGetReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackLogsStatusGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackLogsStatusGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackStaggerStatusInspect inspects an edge stack s parallel update status

**Access policy**: administrator
*/
func (a *Client) EdgeStackStaggerStatusInspect(params *EdgeStackStaggerStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackStaggerStatusInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackStaggerStatusInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackStaggerStatusInspect",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}/stagger/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackStaggerStatusInspectReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackStaggerStatusInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackStaggerStatusInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackStatusDelete deletes an edge stack status

Authorized only if the request is done by an Edge Environment(Endpoint)
*/
func (a *Client) EdgeStackStatusDelete(params *EdgeStackStatusDeleteParams, opts ...ClientOption) (*EdgeStackStatusDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackStatusDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackStatusDelete",
		Method:             "DELETE",
		PathPattern:        "/edge_stacks/{id}/status/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackStatusDeleteReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackStatusDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackStatusDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackStatusUpdate updates an edge stack status

Authorized only if the request is done by an Edge Environment(Endpoint)
*/
func (a *Client) EdgeStackStatusUpdate(params *EdgeStackStatusUpdateParams, opts ...ClientOption) (*EdgeStackStatusUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackStatusUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackStatusUpdate",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackStatusUpdateReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackStatusUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackStatusUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackUpdate updates an edge stack

**Access policy**: administrator
*/
func (a *Client) EdgeStackUpdate(params *EdgeStackUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeStackUpdate",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackUpdateReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeStackUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeStackUpdateFromGit updates git configuration and pulls the repository

**Access policy**: authenticated
*/
func (a *Client) EdgeStackUpdateFromGit(params *EdgeStackUpdateFromGitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeStackUpdateFromGitNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeStackUpdateFromGitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "edgeStackUpdateFromGit",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}/git",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeStackUpdateFromGitReader{formats: a.formats},
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
	success, ok := result.(*EdgeStackUpdateFromGitNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for edgeStackUpdateFromGit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
