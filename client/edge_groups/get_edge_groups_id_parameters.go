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

// NewGetEdgeGroupsIDParams creates a new GetEdgeGroupsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeGroupsIDParams() *GetEdgeGroupsIDParams {
	return &GetEdgeGroupsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeGroupsIDParamsWithTimeout creates a new GetEdgeGroupsIDParams object
// with the ability to set a timeout on a request.
func NewGetEdgeGroupsIDParamsWithTimeout(timeout time.Duration) *GetEdgeGroupsIDParams {
	return &GetEdgeGroupsIDParams{
		timeout: timeout,
	}
}

// NewGetEdgeGroupsIDParamsWithContext creates a new GetEdgeGroupsIDParams object
// with the ability to set a context for a request.
func NewGetEdgeGroupsIDParamsWithContext(ctx context.Context) *GetEdgeGroupsIDParams {
	return &GetEdgeGroupsIDParams{
		Context: ctx,
	}
}

// NewGetEdgeGroupsIDParamsWithHTTPClient creates a new GetEdgeGroupsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeGroupsIDParamsWithHTTPClient(client *http.Client) *GetEdgeGroupsIDParams {
	return &GetEdgeGroupsIDParams{
		HTTPClient: client,
	}
}

/* GetEdgeGroupsIDParams contains all the parameters to send to the API endpoint
   for the get edge groups ID operation.

   Typically these are written to a http.Request.
*/
type GetEdgeGroupsIDParams struct {

	/* ID.

	   EdgeGroup Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge groups ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeGroupsIDParams) WithDefaults() *GetEdgeGroupsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge groups ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeGroupsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) WithTimeout(timeout time.Duration) *GetEdgeGroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) WithContext(ctx context.Context) *GetEdgeGroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) WithHTTPClient(client *http.Client) *GetEdgeGroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) WithID(id int64) *GetEdgeGroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge groups ID params
func (o *GetEdgeGroupsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeGroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
