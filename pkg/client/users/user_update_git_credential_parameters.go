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

// NewUserUpdateGitCredentialParams creates a new UserUpdateGitCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserUpdateGitCredentialParams() *UserUpdateGitCredentialParams {
	return &UserUpdateGitCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateGitCredentialParamsWithTimeout creates a new UserUpdateGitCredentialParams object
// with the ability to set a timeout on a request.
func NewUserUpdateGitCredentialParamsWithTimeout(timeout time.Duration) *UserUpdateGitCredentialParams {
	return &UserUpdateGitCredentialParams{
		timeout: timeout,
	}
}

// NewUserUpdateGitCredentialParamsWithContext creates a new UserUpdateGitCredentialParams object
// with the ability to set a context for a request.
func NewUserUpdateGitCredentialParamsWithContext(ctx context.Context) *UserUpdateGitCredentialParams {
	return &UserUpdateGitCredentialParams{
		Context: ctx,
	}
}

// NewUserUpdateGitCredentialParamsWithHTTPClient creates a new UserUpdateGitCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserUpdateGitCredentialParamsWithHTTPClient(client *http.Client) *UserUpdateGitCredentialParams {
	return &UserUpdateGitCredentialParams{
		HTTPClient: client,
	}
}

/*
UserUpdateGitCredentialParams contains all the parameters to send to the API endpoint

	for the user update git credential operation.

	Typically these are written to a http.Request.
*/
type UserUpdateGitCredentialParams struct {

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

// WithDefaults hydrates default values in the user update git credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserUpdateGitCredentialParams) WithDefaults() *UserUpdateGitCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user update git credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserUpdateGitCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user update git credential params
func (o *UserUpdateGitCredentialParams) WithTimeout(timeout time.Duration) *UserUpdateGitCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update git credential params
func (o *UserUpdateGitCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update git credential params
func (o *UserUpdateGitCredentialParams) WithContext(ctx context.Context) *UserUpdateGitCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update git credential params
func (o *UserUpdateGitCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update git credential params
func (o *UserUpdateGitCredentialParams) WithHTTPClient(client *http.Client) *UserUpdateGitCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update git credential params
func (o *UserUpdateGitCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the user update git credential params
func (o *UserUpdateGitCredentialParams) WithCredentialID(credentialID int64) *UserUpdateGitCredentialParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the user update git credential params
func (o *UserUpdateGitCredentialParams) SetCredentialID(credentialID int64) {
	o.CredentialID = credentialID
}

// WithID adds the id to the user update git credential params
func (o *UserUpdateGitCredentialParams) WithID(id int64) *UserUpdateGitCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user update git credential params
func (o *UserUpdateGitCredentialParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateGitCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
