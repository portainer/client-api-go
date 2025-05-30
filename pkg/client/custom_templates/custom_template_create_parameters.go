// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

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

// NewCustomTemplateCreateParams creates a new CustomTemplateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomTemplateCreateParams() *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomTemplateCreateParamsWithTimeout creates a new CustomTemplateCreateParams object
// with the ability to set a timeout on a request.
func NewCustomTemplateCreateParamsWithTimeout(timeout time.Duration) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		timeout: timeout,
	}
}

// NewCustomTemplateCreateParamsWithContext creates a new CustomTemplateCreateParams object
// with the ability to set a context for a request.
func NewCustomTemplateCreateParamsWithContext(ctx context.Context) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		Context: ctx,
	}
}

// NewCustomTemplateCreateParamsWithHTTPClient creates a new CustomTemplateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomTemplateCreateParamsWithHTTPClient(client *http.Client) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		HTTPClient: client,
	}
}

/*
CustomTemplateCreateParams contains all the parameters to send to the API endpoint

	for the custom template create operation.

	Typically these are written to a http.Request.
*/
type CustomTemplateCreateParams struct {

	/* Body.

	   for body documentation see the relevant /custom_templates/{method} endpoint
	*/
	Body interface{}

	/* Method.

	   method for creating template
	*/
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateParams) WithDefaults() *CustomTemplateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom template create params
func (o *CustomTemplateCreateParams) WithTimeout(timeout time.Duration) *CustomTemplateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom template create params
func (o *CustomTemplateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom template create params
func (o *CustomTemplateCreateParams) WithContext(ctx context.Context) *CustomTemplateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom template create params
func (o *CustomTemplateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom template create params
func (o *CustomTemplateCreateParams) WithHTTPClient(client *http.Client) *CustomTemplateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom template create params
func (o *CustomTemplateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the custom template create params
func (o *CustomTemplateCreateParams) WithBody(body interface{}) *CustomTemplateCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the custom template create params
func (o *CustomTemplateCreateParams) SetBody(body interface{}) {
	o.Body = body
}

// WithMethod adds the method to the custom template create params
func (o *CustomTemplateCreateParams) WithMethod(method string) *CustomTemplateCreateParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the custom template create params
func (o *CustomTemplateCreateParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *CustomTemplateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param method
	qrMethod := o.Method
	qMethod := qrMethod
	if qMethod != "" {

		if err := r.SetQueryParam("method", qMethod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
