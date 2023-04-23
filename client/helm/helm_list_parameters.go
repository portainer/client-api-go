// Code generated by go-swagger; DO NOT EDIT.

package helm

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

// NewHelmListParams creates a new HelmListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHelmListParams() *HelmListParams {
	return &HelmListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHelmListParamsWithTimeout creates a new HelmListParams object
// with the ability to set a timeout on a request.
func NewHelmListParamsWithTimeout(timeout time.Duration) *HelmListParams {
	return &HelmListParams{
		timeout: timeout,
	}
}

// NewHelmListParamsWithContext creates a new HelmListParams object
// with the ability to set a context for a request.
func NewHelmListParamsWithContext(ctx context.Context) *HelmListParams {
	return &HelmListParams{
		Context: ctx,
	}
}

// NewHelmListParamsWithHTTPClient creates a new HelmListParams object
// with the ability to set a custom HTTPClient for a request.
func NewHelmListParamsWithHTTPClient(client *http.Client) *HelmListParams {
	return &HelmListParams{
		HTTPClient: client,
	}
}

/*
HelmListParams contains all the parameters to send to the API endpoint

	for the helm list operation.

	Typically these are written to a http.Request.
*/
type HelmListParams struct {

	/* Filter.

	   specify an optional filter
	*/
	Filter string

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	/* Namespace.

	   specify an optional namespace
	*/
	Namespace string

	/* Selector.

	   specify an optional selector
	*/
	Selector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the helm list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmListParams) WithDefaults() *HelmListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the helm list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the helm list params
func (o *HelmListParams) WithTimeout(timeout time.Duration) *HelmListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the helm list params
func (o *HelmListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the helm list params
func (o *HelmListParams) WithContext(ctx context.Context) *HelmListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the helm list params
func (o *HelmListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the helm list params
func (o *HelmListParams) WithHTTPClient(client *http.Client) *HelmListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the helm list params
func (o *HelmListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the helm list params
func (o *HelmListParams) WithFilter(filter string) *HelmListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the helm list params
func (o *HelmListParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithID adds the id to the helm list params
func (o *HelmListParams) WithID(id int64) *HelmListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the helm list params
func (o *HelmListParams) SetID(id int64) {
	o.ID = id
}

// WithNamespace adds the namespace to the helm list params
func (o *HelmListParams) WithNamespace(namespace string) *HelmListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the helm list params
func (o *HelmListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSelector adds the selector to the helm list params
func (o *HelmListParams) WithSelector(selector string) *HelmListParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the helm list params
func (o *HelmListParams) SetSelector(selector string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *HelmListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter
	if qFilter != "" {

		if err := r.SetQueryParam("filter", qFilter); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {

		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	// query param selector
	qrSelector := o.Selector
	qSelector := qrSelector
	if qSelector != "" {

		if err := r.SetQueryParam("selector", qSelector); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
