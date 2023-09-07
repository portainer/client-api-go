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

	"github.com/portainer/client-api-go/v2/models"
)

// NewTestSSHParams creates a new TestSSHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestSSHParams() *TestSSHParams {
	return &TestSSHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestSSHParamsWithTimeout creates a new TestSSHParams object
// with the ability to set a timeout on a request.
func NewTestSSHParamsWithTimeout(timeout time.Duration) *TestSSHParams {
	return &TestSSHParams{
		timeout: timeout,
	}
}

// NewTestSSHParamsWithContext creates a new TestSSHParams object
// with the ability to set a context for a request.
func NewTestSSHParamsWithContext(ctx context.Context) *TestSSHParams {
	return &TestSSHParams{
		Context: ctx,
	}
}

// NewTestSSHParamsWithHTTPClient creates a new TestSSHParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestSSHParamsWithHTTPClient(client *http.Client) *TestSSHParams {
	return &TestSSHParams{
		HTTPClient: client,
	}
}

/*
TestSSHParams contains all the parameters to send to the API endpoint

	for the test SSH operation.

	Typically these are written to a http.Request.
*/
type TestSSHParams struct {

	/* Body.

	   Node IPs and credential ID
	*/
	Body *models.ProvidersMicrok8sTestSSHPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test SSH params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestSSHParams) WithDefaults() *TestSSHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test SSH params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestSSHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test SSH params
func (o *TestSSHParams) WithTimeout(timeout time.Duration) *TestSSHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test SSH params
func (o *TestSSHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test SSH params
func (o *TestSSHParams) WithContext(ctx context.Context) *TestSSHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test SSH params
func (o *TestSSHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test SSH params
func (o *TestSSHParams) WithHTTPClient(client *http.Client) *TestSSHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test SSH params
func (o *TestSSHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test SSH params
func (o *TestSSHParams) WithBody(body *models.ProvidersMicrok8sTestSSHPayload) *TestSSHParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test SSH params
func (o *TestSSHParams) SetBody(body *models.ProvidersMicrok8sTestSSHPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestSSHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
