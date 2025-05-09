// Code generated by go-swagger; DO NOT EDIT.

package kaas

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

// NewKaasVersionParams creates a new KaasVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKaasVersionParams() *KaasVersionParams {
	return &KaasVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKaasVersionParamsWithTimeout creates a new KaasVersionParams object
// with the ability to set a timeout on a request.
func NewKaasVersionParamsWithTimeout(timeout time.Duration) *KaasVersionParams {
	return &KaasVersionParams{
		timeout: timeout,
	}
}

// NewKaasVersionParamsWithContext creates a new KaasVersionParams object
// with the ability to set a context for a request.
func NewKaasVersionParamsWithContext(ctx context.Context) *KaasVersionParams {
	return &KaasVersionParams{
		Context: ctx,
	}
}

// NewKaasVersionParamsWithHTTPClient creates a new KaasVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewKaasVersionParamsWithHTTPClient(client *http.Client) *KaasVersionParams {
	return &KaasVersionParams{
		HTTPClient: client,
	}
}

/*
KaasVersionParams contains all the parameters to send to the API endpoint

	for the kaas version operation.

	Typically these are written to a http.Request.
*/
type KaasVersionParams struct {

	/* EnvironmentID.

	   Environment(Endpoint) identifier
	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kaas version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KaasVersionParams) WithDefaults() *KaasVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kaas version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KaasVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kaas version params
func (o *KaasVersionParams) WithTimeout(timeout time.Duration) *KaasVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kaas version params
func (o *KaasVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kaas version params
func (o *KaasVersionParams) WithContext(ctx context.Context) *KaasVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kaas version params
func (o *KaasVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kaas version params
func (o *KaasVersionParams) WithHTTPClient(client *http.Client) *KaasVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kaas version params
func (o *KaasVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the kaas version params
func (o *KaasVersionParams) WithEnvironmentID(environmentID int64) *KaasVersionParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the kaas version params
func (o *KaasVersionParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *KaasVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
