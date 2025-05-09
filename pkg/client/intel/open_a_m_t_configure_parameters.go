// Code generated by go-swagger; DO NOT EDIT.

package intel

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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewOpenAMTConfigureParams creates a new OpenAMTConfigureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenAMTConfigureParams() *OpenAMTConfigureParams {
	return &OpenAMTConfigureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenAMTConfigureParamsWithTimeout creates a new OpenAMTConfigureParams object
// with the ability to set a timeout on a request.
func NewOpenAMTConfigureParamsWithTimeout(timeout time.Duration) *OpenAMTConfigureParams {
	return &OpenAMTConfigureParams{
		timeout: timeout,
	}
}

// NewOpenAMTConfigureParamsWithContext creates a new OpenAMTConfigureParams object
// with the ability to set a context for a request.
func NewOpenAMTConfigureParamsWithContext(ctx context.Context) *OpenAMTConfigureParams {
	return &OpenAMTConfigureParams{
		Context: ctx,
	}
}

// NewOpenAMTConfigureParamsWithHTTPClient creates a new OpenAMTConfigureParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenAMTConfigureParamsWithHTTPClient(client *http.Client) *OpenAMTConfigureParams {
	return &OpenAMTConfigureParams{
		HTTPClient: client,
	}
}

/*
OpenAMTConfigureParams contains all the parameters to send to the API endpoint

	for the open a m t configure operation.

	Typically these are written to a http.Request.
*/
type OpenAMTConfigureParams struct {

	/* Body.

	   OpenAMT Settings
	*/
	Body *models.OpenamtOpenAMTConfigurePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the open a m t configure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAMTConfigureParams) WithDefaults() *OpenAMTConfigureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the open a m t configure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAMTConfigureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the open a m t configure params
func (o *OpenAMTConfigureParams) WithTimeout(timeout time.Duration) *OpenAMTConfigureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open a m t configure params
func (o *OpenAMTConfigureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open a m t configure params
func (o *OpenAMTConfigureParams) WithContext(ctx context.Context) *OpenAMTConfigureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open a m t configure params
func (o *OpenAMTConfigureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open a m t configure params
func (o *OpenAMTConfigureParams) WithHTTPClient(client *http.Client) *OpenAMTConfigureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open a m t configure params
func (o *OpenAMTConfigureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the open a m t configure params
func (o *OpenAMTConfigureParams) WithBody(body *models.OpenamtOpenAMTConfigurePayload) *OpenAMTConfigureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the open a m t configure params
func (o *OpenAMTConfigureParams) SetBody(body *models.OpenamtOpenAMTConfigurePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OpenAMTConfigureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
