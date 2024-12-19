// Code generated by go-swagger; DO NOT EDIT.

package edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new edge API client with basic auth credentials.
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

// New creates a new edge API client with a bearer token for authentication.
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
Client for edge API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetEndpointsIDEdgeStacksStackID(params *GetEndpointsIDEdgeStacksStackIDParams, opts ...ClientOption) (*GetEndpointsIDEdgeStacksStackIDOK, error)

	PostEndpointsEdgeGenerateKey(params *PostEndpointsEdgeGenerateKeyParams, opts ...ClientOption) (*PostEndpointsEdgeGenerateKeyOK, error)

	PostEndpointsIDEdgeJobsJobIDLogs(params *PostEndpointsIDEdgeJobsJobIDLogsParams, opts ...ClientOption) (*PostEndpointsIDEdgeJobsJobIDLogsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetEndpointsIDEdgeStacksStackID inspects an edge stack for an environment endpoint

**Access policy**: public
*/
func (a *Client) GetEndpointsIDEdgeStacksStackID(params *GetEndpointsIDEdgeStacksStackIDParams, opts ...ClientOption) (*GetEndpointsIDEdgeStacksStackIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointsIDEdgeStacksStackIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEndpointsIDEdgeStacksStackID",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}/edge/stacks/{stackId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEndpointsIDEdgeStacksStackIDReader{formats: a.formats},
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
	success, ok := result.(*GetEndpointsIDEdgeStacksStackIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointsIDEdgeStacksStackID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostEndpointsEdgeGenerateKey generates an edge key

	Generates a general edge key

**Access policy**: admin
*/
func (a *Client) PostEndpointsEdgeGenerateKey(params *PostEndpointsEdgeGenerateKeyParams, opts ...ClientOption) (*PostEndpointsEdgeGenerateKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEndpointsEdgeGenerateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEndpointsEdgeGenerateKey",
		Method:             "POST",
		PathPattern:        "/endpoints/edge/generate-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEndpointsEdgeGenerateKeyReader{formats: a.formats},
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
	success, ok := result.(*PostEndpointsEdgeGenerateKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEndpointsEdgeGenerateKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostEndpointsIDEdgeJobsJobIDLogs inspects an edge job log

**Access policy**: public
*/
func (a *Client) PostEndpointsIDEdgeJobsJobIDLogs(params *PostEndpointsIDEdgeJobsJobIDLogsParams, opts ...ClientOption) (*PostEndpointsIDEdgeJobsJobIDLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEndpointsIDEdgeJobsJobIDLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEndpointsIDEdgeJobsJobIDLogs",
		Method:             "POST",
		PathPattern:        "/endpoints/{id}/edge/jobs/{jobID}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEndpointsIDEdgeJobsJobIDLogsReader{formats: a.formats},
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
	success, ok := result.(*PostEndpointsIDEdgeJobsJobIDLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEndpointsIDEdgeJobsJobIDLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
