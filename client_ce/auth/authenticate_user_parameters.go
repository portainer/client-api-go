// Code generated by go-swagger; DO NOT EDIT.

package auth

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

	"github.com/portainer/client-api-go/v2/models_ce"
)

// NewAuthenticateUserParams creates a new AuthenticateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticateUserParams() *AuthenticateUserParams {
	return &AuthenticateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticateUserParamsWithTimeout creates a new AuthenticateUserParams object
// with the ability to set a timeout on a request.
func NewAuthenticateUserParamsWithTimeout(timeout time.Duration) *AuthenticateUserParams {
	return &AuthenticateUserParams{
		timeout: timeout,
	}
}

// NewAuthenticateUserParamsWithContext creates a new AuthenticateUserParams object
// with the ability to set a context for a request.
func NewAuthenticateUserParamsWithContext(ctx context.Context) *AuthenticateUserParams {
	return &AuthenticateUserParams{
		Context: ctx,
	}
}

// NewAuthenticateUserParamsWithHTTPClient creates a new AuthenticateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticateUserParamsWithHTTPClient(client *http.Client) *AuthenticateUserParams {
	return &AuthenticateUserParams{
		HTTPClient: client,
	}
}

/*
AuthenticateUserParams contains all the parameters to send to the API endpoint

	for the authenticate user operation.

	Typically these are written to a http.Request.
*/
type AuthenticateUserParams struct {

	/* Body.

	   Credentials used for authentication
	*/
	Body *models_ce.AuthAuthenticatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authenticate user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateUserParams) WithDefaults() *AuthenticateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authenticate user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authenticate user params
func (o *AuthenticateUserParams) WithTimeout(timeout time.Duration) *AuthenticateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authenticate user params
func (o *AuthenticateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authenticate user params
func (o *AuthenticateUserParams) WithContext(ctx context.Context) *AuthenticateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authenticate user params
func (o *AuthenticateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authenticate user params
func (o *AuthenticateUserParams) WithHTTPClient(client *http.Client) *AuthenticateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authenticate user params
func (o *AuthenticateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the authenticate user params
func (o *AuthenticateUserParams) WithBody(body *models_ce.AuthAuthenticatePayload) *AuthenticateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authenticate user params
func (o *AuthenticateUserParams) SetBody(body *models_ce.AuthAuthenticatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
