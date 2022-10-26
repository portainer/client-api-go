// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TeamCreate(params *TeamCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamCreateOK, error)

	TeamDelete(params *TeamDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamDeleteNoContent, error)

	TeamInspect(params *TeamInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamInspectOK, *TeamInspectNoContent, error)

	TeamList(params *TeamListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamListOK, error)

	TeamUpdate(params *TeamUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamUpdateOK, *TeamUpdateNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	TeamCreate creates a new team

	Create a new team.

**Access policy**: administrator
*/
func (a *Client) TeamCreate(params *TeamCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamCreate",
		Method:             "POST",
		PathPattern:        "/team",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamCreateReader{formats: a.formats},
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
	success, ok := result.(*TeamCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamDelete removes a team

	Remove a team.

**Access policy**: administrator
*/
func (a *Client) TeamDelete(params *TeamDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamDelete",
		Method:             "DELETE",
		PathPattern:        "/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamDeleteReader{formats: a.formats},
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
	success, ok := result.(*TeamDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamInspect inspects a team

	Retrieve details about a team. Access is only available for administrator and leaders of that team.

**Access policy**: administrator or team leader
*/
func (a *Client) TeamInspect(params *TeamInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamInspectOK, *TeamInspectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamInspect",
		Method:             "GET",
		PathPattern:        "/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamInspectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TeamInspectOK:
		return value, nil, nil
	case *TeamInspectNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamList lists teams

	List teams. For non-administrator users, will only list the teams they are member of.

**Access policy**: restricted
*/
func (a *Client) TeamList(params *TeamListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamList",
		Method:             "GET",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamListReader{formats: a.formats},
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
	success, ok := result.(*TeamListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamUpdate updates a team

	Update a team.

**Access policy**: administrator
*/
func (a *Client) TeamUpdate(params *TeamUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamUpdateOK, *TeamUpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamUpdate",
		Method:             "PUT",
		PathPattern:        "/team/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TeamUpdateOK:
		return value, nil, nil
	case *TeamUpdateNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
