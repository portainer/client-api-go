// Code generated by go-swagger; DO NOT EDIT.

package edge_configs

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

// NewEdgeConfigStateParams creates a new EdgeConfigStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeConfigStateParams() *EdgeConfigStateParams {
	return &EdgeConfigStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeConfigStateParamsWithTimeout creates a new EdgeConfigStateParams object
// with the ability to set a timeout on a request.
func NewEdgeConfigStateParamsWithTimeout(timeout time.Duration) *EdgeConfigStateParams {
	return &EdgeConfigStateParams{
		timeout: timeout,
	}
}

// NewEdgeConfigStateParamsWithContext creates a new EdgeConfigStateParams object
// with the ability to set a context for a request.
func NewEdgeConfigStateParamsWithContext(ctx context.Context) *EdgeConfigStateParams {
	return &EdgeConfigStateParams{
		Context: ctx,
	}
}

// NewEdgeConfigStateParamsWithHTTPClient creates a new EdgeConfigStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeConfigStateParamsWithHTTPClient(client *http.Client) *EdgeConfigStateParams {
	return &EdgeConfigStateParams{
		HTTPClient: client,
	}
}

/*
EdgeConfigStateParams contains all the parameters to send to the API endpoint

	for the edge config state operation.

	Typically these are written to a http.Request.
*/
type EdgeConfigStateParams struct {

	/* ID.

	   Edge configuration identifier
	*/
	ID int64

	/* State.

	   Edge configuration state
	*/
	State int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge config state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeConfigStateParams) WithDefaults() *EdgeConfigStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge config state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeConfigStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge config state params
func (o *EdgeConfigStateParams) WithTimeout(timeout time.Duration) *EdgeConfigStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge config state params
func (o *EdgeConfigStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge config state params
func (o *EdgeConfigStateParams) WithContext(ctx context.Context) *EdgeConfigStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge config state params
func (o *EdgeConfigStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge config state params
func (o *EdgeConfigStateParams) WithHTTPClient(client *http.Client) *EdgeConfigStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge config state params
func (o *EdgeConfigStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the edge config state params
func (o *EdgeConfigStateParams) WithID(id int64) *EdgeConfigStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge config state params
func (o *EdgeConfigStateParams) SetID(id int64) {
	o.ID = id
}

// WithState adds the state to the edge config state params
func (o *EdgeConfigStateParams) WithState(state int64) *EdgeConfigStateParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the edge config state params
func (o *EdgeConfigStateParams) SetState(state int64) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeConfigStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param state
	if err := r.SetPathParam("state", swag.FormatInt64(o.State)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
