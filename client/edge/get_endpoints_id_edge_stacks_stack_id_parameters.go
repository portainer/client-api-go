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

// NewGetEndpointsIDEdgeStacksStackIDParams creates a new GetEndpointsIDEdgeStacksStackIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEndpointsIDEdgeStacksStackIDParams() *GetEndpointsIDEdgeStacksStackIDParams {
	return &GetEndpointsIDEdgeStacksStackIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointsIDEdgeStacksStackIDParamsWithTimeout creates a new GetEndpointsIDEdgeStacksStackIDParams object
// with the ability to set a timeout on a request.
func NewGetEndpointsIDEdgeStacksStackIDParamsWithTimeout(timeout time.Duration) *GetEndpointsIDEdgeStacksStackIDParams {
	return &GetEndpointsIDEdgeStacksStackIDParams{
		timeout: timeout,
	}
}

// NewGetEndpointsIDEdgeStacksStackIDParamsWithContext creates a new GetEndpointsIDEdgeStacksStackIDParams object
// with the ability to set a context for a request.
func NewGetEndpointsIDEdgeStacksStackIDParamsWithContext(ctx context.Context) *GetEndpointsIDEdgeStacksStackIDParams {
	return &GetEndpointsIDEdgeStacksStackIDParams{
		Context: ctx,
	}
}

// NewGetEndpointsIDEdgeStacksStackIDParamsWithHTTPClient creates a new GetEndpointsIDEdgeStacksStackIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEndpointsIDEdgeStacksStackIDParamsWithHTTPClient(client *http.Client) *GetEndpointsIDEdgeStacksStackIDParams {
	return &GetEndpointsIDEdgeStacksStackIDParams{
		HTTPClient: client,
	}
}

/*
GetEndpointsIDEdgeStacksStackIDParams contains all the parameters to send to the API endpoint

	for the get endpoints ID edge stacks stack ID operation.

	Typically these are written to a http.Request.
*/
type GetEndpointsIDEdgeStacksStackIDParams struct {

	/* ID.

	   Environment(Endpoint) Id
	*/
	ID string

	/* StackID.

	   EdgeStack Id
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get endpoints ID edge stacks stack ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithDefaults() *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get endpoints ID edge stacks stack ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithTimeout(timeout time.Duration) *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithContext(ctx context.Context) *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithHTTPClient(client *http.Client) *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithID(id string) *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetID(id string) {
	o.ID = id
}

// WithStackID adds the stackID to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) WithStackID(stackID string) *GetEndpointsIDEdgeStacksStackIDParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get endpoints ID edge stacks stack ID params
func (o *GetEndpointsIDEdgeStacksStackIDParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointsIDEdgeStacksStackIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param stackId
	if err := r.SetPathParam("stackId", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
