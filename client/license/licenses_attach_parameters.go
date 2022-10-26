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

	"github.com/portainer/client-api-go/models"
)

// NewLicensesAttachParams creates a new LicensesAttachParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicensesAttachParams() *LicensesAttachParams {
	return &LicensesAttachParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicensesAttachParamsWithTimeout creates a new LicensesAttachParams object
// with the ability to set a timeout on a request.
func NewLicensesAttachParamsWithTimeout(timeout time.Duration) *LicensesAttachParams {
	return &LicensesAttachParams{
		timeout: timeout,
	}
}

// NewLicensesAttachParamsWithContext creates a new LicensesAttachParams object
// with the ability to set a context for a request.
func NewLicensesAttachParamsWithContext(ctx context.Context) *LicensesAttachParams {
	return &LicensesAttachParams{
		Context: ctx,
	}
}

// NewLicensesAttachParamsWithHTTPClient creates a new LicensesAttachParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicensesAttachParamsWithHTTPClient(client *http.Client) *LicensesAttachParams {
	return &LicensesAttachParams{
		HTTPClient: client,
	}
}

/*
LicensesAttachParams contains all the parameters to send to the API endpoint

	for the licenses attach operation.

	Typically these are written to a http.Request.
*/
type LicensesAttachParams struct {

	/* Body.

	   list of licenses keys to attach
	*/
	Body *models.LicensesAttachPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the licenses attach params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicensesAttachParams) WithDefaults() *LicensesAttachParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the licenses attach params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicensesAttachParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the licenses attach params
func (o *LicensesAttachParams) WithTimeout(timeout time.Duration) *LicensesAttachParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the licenses attach params
func (o *LicensesAttachParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the licenses attach params
func (o *LicensesAttachParams) WithContext(ctx context.Context) *LicensesAttachParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the licenses attach params
func (o *LicensesAttachParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the licenses attach params
func (o *LicensesAttachParams) WithHTTPClient(client *http.Client) *LicensesAttachParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the licenses attach params
func (o *LicensesAttachParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the licenses attach params
func (o *LicensesAttachParams) WithBody(body *models.LicensesAttachPayload) *LicensesAttachParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the licenses attach params
func (o *LicensesAttachParams) SetBody(body *models.LicensesAttachPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LicensesAttachParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
