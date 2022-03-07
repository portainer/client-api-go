// Code generated by go-swagger; DO NOT EDIT.

package dockerhub

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

	"github.com/portainer/client-api/models"
)

// NewDockerHubUpdateParams creates a new DockerHubUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDockerHubUpdateParams() *DockerHubUpdateParams {
	return &DockerHubUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDockerHubUpdateParamsWithTimeout creates a new DockerHubUpdateParams object
// with the ability to set a timeout on a request.
func NewDockerHubUpdateParamsWithTimeout(timeout time.Duration) *DockerHubUpdateParams {
	return &DockerHubUpdateParams{
		timeout: timeout,
	}
}

// NewDockerHubUpdateParamsWithContext creates a new DockerHubUpdateParams object
// with the ability to set a context for a request.
func NewDockerHubUpdateParamsWithContext(ctx context.Context) *DockerHubUpdateParams {
	return &DockerHubUpdateParams{
		Context: ctx,
	}
}

// NewDockerHubUpdateParamsWithHTTPClient creates a new DockerHubUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDockerHubUpdateParamsWithHTTPClient(client *http.Client) *DockerHubUpdateParams {
	return &DockerHubUpdateParams{
		HTTPClient: client,
	}
}

/* DockerHubUpdateParams contains all the parameters to send to the API endpoint
   for the docker hub update operation.

   Typically these are written to a http.Request.
*/
type DockerHubUpdateParams struct {

	/* Body.

	   DockerHub information
	*/
	Body *models.DockerhubDockerhubUpdatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the docker hub update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DockerHubUpdateParams) WithDefaults() *DockerHubUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the docker hub update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DockerHubUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the docker hub update params
func (o *DockerHubUpdateParams) WithTimeout(timeout time.Duration) *DockerHubUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the docker hub update params
func (o *DockerHubUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the docker hub update params
func (o *DockerHubUpdateParams) WithContext(ctx context.Context) *DockerHubUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the docker hub update params
func (o *DockerHubUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the docker hub update params
func (o *DockerHubUpdateParams) WithHTTPClient(client *http.Client) *DockerHubUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the docker hub update params
func (o *DockerHubUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the docker hub update params
func (o *DockerHubUpdateParams) WithBody(body *models.DockerhubDockerhubUpdatePayload) *DockerHubUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the docker hub update params
func (o *DockerHubUpdateParams) SetBody(body *models.DockerhubDockerhubUpdatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DockerHubUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
