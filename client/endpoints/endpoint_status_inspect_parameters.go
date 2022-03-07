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

// NewEndpointStatusInspectParams creates a new EndpointStatusInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointStatusInspectParams() *EndpointStatusInspectParams {
	return &EndpointStatusInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointStatusInspectParamsWithTimeout creates a new EndpointStatusInspectParams object
// with the ability to set a timeout on a request.
func NewEndpointStatusInspectParamsWithTimeout(timeout time.Duration) *EndpointStatusInspectParams {
	return &EndpointStatusInspectParams{
		timeout: timeout,
	}
}

// NewEndpointStatusInspectParamsWithContext creates a new EndpointStatusInspectParams object
// with the ability to set a context for a request.
func NewEndpointStatusInspectParamsWithContext(ctx context.Context) *EndpointStatusInspectParams {
	return &EndpointStatusInspectParams{
		Context: ctx,
	}
}

// NewEndpointStatusInspectParamsWithHTTPClient creates a new EndpointStatusInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointStatusInspectParamsWithHTTPClient(client *http.Client) *EndpointStatusInspectParams {
	return &EndpointStatusInspectParams{
		HTTPClient: client,
	}
}

/* EndpointStatusInspectParams contains all the parameters to send to the API endpoint
   for the endpoint status inspect operation.

   Typically these are written to a http.Request.
*/
type EndpointStatusInspectParams struct {

	/* ID.

	   Endpoint identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint status inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointStatusInspectParams) WithDefaults() *EndpointStatusInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint status inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointStatusInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint status inspect params
func (o *EndpointStatusInspectParams) WithTimeout(timeout time.Duration) *EndpointStatusInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint status inspect params
func (o *EndpointStatusInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint status inspect params
func (o *EndpointStatusInspectParams) WithContext(ctx context.Context) *EndpointStatusInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint status inspect params
func (o *EndpointStatusInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint status inspect params
func (o *EndpointStatusInspectParams) WithHTTPClient(client *http.Client) *EndpointStatusInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint status inspect params
func (o *EndpointStatusInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the endpoint status inspect params
func (o *EndpointStatusInspectParams) WithID(id int64) *EndpointStatusInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint status inspect params
func (o *EndpointStatusInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointStatusInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
