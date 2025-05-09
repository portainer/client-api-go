// Code generated by go-swagger; DO NOT EDIT.

package edge_groups

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

// NewEdgeGroupInspectParams creates a new EdgeGroupInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeGroupInspectParams() *EdgeGroupInspectParams {
	return &EdgeGroupInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeGroupInspectParamsWithTimeout creates a new EdgeGroupInspectParams object
// with the ability to set a timeout on a request.
func NewEdgeGroupInspectParamsWithTimeout(timeout time.Duration) *EdgeGroupInspectParams {
	return &EdgeGroupInspectParams{
		timeout: timeout,
	}
}

// NewEdgeGroupInspectParamsWithContext creates a new EdgeGroupInspectParams object
// with the ability to set a context for a request.
func NewEdgeGroupInspectParamsWithContext(ctx context.Context) *EdgeGroupInspectParams {
	return &EdgeGroupInspectParams{
		Context: ctx,
	}
}

// NewEdgeGroupInspectParamsWithHTTPClient creates a new EdgeGroupInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeGroupInspectParamsWithHTTPClient(client *http.Client) *EdgeGroupInspectParams {
	return &EdgeGroupInspectParams{
		HTTPClient: client,
	}
}

/*
EdgeGroupInspectParams contains all the parameters to send to the API endpoint

	for the edge group inspect operation.

	Typically these are written to a http.Request.
*/
type EdgeGroupInspectParams struct {

	/* ID.

	   EdgeGroup Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge group inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeGroupInspectParams) WithDefaults() *EdgeGroupInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge group inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeGroupInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge group inspect params
func (o *EdgeGroupInspectParams) WithTimeout(timeout time.Duration) *EdgeGroupInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge group inspect params
func (o *EdgeGroupInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge group inspect params
func (o *EdgeGroupInspectParams) WithContext(ctx context.Context) *EdgeGroupInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge group inspect params
func (o *EdgeGroupInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge group inspect params
func (o *EdgeGroupInspectParams) WithHTTPClient(client *http.Client) *EdgeGroupInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge group inspect params
func (o *EdgeGroupInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge group inspect params
func (o *EdgeGroupInspectParams) WithID(id int64) *EdgeGroupInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge group inspect params
func (o *EdgeGroupInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeGroupInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
