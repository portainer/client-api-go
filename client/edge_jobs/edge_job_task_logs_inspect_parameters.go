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

// NewEdgeJobTaskLogsInspectParams creates a new EdgeJobTaskLogsInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeJobTaskLogsInspectParams() *EdgeJobTaskLogsInspectParams {
	return &EdgeJobTaskLogsInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeJobTaskLogsInspectParamsWithTimeout creates a new EdgeJobTaskLogsInspectParams object
// with the ability to set a timeout on a request.
func NewEdgeJobTaskLogsInspectParamsWithTimeout(timeout time.Duration) *EdgeJobTaskLogsInspectParams {
	return &EdgeJobTaskLogsInspectParams{
		timeout: timeout,
	}
}

// NewEdgeJobTaskLogsInspectParamsWithContext creates a new EdgeJobTaskLogsInspectParams object
// with the ability to set a context for a request.
func NewEdgeJobTaskLogsInspectParamsWithContext(ctx context.Context) *EdgeJobTaskLogsInspectParams {
	return &EdgeJobTaskLogsInspectParams{
		Context: ctx,
	}
}

// NewEdgeJobTaskLogsInspectParamsWithHTTPClient creates a new EdgeJobTaskLogsInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeJobTaskLogsInspectParamsWithHTTPClient(client *http.Client) *EdgeJobTaskLogsInspectParams {
	return &EdgeJobTaskLogsInspectParams{
		HTTPClient: client,
	}
}

/*
EdgeJobTaskLogsInspectParams contains all the parameters to send to the API endpoint

	for the edge job task logs inspect operation.

	Typically these are written to a http.Request.
*/
type EdgeJobTaskLogsInspectParams struct {

	/* ID.

	   EdgeJob Id
	*/
	ID string

	/* TaskID.

	   Task Id
	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge job task logs inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobTaskLogsInspectParams) WithDefaults() *EdgeJobTaskLogsInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge job task logs inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeJobTaskLogsInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) WithTimeout(timeout time.Duration) *EdgeJobTaskLogsInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) WithContext(ctx context.Context) *EdgeJobTaskLogsInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) WithHTTPClient(client *http.Client) *EdgeJobTaskLogsInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) WithID(id string) *EdgeJobTaskLogsInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) SetID(id string) {
	o.ID = id
}

// WithTaskID adds the taskID to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) WithTaskID(taskID string) *EdgeJobTaskLogsInspectParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the edge job task logs inspect params
func (o *EdgeJobTaskLogsInspectParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeJobTaskLogsInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param taskID
	if err := r.SetPathParam("taskID", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
