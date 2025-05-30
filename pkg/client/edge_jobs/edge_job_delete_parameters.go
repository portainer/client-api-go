// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

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

// NewEdgeJobDeleteParams creates a new EdgeJobDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeJobDeleteParams() *EdgeJobDeleteParams {
	return &EdgeJobDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeJobDeleteParamsWithTimeout creates a new EdgeJobDeleteParams object
// with the ability to set a timeout on a request.
func NewEdgeJobDeleteParamsWithTimeout(timeout time.Duration) *EdgeJobDeleteParams {
	return &EdgeJobDeleteParams{
		timeout: timeout,
	}
}

// NewEdgeJobDeleteParamsWithContext creates a new EdgeJobDeleteParams object
// with the ability to set a context for a request.
func NewEdgeJobDeleteParamsWithContext(ctx context.Context) *EdgeJobDeleteParams {
	return &EdgeJobDeleteParams{
		Context: ctx,
	}
}

// NewEdgeJobDeleteParamsWithHTTPClient creates a new EdgeJobDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeJobDeleteParamsWithHTTPClient(client *http.Client) *EdgeJobDeleteParams {
	return &EdgeJobDeleteParams{
		HTTPClient: client,
	}
}

/*
EdgeJobDeleteParams contains all the parameters to send to the API endpoint

	for the edge job delete operation.

	Typically these are written to a http.Request.
*/
type EdgeJobDeleteParams struct {

	/* ID.

	   EdgeJob Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge job delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobDeleteParams) WithDefaults() *EdgeJobDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge job delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge job delete params
func (o *EdgeJobDeleteParams) WithTimeout(timeout time.Duration) *EdgeJobDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge job delete params
func (o *EdgeJobDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge job delete params
func (o *EdgeJobDeleteParams) WithContext(ctx context.Context) *EdgeJobDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge job delete params
func (o *EdgeJobDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge job delete params
func (o *EdgeJobDeleteParams) WithHTTPClient(client *http.Client) *EdgeJobDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge job delete params
func (o *EdgeJobDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge job delete params
func (o *EdgeJobDeleteParams) WithID(id int64) *EdgeJobDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge job delete params
func (o *EdgeJobDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeJobDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
