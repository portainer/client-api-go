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

// NewEndpointExtensionRemoveParams creates a new EndpointExtensionRemoveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointExtensionRemoveParams() *EndpointExtensionRemoveParams {
	return &EndpointExtensionRemoveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointExtensionRemoveParamsWithTimeout creates a new EndpointExtensionRemoveParams object
// with the ability to set a timeout on a request.
func NewEndpointExtensionRemoveParamsWithTimeout(timeout time.Duration) *EndpointExtensionRemoveParams {
	return &EndpointExtensionRemoveParams{
		timeout: timeout,
	}
}

// NewEndpointExtensionRemoveParamsWithContext creates a new EndpointExtensionRemoveParams object
// with the ability to set a context for a request.
func NewEndpointExtensionRemoveParamsWithContext(ctx context.Context) *EndpointExtensionRemoveParams {
	return &EndpointExtensionRemoveParams{
		Context: ctx,
	}
}

// NewEndpointExtensionRemoveParamsWithHTTPClient creates a new EndpointExtensionRemoveParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointExtensionRemoveParamsWithHTTPClient(client *http.Client) *EndpointExtensionRemoveParams {
	return &EndpointExtensionRemoveParams{
		HTTPClient: client,
	}
}

/* EndpointExtensionRemoveParams contains all the parameters to send to the API endpoint
   for the endpoint extension remove operation.

   Typically these are written to a http.Request.
*/
type EndpointExtensionRemoveParams struct {

	/* ExtensionType.

	   Extension Type
	*/
	ExtensionType string

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint extension remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointExtensionRemoveParams) WithDefaults() *EndpointExtensionRemoveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint extension remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointExtensionRemoveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) WithTimeout(timeout time.Duration) *EndpointExtensionRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) WithContext(ctx context.Context) *EndpointExtensionRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) WithHTTPClient(client *http.Client) *EndpointExtensionRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionType adds the extensionType to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) WithExtensionType(extensionType string) *EndpointExtensionRemoveParams {
	o.SetExtensionType(extensionType)
	return o
}

// SetExtensionType adds the extensionType to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) SetExtensionType(extensionType string) {
	o.ExtensionType = extensionType
}

// WithID adds the id to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) WithID(id int64) *EndpointExtensionRemoveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint extension remove params
func (o *EndpointExtensionRemoveParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointExtensionRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extensionType
	if err := r.SetPathParam("extensionType", o.ExtensionType); err != nil {
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
