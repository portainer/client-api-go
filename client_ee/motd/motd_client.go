// Code generated by go-swagger; DO NOT EDIT.

package motd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new motd API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for motd API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MOTD(params *MOTDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MOTDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
MOTD fetches the message of the day

**Access policy**: restricted
*/
func (a *Client) MOTD(params *MOTDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MOTDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMOTDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MOTD",
		Method:             "GET",
		PathPattern:        "/motd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MOTDReader{formats: a.formats},
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
	success, ok := result.(*MOTDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for MOTD: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
