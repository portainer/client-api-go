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

	"github.com/portainer/client-api-go/v2/models"
)

// NewStackCreateDockerSwarmRepositoryParams creates a new StackCreateDockerSwarmRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackCreateDockerSwarmRepositoryParams() *StackCreateDockerSwarmRepositoryParams {
	return &StackCreateDockerSwarmRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackCreateDockerSwarmRepositoryParamsWithTimeout creates a new StackCreateDockerSwarmRepositoryParams object
// with the ability to set a timeout on a request.
func NewStackCreateDockerSwarmRepositoryParamsWithTimeout(timeout time.Duration) *StackCreateDockerSwarmRepositoryParams {
	return &StackCreateDockerSwarmRepositoryParams{
		timeout: timeout,
	}
}

// NewStackCreateDockerSwarmRepositoryParamsWithContext creates a new StackCreateDockerSwarmRepositoryParams object
// with the ability to set a context for a request.
func NewStackCreateDockerSwarmRepositoryParamsWithContext(ctx context.Context) *StackCreateDockerSwarmRepositoryParams {
	return &StackCreateDockerSwarmRepositoryParams{
		Context: ctx,
	}
}

// NewStackCreateDockerSwarmRepositoryParamsWithHTTPClient creates a new StackCreateDockerSwarmRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackCreateDockerSwarmRepositoryParamsWithHTTPClient(client *http.Client) *StackCreateDockerSwarmRepositoryParams {
	return &StackCreateDockerSwarmRepositoryParams{
		HTTPClient: client,
	}
}

/*
StackCreateDockerSwarmRepositoryParams contains all the parameters to send to the API endpoint

	for the stack create docker swarm repository operation.

	Typically these are written to a http.Request.
*/
type StackCreateDockerSwarmRepositoryParams struct {

	/* Body.

	   stack config
	*/
	Body *models.StacksSwarmStackFromGitRepositoryPayload

	/* EndpointID.

	   Identifier of the environment that will be used to deploy the stack
	*/
	EndpointID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack create docker swarm repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerSwarmRepositoryParams) WithDefaults() *StackCreateDockerSwarmRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack create docker swarm repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerSwarmRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) WithTimeout(timeout time.Duration) *StackCreateDockerSwarmRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) WithContext(ctx context.Context) *StackCreateDockerSwarmRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) WithHTTPClient(client *http.Client) *StackCreateDockerSwarmRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) WithBody(body *models.StacksSwarmStackFromGitRepositoryPayload) *StackCreateDockerSwarmRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) SetBody(body *models.StacksSwarmStackFromGitRepositoryPayload) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) WithEndpointID(endpointID int64) *StackCreateDockerSwarmRepositoryParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack create docker swarm repository params
func (o *StackCreateDockerSwarmRepositoryParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WriteToRequest writes these params to a swagger request
func (o *StackCreateDockerSwarmRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
