// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewCurrentUserEndpointAuthorizationsInspectParams creates a new CurrentUserEndpointAuthorizationsInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCurrentUserEndpointAuthorizationsInspectParams() *CurrentUserEndpointAuthorizationsInspectParams {
	return &CurrentUserEndpointAuthorizationsInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCurrentUserEndpointAuthorizationsInspectParamsWithTimeout creates a new CurrentUserEndpointAuthorizationsInspectParams object
// with the ability to set a timeout on a request.
func NewCurrentUserEndpointAuthorizationsInspectParamsWithTimeout(timeout time.Duration) *CurrentUserEndpointAuthorizationsInspectParams {
	return &CurrentUserEndpointAuthorizationsInspectParams{
		timeout: timeout,
	}
}

// NewCurrentUserEndpointAuthorizationsInspectParamsWithContext creates a new CurrentUserEndpointAuthorizationsInspectParams object
// with the ability to set a context for a request.
func NewCurrentUserEndpointAuthorizationsInspectParamsWithContext(ctx context.Context) *CurrentUserEndpointAuthorizationsInspectParams {
	return &CurrentUserEndpointAuthorizationsInspectParams{
		Context: ctx,
	}
}

// NewCurrentUserEndpointAuthorizationsInspectParamsWithHTTPClient creates a new CurrentUserEndpointAuthorizationsInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCurrentUserEndpointAuthorizationsInspectParamsWithHTTPClient(client *http.Client) *CurrentUserEndpointAuthorizationsInspectParams {
	return &CurrentUserEndpointAuthorizationsInspectParams{
		HTTPClient: client,
	}
}

/*
CurrentUserEndpointAuthorizationsInspectParams contains all the parameters to send to the API endpoint

	for the current user endpoint authorizations inspect operation.

	Typically these are written to a http.Request.
*/
type CurrentUserEndpointAuthorizationsInspectParams struct {

	/* EndpointID.

	   Environment identifier
	*/
	EndpointID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the current user endpoint authorizations inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentUserEndpointAuthorizationsInspectParams) WithDefaults() *CurrentUserEndpointAuthorizationsInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the current user endpoint authorizations inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentUserEndpointAuthorizationsInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) WithTimeout(timeout time.Duration) *CurrentUserEndpointAuthorizationsInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) WithContext(ctx context.Context) *CurrentUserEndpointAuthorizationsInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) WithHTTPClient(client *http.Client) *CurrentUserEndpointAuthorizationsInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointID adds the endpointID to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) WithEndpointID(endpointID int64) *CurrentUserEndpointAuthorizationsInspectParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the current user endpoint authorizations inspect params
func (o *CurrentUserEndpointAuthorizationsInspectParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WriteToRequest writes these params to a swagger request
func (o *CurrentUserEndpointAuthorizationsInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointID
	if err := r.SetPathParam("endpointID", swag.FormatInt64(o.EndpointID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
