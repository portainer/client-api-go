// Code generated by go-swagger; DO NOT EDIT.

package ssl

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

// NewSSLUpdateParams creates a new SSLUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSLUpdateParams() *SSLUpdateParams {
	return &SSLUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSLUpdateParamsWithTimeout creates a new SSLUpdateParams object
// with the ability to set a timeout on a request.
func NewSSLUpdateParamsWithTimeout(timeout time.Duration) *SSLUpdateParams {
	return &SSLUpdateParams{
		timeout: timeout,
	}
}

// NewSSLUpdateParamsWithContext creates a new SSLUpdateParams object
// with the ability to set a context for a request.
func NewSSLUpdateParamsWithContext(ctx context.Context) *SSLUpdateParams {
	return &SSLUpdateParams{
		Context: ctx,
	}
}

// NewSSLUpdateParamsWithHTTPClient creates a new SSLUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSLUpdateParamsWithHTTPClient(client *http.Client) *SSLUpdateParams {
	return &SSLUpdateParams{
		HTTPClient: client,
	}
}

/*
SSLUpdateParams contains all the parameters to send to the API endpoint

	for the s s l update operation.

	Typically these are written to a http.Request.
*/
type SSLUpdateParams struct {

	/* Body.

	   SSL Settings
	*/
	Body *models.SslSslUpdatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s s l update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSLUpdateParams) WithDefaults() *SSLUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s s l update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSLUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s s l update params
func (o *SSLUpdateParams) WithTimeout(timeout time.Duration) *SSLUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s s l update params
func (o *SSLUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s s l update params
func (o *SSLUpdateParams) WithContext(ctx context.Context) *SSLUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s s l update params
func (o *SSLUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s s l update params
func (o *SSLUpdateParams) WithHTTPClient(client *http.Client) *SSLUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s s l update params
func (o *SSLUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s s l update params
func (o *SSLUpdateParams) WithBody(body *models.SslSslUpdatePayload) *SSLUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s s l update params
func (o *SSLUpdateParams) SetBody(body *models.SslSslUpdatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SSLUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
