// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewGetKubernetesMetricsForNodeParams creates a new GetKubernetesMetricsForNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesMetricsForNodeParams() *GetKubernetesMetricsForNodeParams {
	return &GetKubernetesMetricsForNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesMetricsForNodeParamsWithTimeout creates a new GetKubernetesMetricsForNodeParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesMetricsForNodeParamsWithTimeout(timeout time.Duration) *GetKubernetesMetricsForNodeParams {
	return &GetKubernetesMetricsForNodeParams{
		timeout: timeout,
	}
}

// NewGetKubernetesMetricsForNodeParamsWithContext creates a new GetKubernetesMetricsForNodeParams object
// with the ability to set a context for a request.
func NewGetKubernetesMetricsForNodeParamsWithContext(ctx context.Context) *GetKubernetesMetricsForNodeParams {
	return &GetKubernetesMetricsForNodeParams{
		Context: ctx,
	}
}

// NewGetKubernetesMetricsForNodeParamsWithHTTPClient creates a new GetKubernetesMetricsForNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesMetricsForNodeParamsWithHTTPClient(client *http.Client) *GetKubernetesMetricsForNodeParams {
	return &GetKubernetesMetricsForNodeParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesMetricsForNodeParams contains all the parameters to send to the API endpoint

	for the get kubernetes metrics for node operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesMetricsForNodeParams struct {

	/* ID.

	   Environment (Endpoint) identifier
	*/
	ID int64

	/* Name.

	   Node identifier
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes metrics for node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesMetricsForNodeParams) WithDefaults() *GetKubernetesMetricsForNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes metrics for node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesMetricsForNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) WithTimeout(timeout time.Duration) *GetKubernetesMetricsForNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) WithContext(ctx context.Context) *GetKubernetesMetricsForNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) WithHTTPClient(client *http.Client) *GetKubernetesMetricsForNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) WithID(id int64) *GetKubernetesMetricsForNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) WithName(name string) *GetKubernetesMetricsForNodeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get kubernetes metrics for node params
func (o *GetKubernetesMetricsForNodeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesMetricsForNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
