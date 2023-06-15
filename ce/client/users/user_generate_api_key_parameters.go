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

	"github.com/portainer/client-api-go/ce/v2/models"
)

// NewUserGenerateAPIKeyParams creates a new UserGenerateAPIKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGenerateAPIKeyParams() *UserGenerateAPIKeyParams {
	return &UserGenerateAPIKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGenerateAPIKeyParamsWithTimeout creates a new UserGenerateAPIKeyParams object
// with the ability to set a timeout on a request.
func NewUserGenerateAPIKeyParamsWithTimeout(timeout time.Duration) *UserGenerateAPIKeyParams {
	return &UserGenerateAPIKeyParams{
		timeout: timeout,
	}
}

// NewUserGenerateAPIKeyParamsWithContext creates a new UserGenerateAPIKeyParams object
// with the ability to set a context for a request.
func NewUserGenerateAPIKeyParamsWithContext(ctx context.Context) *UserGenerateAPIKeyParams {
	return &UserGenerateAPIKeyParams{
		Context: ctx,
	}
}

// NewUserGenerateAPIKeyParamsWithHTTPClient creates a new UserGenerateAPIKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGenerateAPIKeyParamsWithHTTPClient(client *http.Client) *UserGenerateAPIKeyParams {
	return &UserGenerateAPIKeyParams{
		HTTPClient: client,
	}
}

/*
UserGenerateAPIKeyParams contains all the parameters to send to the API endpoint

	for the user generate API key operation.

	Typically these are written to a http.Request.
*/
type UserGenerateAPIKeyParams struct {

	/* Body.

	   details
	*/
	Body *models.UsersUserAccessTokenCreatePayload

	/* ID.

	   User identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user generate API key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGenerateAPIKeyParams) WithDefaults() *UserGenerateAPIKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user generate API key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGenerateAPIKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user generate API key params
func (o *UserGenerateAPIKeyParams) WithTimeout(timeout time.Duration) *UserGenerateAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user generate API key params
func (o *UserGenerateAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user generate API key params
func (o *UserGenerateAPIKeyParams) WithContext(ctx context.Context) *UserGenerateAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user generate API key params
func (o *UserGenerateAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user generate API key params
func (o *UserGenerateAPIKeyParams) WithHTTPClient(client *http.Client) *UserGenerateAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user generate API key params
func (o *UserGenerateAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user generate API key params
func (o *UserGenerateAPIKeyParams) WithBody(body *models.UsersUserAccessTokenCreatePayload) *UserGenerateAPIKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user generate API key params
func (o *UserGenerateAPIKeyParams) SetBody(body *models.UsersUserAccessTokenCreatePayload) {
	o.Body = body
}

// WithID adds the id to the user generate API key params
func (o *UserGenerateAPIKeyParams) WithID(id int64) *UserGenerateAPIKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user generate API key params
func (o *UserGenerateAPIKeyParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserGenerateAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
