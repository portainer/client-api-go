// Code generated by go-swagger; DO NOT EDIT.

package settings

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

	"github.com/portainer/client-api/models"
)

// NewSettingsUpdateParams creates a new SettingsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsUpdateParams() *SettingsUpdateParams {
	return &SettingsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsUpdateParamsWithTimeout creates a new SettingsUpdateParams object
// with the ability to set a timeout on a request.
func NewSettingsUpdateParamsWithTimeout(timeout time.Duration) *SettingsUpdateParams {
	return &SettingsUpdateParams{
		timeout: timeout,
	}
}

// NewSettingsUpdateParamsWithContext creates a new SettingsUpdateParams object
// with the ability to set a context for a request.
func NewSettingsUpdateParamsWithContext(ctx context.Context) *SettingsUpdateParams {
	return &SettingsUpdateParams{
		Context: ctx,
	}
}

// NewSettingsUpdateParamsWithHTTPClient creates a new SettingsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsUpdateParamsWithHTTPClient(client *http.Client) *SettingsUpdateParams {
	return &SettingsUpdateParams{
		HTTPClient: client,
	}
}

/* SettingsUpdateParams contains all the parameters to send to the API endpoint
   for the settings update operation.

   Typically these are written to a http.Request.
*/
type SettingsUpdateParams struct {

	/* Body.

	   New settings
	*/
	Body *models.SettingsSettingsUpdatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsUpdateParams) WithDefaults() *SettingsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings update params
func (o *SettingsUpdateParams) WithTimeout(timeout time.Duration) *SettingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings update params
func (o *SettingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings update params
func (o *SettingsUpdateParams) WithContext(ctx context.Context) *SettingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings update params
func (o *SettingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings update params
func (o *SettingsUpdateParams) WithHTTPClient(client *http.Client) *SettingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings update params
func (o *SettingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the settings update params
func (o *SettingsUpdateParams) WithBody(body *models.SettingsSettingsUpdatePayload) *SettingsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the settings update params
func (o *SettingsUpdateParams) SetBody(body *models.SettingsSettingsUpdatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
