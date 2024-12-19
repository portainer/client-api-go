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

// NewProvisionClusterAzureParams creates a new ProvisionClusterAzureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionClusterAzureParams() *ProvisionClusterAzureParams {
	return &ProvisionClusterAzureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionClusterAzureParamsWithTimeout creates a new ProvisionClusterAzureParams object
// with the ability to set a timeout on a request.
func NewProvisionClusterAzureParamsWithTimeout(timeout time.Duration) *ProvisionClusterAzureParams {
	return &ProvisionClusterAzureParams{
		timeout: timeout,
	}
}

// NewProvisionClusterAzureParamsWithContext creates a new ProvisionClusterAzureParams object
// with the ability to set a context for a request.
func NewProvisionClusterAzureParamsWithContext(ctx context.Context) *ProvisionClusterAzureParams {
	return &ProvisionClusterAzureParams{
		Context: ctx,
	}
}

// NewProvisionClusterAzureParamsWithHTTPClient creates a new ProvisionClusterAzureParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionClusterAzureParamsWithHTTPClient(client *http.Client) *ProvisionClusterAzureParams {
	return &ProvisionClusterAzureParams{
		HTTPClient: client,
	}
}

/*
ProvisionClusterAzureParams contains all the parameters to send to the API endpoint

	for the provision cluster azure operation.

	Typically these are written to a http.Request.
*/
type ProvisionClusterAzureParams struct {

	/* Body.

	   KaaS cluster provisioning details
	*/
	Body *models.ProvidersAzureProvisionPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provision cluster azure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionClusterAzureParams) WithDefaults() *ProvisionClusterAzureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provision cluster azure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionClusterAzureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provision cluster azure params
func (o *ProvisionClusterAzureParams) WithTimeout(timeout time.Duration) *ProvisionClusterAzureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provision cluster azure params
func (o *ProvisionClusterAzureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provision cluster azure params
func (o *ProvisionClusterAzureParams) WithContext(ctx context.Context) *ProvisionClusterAzureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provision cluster azure params
func (o *ProvisionClusterAzureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provision cluster azure params
func (o *ProvisionClusterAzureParams) WithHTTPClient(client *http.Client) *ProvisionClusterAzureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provision cluster azure params
func (o *ProvisionClusterAzureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the provision cluster azure params
func (o *ProvisionClusterAzureParams) WithBody(body *models.ProvidersAzureProvisionPayload) *ProvisionClusterAzureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the provision cluster azure params
func (o *ProvisionClusterAzureParams) SetBody(body *models.ProvidersAzureProvisionPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionClusterAzureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
