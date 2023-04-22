// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

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

// NewEdgeStackInspectParams creates a new EdgeStackInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeStackInspectParams() *EdgeStackInspectParams {
	return &EdgeStackInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeStackInspectParamsWithTimeout creates a new EdgeStackInspectParams object
// with the ability to set a timeout on a request.
func NewEdgeStackInspectParamsWithTimeout(timeout time.Duration) *EdgeStackInspectParams {
	return &EdgeStackInspectParams{
		timeout: timeout,
	}
}

// NewEdgeStackInspectParamsWithContext creates a new EdgeStackInspectParams object
// with the ability to set a context for a request.
func NewEdgeStackInspectParamsWithContext(ctx context.Context) *EdgeStackInspectParams {
	return &EdgeStackInspectParams{
		Context: ctx,
	}
}

// NewEdgeStackInspectParamsWithHTTPClient creates a new EdgeStackInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeStackInspectParamsWithHTTPClient(client *http.Client) *EdgeStackInspectParams {
	return &EdgeStackInspectParams{
		HTTPClient: client,
	}
}

/*
EdgeStackInspectParams contains all the parameters to send to the API endpoint

	for the edge stack inspect operation.

	Typically these are written to a http.Request.
*/
type EdgeStackInspectParams struct {

	/* ID.

	   EdgeStack Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge stack inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackInspectParams) WithDefaults() *EdgeStackInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge stack inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge stack inspect params
func (o *EdgeStackInspectParams) WithTimeout(timeout time.Duration) *EdgeStackInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge stack inspect params
func (o *EdgeStackInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge stack inspect params
func (o *EdgeStackInspectParams) WithContext(ctx context.Context) *EdgeStackInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge stack inspect params
func (o *EdgeStackInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge stack inspect params
func (o *EdgeStackInspectParams) WithHTTPClient(client *http.Client) *EdgeStackInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge stack inspect params
func (o *EdgeStackInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge stack inspect params
func (o *EdgeStackInspectParams) WithID(id int64) *EdgeStackInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge stack inspect params
func (o *EdgeStackInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeStackInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
