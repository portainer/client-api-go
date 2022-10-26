// Code generated by go-swagger; DO NOT EDIT.

package docker

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

// NewDockerContainerGpusInspectParams creates a new DockerContainerGpusInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDockerContainerGpusInspectParams() *DockerContainerGpusInspectParams {
	return &DockerContainerGpusInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDockerContainerGpusInspectParamsWithTimeout creates a new DockerContainerGpusInspectParams object
// with the ability to set a timeout on a request.
func NewDockerContainerGpusInspectParamsWithTimeout(timeout time.Duration) *DockerContainerGpusInspectParams {
	return &DockerContainerGpusInspectParams{
		timeout: timeout,
	}
}

// NewDockerContainerGpusInspectParamsWithContext creates a new DockerContainerGpusInspectParams object
// with the ability to set a context for a request.
func NewDockerContainerGpusInspectParamsWithContext(ctx context.Context) *DockerContainerGpusInspectParams {
	return &DockerContainerGpusInspectParams{
		Context: ctx,
	}
}

// NewDockerContainerGpusInspectParamsWithHTTPClient creates a new DockerContainerGpusInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDockerContainerGpusInspectParamsWithHTTPClient(client *http.Client) *DockerContainerGpusInspectParams {
	return &DockerContainerGpusInspectParams{
		HTTPClient: client,
	}
}

/*
DockerContainerGpusInspectParams contains all the parameters to send to the API endpoint

	for the docker container gpus inspect operation.

	Typically these are written to a http.Request.
*/
type DockerContainerGpusInspectParams struct {

	/* ContainerID.

	   Container identifier
	*/
	ContainerID int64

	/* EnvironmentID.

	   Environment identifier
	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the docker container gpus inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DockerContainerGpusInspectParams) WithDefaults() *DockerContainerGpusInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the docker container gpus inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DockerContainerGpusInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) WithTimeout(timeout time.Duration) *DockerContainerGpusInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) WithContext(ctx context.Context) *DockerContainerGpusInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) WithHTTPClient(client *http.Client) *DockerContainerGpusInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainerID adds the containerID to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) WithContainerID(containerID int64) *DockerContainerGpusInspectParams {
	o.SetContainerID(containerID)
	return o
}

// SetContainerID adds the containerId to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) SetContainerID(containerID int64) {
	o.ContainerID = containerID
}

// WithEnvironmentID adds the environmentID to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) WithEnvironmentID(environmentID int64) *DockerContainerGpusInspectParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the docker container gpus inspect params
func (o *DockerContainerGpusInspectParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DockerContainerGpusInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param containerId
	if err := r.SetPathParam("containerId", swag.FormatInt64(o.ContainerID)); err != nil {
		return err
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
