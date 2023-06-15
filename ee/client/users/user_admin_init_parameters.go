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

	"github.com/portainer/client-api-go/ee/v2/models"
)

// NewUserAdminInitParams creates a new UserAdminInitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserAdminInitParams() *UserAdminInitParams {
	return &UserAdminInitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserAdminInitParamsWithTimeout creates a new UserAdminInitParams object
// with the ability to set a timeout on a request.
func NewUserAdminInitParamsWithTimeout(timeout time.Duration) *UserAdminInitParams {
	return &UserAdminInitParams{
		timeout: timeout,
	}
}

// NewUserAdminInitParamsWithContext creates a new UserAdminInitParams object
// with the ability to set a context for a request.
func NewUserAdminInitParamsWithContext(ctx context.Context) *UserAdminInitParams {
	return &UserAdminInitParams{
		Context: ctx,
	}
}

// NewUserAdminInitParamsWithHTTPClient creates a new UserAdminInitParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserAdminInitParamsWithHTTPClient(client *http.Client) *UserAdminInitParams {
	return &UserAdminInitParams{
		HTTPClient: client,
	}
}

/*
UserAdminInitParams contains all the parameters to send to the API endpoint

	for the user admin init operation.

	Typically these are written to a http.Request.
*/
type UserAdminInitParams struct {

	/* Body.

	   User details
	*/
	Body *models.UsersAdminInitPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user admin init params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAdminInitParams) WithDefaults() *UserAdminInitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user admin init params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAdminInitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user admin init params
func (o *UserAdminInitParams) WithTimeout(timeout time.Duration) *UserAdminInitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user admin init params
func (o *UserAdminInitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user admin init params
func (o *UserAdminInitParams) WithContext(ctx context.Context) *UserAdminInitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user admin init params
func (o *UserAdminInitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user admin init params
func (o *UserAdminInitParams) WithHTTPClient(client *http.Client) *UserAdminInitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user admin init params
func (o *UserAdminInitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user admin init params
func (o *UserAdminInitParams) WithBody(body *models.UsersAdminInitPayload) *UserAdminInitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user admin init params
func (o *UserAdminInitParams) SetBody(body *models.UsersAdminInitPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserAdminInitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
