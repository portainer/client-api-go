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
	"github.com/go-openapi/swag"
)

// NewCustomTemplateFileParams creates a new CustomTemplateFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomTemplateFileParams() *CustomTemplateFileParams {
	return &CustomTemplateFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomTemplateFileParamsWithTimeout creates a new CustomTemplateFileParams object
// with the ability to set a timeout on a request.
func NewCustomTemplateFileParamsWithTimeout(timeout time.Duration) *CustomTemplateFileParams {
	return &CustomTemplateFileParams{
		timeout: timeout,
	}
}

// NewCustomTemplateFileParamsWithContext creates a new CustomTemplateFileParams object
// with the ability to set a context for a request.
func NewCustomTemplateFileParamsWithContext(ctx context.Context) *CustomTemplateFileParams {
	return &CustomTemplateFileParams{
		Context: ctx,
	}
}

// NewCustomTemplateFileParamsWithHTTPClient creates a new CustomTemplateFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomTemplateFileParamsWithHTTPClient(client *http.Client) *CustomTemplateFileParams {
	return &CustomTemplateFileParams{
		HTTPClient: client,
	}
}

/*
CustomTemplateFileParams contains all the parameters to send to the API endpoint

	for the custom template file operation.

	Typically these are written to a http.Request.
*/
type CustomTemplateFileParams struct {

	/* ID.

	   Template identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom template file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateFileParams) WithDefaults() *CustomTemplateFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom template file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom template file params
func (o *CustomTemplateFileParams) WithTimeout(timeout time.Duration) *CustomTemplateFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom template file params
func (o *CustomTemplateFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom template file params
func (o *CustomTemplateFileParams) WithContext(ctx context.Context) *CustomTemplateFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom template file params
func (o *CustomTemplateFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom template file params
func (o *CustomTemplateFileParams) WithHTTPClient(client *http.Client) *CustomTemplateFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom template file params
func (o *CustomTemplateFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom template file params
func (o *CustomTemplateFileParams) WithID(id int64) *CustomTemplateFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom template file params
func (o *CustomTemplateFileParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomTemplateFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
