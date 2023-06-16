// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

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

// NewEndpointGroupDeleteParams creates a new EndpointGroupDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointGroupDeleteParams() *EndpointGroupDeleteParams {
	return &EndpointGroupDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointGroupDeleteParamsWithTimeout creates a new EndpointGroupDeleteParams object
// with the ability to set a timeout on a request.
func NewEndpointGroupDeleteParamsWithTimeout(timeout time.Duration) *EndpointGroupDeleteParams {
	return &EndpointGroupDeleteParams{
		timeout: timeout,
	}
}

// NewEndpointGroupDeleteParamsWithContext creates a new EndpointGroupDeleteParams object
// with the ability to set a context for a request.
func NewEndpointGroupDeleteParamsWithContext(ctx context.Context) *EndpointGroupDeleteParams {
	return &EndpointGroupDeleteParams{
		Context: ctx,
	}
}

// NewEndpointGroupDeleteParamsWithHTTPClient creates a new EndpointGroupDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointGroupDeleteParamsWithHTTPClient(client *http.Client) *EndpointGroupDeleteParams {
	return &EndpointGroupDeleteParams{
		HTTPClient: client,
	}
}

/*
EndpointGroupDeleteParams contains all the parameters to send to the API endpoint

	for the endpoint group delete operation.

	Typically these are written to a http.Request.
*/
type EndpointGroupDeleteParams struct {

	/* ID.

	   EndpointGroup identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupDeleteParams) WithDefaults() *EndpointGroupDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint group delete params
func (o *EndpointGroupDeleteParams) WithTimeout(timeout time.Duration) *EndpointGroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint group delete params
func (o *EndpointGroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint group delete params
func (o *EndpointGroupDeleteParams) WithContext(ctx context.Context) *EndpointGroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint group delete params
func (o *EndpointGroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint group delete params
func (o *EndpointGroupDeleteParams) WithHTTPClient(client *http.Client) *EndpointGroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint group delete params
func (o *EndpointGroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the endpoint group delete params
func (o *EndpointGroupDeleteParams) WithID(id int64) *EndpointGroupDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint group delete params
func (o *EndpointGroupDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointGroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
