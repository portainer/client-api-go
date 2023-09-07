// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewStackCreateDockerSwarmFileParams creates a new StackCreateDockerSwarmFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackCreateDockerSwarmFileParams() *StackCreateDockerSwarmFileParams {
	return &StackCreateDockerSwarmFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackCreateDockerSwarmFileParamsWithTimeout creates a new StackCreateDockerSwarmFileParams object
// with the ability to set a timeout on a request.
func NewStackCreateDockerSwarmFileParamsWithTimeout(timeout time.Duration) *StackCreateDockerSwarmFileParams {
	return &StackCreateDockerSwarmFileParams{
		timeout: timeout,
	}
}

// NewStackCreateDockerSwarmFileParamsWithContext creates a new StackCreateDockerSwarmFileParams object
// with the ability to set a context for a request.
func NewStackCreateDockerSwarmFileParamsWithContext(ctx context.Context) *StackCreateDockerSwarmFileParams {
	return &StackCreateDockerSwarmFileParams{
		Context: ctx,
	}
}

// NewStackCreateDockerSwarmFileParamsWithHTTPClient creates a new StackCreateDockerSwarmFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackCreateDockerSwarmFileParamsWithHTTPClient(client *http.Client) *StackCreateDockerSwarmFileParams {
	return &StackCreateDockerSwarmFileParams{
		HTTPClient: client,
	}
}

/*
StackCreateDockerSwarmFileParams contains all the parameters to send to the API endpoint

	for the stack create docker swarm file operation.

	Typically these are written to a http.Request.
*/
type StackCreateDockerSwarmFileParams struct {

	/* Env.

	   Environment variables passed during deployment, represented as a JSON array [{'name': 'name', 'value': 'value'}]. Optional
	*/
	Env *string

	/* Name.

	   Name of the stack
	*/
	Name *string

	/* SwarmID.

	   Swarm cluster identifier.
	*/
	SwarmID *string

	/* EndpointID.

	   Identifier of the environment that will be used to deploy the stack
	*/
	EndpointID int64

	/* File.

	   Stack file
	*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack create docker swarm file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerSwarmFileParams) WithDefaults() *StackCreateDockerSwarmFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack create docker swarm file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerSwarmFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithTimeout(timeout time.Duration) *StackCreateDockerSwarmFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithContext(ctx context.Context) *StackCreateDockerSwarmFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithHTTPClient(client *http.Client) *StackCreateDockerSwarmFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnv adds the env to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithEnv(env *string) *StackCreateDockerSwarmFileParams {
	o.SetEnv(env)
	return o
}

// SetEnv adds the env to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetEnv(env *string) {
	o.Env = env
}

// WithName adds the name to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithName(name *string) *StackCreateDockerSwarmFileParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetName(name *string) {
	o.Name = name
}

// WithSwarmID adds the swarmID to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithSwarmID(swarmID *string) *StackCreateDockerSwarmFileParams {
	o.SetSwarmID(swarmID)
	return o
}

// SetSwarmID adds the swarmId to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetSwarmID(swarmID *string) {
	o.SwarmID = swarmID
}

// WithEndpointID adds the endpointID to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithEndpointID(endpointID int64) *StackCreateDockerSwarmFileParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WithFile adds the file to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) WithFile(file runtime.NamedReadCloser) *StackCreateDockerSwarmFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the stack create docker swarm file params
func (o *StackCreateDockerSwarmFileParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *StackCreateDockerSwarmFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Env != nil {

		// form param Env
		var frEnv string
		if o.Env != nil {
			frEnv = *o.Env
		}
		fEnv := frEnv
		if fEnv != "" {
			if err := r.SetFormParam("Env", fEnv); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// form param Name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("Name", fName); err != nil {
				return err
			}
		}
	}

	if o.SwarmID != nil {

		// form param SwarmID
		var frSwarmID string
		if o.SwarmID != nil {
			frSwarmID = *o.SwarmID
		}
		fSwarmID := frSwarmID
		if fSwarmID != "" {
			if err := r.SetFormParam("SwarmID", fSwarmID); err != nil {
				return err
			}
		}
	}

	// query param endpointId
	qrEndpointID := o.EndpointID
	qEndpointID := swag.FormatInt64(qrEndpointID)
	if qEndpointID != "" {

		if err := r.SetQueryParam("endpointId", qEndpointID); err != nil {
			return err
		}
	}

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
