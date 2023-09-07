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

	"github.com/portainer/client-api-go/v2/models"
)

// NewCustomTemplateCreateRepositoryParams creates a new CustomTemplateCreateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomTemplateCreateRepositoryParams() *CustomTemplateCreateRepositoryParams {
	return &CustomTemplateCreateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomTemplateCreateRepositoryParamsWithTimeout creates a new CustomTemplateCreateRepositoryParams object
// with the ability to set a timeout on a request.
func NewCustomTemplateCreateRepositoryParamsWithTimeout(timeout time.Duration) *CustomTemplateCreateRepositoryParams {
	return &CustomTemplateCreateRepositoryParams{
		timeout: timeout,
	}
}

// NewCustomTemplateCreateRepositoryParamsWithContext creates a new CustomTemplateCreateRepositoryParams object
// with the ability to set a context for a request.
func NewCustomTemplateCreateRepositoryParamsWithContext(ctx context.Context) *CustomTemplateCreateRepositoryParams {
	return &CustomTemplateCreateRepositoryParams{
		Context: ctx,
	}
}

// NewCustomTemplateCreateRepositoryParamsWithHTTPClient creates a new CustomTemplateCreateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomTemplateCreateRepositoryParamsWithHTTPClient(client *http.Client) *CustomTemplateCreateRepositoryParams {
	return &CustomTemplateCreateRepositoryParams{
		HTTPClient: client,
	}
}

/*
CustomTemplateCreateRepositoryParams contains all the parameters to send to the API endpoint

	for the custom template create repository operation.

	Typically these are written to a http.Request.
*/
type CustomTemplateCreateRepositoryParams struct {

	/* Body.

	   Required when using method=repository
	*/
	Body *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom template create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateRepositoryParams) WithDefaults() *CustomTemplateCreateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom template create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) WithTimeout(timeout time.Duration) *CustomTemplateCreateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) WithContext(ctx context.Context) *CustomTemplateCreateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) WithHTTPClient(client *http.Client) *CustomTemplateCreateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) WithBody(body *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload) *CustomTemplateCreateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the custom template create repository params
func (o *CustomTemplateCreateRepositoryParams) SetBody(body *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CustomTemplateCreateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
