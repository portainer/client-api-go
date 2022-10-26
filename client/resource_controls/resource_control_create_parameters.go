// Code generated by go-swagger; DO NOT EDIT.

package resource_controls

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

// NewResourceControlCreateParams creates a new ResourceControlCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceControlCreateParams() *ResourceControlCreateParams {
	return &ResourceControlCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceControlCreateParamsWithTimeout creates a new ResourceControlCreateParams object
// with the ability to set a timeout on a request.
func NewResourceControlCreateParamsWithTimeout(timeout time.Duration) *ResourceControlCreateParams {
	return &ResourceControlCreateParams{
		timeout: timeout,
	}
}

// NewResourceControlCreateParamsWithContext creates a new ResourceControlCreateParams object
// with the ability to set a context for a request.
func NewResourceControlCreateParamsWithContext(ctx context.Context) *ResourceControlCreateParams {
	return &ResourceControlCreateParams{
		Context: ctx,
	}
}

// NewResourceControlCreateParamsWithHTTPClient creates a new ResourceControlCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceControlCreateParamsWithHTTPClient(client *http.Client) *ResourceControlCreateParams {
	return &ResourceControlCreateParams{
		HTTPClient: client,
	}
}

/*
ResourceControlCreateParams contains all the parameters to send to the API endpoint

	for the resource control create operation.

	Typically these are written to a http.Request.
*/
type ResourceControlCreateParams struct {

	/* Body.

	   Resource control details
	*/
	Body *models.ResourcecontrolsResourceControlCreatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource control create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceControlCreateParams) WithDefaults() *ResourceControlCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource control create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceControlCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource control create params
func (o *ResourceControlCreateParams) WithTimeout(timeout time.Duration) *ResourceControlCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource control create params
func (o *ResourceControlCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource control create params
func (o *ResourceControlCreateParams) WithContext(ctx context.Context) *ResourceControlCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource control create params
func (o *ResourceControlCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource control create params
func (o *ResourceControlCreateParams) WithHTTPClient(client *http.Client) *ResourceControlCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource control create params
func (o *ResourceControlCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resource control create params
func (o *ResourceControlCreateParams) WithBody(body *models.ResourcecontrolsResourceControlCreatePayload) *ResourceControlCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resource control create params
func (o *ResourceControlCreateParams) SetBody(body *models.ResourcecontrolsResourceControlCreatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceControlCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
