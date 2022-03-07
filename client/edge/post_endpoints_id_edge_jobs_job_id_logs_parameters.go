// Code generated by go-swagger; DO NOT EDIT.

package edge

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

// NewPostEndpointsIDEdgeJobsJobIDLogsParams creates a new PostEndpointsIDEdgeJobsJobIDLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEndpointsIDEdgeJobsJobIDLogsParams() *PostEndpointsIDEdgeJobsJobIDLogsParams {
	return &PostEndpointsIDEdgeJobsJobIDLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithTimeout creates a new PostEndpointsIDEdgeJobsJobIDLogsParams object
// with the ability to set a timeout on a request.
func NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithTimeout(timeout time.Duration) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	return &PostEndpointsIDEdgeJobsJobIDLogsParams{
		timeout: timeout,
	}
}

// NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithContext creates a new PostEndpointsIDEdgeJobsJobIDLogsParams object
// with the ability to set a context for a request.
func NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithContext(ctx context.Context) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	return &PostEndpointsIDEdgeJobsJobIDLogsParams{
		Context: ctx,
	}
}

// NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithHTTPClient creates a new PostEndpointsIDEdgeJobsJobIDLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEndpointsIDEdgeJobsJobIDLogsParamsWithHTTPClient(client *http.Client) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	return &PostEndpointsIDEdgeJobsJobIDLogsParams{
		HTTPClient: client,
	}
}

/* PostEndpointsIDEdgeJobsJobIDLogsParams contains all the parameters to send to the API endpoint
   for the post endpoints ID edge jobs job ID logs operation.

   Typically these are written to a http.Request.
*/
type PostEndpointsIDEdgeJobsJobIDLogsParams struct {

	/* ID.

	   Endpoint Id
	*/
	ID string

	/* JobID.

	   Job Id
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post endpoints ID edge jobs job ID logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithDefaults() *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post endpoints ID edge jobs job ID logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithTimeout(timeout time.Duration) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithContext(ctx context.Context) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithHTTPClient(client *http.Client) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithID(id string) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetID(id string) {
	o.ID = id
}

// WithJobID adds the jobID to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WithJobID(jobID string) *PostEndpointsIDEdgeJobsJobIDLogsParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the post endpoints ID edge jobs job ID logs params
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *PostEndpointsIDEdgeJobsJobIDLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param jobID
	if err := r.SetPathParam("jobID", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
