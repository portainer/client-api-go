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
	EndpointAssociationDelete(params *EndpointAssociationDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointAssociationDeleteOK, error)

	EndpointCreate(params *EndpointCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointCreateOK, error)

	EndpointCreateGlobalKey(params *EndpointCreateGlobalKeyParams, opts ...ClientOption) (*EndpointCreateGlobalKeyOK, error)

	EndpointDelete(params *EndpointDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointDeleteNoContent, error)

	EndpointEdgeStatusInspect(params *EndpointEdgeStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointEdgeStatusInspectOK, error)

	EndpointInspect(params *EndpointInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointInspectOK, error)

	EndpointList(params *EndpointListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointListOK, error)

	EndpointSettingsUpdate(params *EndpointSettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSettingsUpdateOK, error)

	EndpointSnapshot(params *EndpointSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotNoContent, error)

	EndpointSnapshots(params *EndpointSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSnapshotsNoContent, error)

	EndpointUpdate(params *EndpointUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointUpdateOK, error)

	EndpointDockerhubStatus(params *EndpointDockerhubStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointDockerhubStatusOK, error)

	EndpointEdgeAsync(params *EndpointEdgeAsyncParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointEdgeAsyncOK, error)

	EndpointForceUpdateService(params *EndpointForceUpdateServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointForceUpdateServiceOK, error)

	EndpointPoolsAccessUpdate(params *EndpointPoolsAccessUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointPoolsAccessUpdateNoContent, error)

	EndpointRegistriesList(params *EndpointRegistriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointRegistriesListOK, error)

	EndpointRegistryAccess(params *EndpointRegistryAccessParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointRegistryAccessNoContent, error)

	SnapshotContainerInspect(params *SnapshotContainerInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SnapshotContainerInspectOK, error)

	SnapshotContainersList(params *SnapshotContainersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SnapshotContainersListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	EndpointAssociationDelete des association an edge environment endpoint

	De-association an edge environment(endpoint).

**Access policy**: administrator
*/
func (a *Client) EndpointAssociationDelete(params *EndpointAssociationDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointAssociationDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointAssociationDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointAssociationDelete",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/association",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointAssociationDeleteReader{formats: a.formats},
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
	success, ok := result.(*EndpointAssociationDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointAssociationDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointCreate creates a new environment endpoint

	Create a new environment(endpoint) that will be used to manage an environment(endpoint).

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
EndpointCreateGlobalKey creates or retrieve the endpoint for an edge ID
*/
func (a *Client) EndpointCreateGlobalKey(params *EndpointCreateGlobalKeyParams, opts ...ClientOption) (*EndpointCreateGlobalKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointCreateGlobalKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointCreateGlobalKey",
		Method:             "POST",
		PathPattern:        "/endpoints/global-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointCreateGlobalKeyReader{formats: a.formats},
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
	success, ok := result.(*EndpointCreateGlobalKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointCreateGlobalKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointDelete removes an environment endpoint

	Remove an environment(endpoint).

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
	EndpointEdgeStatusInspect gets environment endpoint status

	environment(endpoint) for edge agent to check status of environment(endpoint)

**Access policy**: restricted only to Edge environments(endpoints)
*/
func (a *Client) EndpointEdgeStatusInspect(params *EndpointEdgeStatusInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointEdgeStatusInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointEdgeStatusInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointEdgeStatusInspect",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}/edge/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointEdgeStatusInspectReader{formats: a.formats},
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
	success, ok := result.(*EndpointEdgeStatusInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointEdgeStatusInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointInspect inspects an environment endpoint

	Retrieve details about an environment(endpoint).

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
	EndpointList lists environments endpoints

	List all environments(endpoints) based on the current user authorizations. Will

return all environments(endpoints) if using an administrator or team leader account otherwise it will
only return authorized environments(endpoints).
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
	EndpointSettingsUpdate updates settings for an environment endpoint

	Update settings for an environment(endpoint).

**Access policy**: authenticated
*/
func (a *Client) EndpointSettingsUpdate(params *EndpointSettingsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointSettingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointSettingsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "EndpointSettingsUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointSettingsUpdateReader{formats: a.formats},
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
	success, ok := result.(*EndpointSettingsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EndpointSettingsUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointSnapshot snapshots an environment endpoint

	Snapshots an environment(endpoint)

**Access policy**: administrator
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
	EndpointSnapshots snapshots all environment endpoint

	Snapshot all environments(endpoints)

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
	EndpointUpdate updates an environment endpoint

	Update an environment(endpoint).

**Access policy**: authenticated
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

/*
	EndpointDockerhubStatus fetches docker pull limits

	get docker pull limits for a docker hub registry in the environment

**Access policy**:
*/
func (a *Client) EndpointDockerhubStatus(params *EndpointDockerhubStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointDockerhubStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointDockerhubStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointDockerhubStatus",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}/dockerhub/{registryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointDockerhubStatusReader{formats: a.formats},
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
	success, ok := result.(*EndpointDockerhubStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointDockerhubStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointEdgeAsync gets environment endpoint status

	Environment(Endpoint) for edge agent to check status of environment(endpoint)

**Access policy**: restricted only to Edge environments(endpoints)
*/
func (a *Client) EndpointEdgeAsync(params *EndpointEdgeAsyncParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointEdgeAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointEdgeAsyncParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointEdgeAsync",
		Method:             "POST",
		PathPattern:        "/endpoints/edge/async",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointEdgeAsyncReader{formats: a.formats},
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
	success, ok := result.(*EndpointEdgeAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointEdgeAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointForceUpdateService forces update a docker service

	force update a docker service

**Access policy**: authenticated
*/
func (a *Client) EndpointForceUpdateService(params *EndpointForceUpdateServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointForceUpdateServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointForceUpdateServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointForceUpdateService",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/forceupdateservice",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointForceUpdateServiceReader{formats: a.formats},
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
	success, ok := result.(*EndpointForceUpdateServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointForceUpdateService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointPoolsAccessUpdate updates resource pool access

	update the access on the resource pool in the current environment

**Access policy**: restricted
*/
func (a *Client) EndpointPoolsAccessUpdate(params *EndpointPoolsAccessUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointPoolsAccessUpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointPoolsAccessUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointPoolsAccessUpdate",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/pools/{rpn}/access",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointPoolsAccessUpdateReader{formats: a.formats},
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
	success, ok := result.(*EndpointPoolsAccessUpdateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointPoolsAccessUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EndpointRegistriesList lists registries on environment

	List all registries based on the current user authorizations in current environment.

**Access policy**: authenticated
*/
func (a *Client) EndpointRegistriesList(params *EndpointRegistriesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointRegistriesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointRegistriesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointRegistriesList",
		Method:             "GET",
		PathPattern:        "/endpoints/{id}/registries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointRegistriesListReader{formats: a.formats},
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
	success, ok := result.(*EndpointRegistriesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointRegistriesList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EndpointRegistryAccess updates registry access for environment

**Access policy**: authenticated
*/
func (a *Client) EndpointRegistryAccess(params *EndpointRegistryAccessParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EndpointRegistryAccessNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointRegistryAccessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "endpointRegistryAccess",
		Method:             "PUT",
		PathPattern:        "/endpoints/{id}/registries/{registryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EndpointRegistryAccessReader{formats: a.formats},
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
	success, ok := result.(*EndpointRegistryAccessNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpointRegistryAccess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SnapshotContainerInspect fetches container from a snapshot

**Access policy**:
*/
func (a *Client) SnapshotContainerInspect(params *SnapshotContainerInspectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SnapshotContainerInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotContainerInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "snapshotContainerInspect",
		Method:             "GET",
		PathPattern:        "/docker/{environmentId}/snapshot/container/{containerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SnapshotContainerInspectReader{formats: a.formats},
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
	success, ok := result.(*SnapshotContainerInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotContainerInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SnapshotContainersList fetches containers list from a snapshot

**Access policy**:
*/
func (a *Client) SnapshotContainersList(params *SnapshotContainersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SnapshotContainersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotContainersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "snapshotContainersList",
		Method:             "GET",
		PathPattern:        "/docker/{environmentId}/snapshot/containers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SnapshotContainersListReader{formats: a.formats},
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
	success, ok := result.(*SnapshotContainersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for snapshotContainersList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
