// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewLicensesListParams creates a new LicensesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicensesListParams() *LicensesListParams {
	return &LicensesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicensesListParamsWithTimeout creates a new LicensesListParams object
// with the ability to set a timeout on a request.
func NewLicensesListParamsWithTimeout(timeout time.Duration) *LicensesListParams {
	return &LicensesListParams{
		timeout: timeout,
	}
}

// NewLicensesListParamsWithContext creates a new LicensesListParams object
// with the ability to set a context for a request.
func NewLicensesListParamsWithContext(ctx context.Context) *LicensesListParams {
	return &LicensesListParams{
		Context: ctx,
	}
}

// NewLicensesListParamsWithHTTPClient creates a new LicensesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicensesListParamsWithHTTPClient(client *http.Client) *LicensesListParams {
	return &LicensesListParams{
		HTTPClient: client,
	}
}

/*
LicensesListParams contains all the parameters to send to the API endpoint

	for the licenses list operation.

	Typically these are written to a http.Request.
*/
type LicensesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the licenses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicensesListParams) WithDefaults() *LicensesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the licenses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicensesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the licenses list params
func (o *LicensesListParams) WithTimeout(timeout time.Duration) *LicensesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the licenses list params
func (o *LicensesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the licenses list params
func (o *LicensesListParams) WithContext(ctx context.Context) *LicensesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the licenses list params
func (o *LicensesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the licenses list params
func (o *LicensesListParams) WithHTTPClient(client *http.Client) *LicensesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the licenses list params
func (o *LicensesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LicensesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
