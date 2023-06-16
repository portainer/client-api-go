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
)

// NewEdgeJobListParams creates a new EdgeJobListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeJobListParams() *EdgeJobListParams {
	return &EdgeJobListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeJobListParamsWithTimeout creates a new EdgeJobListParams object
// with the ability to set a timeout on a request.
func NewEdgeJobListParamsWithTimeout(timeout time.Duration) *EdgeJobListParams {
	return &EdgeJobListParams{
		timeout: timeout,
	}
}

// NewEdgeJobListParamsWithContext creates a new EdgeJobListParams object
// with the ability to set a context for a request.
func NewEdgeJobListParamsWithContext(ctx context.Context) *EdgeJobListParams {
	return &EdgeJobListParams{
		Context: ctx,
	}
}

// NewEdgeJobListParamsWithHTTPClient creates a new EdgeJobListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeJobListParamsWithHTTPClient(client *http.Client) *EdgeJobListParams {
	return &EdgeJobListParams{
		HTTPClient: client,
	}
}

/*
EdgeJobListParams contains all the parameters to send to the API endpoint

	for the edge job list operation.

	Typically these are written to a http.Request.
*/
type EdgeJobListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge job list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobListParams) WithDefaults() *EdgeJobListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge job list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge job list params
func (o *EdgeJobListParams) WithTimeout(timeout time.Duration) *EdgeJobListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge job list params
func (o *EdgeJobListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge job list params
func (o *EdgeJobListParams) WithContext(ctx context.Context) *EdgeJobListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge job list params
func (o *EdgeJobListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge job list params
func (o *EdgeJobListParams) WithHTTPClient(client *http.Client) *EdgeJobListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge job list params
func (o *EdgeJobListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeJobListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
