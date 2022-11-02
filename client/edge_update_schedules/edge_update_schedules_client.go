// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge update schedules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge update schedules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EdgeUpdatePreviousVersions(params *EdgeUpdatePreviousVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdatePreviousVersionsOK, error)

	EdgeUpdateScheduleActiveSchedulesList(params *EdgeUpdateScheduleActiveSchedulesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleActiveSchedulesListOK, error)

	EdgeUpdateScheduleDelete(params *EdgeUpdateScheduleDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleDeleteNoContent, error)

	EdgeUpdateScheduleInspect(params *EdgeUpdateScheduleInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleInspectOK, error)

	EdgeUpdateScheduleList(params *EdgeUpdateScheduleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleListOK, error)

	EdgeUpdateScheduleUpdate(params *EdgeUpdateScheduleUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleUpdateNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
EdgeUpdatePreviousVersions fetches the previous versions of updated agents

**Access policy**: authenticated
*/
func (a *Client) EdgeUpdatePreviousVersions(params *EdgeUpdatePreviousVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdatePreviousVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdatePreviousVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdatePreviousVersions",
		Method:             "GET",
		PathPattern:        "/edge_update_schedules/agent_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdatePreviousVersionsReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdatePreviousVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdatePreviousVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeUpdateScheduleActiveSchedulesList fetches the list of active edge update schedules

**Access policy**: administrator
*/
func (a *Client) EdgeUpdateScheduleActiveSchedulesList(params *EdgeUpdateScheduleActiveSchedulesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleActiveSchedulesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateScheduleActiveSchedulesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdateScheduleActiveSchedulesList",
		Method:             "GET",
		PathPattern:        "/edge_update_schedules/active",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdateScheduleActiveSchedulesListReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdateScheduleActiveSchedulesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdateScheduleActiveSchedulesList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeUpdateScheduleDelete deletes an edge update schedule

**Access policy**: administrator
*/
func (a *Client) EdgeUpdateScheduleDelete(params *EdgeUpdateScheduleDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateScheduleDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdateScheduleDelete",
		Method:             "DELETE",
		PathPattern:        "/edge_update_schedules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdateScheduleDeleteReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdateScheduleDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdateScheduleDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeUpdateScheduleInspect returns the edge update schedule with the given ID

**Access policy**: administrator
*/
func (a *Client) EdgeUpdateScheduleInspect(params *EdgeUpdateScheduleInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateScheduleInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdateScheduleInspect",
		Method:             "GET",
		PathPattern:        "/edge_update_schedules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdateScheduleInspectReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdateScheduleInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdateScheduleInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeUpdateScheduleList fetches the list of edge update schedules

**Access policy**: administrator
*/
func (a *Client) EdgeUpdateScheduleList(params *EdgeUpdateScheduleListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateScheduleListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdateScheduleList",
		Method:             "GET",
		PathPattern:        "/edge_update_schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdateScheduleListReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdateScheduleListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdateScheduleList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EdgeUpdateScheduleUpdate updates an edge update schedule

**Access policy**: administrator
*/
func (a *Client) EdgeUpdateScheduleUpdate(params *EdgeUpdateScheduleUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EdgeUpdateScheduleUpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateScheduleUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EdgeUpdateScheduleUpdate",
		Method:             "POST",
		PathPattern:        "/edge_update_schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EdgeUpdateScheduleUpdateReader{formats: a.formats},
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
	success, ok := result.(*EdgeUpdateScheduleUpdateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EdgeUpdateScheduleUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
