// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new endpoint groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoint groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EndpointGroupAddEndpoint(params *EndpointGroupAddEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupAddEndpointNoContent, error)

	EndpointGroupDelete(params *EndpointGroupDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupDeleteNoContent, error)

	EndpointGroupDeleteEndpoint(params *EndpointGroupDeleteEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupDeleteEndpointNoContent, error)

	EndpointGroupList(params *EndpointGroupListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupListOK, error)

	EndpointGroupUpdate(params *EndpointGroupUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupUpdateOK, error)

	GetEndpointGroupsID(params *GetEndpointGroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEndpointGroupsIDOK, error)

	PostEndpointGroups(params *PostEndpointGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEndpointGroupsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EndpointGroupAddEndpoint adds an endpoint to an endpoint group

  Add an endpoint to an endpoint group
**Access policy**: administrator
*/
func (a *Client) EndpointGroupAddEndpoint(params *EndpointGroupAddEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupAddEndpointNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupAddEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointGroupAddEndpoint",
		Method:             "PUT",
		PathPattern:        "/endpoint_groups/{id}/endpoints/{endpointId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupAddEndpointReader{formats: a.formats},
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
	success, ok := result.(*EndpointGroupAddEndpointNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointGroupAddEndpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointGroupDelete removes an endpoint group

  Remove an endpoint group.
**Access policy**: administrator
*/
func (a *Client) EndpointGroupDelete(params *EndpointGroupDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointGroupDelete",
		Method:             "DELETE",
		PathPattern:        "/endpoint_groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupDeleteReader{formats: a.formats},
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
	success, ok := result.(*EndpointGroupDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointGroupDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointGroupDeleteEndpoint removes endpoint from an endpoint group

  **Access policy**: administrator
*/
func (a *Client) EndpointGroupDeleteEndpoint(params *EndpointGroupDeleteEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupDeleteEndpointNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupDeleteEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointGroupDeleteEndpoint",
		Method:             "DELETE",
		PathPattern:        "/endpoint_groups/{id}/endpoints/{endpointId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupDeleteEndpointReader{formats: a.formats},
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
	success, ok := result.(*EndpointGroupDeleteEndpointNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointGroupDeleteEndpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointGroupList lists endpoint groups

  List all endpoint groups based on the current user authorizations. Will
return all endpoint groups if using an administrator account otherwise it will
only return authorized endpoint groups.
**Access policy**: restricted
*/
func (a *Client) EndpointGroupList(params *EndpointGroupListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointGroupList",
		Method:             "GET",
		PathPattern:        "/endpoint_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupListReader{formats: a.formats},
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
	success, ok := result.(*EndpointGroupListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointGroupList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EndpointGroupUpdate updates an endpoint group

  Update an endpoint group.
**Access policy**: administrator
*/
func (a *Client) EndpointGroupUpdate(params *EndpointGroupUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointGroupUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointGroupUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointGroupUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoint_groups/:id",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointGroupUpdateReader{formats: a.formats},
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
	success, ok := result.(*EndpointGroupUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointGroupUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEndpointGroupsID inspects an endpoint group

  Retrieve details abont an endpoint group.
**Access policy**: administrator
*/
func (a *Client) GetEndpointGroupsID(params *GetEndpointGroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEndpointGroupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointGroupsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEndpointGroupsID",
		Method:             "GET",
		PathPattern:        "/endpoint_groups/:id",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEndpointGroupsIDReader{formats: a.formats},
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
	success, ok := result.(*GetEndpointGroupsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointGroupsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostEndpointGroups creates an endpoint group

  Create a new endpoint group.
**Access policy**: administrator
*/
func (a *Client) PostEndpointGroups(params *PostEndpointGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEndpointGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEndpointGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEndpointGroups",
		Method:             "POST",
		PathPattern:        "/endpoint_groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEndpointGroupsReader{formats: a.formats},
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
	success, ok := result.(*PostEndpointGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEndpointGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
