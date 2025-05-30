// Code generated by go-swagger; DO NOT EDIT.

package websocket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new websocket API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new websocket API client with basic auth credentials.
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

// New creates a new websocket API client with a bearer token for authentication.
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
Client for websocket API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetWebsocketAttach(params *GetWebsocketAttachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketAttachOK, error)

	GetWebsocketExec(params *GetWebsocketExecParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketExecOK, error)

	GetWebsocketKubernetesShell(params *GetWebsocketKubernetesShellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketKubernetesShellOK, error)

	GetWebsocketMicrok8sShell(params *GetWebsocketMicrok8sShellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketMicrok8sShellOK, error)

	GetWebsocketPod(params *GetWebsocketPodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketPodOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetWebsocketAttach attaches a websocket

	If the nodeName query parameter is present, the request will be proxied to the underlying agent environment(endpoint).

If the nodeName query parameter is not specified, the request will be upgraded to the websocket protocol and
an AttachStart operation HTTP request will be created and hijacked.
**Access policy**: authenticated
*/
func (a *Client) GetWebsocketAttach(params *GetWebsocketAttachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketAttachOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebsocketAttachParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWebsocketAttach",
		Method:             "GET",
		PathPattern:        "/websocket/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebsocketAttachReader{formats: a.formats},
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
	success, ok := result.(*GetWebsocketAttachOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWebsocketAttach: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetWebsocketExec executes a websocket

	If the nodeName query parameter is present, the request will be proxied to the underlying agent environment(endpoint).

If the nodeName query parameter is not specified, the request will be upgraded to the websocket protocol and
an ExecStart operation HTTP request will be created and hijacked.
*/
func (a *Client) GetWebsocketExec(params *GetWebsocketExecParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketExecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebsocketExecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWebsocketExec",
		Method:             "GET",
		PathPattern:        "/websocket/exec",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebsocketExecReader{formats: a.formats},
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
	success, ok := result.(*GetWebsocketExecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWebsocketExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetWebsocketKubernetesShell executes a websocket on kubectl shell pod

	The request will be upgraded to the websocket protocol. The request will proxy input from the client to the pod via long-lived websocket connection.

**Access policy**: authenticated
*/
func (a *Client) GetWebsocketKubernetesShell(params *GetWebsocketKubernetesShellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketKubernetesShellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebsocketKubernetesShellParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWebsocketKubernetesShell",
		Method:             "GET",
		PathPattern:        "/websocket/kubernetes-shell",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebsocketKubernetesShellReader{formats: a.formats},
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
	success, ok := result.(*GetWebsocketKubernetesShellOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWebsocketKubernetesShell: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetWebsocketMicrok8sShell connects to a remote SSH shell via a websocket

	When called, an SSH session to a microk8s cluster node will be created

an ssh session will be created and hijacked.
**Access policy**: authenticated
*/
func (a *Client) GetWebsocketMicrok8sShell(params *GetWebsocketMicrok8sShellParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketMicrok8sShellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebsocketMicrok8sShellParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWebsocketMicrok8sShell",
		Method:             "GET",
		PathPattern:        "/websocket/microk8s-shell",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebsocketMicrok8sShellReader{formats: a.formats},
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
	success, ok := result.(*GetWebsocketMicrok8sShellOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWebsocketMicrok8sShell: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetWebsocketPod executes a websocket on pod

	The request will be upgraded to the websocket protocol.

**Access policy**: authenticated
*/
func (a *Client) GetWebsocketPod(params *GetWebsocketPodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebsocketPodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebsocketPodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWebsocketPod",
		Method:             "GET",
		PathPattern:        "/websocket/pod",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebsocketPodReader{formats: a.formats},
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
	success, ok := result.(*GetWebsocketPodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWebsocketPod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
