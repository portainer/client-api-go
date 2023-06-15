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
)

// NewEndpointDockerhubStatusParams creates a new EndpointDockerhubStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointDockerhubStatusParams() *EndpointDockerhubStatusParams {
	return &EndpointDockerhubStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointDockerhubStatusParamsWithTimeout creates a new EndpointDockerhubStatusParams object
// with the ability to set a timeout on a request.
func NewEndpointDockerhubStatusParamsWithTimeout(timeout time.Duration) *EndpointDockerhubStatusParams {
	return &EndpointDockerhubStatusParams{
		timeout: timeout,
	}
}

// NewEndpointDockerhubStatusParamsWithContext creates a new EndpointDockerhubStatusParams object
// with the ability to set a context for a request.
func NewEndpointDockerhubStatusParamsWithContext(ctx context.Context) *EndpointDockerhubStatusParams {
	return &EndpointDockerhubStatusParams{
		Context: ctx,
	}
}

// NewEndpointDockerhubStatusParamsWithHTTPClient creates a new EndpointDockerhubStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointDockerhubStatusParamsWithHTTPClient(client *http.Client) *EndpointDockerhubStatusParams {
	return &EndpointDockerhubStatusParams{
		HTTPClient: client,
	}
}

/*
EndpointDockerhubStatusParams contains all the parameters to send to the API endpoint

	for the endpoint dockerhub status operation.

	Typically these are written to a http.Request.
*/
type EndpointDockerhubStatusParams struct {

	/* ID.

	   endpoint ID
	*/
	ID int64

	/* RegistryID.

	   registry ID
	*/
	RegistryID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint dockerhub status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointDockerhubStatusParams) WithDefaults() *EndpointDockerhubStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint dockerhub status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointDockerhubStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) WithTimeout(timeout time.Duration) *EndpointDockerhubStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) WithContext(ctx context.Context) *EndpointDockerhubStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) WithHTTPClient(client *http.Client) *EndpointDockerhubStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) WithID(id int64) *EndpointDockerhubStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) SetID(id int64) {
	o.ID = id
}

// WithRegistryID adds the registryID to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) WithRegistryID(registryID int64) *EndpointDockerhubStatusParams {
	o.SetRegistryID(registryID)
	return o
}

// SetRegistryID adds the registryId to the endpoint dockerhub status params
func (o *EndpointDockerhubStatusParams) SetRegistryID(registryID int64) {
	o.RegistryID = registryID
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointDockerhubStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
