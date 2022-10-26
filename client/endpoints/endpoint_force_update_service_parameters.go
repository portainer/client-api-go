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
	"github.com/go-openapi/swag"

	"github.com/portainer/client-api-go/models"
)

// NewEndpointForceUpdateServiceParams creates a new EndpointForceUpdateServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointForceUpdateServiceParams() *EndpointForceUpdateServiceParams {
	return &EndpointForceUpdateServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointForceUpdateServiceParamsWithTimeout creates a new EndpointForceUpdateServiceParams object
// with the ability to set a timeout on a request.
func NewEndpointForceUpdateServiceParamsWithTimeout(timeout time.Duration) *EndpointForceUpdateServiceParams {
	return &EndpointForceUpdateServiceParams{
		timeout: timeout,
	}
}

// NewEndpointForceUpdateServiceParamsWithContext creates a new EndpointForceUpdateServiceParams object
// with the ability to set a context for a request.
func NewEndpointForceUpdateServiceParamsWithContext(ctx context.Context) *EndpointForceUpdateServiceParams {
	return &EndpointForceUpdateServiceParams{
		Context: ctx,
	}
}

// NewEndpointForceUpdateServiceParamsWithHTTPClient creates a new EndpointForceUpdateServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointForceUpdateServiceParamsWithHTTPClient(client *http.Client) *EndpointForceUpdateServiceParams {
	return &EndpointForceUpdateServiceParams{
		HTTPClient: client,
	}
}

/*
EndpointForceUpdateServiceParams contains all the parameters to send to the API endpoint

	for the endpoint force update service operation.

	Typically these are written to a http.Request.
*/
type EndpointForceUpdateServiceParams struct {

	/* Body.

	   details
	*/
	Body *models.EndpointsForceUpdateServicePayload

	/* ID.

	   endpoint identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint force update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointForceUpdateServiceParams) WithDefaults() *EndpointForceUpdateServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint force update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointForceUpdateServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) WithTimeout(timeout time.Duration) *EndpointForceUpdateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) WithContext(ctx context.Context) *EndpointForceUpdateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) WithHTTPClient(client *http.Client) *EndpointForceUpdateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) WithBody(body *models.EndpointsForceUpdateServicePayload) *EndpointForceUpdateServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) SetBody(body *models.EndpointsForceUpdateServicePayload) {
	o.Body = body
}

// WithID adds the id to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) WithID(id int64) *EndpointForceUpdateServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint force update service params
func (o *EndpointForceUpdateServiceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointForceUpdateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
