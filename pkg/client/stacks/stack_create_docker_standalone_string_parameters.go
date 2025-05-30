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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewStackCreateDockerStandaloneStringParams creates a new StackCreateDockerStandaloneStringParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackCreateDockerStandaloneStringParams() *StackCreateDockerStandaloneStringParams {
	return &StackCreateDockerStandaloneStringParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackCreateDockerStandaloneStringParamsWithTimeout creates a new StackCreateDockerStandaloneStringParams object
// with the ability to set a timeout on a request.
func NewStackCreateDockerStandaloneStringParamsWithTimeout(timeout time.Duration) *StackCreateDockerStandaloneStringParams {
	return &StackCreateDockerStandaloneStringParams{
		timeout: timeout,
	}
}

// NewStackCreateDockerStandaloneStringParamsWithContext creates a new StackCreateDockerStandaloneStringParams object
// with the ability to set a context for a request.
func NewStackCreateDockerStandaloneStringParamsWithContext(ctx context.Context) *StackCreateDockerStandaloneStringParams {
	return &StackCreateDockerStandaloneStringParams{
		Context: ctx,
	}
}

// NewStackCreateDockerStandaloneStringParamsWithHTTPClient creates a new StackCreateDockerStandaloneStringParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackCreateDockerStandaloneStringParamsWithHTTPClient(client *http.Client) *StackCreateDockerStandaloneStringParams {
	return &StackCreateDockerStandaloneStringParams{
		HTTPClient: client,
	}
}

/*
StackCreateDockerStandaloneStringParams contains all the parameters to send to the API endpoint

	for the stack create docker standalone string operation.

	Typically these are written to a http.Request.
*/
type StackCreateDockerStandaloneStringParams struct {

	/* Body.

	   stack config
	*/
	Body *models.StacksComposeStackFromFileContentPayload

	/* EndpointID.

	   Identifier of the environment that will be used to deploy the stack
	*/
	EndpointID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack create docker standalone string params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerStandaloneStringParams) WithDefaults() *StackCreateDockerStandaloneStringParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack create docker standalone string params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateDockerStandaloneStringParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) WithTimeout(timeout time.Duration) *StackCreateDockerStandaloneStringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) WithContext(ctx context.Context) *StackCreateDockerStandaloneStringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) WithHTTPClient(client *http.Client) *StackCreateDockerStandaloneStringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) WithBody(body *models.StacksComposeStackFromFileContentPayload) *StackCreateDockerStandaloneStringParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) SetBody(body *models.StacksComposeStackFromFileContentPayload) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) WithEndpointID(endpointID int64) *StackCreateDockerStandaloneStringParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack create docker standalone string params
func (o *StackCreateDockerStandaloneStringParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WriteToRequest writes these params to a swagger request
func (o *StackCreateDockerStandaloneStringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
