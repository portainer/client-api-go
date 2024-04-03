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

// NewUpgradeParams creates a new UpgradeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeParams() *UpgradeParams {
	return &UpgradeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeParamsWithTimeout creates a new UpgradeParams object
// with the ability to set a timeout on a request.
func NewUpgradeParamsWithTimeout(timeout time.Duration) *UpgradeParams {
	return &UpgradeParams{
		timeout: timeout,
	}
}

// NewUpgradeParamsWithContext creates a new UpgradeParams object
// with the ability to set a context for a request.
func NewUpgradeParamsWithContext(ctx context.Context) *UpgradeParams {
	return &UpgradeParams{
		Context: ctx,
	}
}

// NewUpgradeParamsWithHTTPClient creates a new UpgradeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeParamsWithHTTPClient(client *http.Client) *UpgradeParams {
	return &UpgradeParams{
		HTTPClient: client,
	}
}

/*
UpgradeParams contains all the parameters to send to the API endpoint

	for the upgrade operation.

	Typically these are written to a http.Request.
*/
type UpgradeParams struct {

	/* EnvironmentID.

	   Environment(Endpoint) identifier
	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeParams) WithDefaults() *UpgradeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade params
func (o *UpgradeParams) WithTimeout(timeout time.Duration) *UpgradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade params
func (o *UpgradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade params
func (o *UpgradeParams) WithContext(ctx context.Context) *UpgradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade params
func (o *UpgradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade params
func (o *UpgradeParams) WithHTTPClient(client *http.Client) *UpgradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade params
func (o *UpgradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the upgrade params
func (o *UpgradeParams) WithEnvironmentID(environmentID int64) *UpgradeParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the upgrade params
func (o *UpgradeParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
