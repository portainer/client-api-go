// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewEndpointInspectParams creates a new EndpointInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointInspectParams() *EndpointInspectParams {
	return &EndpointInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointInspectParamsWithTimeout creates a new EndpointInspectParams object
// with the ability to set a timeout on a request.
func NewEndpointInspectParamsWithTimeout(timeout time.Duration) *EndpointInspectParams {
	return &EndpointInspectParams{
		timeout: timeout,
	}
}

// NewEndpointInspectParamsWithContext creates a new EndpointInspectParams object
// with the ability to set a context for a request.
func NewEndpointInspectParamsWithContext(ctx context.Context) *EndpointInspectParams {
	return &EndpointInspectParams{
		Context: ctx,
	}
}

// NewEndpointInspectParamsWithHTTPClient creates a new EndpointInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointInspectParamsWithHTTPClient(client *http.Client) *EndpointInspectParams {
	return &EndpointInspectParams{
		HTTPClient: client,
	}
}

/*
EndpointInspectParams contains all the parameters to send to the API endpoint

	for the endpoint inspect operation.

	Typically these are written to a http.Request.
*/
type EndpointInspectParams struct {

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointInspectParams) WithDefaults() *EndpointInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint inspect params
func (o *EndpointInspectParams) WithTimeout(timeout time.Duration) *EndpointInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint inspect params
func (o *EndpointInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint inspect params
func (o *EndpointInspectParams) WithContext(ctx context.Context) *EndpointInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint inspect params
func (o *EndpointInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint inspect params
func (o *EndpointInspectParams) WithHTTPClient(client *http.Client) *EndpointInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint inspect params
func (o *EndpointInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the endpoint inspect params
func (o *EndpointInspectParams) WithID(id int64) *EndpointInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint inspect params
func (o *EndpointInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
