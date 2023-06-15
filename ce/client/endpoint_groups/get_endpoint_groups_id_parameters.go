// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

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

// NewGetEndpointGroupsIDParams creates a new GetEndpointGroupsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEndpointGroupsIDParams() *GetEndpointGroupsIDParams {
	return &GetEndpointGroupsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointGroupsIDParamsWithTimeout creates a new GetEndpointGroupsIDParams object
// with the ability to set a timeout on a request.
func NewGetEndpointGroupsIDParamsWithTimeout(timeout time.Duration) *GetEndpointGroupsIDParams {
	return &GetEndpointGroupsIDParams{
		timeout: timeout,
	}
}

// NewGetEndpointGroupsIDParamsWithContext creates a new GetEndpointGroupsIDParams object
// with the ability to set a context for a request.
func NewGetEndpointGroupsIDParamsWithContext(ctx context.Context) *GetEndpointGroupsIDParams {
	return &GetEndpointGroupsIDParams{
		Context: ctx,
	}
}

// NewGetEndpointGroupsIDParamsWithHTTPClient creates a new GetEndpointGroupsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEndpointGroupsIDParamsWithHTTPClient(client *http.Client) *GetEndpointGroupsIDParams {
	return &GetEndpointGroupsIDParams{
		HTTPClient: client,
	}
}

/*
GetEndpointGroupsIDParams contains all the parameters to send to the API endpoint

	for the get endpoint groups ID operation.

	Typically these are written to a http.Request.
*/
type GetEndpointGroupsIDParams struct {

	/* ID.

	   Environment(Endpoint) group identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get endpoint groups ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointGroupsIDParams) WithDefaults() *GetEndpointGroupsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get endpoint groups ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEndpointGroupsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) WithTimeout(timeout time.Duration) *GetEndpointGroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) WithContext(ctx context.Context) *GetEndpointGroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) WithHTTPClient(client *http.Client) *GetEndpointGroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) WithID(id int64) *GetEndpointGroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get endpoint groups ID params
func (o *GetEndpointGroupsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointGroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
