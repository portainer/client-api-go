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

// NewSettingsExperimentalInspectParams creates a new SettingsExperimentalInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsExperimentalInspectParams() *SettingsExperimentalInspectParams {
	return &SettingsExperimentalInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsExperimentalInspectParamsWithTimeout creates a new SettingsExperimentalInspectParams object
// with the ability to set a timeout on a request.
func NewSettingsExperimentalInspectParamsWithTimeout(timeout time.Duration) *SettingsExperimentalInspectParams {
	return &SettingsExperimentalInspectParams{
		timeout: timeout,
	}
}

// NewSettingsExperimentalInspectParamsWithContext creates a new SettingsExperimentalInspectParams object
// with the ability to set a context for a request.
func NewSettingsExperimentalInspectParamsWithContext(ctx context.Context) *SettingsExperimentalInspectParams {
	return &SettingsExperimentalInspectParams{
		Context: ctx,
	}
}

// NewSettingsExperimentalInspectParamsWithHTTPClient creates a new SettingsExperimentalInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsExperimentalInspectParamsWithHTTPClient(client *http.Client) *SettingsExperimentalInspectParams {
	return &SettingsExperimentalInspectParams{
		HTTPClient: client,
	}
}

/*
SettingsExperimentalInspectParams contains all the parameters to send to the API endpoint

	for the settings experimental inspect operation.

	Typically these are written to a http.Request.
*/
type SettingsExperimentalInspectParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings experimental inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsExperimentalInspectParams) WithDefaults() *SettingsExperimentalInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings experimental inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsExperimentalInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) WithTimeout(timeout time.Duration) *SettingsExperimentalInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) WithContext(ctx context.Context) *SettingsExperimentalInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) WithHTTPClient(client *http.Client) *SettingsExperimentalInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings experimental inspect params
func (o *SettingsExperimentalInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsExperimentalInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
