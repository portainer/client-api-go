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

// NewEndpointGroupDeleteEndpointParams creates a new EndpointGroupDeleteEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointGroupDeleteEndpointParams() *EndpointGroupDeleteEndpointParams {
	return &EndpointGroupDeleteEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointGroupDeleteEndpointParamsWithTimeout creates a new EndpointGroupDeleteEndpointParams object
// with the ability to set a timeout on a request.
func NewEndpointGroupDeleteEndpointParamsWithTimeout(timeout time.Duration) *EndpointGroupDeleteEndpointParams {
	return &EndpointGroupDeleteEndpointParams{
		timeout: timeout,
	}
}

// NewEndpointGroupDeleteEndpointParamsWithContext creates a new EndpointGroupDeleteEndpointParams object
// with the ability to set a context for a request.
func NewEndpointGroupDeleteEndpointParamsWithContext(ctx context.Context) *EndpointGroupDeleteEndpointParams {
	return &EndpointGroupDeleteEndpointParams{
		Context: ctx,
	}
}

// NewEndpointGroupDeleteEndpointParamsWithHTTPClient creates a new EndpointGroupDeleteEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointGroupDeleteEndpointParamsWithHTTPClient(client *http.Client) *EndpointGroupDeleteEndpointParams {
	return &EndpointGroupDeleteEndpointParams{
		HTTPClient: client,
	}
}

/* EndpointGroupDeleteEndpointParams contains all the parameters to send to the API endpoint
   for the endpoint group delete endpoint operation.

   Typically these are written to a http.Request.
*/
type EndpointGroupDeleteEndpointParams struct {

	/* EndpointID.

	   Environment(Endpoint) identifier
	*/
	EndpointID int64

	/* ID.

	   EndpointGroup identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint group delete endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupDeleteEndpointParams) WithDefaults() *EndpointGroupDeleteEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint group delete endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupDeleteEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) WithTimeout(timeout time.Duration) *EndpointGroupDeleteEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) WithContext(ctx context.Context) *EndpointGroupDeleteEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) WithHTTPClient(client *http.Client) *EndpointGroupDeleteEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointID adds the endpointID to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) WithEndpointID(endpointID int64) *EndpointGroupDeleteEndpointParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WithID adds the id to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) WithID(id int64) *EndpointGroupDeleteEndpointParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint group delete endpoint params
func (o *EndpointGroupDeleteEndpointParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointGroupDeleteEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointId
	if err := r.SetPathParam("endpointId", swag.FormatInt64(o.EndpointID)); err != nil {
		return err
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
