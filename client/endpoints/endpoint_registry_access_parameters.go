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

	"github.com/portainer/client-api-go/v2/models"
)

// NewEndpointRegistryAccessParams creates a new EndpointRegistryAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointRegistryAccessParams() *EndpointRegistryAccessParams {
	return &EndpointRegistryAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointRegistryAccessParamsWithTimeout creates a new EndpointRegistryAccessParams object
// with the ability to set a timeout on a request.
func NewEndpointRegistryAccessParamsWithTimeout(timeout time.Duration) *EndpointRegistryAccessParams {
	return &EndpointRegistryAccessParams{
		timeout: timeout,
	}
}

// NewEndpointRegistryAccessParamsWithContext creates a new EndpointRegistryAccessParams object
// with the ability to set a context for a request.
func NewEndpointRegistryAccessParamsWithContext(ctx context.Context) *EndpointRegistryAccessParams {
	return &EndpointRegistryAccessParams{
		Context: ctx,
	}
}

// NewEndpointRegistryAccessParamsWithHTTPClient creates a new EndpointRegistryAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointRegistryAccessParamsWithHTTPClient(client *http.Client) *EndpointRegistryAccessParams {
	return &EndpointRegistryAccessParams{
		HTTPClient: client,
	}
}

/*
EndpointRegistryAccessParams contains all the parameters to send to the API endpoint

	for the endpoint registry access operation.

	Typically these are written to a http.Request.
*/
type EndpointRegistryAccessParams struct {

	/* Body.

	   details
	*/
	Body *models.EndpointsRegistryAccessPayload

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	/* RegistryID.

	   Registry identifier
	*/
	RegistryID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint registry access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointRegistryAccessParams) WithDefaults() *EndpointRegistryAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint registry access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointRegistryAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithTimeout(timeout time.Duration) *EndpointRegistryAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithContext(ctx context.Context) *EndpointRegistryAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithHTTPClient(client *http.Client) *EndpointRegistryAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithBody(body *models.EndpointsRegistryAccessPayload) *EndpointRegistryAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetBody(body *models.EndpointsRegistryAccessPayload) {
	o.Body = body
}

// WithID adds the id to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithID(id int64) *EndpointRegistryAccessParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetID(id int64) {
	o.ID = id
}

// WithRegistryID adds the registryID to the endpoint registry access params
func (o *EndpointRegistryAccessParams) WithRegistryID(registryID int64) *EndpointRegistryAccessParams {
	o.SetRegistryID(registryID)
	return o
}

// SetRegistryID adds the registryId to the endpoint registry access params
func (o *EndpointRegistryAccessParams) SetRegistryID(registryID int64) {
	o.RegistryID = registryID
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointRegistryAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param registryId
	if err := r.SetPathParam("registryId", swag.FormatInt64(o.RegistryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
