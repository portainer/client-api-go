// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new team memberships API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new team memberships API client with basic auth credentials.
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

// New creates a new team memberships API client with a bearer token for authentication.
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
Client for team memberships API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TeamMembershipCreate(params *TeamMembershipCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipCreateOK, error)

	TeamMembershipDelete(params *TeamMembershipDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipDeleteNoContent, error)

	TeamMembershipList(params *TeamMembershipListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipListOK, error)

	TeamMembershipUpdate(params *TeamMembershipUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipUpdateOK, error)

	TeamMemberships(params *TeamMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	TeamMembershipCreate creates a new team membership

	Create a new team memberships. Access is only available to administrators leaders of the associated team.

**Access policy**: administrator
*/
func (a *Client) TeamMembershipCreate(params *TeamMembershipCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamMembershipCreate",
		Method:             "POST",
		PathPattern:        "/team_memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipCreateReader{formats: a.formats},
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
	success, ok := result.(*TeamMembershipCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamMembershipCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamMembershipDelete removes a team membership

	Remove a team membership. Access is only available to administrators leaders of the associated team.

**Access policy**: administrator
*/
func (a *Client) TeamMembershipDelete(params *TeamMembershipDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamMembershipDelete",
		Method:             "DELETE",
		PathPattern:        "/team_memberships/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipDeleteReader{formats: a.formats},
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
	success, ok := result.(*TeamMembershipDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamMembershipDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamMembershipList lists team memberships

	List team memberships. Access is only available to administrators and team leaders.

**Access policy**: administrator
*/
func (a *Client) TeamMembershipList(params *TeamMembershipListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamMembershipList",
		Method:             "GET",
		PathPattern:        "/team_memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipListReader{formats: a.formats},
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
	success, ok := result.(*TeamMembershipListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamMembershipList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamMembershipUpdate updates a team membership

	Update a team membership. Access is only available to administrators or leaders of the associated team.

**Access policy**: administrator or leaders of the associated team
*/
func (a *Client) TeamMembershipUpdate(params *TeamMembershipUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamMembershipUpdate",
		Method:             "PUT",
		PathPattern:        "/team_memberships/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipUpdateReader{formats: a.formats},
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
	success, ok := result.(*TeamMembershipUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamMembershipUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TeamMemberships lists team memberships

	List team memberships. Access is only available to administrators and team leaders.

**Access policy**: restricted
*/
func (a *Client) TeamMemberships(params *TeamMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TeamMembershipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamMembershipsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TeamMemberships",
		Method:             "GET",
		PathPattern:        "/teams/{id}/memberships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TeamMembershipsReader{formats: a.formats},
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
	success, ok := result.(*TeamMembershipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TeamMemberships: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
