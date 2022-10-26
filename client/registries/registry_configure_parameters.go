// Code generated by go-swagger; DO NOT EDIT.

package registries

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

	"github.com/portainer/client-api-go/v2/models"
)

// NewRegistryConfigureParams creates a new RegistryConfigureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegistryConfigureParams() *RegistryConfigureParams {
	return &RegistryConfigureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegistryConfigureParamsWithTimeout creates a new RegistryConfigureParams object
// with the ability to set a timeout on a request.
func NewRegistryConfigureParamsWithTimeout(timeout time.Duration) *RegistryConfigureParams {
	return &RegistryConfigureParams{
		timeout: timeout,
	}
}

// NewRegistryConfigureParamsWithContext creates a new RegistryConfigureParams object
// with the ability to set a context for a request.
func NewRegistryConfigureParamsWithContext(ctx context.Context) *RegistryConfigureParams {
	return &RegistryConfigureParams{
		Context: ctx,
	}
}

// NewRegistryConfigureParamsWithHTTPClient creates a new RegistryConfigureParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegistryConfigureParamsWithHTTPClient(client *http.Client) *RegistryConfigureParams {
	return &RegistryConfigureParams{
		HTTPClient: client,
	}
}

/*
RegistryConfigureParams contains all the parameters to send to the API endpoint

	for the registry configure operation.

	Typically these are written to a http.Request.
*/
type RegistryConfigureParams struct {

	/* Body.

	   Registry configuration
	*/
	Body *models.RegistriesRegistryConfigurePayload

	/* ID.

	   Registry identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the registry configure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegistryConfigureParams) WithDefaults() *RegistryConfigureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the registry configure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegistryConfigureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the registry configure params
func (o *RegistryConfigureParams) WithTimeout(timeout time.Duration) *RegistryConfigureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the registry configure params
func (o *RegistryConfigureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the registry configure params
func (o *RegistryConfigureParams) WithContext(ctx context.Context) *RegistryConfigureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the registry configure params
func (o *RegistryConfigureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the registry configure params
func (o *RegistryConfigureParams) WithHTTPClient(client *http.Client) *RegistryConfigureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the registry configure params
func (o *RegistryConfigureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the registry configure params
func (o *RegistryConfigureParams) WithBody(body *models.RegistriesRegistryConfigurePayload) *RegistryConfigureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the registry configure params
func (o *RegistryConfigureParams) SetBody(body *models.RegistriesRegistryConfigurePayload) {
	o.Body = body
}

// WithID adds the id to the registry configure params
func (o *RegistryConfigureParams) WithID(id int64) *RegistryConfigureParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the registry configure params
func (o *RegistryConfigureParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RegistryConfigureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
