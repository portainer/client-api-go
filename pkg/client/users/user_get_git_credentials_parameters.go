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

// NewUserGetGitCredentialsParams creates a new UserGetGitCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGetGitCredentialsParams() *UserGetGitCredentialsParams {
	return &UserGetGitCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetGitCredentialsParamsWithTimeout creates a new UserGetGitCredentialsParams object
// with the ability to set a timeout on a request.
func NewUserGetGitCredentialsParamsWithTimeout(timeout time.Duration) *UserGetGitCredentialsParams {
	return &UserGetGitCredentialsParams{
		timeout: timeout,
	}
}

// NewUserGetGitCredentialsParamsWithContext creates a new UserGetGitCredentialsParams object
// with the ability to set a context for a request.
func NewUserGetGitCredentialsParamsWithContext(ctx context.Context) *UserGetGitCredentialsParams {
	return &UserGetGitCredentialsParams{
		Context: ctx,
	}
}

// NewUserGetGitCredentialsParamsWithHTTPClient creates a new UserGetGitCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGetGitCredentialsParamsWithHTTPClient(client *http.Client) *UserGetGitCredentialsParams {
	return &UserGetGitCredentialsParams{
		HTTPClient: client,
	}
}

/*
UserGetGitCredentialsParams contains all the parameters to send to the API endpoint

	for the user get git credentials operation.

	Typically these are written to a http.Request.
*/
type UserGetGitCredentialsParams struct {

	/* ID.

	   User identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user get git credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGetGitCredentialsParams) WithDefaults() *UserGetGitCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user get git credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGetGitCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user get git credentials params
func (o *UserGetGitCredentialsParams) WithTimeout(timeout time.Duration) *UserGetGitCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get git credentials params
func (o *UserGetGitCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get git credentials params
func (o *UserGetGitCredentialsParams) WithContext(ctx context.Context) *UserGetGitCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get git credentials params
func (o *UserGetGitCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get git credentials params
func (o *UserGetGitCredentialsParams) WithHTTPClient(client *http.Client) *UserGetGitCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get git credentials params
func (o *UserGetGitCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user get git credentials params
func (o *UserGetGitCredentialsParams) WithID(id int64) *UserGetGitCredentialsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user get git credentials params
func (o *UserGetGitCredentialsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetGitCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
