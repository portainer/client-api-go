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

// NewStackUpdateGitParams creates a new StackUpdateGitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackUpdateGitParams() *StackUpdateGitParams {
	return &StackUpdateGitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackUpdateGitParamsWithTimeout creates a new StackUpdateGitParams object
// with the ability to set a timeout on a request.
func NewStackUpdateGitParamsWithTimeout(timeout time.Duration) *StackUpdateGitParams {
	return &StackUpdateGitParams{
		timeout: timeout,
	}
}

// NewStackUpdateGitParamsWithContext creates a new StackUpdateGitParams object
// with the ability to set a context for a request.
func NewStackUpdateGitParamsWithContext(ctx context.Context) *StackUpdateGitParams {
	return &StackUpdateGitParams{
		Context: ctx,
	}
}

// NewStackUpdateGitParamsWithHTTPClient creates a new StackUpdateGitParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackUpdateGitParamsWithHTTPClient(client *http.Client) *StackUpdateGitParams {
	return &StackUpdateGitParams{
		HTTPClient: client,
	}
}

/*
StackUpdateGitParams contains all the parameters to send to the API endpoint

	for the stack update git operation.

	Typically these are written to a http.Request.
*/
type StackUpdateGitParams struct {

	/* Body.

	   Git configs for pull and redeploy a stack
	*/
	Body *models.StacksStackGitUpdatePayload

	/* EndpointID.

	   Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack.
	*/
	EndpointID *int64

	/* ID.

	   Stack identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack update git params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackUpdateGitParams) WithDefaults() *StackUpdateGitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack update git params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackUpdateGitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack update git params
func (o *StackUpdateGitParams) WithTimeout(timeout time.Duration) *StackUpdateGitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack update git params
func (o *StackUpdateGitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack update git params
func (o *StackUpdateGitParams) WithContext(ctx context.Context) *StackUpdateGitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack update git params
func (o *StackUpdateGitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack update git params
func (o *StackUpdateGitParams) WithHTTPClient(client *http.Client) *StackUpdateGitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack update git params
func (o *StackUpdateGitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack update git params
func (o *StackUpdateGitParams) WithBody(body *models.StacksStackGitUpdatePayload) *StackUpdateGitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack update git params
func (o *StackUpdateGitParams) SetBody(body *models.StacksStackGitUpdatePayload) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the stack update git params
func (o *StackUpdateGitParams) WithEndpointID(endpointID *int64) *StackUpdateGitParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack update git params
func (o *StackUpdateGitParams) SetEndpointID(endpointID *int64) {
	o.EndpointID = endpointID
}

// WithID adds the id to the stack update git params
func (o *StackUpdateGitParams) WithID(id int64) *StackUpdateGitParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stack update git params
func (o *StackUpdateGitParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StackUpdateGitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.EndpointID != nil {

		// query param endpointId
		var qrEndpointID int64

		if o.EndpointID != nil {
			qrEndpointID = *o.EndpointID
		}
		qEndpointID := swag.FormatInt64(qrEndpointID)
		if qEndpointID != "" {

			if err := r.SetQueryParam("endpointId", qEndpointID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
