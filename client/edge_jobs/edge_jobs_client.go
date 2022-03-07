// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new edge jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteEdgeJobsID(params *DeleteEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeJobsIDNoContent, error)

	DeleteEdgeJobsIDTasksTaskIDLogs(params *DeleteEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeJobsIDTasksTaskIDLogsNoContent, error)

	GetEdgeJobs(params *GetEdgeJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsOK, error)

	GetEdgeJobsID(params *GetEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDOK, error)

	GetEdgeJobsIDFile(params *GetEdgeJobsIDFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDFileOK, error)

	GetEdgeJobsIDTasks(params *GetEdgeJobsIDTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDTasksOK, error)

	GetEdgeJobsIDTasksTaskIDLogs(params *GetEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDTasksTaskIDLogsOK, error)

	PostEdgeJobs(params *PostEdgeJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsOK, error)

	PostEdgeJobsID(params *PostEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsIDOK, error)

	PostEdgeJobsIDTasksTaskIDLogs(params *PostEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsIDTasksTaskIDLogsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteEdgeJobsID deletes an edge job
*/
func (a *Client) DeleteEdgeJobsID(params *DeleteEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeJobsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEdgeJobsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEdgeJobsID",
		Method:             "DELETE",
		PathPattern:        "/edge_jobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEdgeJobsIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteEdgeJobsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEdgeJobsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEdgeJobsIDTasksTaskIDLogs clears the log for a specifc task on an edge job
*/
func (a *Client) DeleteEdgeJobsIDTasksTaskIDLogs(params *DeleteEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEdgeJobsIDTasksTaskIDLogsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEdgeJobsIDTasksTaskIDLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEdgeJobsIDTasksTaskIDLogs",
		Method:             "DELETE",
		PathPattern:        "/edge_jobs/{id}/tasks/{taskID}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEdgeJobsIDTasksTaskIDLogsReader{formats: a.formats},
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
	success, ok := result.(*DeleteEdgeJobsIDTasksTaskIDLogsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEdgeJobsIDTasksTaskIDLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeJobs fetches edge jobs list
*/
func (a *Client) GetEdgeJobs(params *GetEdgeJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeJobs",
		Method:             "GET",
		PathPattern:        "/edge_jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeJobsReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeJobsID inspects an edge job
*/
func (a *Client) GetEdgeJobsID(params *GetEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeJobsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeJobsID",
		Method:             "GET",
		PathPattern:        "/edge_jobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeJobsIDReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeJobsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeJobsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeJobsIDFile fetches a file of an edge job
*/
func (a *Client) GetEdgeJobsIDFile(params *GetEdgeJobsIDFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeJobsIDFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeJobsIDFile",
		Method:             "GET",
		PathPattern:        "/edge_jobs/{id}/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeJobsIDFileReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeJobsIDFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeJobsIDFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeJobsIDTasks fetches the list of tasks on an edge job
*/
func (a *Client) GetEdgeJobsIDTasks(params *GetEdgeJobsIDTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeJobsIDTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeJobsIDTasks",
		Method:             "GET",
		PathPattern:        "/edge_jobs/{id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeJobsIDTasksReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeJobsIDTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeJobsIDTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEdgeJobsIDTasksTaskIDLogs fetches the log for a specifc task on an edge job
*/
func (a *Client) GetEdgeJobsIDTasksTaskIDLogs(params *GetEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEdgeJobsIDTasksTaskIDLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeJobsIDTasksTaskIDLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEdgeJobsIDTasksTaskIDLogs",
		Method:             "GET",
		PathPattern:        "/edge_jobs/{id}/tasks/{taskID}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEdgeJobsIDTasksTaskIDLogsReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeJobsIDTasksTaskIDLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEdgeJobsIDTasksTaskIDLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostEdgeJobs creates an edge job
*/
func (a *Client) PostEdgeJobs(params *PostEdgeJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEdgeJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEdgeJobs",
		Method:             "POST",
		PathPattern:        "/edge_jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEdgeJobsReader{formats: a.formats},
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
	success, ok := result.(*PostEdgeJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEdgeJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostEdgeJobsID updates an edge job
*/
func (a *Client) PostEdgeJobsID(params *PostEdgeJobsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEdgeJobsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEdgeJobsID",
		Method:             "POST",
		PathPattern:        "/edge_jobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEdgeJobsIDReader{formats: a.formats},
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
	success, ok := result.(*PostEdgeJobsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEdgeJobsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostEdgeJobsIDTasksTaskIDLogs collects the log for a specifc task on an edge job
*/
func (a *Client) PostEdgeJobsIDTasksTaskIDLogs(params *PostEdgeJobsIDTasksTaskIDLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEdgeJobsIDTasksTaskIDLogsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEdgeJobsIDTasksTaskIDLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEdgeJobsIDTasksTaskIDLogs",
		Method:             "POST",
		PathPattern:        "/edge_jobs/{id}/tasks/{taskID}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostEdgeJobsIDTasksTaskIDLogsReader{formats: a.formats},
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
	success, ok := result.(*PostEdgeJobsIDTasksTaskIDLogsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEdgeJobsIDTasksTaskIDLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
