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

// NewEdgeGroupDeleteParams creates a new EdgeGroupDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeGroupDeleteParams() *EdgeGroupDeleteParams {
	return &EdgeGroupDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeGroupDeleteParamsWithTimeout creates a new EdgeGroupDeleteParams object
// with the ability to set a timeout on a request.
func NewEdgeGroupDeleteParamsWithTimeout(timeout time.Duration) *EdgeGroupDeleteParams {
	return &EdgeGroupDeleteParams{
		timeout: timeout,
	}
}

// NewEdgeGroupDeleteParamsWithContext creates a new EdgeGroupDeleteParams object
// with the ability to set a context for a request.
func NewEdgeGroupDeleteParamsWithContext(ctx context.Context) *EdgeGroupDeleteParams {
	return &EdgeGroupDeleteParams{
		Context: ctx,
	}
}

// NewEdgeGroupDeleteParamsWithHTTPClient creates a new EdgeGroupDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeGroupDeleteParamsWithHTTPClient(client *http.Client) *EdgeGroupDeleteParams {
	return &EdgeGroupDeleteParams{
		HTTPClient: client,
	}
}

/*
EdgeGroupDeleteParams contains all the parameters to send to the API endpoint

	for the edge group delete operation.

	Typically these are written to a http.Request.
*/
type EdgeGroupDeleteParams struct {

	/* ID.

	   EdgeGroup Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeGroupDeleteParams) WithDefaults() *EdgeGroupDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeGroupDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge group delete params
func (o *EdgeGroupDeleteParams) WithTimeout(timeout time.Duration) *EdgeGroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge group delete params
func (o *EdgeGroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge group delete params
func (o *EdgeGroupDeleteParams) WithContext(ctx context.Context) *EdgeGroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge group delete params
func (o *EdgeGroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge group delete params
func (o *EdgeGroupDeleteParams) WithHTTPClient(client *http.Client) *EdgeGroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge group delete params
func (o *EdgeGroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge group delete params
func (o *EdgeGroupDeleteParams) WithID(id int64) *EdgeGroupDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge group delete params
func (o *EdgeGroupDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeGroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
