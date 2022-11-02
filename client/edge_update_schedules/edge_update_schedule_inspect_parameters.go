// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

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

// NewEdgeUpdateScheduleInspectParams creates a new EdgeUpdateScheduleInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeUpdateScheduleInspectParams() *EdgeUpdateScheduleInspectParams {
	return &EdgeUpdateScheduleInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeUpdateScheduleInspectParamsWithTimeout creates a new EdgeUpdateScheduleInspectParams object
// with the ability to set a timeout on a request.
func NewEdgeUpdateScheduleInspectParamsWithTimeout(timeout time.Duration) *EdgeUpdateScheduleInspectParams {
	return &EdgeUpdateScheduleInspectParams{
		timeout: timeout,
	}
}

// NewEdgeUpdateScheduleInspectParamsWithContext creates a new EdgeUpdateScheduleInspectParams object
// with the ability to set a context for a request.
func NewEdgeUpdateScheduleInspectParamsWithContext(ctx context.Context) *EdgeUpdateScheduleInspectParams {
	return &EdgeUpdateScheduleInspectParams{
		Context: ctx,
	}
}

// NewEdgeUpdateScheduleInspectParamsWithHTTPClient creates a new EdgeUpdateScheduleInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeUpdateScheduleInspectParamsWithHTTPClient(client *http.Client) *EdgeUpdateScheduleInspectParams {
	return &EdgeUpdateScheduleInspectParams{
		HTTPClient: client,
	}
}

/*
EdgeUpdateScheduleInspectParams contains all the parameters to send to the API endpoint

	for the edge update schedule inspect operation.

	Typically these are written to a http.Request.
*/
type EdgeUpdateScheduleInspectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge update schedule inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeUpdateScheduleInspectParams) WithDefaults() *EdgeUpdateScheduleInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge update schedule inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeUpdateScheduleInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) WithTimeout(timeout time.Duration) *EdgeUpdateScheduleInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) WithContext(ctx context.Context) *EdgeUpdateScheduleInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) WithHTTPClient(client *http.Client) *EdgeUpdateScheduleInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge update schedule inspect params
func (o *EdgeUpdateScheduleInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeUpdateScheduleInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
