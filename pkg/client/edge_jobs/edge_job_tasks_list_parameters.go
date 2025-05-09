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

// NewEdgeJobTasksListParams creates a new EdgeJobTasksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeJobTasksListParams() *EdgeJobTasksListParams {
	return &EdgeJobTasksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeJobTasksListParamsWithTimeout creates a new EdgeJobTasksListParams object
// with the ability to set a timeout on a request.
func NewEdgeJobTasksListParamsWithTimeout(timeout time.Duration) *EdgeJobTasksListParams {
	return &EdgeJobTasksListParams{
		timeout: timeout,
	}
}

// NewEdgeJobTasksListParamsWithContext creates a new EdgeJobTasksListParams object
// with the ability to set a context for a request.
func NewEdgeJobTasksListParamsWithContext(ctx context.Context) *EdgeJobTasksListParams {
	return &EdgeJobTasksListParams{
		Context: ctx,
	}
}

// NewEdgeJobTasksListParamsWithHTTPClient creates a new EdgeJobTasksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeJobTasksListParamsWithHTTPClient(client *http.Client) *EdgeJobTasksListParams {
	return &EdgeJobTasksListParams{
		HTTPClient: client,
	}
}

/*
EdgeJobTasksListParams contains all the parameters to send to the API endpoint

	for the edge job tasks list operation.

	Typically these are written to a http.Request.
*/
type EdgeJobTasksListParams struct {

	/* ID.

	   EdgeJob Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge job tasks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobTasksListParams) WithDefaults() *EdgeJobTasksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge job tasks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobTasksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge job tasks list params
func (o *EdgeJobTasksListParams) WithTimeout(timeout time.Duration) *EdgeJobTasksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge job tasks list params
func (o *EdgeJobTasksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge job tasks list params
func (o *EdgeJobTasksListParams) WithContext(ctx context.Context) *EdgeJobTasksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge job tasks list params
func (o *EdgeJobTasksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge job tasks list params
func (o *EdgeJobTasksListParams) WithHTTPClient(client *http.Client) *EdgeJobTasksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge job tasks list params
func (o *EdgeJobTasksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge job tasks list params
func (o *EdgeJobTasksListParams) WithID(id int64) *EdgeJobTasksListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge job tasks list params
func (o *EdgeJobTasksListParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeJobTasksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
