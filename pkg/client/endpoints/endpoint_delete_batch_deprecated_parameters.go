// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewEndpointDeleteBatchDeprecatedParams creates a new EndpointDeleteBatchDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointDeleteBatchDeprecatedParams() *EndpointDeleteBatchDeprecatedParams {
	return &EndpointDeleteBatchDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointDeleteBatchDeprecatedParamsWithTimeout creates a new EndpointDeleteBatchDeprecatedParams object
// with the ability to set a timeout on a request.
func NewEndpointDeleteBatchDeprecatedParamsWithTimeout(timeout time.Duration) *EndpointDeleteBatchDeprecatedParams {
	return &EndpointDeleteBatchDeprecatedParams{
		timeout: timeout,
	}
}

// NewEndpointDeleteBatchDeprecatedParamsWithContext creates a new EndpointDeleteBatchDeprecatedParams object
// with the ability to set a context for a request.
func NewEndpointDeleteBatchDeprecatedParamsWithContext(ctx context.Context) *EndpointDeleteBatchDeprecatedParams {
	return &EndpointDeleteBatchDeprecatedParams{
		Context: ctx,
	}
}

// NewEndpointDeleteBatchDeprecatedParamsWithHTTPClient creates a new EndpointDeleteBatchDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointDeleteBatchDeprecatedParamsWithHTTPClient(client *http.Client) *EndpointDeleteBatchDeprecatedParams {
	return &EndpointDeleteBatchDeprecatedParams{
		HTTPClient: client,
	}
}

/*
EndpointDeleteBatchDeprecatedParams contains all the parameters to send to the API endpoint

	for the endpoint delete batch deprecated operation.

	Typically these are written to a http.Request.
*/
type EndpointDeleteBatchDeprecatedParams struct {

	/* Body.

	   List of environments to delete, with optional deleteCluster flag to clean-up associated resources (cloud environments only)
	*/
	Body *models.EndpointsEndpointDeleteBatchPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint delete batch deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointDeleteBatchDeprecatedParams) WithDefaults() *EndpointDeleteBatchDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint delete batch deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointDeleteBatchDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) WithTimeout(timeout time.Duration) *EndpointDeleteBatchDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) WithContext(ctx context.Context) *EndpointDeleteBatchDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) WithHTTPClient(client *http.Client) *EndpointDeleteBatchDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) WithBody(body *models.EndpointsEndpointDeleteBatchPayload) *EndpointDeleteBatchDeprecatedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint delete batch deprecated params
func (o *EndpointDeleteBatchDeprecatedParams) SetBody(body *models.EndpointsEndpointDeleteBatchPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointDeleteBatchDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
