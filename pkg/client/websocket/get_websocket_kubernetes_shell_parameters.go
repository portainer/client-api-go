// Code generated by go-swagger; DO NOT EDIT.

package websocket

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

// NewGetWebsocketKubernetesShellParams creates a new GetWebsocketKubernetesShellParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsocketKubernetesShellParams() *GetWebsocketKubernetesShellParams {
	return &GetWebsocketKubernetesShellParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsocketKubernetesShellParamsWithTimeout creates a new GetWebsocketKubernetesShellParams object
// with the ability to set a timeout on a request.
func NewGetWebsocketKubernetesShellParamsWithTimeout(timeout time.Duration) *GetWebsocketKubernetesShellParams {
	return &GetWebsocketKubernetesShellParams{
		timeout: timeout,
	}
}

// NewGetWebsocketKubernetesShellParamsWithContext creates a new GetWebsocketKubernetesShellParams object
// with the ability to set a context for a request.
func NewGetWebsocketKubernetesShellParamsWithContext(ctx context.Context) *GetWebsocketKubernetesShellParams {
	return &GetWebsocketKubernetesShellParams{
		Context: ctx,
	}
}

// NewGetWebsocketKubernetesShellParamsWithHTTPClient creates a new GetWebsocketKubernetesShellParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsocketKubernetesShellParamsWithHTTPClient(client *http.Client) *GetWebsocketKubernetesShellParams {
	return &GetWebsocketKubernetesShellParams{
		HTTPClient: client,
	}
}

/*
GetWebsocketKubernetesShellParams contains all the parameters to send to the API endpoint

	for the get websocket kubernetes shell operation.

	Typically these are written to a http.Request.
*/
type GetWebsocketKubernetesShellParams struct {

	/* EndpointID.

	   environment(endpoint) ID of the environment(endpoint) where the resource is located
	*/
	EndpointID int64

	/* Token.

	   JWT token used for authentication against this environment(endpoint)
	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get websocket kubernetes shell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsocketKubernetesShellParams) WithDefaults() *GetWebsocketKubernetesShellParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get websocket kubernetes shell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsocketKubernetesShellParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) WithTimeout(timeout time.Duration) *GetWebsocketKubernetesShellParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) WithContext(ctx context.Context) *GetWebsocketKubernetesShellParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) WithHTTPClient(client *http.Client) *GetWebsocketKubernetesShellParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointID adds the endpointID to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) WithEndpointID(endpointID int64) *GetWebsocketKubernetesShellParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WithToken adds the token to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) WithToken(token string) *GetWebsocketKubernetesShellParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get websocket kubernetes shell params
func (o *GetWebsocketKubernetesShellParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsocketKubernetesShellParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endpointId
	qrEndpointID := o.EndpointID
	qEndpointID := swag.FormatInt64(qrEndpointID)
	if qEndpointID != "" {

		if err := r.SetQueryParam("endpointId", qEndpointID); err != nil {
			return err
		}
	}

	// query param token
	qrToken := o.Token
	qToken := qrToken
	if qToken != "" {

		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
