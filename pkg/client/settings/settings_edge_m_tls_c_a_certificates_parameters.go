// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewSettingsEdgeMTLSCACertificatesParams creates a new SettingsEdgeMTLSCACertificatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsEdgeMTLSCACertificatesParams() *SettingsEdgeMTLSCACertificatesParams {
	return &SettingsEdgeMTLSCACertificatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsEdgeMTLSCACertificatesParamsWithTimeout creates a new SettingsEdgeMTLSCACertificatesParams object
// with the ability to set a timeout on a request.
func NewSettingsEdgeMTLSCACertificatesParamsWithTimeout(timeout time.Duration) *SettingsEdgeMTLSCACertificatesParams {
	return &SettingsEdgeMTLSCACertificatesParams{
		timeout: timeout,
	}
}

// NewSettingsEdgeMTLSCACertificatesParamsWithContext creates a new SettingsEdgeMTLSCACertificatesParams object
// with the ability to set a context for a request.
func NewSettingsEdgeMTLSCACertificatesParamsWithContext(ctx context.Context) *SettingsEdgeMTLSCACertificatesParams {
	return &SettingsEdgeMTLSCACertificatesParams{
		Context: ctx,
	}
}

// NewSettingsEdgeMTLSCACertificatesParamsWithHTTPClient creates a new SettingsEdgeMTLSCACertificatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsEdgeMTLSCACertificatesParamsWithHTTPClient(client *http.Client) *SettingsEdgeMTLSCACertificatesParams {
	return &SettingsEdgeMTLSCACertificatesParams{
		HTTPClient: client,
	}
}

/*
SettingsEdgeMTLSCACertificatesParams contains all the parameters to send to the API endpoint

	for the settings edge m TLS c a certificates operation.

	Typically these are written to a http.Request.
*/
type SettingsEdgeMTLSCACertificatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings edge m TLS c a certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsEdgeMTLSCACertificatesParams) WithDefaults() *SettingsEdgeMTLSCACertificatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings edge m TLS c a certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsEdgeMTLSCACertificatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) WithTimeout(timeout time.Duration) *SettingsEdgeMTLSCACertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) WithContext(ctx context.Context) *SettingsEdgeMTLSCACertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) WithHTTPClient(client *http.Client) *SettingsEdgeMTLSCACertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings edge m TLS c a certificates params
func (o *SettingsEdgeMTLSCACertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsEdgeMTLSCACertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
