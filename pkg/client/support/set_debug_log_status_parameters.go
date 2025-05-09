// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewSetDebugLogStatusParams creates a new SetDebugLogStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDebugLogStatusParams() *SetDebugLogStatusParams {
	return &SetDebugLogStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDebugLogStatusParamsWithTimeout creates a new SetDebugLogStatusParams object
// with the ability to set a timeout on a request.
func NewSetDebugLogStatusParamsWithTimeout(timeout time.Duration) *SetDebugLogStatusParams {
	return &SetDebugLogStatusParams{
		timeout: timeout,
	}
}

// NewSetDebugLogStatusParamsWithContext creates a new SetDebugLogStatusParams object
// with the ability to set a context for a request.
func NewSetDebugLogStatusParamsWithContext(ctx context.Context) *SetDebugLogStatusParams {
	return &SetDebugLogStatusParams{
		Context: ctx,
	}
}

// NewSetDebugLogStatusParamsWithHTTPClient creates a new SetDebugLogStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDebugLogStatusParamsWithHTTPClient(client *http.Client) *SetDebugLogStatusParams {
	return &SetDebugLogStatusParams{
		HTTPClient: client,
	}
}

/*
SetDebugLogStatusParams contains all the parameters to send to the API endpoint

	for the set debug log status operation.

	Typically these are written to a http.Request.
*/
type SetDebugLogStatusParams struct {

	/* Body.

	   Debug log should be set to enabled or disabled
	*/
	Body *models.SupportDebugLogPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set debug log status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDebugLogStatusParams) WithDefaults() *SetDebugLogStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set debug log status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDebugLogStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set debug log status params
func (o *SetDebugLogStatusParams) WithTimeout(timeout time.Duration) *SetDebugLogStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set debug log status params
func (o *SetDebugLogStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set debug log status params
func (o *SetDebugLogStatusParams) WithContext(ctx context.Context) *SetDebugLogStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set debug log status params
func (o *SetDebugLogStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set debug log status params
func (o *SetDebugLogStatusParams) WithHTTPClient(client *http.Client) *SetDebugLogStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set debug log status params
func (o *SetDebugLogStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set debug log status params
func (o *SetDebugLogStatusParams) WithBody(body *models.SupportDebugLogPayload) *SetDebugLogStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set debug log status params
func (o *SetDebugLogStatusParams) SetBody(body *models.SupportDebugLogPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetDebugLogStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
