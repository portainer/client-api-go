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

// NewSettingsEdgeMTLSCertificatesParams creates a new SettingsEdgeMTLSCertificatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsEdgeMTLSCertificatesParams() *SettingsEdgeMTLSCertificatesParams {
	return &SettingsEdgeMTLSCertificatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsEdgeMTLSCertificatesParamsWithTimeout creates a new SettingsEdgeMTLSCertificatesParams object
// with the ability to set a timeout on a request.
func NewSettingsEdgeMTLSCertificatesParamsWithTimeout(timeout time.Duration) *SettingsEdgeMTLSCertificatesParams {
	return &SettingsEdgeMTLSCertificatesParams{
		timeout: timeout,
	}
}

// NewSettingsEdgeMTLSCertificatesParamsWithContext creates a new SettingsEdgeMTLSCertificatesParams object
// with the ability to set a context for a request.
func NewSettingsEdgeMTLSCertificatesParamsWithContext(ctx context.Context) *SettingsEdgeMTLSCertificatesParams {
	return &SettingsEdgeMTLSCertificatesParams{
		Context: ctx,
	}
}

// NewSettingsEdgeMTLSCertificatesParamsWithHTTPClient creates a new SettingsEdgeMTLSCertificatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsEdgeMTLSCertificatesParamsWithHTTPClient(client *http.Client) *SettingsEdgeMTLSCertificatesParams {
	return &SettingsEdgeMTLSCertificatesParams{
		HTTPClient: client,
	}
}

/*
SettingsEdgeMTLSCertificatesParams contains all the parameters to send to the API endpoint

	for the settings edge m TLS certificates operation.

	Typically these are written to a http.Request.
*/
type SettingsEdgeMTLSCertificatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings edge m TLS certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsEdgeMTLSCertificatesParams) WithDefaults() *SettingsEdgeMTLSCertificatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings edge m TLS certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsEdgeMTLSCertificatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) WithTimeout(timeout time.Duration) *SettingsEdgeMTLSCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) WithContext(ctx context.Context) *SettingsEdgeMTLSCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) WithHTTPClient(client *http.Client) *SettingsEdgeMTLSCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings edge m TLS certificates params
func (o *SettingsEdgeMTLSCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsEdgeMTLSCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
