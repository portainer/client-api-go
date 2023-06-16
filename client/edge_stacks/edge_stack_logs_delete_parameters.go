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
)

// NewEdgeStackLogsDeleteParams creates a new EdgeStackLogsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeStackLogsDeleteParams() *EdgeStackLogsDeleteParams {
	return &EdgeStackLogsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeStackLogsDeleteParamsWithTimeout creates a new EdgeStackLogsDeleteParams object
// with the ability to set a timeout on a request.
func NewEdgeStackLogsDeleteParamsWithTimeout(timeout time.Duration) *EdgeStackLogsDeleteParams {
	return &EdgeStackLogsDeleteParams{
		timeout: timeout,
	}
}

// NewEdgeStackLogsDeleteParamsWithContext creates a new EdgeStackLogsDeleteParams object
// with the ability to set a context for a request.
func NewEdgeStackLogsDeleteParamsWithContext(ctx context.Context) *EdgeStackLogsDeleteParams {
	return &EdgeStackLogsDeleteParams{
		Context: ctx,
	}
}

// NewEdgeStackLogsDeleteParamsWithHTTPClient creates a new EdgeStackLogsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeStackLogsDeleteParamsWithHTTPClient(client *http.Client) *EdgeStackLogsDeleteParams {
	return &EdgeStackLogsDeleteParams{
		HTTPClient: client,
	}
}

/*
EdgeStackLogsDeleteParams contains all the parameters to send to the API endpoint

	for the edge stack logs delete operation.

	Typically these are written to a http.Request.
*/
type EdgeStackLogsDeleteParams struct {

	/* EndpointID.

	   Endpoint Id
	*/
	EndpointID string

	/* ID.

	   EdgeStack Id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge stack logs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackLogsDeleteParams) WithDefaults() *EdgeStackLogsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge stack logs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackLogsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) WithTimeout(timeout time.Duration) *EdgeStackLogsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) WithContext(ctx context.Context) *EdgeStackLogsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) WithHTTPClient(client *http.Client) *EdgeStackLogsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointID adds the endpointID to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) WithEndpointID(endpointID string) *EdgeStackLogsDeleteParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) SetEndpointID(endpointID string) {
	o.EndpointID = endpointID
}

// WithID adds the id to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) WithID(id string) *EdgeStackLogsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge stack logs delete params
func (o *EdgeStackLogsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeStackLogsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpoint_id
	if err := r.SetPathParam("endpoint_id", o.EndpointID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
