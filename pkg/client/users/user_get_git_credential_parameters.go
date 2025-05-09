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

// NewUserGetGitCredentialParams creates a new UserGetGitCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGetGitCredentialParams() *UserGetGitCredentialParams {
	return &UserGetGitCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetGitCredentialParamsWithTimeout creates a new UserGetGitCredentialParams object
// with the ability to set a timeout on a request.
func NewUserGetGitCredentialParamsWithTimeout(timeout time.Duration) *UserGetGitCredentialParams {
	return &UserGetGitCredentialParams{
		timeout: timeout,
	}
}

// NewUserGetGitCredentialParamsWithContext creates a new UserGetGitCredentialParams object
// with the ability to set a context for a request.
func NewUserGetGitCredentialParamsWithContext(ctx context.Context) *UserGetGitCredentialParams {
	return &UserGetGitCredentialParams{
		Context: ctx,
	}
}

// NewUserGetGitCredentialParamsWithHTTPClient creates a new UserGetGitCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGetGitCredentialParamsWithHTTPClient(client *http.Client) *UserGetGitCredentialParams {
	return &UserGetGitCredentialParams{
		HTTPClient: client,
	}
}

/*
UserGetGitCredentialParams contains all the parameters to send to the API endpoint

	for the user get git credential operation.

	Typically these are written to a http.Request.
*/
type UserGetGitCredentialParams struct {

	/* CredentialID.

	   Git Credential identifier
	*/
	CredentialID int64

	/* ID.

	   User identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user get git credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGetGitCredentialParams) WithDefaults() *UserGetGitCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user get git credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGetGitCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user get git credential params
func (o *UserGetGitCredentialParams) WithTimeout(timeout time.Duration) *UserGetGitCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get git credential params
func (o *UserGetGitCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get git credential params
func (o *UserGetGitCredentialParams) WithContext(ctx context.Context) *UserGetGitCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get git credential params
func (o *UserGetGitCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get git credential params
func (o *UserGetGitCredentialParams) WithHTTPClient(client *http.Client) *UserGetGitCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get git credential params
func (o *UserGetGitCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the user get git credential params
func (o *UserGetGitCredentialParams) WithCredentialID(credentialID int64) *UserGetGitCredentialParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the user get git credential params
func (o *UserGetGitCredentialParams) SetCredentialID(credentialID int64) {
	o.CredentialID = credentialID
}

// WithID adds the id to the user get git credential params
func (o *UserGetGitCredentialParams) WithID(id int64) *UserGetGitCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user get git credential params
func (o *UserGetGitCredentialParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetGitCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credentialID
	if err := r.SetPathParam("credentialID", swag.FormatInt64(o.CredentialID)); err != nil {
		return err
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
