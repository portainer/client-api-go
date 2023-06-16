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

	"github.com/portainer/client-api-go/v2/models_ce"
)

// NewStackMigrateParams creates a new StackMigrateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackMigrateParams() *StackMigrateParams {
	return &StackMigrateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackMigrateParamsWithTimeout creates a new StackMigrateParams object
// with the ability to set a timeout on a request.
func NewStackMigrateParamsWithTimeout(timeout time.Duration) *StackMigrateParams {
	return &StackMigrateParams{
		timeout: timeout,
	}
}

// NewStackMigrateParamsWithContext creates a new StackMigrateParams object
// with the ability to set a context for a request.
func NewStackMigrateParamsWithContext(ctx context.Context) *StackMigrateParams {
	return &StackMigrateParams{
		Context: ctx,
	}
}

// NewStackMigrateParamsWithHTTPClient creates a new StackMigrateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackMigrateParamsWithHTTPClient(client *http.Client) *StackMigrateParams {
	return &StackMigrateParams{
		HTTPClient: client,
	}
}

/*
StackMigrateParams contains all the parameters to send to the API endpoint

	for the stack migrate operation.

	Typically these are written to a http.Request.
*/
type StackMigrateParams struct {

	/* Body.

	   Stack migration details
	*/
	Body *models_ce.StacksStackMigratePayload

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

// WithDefaults hydrates default values in the stack migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackMigrateParams) WithDefaults() *StackMigrateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackMigrateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack migrate params
func (o *StackMigrateParams) WithTimeout(timeout time.Duration) *StackMigrateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack migrate params
func (o *StackMigrateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack migrate params
func (o *StackMigrateParams) WithContext(ctx context.Context) *StackMigrateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack migrate params
func (o *StackMigrateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack migrate params
func (o *StackMigrateParams) WithHTTPClient(client *http.Client) *StackMigrateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack migrate params
func (o *StackMigrateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack migrate params
func (o *StackMigrateParams) WithBody(body *models_ce.StacksStackMigratePayload) *StackMigrateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack migrate params
func (o *StackMigrateParams) SetBody(body *models_ce.StacksStackMigratePayload) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the stack migrate params
func (o *StackMigrateParams) WithEndpointID(endpointID *int64) *StackMigrateParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack migrate params
func (o *StackMigrateParams) SetEndpointID(endpointID *int64) {
	o.EndpointID = endpointID
}

// WithID adds the id to the stack migrate params
func (o *StackMigrateParams) WithID(id int64) *StackMigrateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stack migrate params
func (o *StackMigrateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StackMigrateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
