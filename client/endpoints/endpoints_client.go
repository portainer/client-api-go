// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EndpointCreate(params *EndpointCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointCreateOK, error)

	EndpointDelete(params *EndpointDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointDeleteNoContent, error)

	EndpointInspect(params *EndpointInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointInspectOK, error)

	EndpointList(params *EndpointListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointListOK, error)

	EndpointSnapshot(params *EndpointSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotNoContent, error)

	EndpointSnapshots(params *EndpointSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotsNoContent, error)

	EndpointStatusInspect(params *EndpointStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointStatusInspectOK, error)

	EndpointUpdate(params *EndpointUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EndpointCreate creates a new endpoint

  Create a new endpoint that will be used to manage an environment.
**Access policy**: administrator
*/
func (a *Client) EndpointCreate(params *EndpointCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointCreate",
		Method:             "POST",
		PathPattern:        "/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointCreateReader{formats: a.formats},
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
	success, ok := result.(*EndpointCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointDelete removes an endpoint

  Remove an endpoint.
**Access policy**: administrator
*/
func (a *Client) EndpointDelete(params *EndpointDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointDelete",
		Method:             "DELETE",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointDeleteReader{formats: a.formats},
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
	success, ok := result.(*EndpointDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointInspect inspects an endpoint

  Retrieve details about an endpoint.
**Access policy**: restricted
*/
func (a *Client) EndpointInspect(params *EndpointInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointInspect",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointInspectReader{formats: a.formats},
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
	success, ok := result.(*EndpointInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointList lists endpoints

  List all endpoints based on the current user authorizations. Will
return all endpoints if using an administrator account otherwise it will
only return authorized endpoints.
**Access policy**: restricted
*/
func (a *Client) EndpointList(params *EndpointListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointList",
		Method:             "GET",
		PathPattern:        "/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointListReader{formats: a.formats},
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
	success, ok := result.(*EndpointListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointSnapshot snapshots an endpoint

  Snapshots an endpoint
**Access policy**: restricted
*/
func (a *Client) EndpointSnapshot(params *EndpointSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointSnapshot",
		Method:             "POST",
		PathPattern:        "/endpoints/{id}/snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointSnapshotReader{formats: a.formats},
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
	success, ok := result.(*EndpointSnapshotNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointSnapshots snapshots all endpoints

  Snapshot all endpoints
**Access policy**: administrator
*/
func (a *Client) EndpointSnapshots(params *EndpointSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointSnapshots",
		Method:             "POST",
		PathPattern:        "/endpoints/snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*EndpointSnapshotsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointSnapshots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointStatusInspect gets endpoint status

  Endpoint for edge agent to check status of environment
**Access policy**: restricted only to Edge endpoints
*/
func (a *Client) EndpointStatusInspect(params *EndpointStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointStatusInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointStatusInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointStatusInspect",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointStatusInspectReader{formats: a.formats},
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
	success, ok := result.(*EndpointStatusInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointStatusInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointUpdate updates an endpoint

  Update an endpoint.
**Access policy**: administrator
*/
func (a *Client) EndpointUpdate(params *EndpointUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointUpdateReader{formats: a.formats},
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
	success, ok := result.(*EndpointUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
