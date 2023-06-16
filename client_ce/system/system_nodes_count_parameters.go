// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewSystemNodesCountParams creates a new SystemNodesCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemNodesCountParams() *SystemNodesCountParams {
	return &SystemNodesCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemNodesCountParamsWithTimeout creates a new SystemNodesCountParams object
// with the ability to set a timeout on a request.
func NewSystemNodesCountParamsWithTimeout(timeout time.Duration) *SystemNodesCountParams {
	return &SystemNodesCountParams{
		timeout: timeout,
	}
}

// NewSystemNodesCountParamsWithContext creates a new SystemNodesCountParams object
// with the ability to set a context for a request.
func NewSystemNodesCountParamsWithContext(ctx context.Context) *SystemNodesCountParams {
	return &SystemNodesCountParams{
		Context: ctx,
	}
}

// NewSystemNodesCountParamsWithHTTPClient creates a new SystemNodesCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemNodesCountParamsWithHTTPClient(client *http.Client) *SystemNodesCountParams {
	return &SystemNodesCountParams{
		HTTPClient: client,
	}
}

/*
SystemNodesCountParams contains all the parameters to send to the API endpoint

	for the system nodes count operation.

	Typically these are written to a http.Request.
*/
type SystemNodesCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system nodes count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemNodesCountParams) WithDefaults() *SystemNodesCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system nodes count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemNodesCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system nodes count params
func (o *SystemNodesCountParams) WithTimeout(timeout time.Duration) *SystemNodesCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system nodes count params
func (o *SystemNodesCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system nodes count params
func (o *SystemNodesCountParams) WithContext(ctx context.Context) *SystemNodesCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system nodes count params
func (o *SystemNodesCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system nodes count params
func (o *SystemNodesCountParams) WithHTTPClient(client *http.Client) *SystemNodesCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system nodes count params
func (o *SystemNodesCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemNodesCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
