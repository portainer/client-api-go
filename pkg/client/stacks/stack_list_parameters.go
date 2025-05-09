// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// NewStackListParams creates a new StackListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackListParams() *StackListParams {
	return &StackListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackListParamsWithTimeout creates a new StackListParams object
// with the ability to set a timeout on a request.
func NewStackListParamsWithTimeout(timeout time.Duration) *StackListParams {
	return &StackListParams{
		timeout: timeout,
	}
}

// NewStackListParamsWithContext creates a new StackListParams object
// with the ability to set a context for a request.
func NewStackListParamsWithContext(ctx context.Context) *StackListParams {
	return &StackListParams{
		Context: ctx,
	}
}

// NewStackListParamsWithHTTPClient creates a new StackListParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackListParamsWithHTTPClient(client *http.Client) *StackListParams {
	return &StackListParams{
		HTTPClient: client,
	}
}

/*
StackListParams contains all the parameters to send to the API endpoint

	for the stack list operation.

	Typically these are written to a http.Request.
*/
type StackListParams struct {

	/* Filters.

	   Filters to process on the stack list. Encoded as JSON (a map[string]string). For example, {'SwarmID': 'jpofkc0i9uo9wtx1zesuk649w'} will only return stacks that are part of the specified Swarm cluster. Available filters: EndpointID, SwarmID.
	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackListParams) WithDefaults() *StackListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack list params
func (o *StackListParams) WithTimeout(timeout time.Duration) *StackListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack list params
func (o *StackListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack list params
func (o *StackListParams) WithContext(ctx context.Context) *StackListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack list params
func (o *StackListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack list params
func (o *StackListParams) WithHTTPClient(client *http.Client) *StackListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack list params
func (o *StackListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the stack list params
func (o *StackListParams) WithFilters(filters *string) *StackListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the stack list params
func (o *StackListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *StackListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
