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
	DeleteEdgeStacksID(params *DeleteEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeStacksIDNoContent, error)

	GetEdgeStacks(params *GetEdgeStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksOK, error)

	GetEdgeStacksID(params *GetEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksIDOK, error)

	GetEdgeStacksIDFile(params *GetEdgeStacksIDFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksIDFileOK, error)

	PostEdgeStacks(params *PostEdgeStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeStacksOK, error)

	PutEdgeStacksID(params *PutEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEdgeStacksIDOK, error)

	PutEdgeStacksIDStatus(params *PutEdgeStacksIDStatusParams, opts ...ClientOption) (*PutEdgeStacksIDStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteEdgeStacksID deletes an edge stack
*/
func (a *Client) DeleteEdgeStacksID(params *DeleteEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeStacksIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEdgeStacksIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEdgeStacksID",
		Method:             "DELETE",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEdgeStacksIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteEdgeStacksIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEdgeStacksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeStacks fetches the list of edge stacks
*/
func (a *Client) GetEdgeStacks(params *GetEdgeStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeStacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeStacks",
		Method:             "GET",
		PathPattern:        "/edge_stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeStacksReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeStacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeStacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeStacksID inspects an edge stack
*/
func (a *Client) GetEdgeStacksID(params *GetEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeStacksIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeStacksID",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeStacksIDReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeStacksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeStacksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeStacksIDFile fetches the stack file for an edge stack
*/
func (a *Client) GetEdgeStacksIDFile(params *GetEdgeStacksIDFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeStacksIDFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeStacksIDFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeStacksIDFile",
		Method:             "GET",
		PathPattern:        "/edge_stacks/{id}/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeStacksIDFileReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeStacksIDFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeStacksIDFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostEdgeStacks creates an edge stack
*/
func (a *Client) PostEdgeStacks(params *PostEdgeStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeStacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEdgeStacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEdgeStacks",
		Method:             "POST",
		PathPattern:        "/edge_stacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEdgeStacksReader{formats: a.formats},
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
	success, ok := result.(*PostEdgeStacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEdgeStacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutEdgeStacksID updates an edge stack
*/
func (a *Client) PutEdgeStacksID(params *PutEdgeStacksIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEdgeStacksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEdgeStacksIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutEdgeStacksID",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutEdgeStacksIDReader{formats: a.formats},
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
	success, ok := result.(*PutEdgeStacksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutEdgeStacksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutEdgeStacksIDStatus updates an edge stack status

  Authorized only if the request is done by an Edge Endpoint
*/
func (a *Client) PutEdgeStacksIDStatus(params *PutEdgeStacksIDStatusParams, opts ...ClientOption) (*PutEdgeStacksIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEdgeStacksIDStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutEdgeStacksIDStatus",
		Method:             "PUT",
		PathPattern:        "/edge_stacks/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutEdgeStacksIDStatusReader{formats: a.formats},
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
	success, ok := result.(*PutEdgeStacksIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutEdgeStacksIDStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
