// Code generated by go-swagger; DO NOT EDIT.

package helm

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

	"github.com/portainer/client-api/models"
)

// NewHelmUserRepositoryCreateParams creates a new HelmUserRepositoryCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHelmUserRepositoryCreateParams() *HelmUserRepositoryCreateParams {
	return &HelmUserRepositoryCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHelmUserRepositoryCreateParamsWithTimeout creates a new HelmUserRepositoryCreateParams object
// with the ability to set a timeout on a request.
func NewHelmUserRepositoryCreateParamsWithTimeout(timeout time.Duration) *HelmUserRepositoryCreateParams {
	return &HelmUserRepositoryCreateParams{
		timeout: timeout,
	}
}

// NewHelmUserRepositoryCreateParamsWithContext creates a new HelmUserRepositoryCreateParams object
// with the ability to set a context for a request.
func NewHelmUserRepositoryCreateParamsWithContext(ctx context.Context) *HelmUserRepositoryCreateParams {
	return &HelmUserRepositoryCreateParams{
		Context: ctx,
	}
}

// NewHelmUserRepositoryCreateParamsWithHTTPClient creates a new HelmUserRepositoryCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHelmUserRepositoryCreateParamsWithHTTPClient(client *http.Client) *HelmUserRepositoryCreateParams {
	return &HelmUserRepositoryCreateParams{
		HTTPClient: client,
	}
}

/* HelmUserRepositoryCreateParams contains all the parameters to send to the API endpoint
   for the helm user repository create operation.

   Typically these are written to a http.Request.
*/
type HelmUserRepositoryCreateParams struct {

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	/* Payload.

	   Helm Repository
	*/
	Payload *models.HelmAddHelmRepoURLPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the helm user repository create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmUserRepositoryCreateParams) WithDefaults() *HelmUserRepositoryCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the helm user repository create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmUserRepositoryCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) WithTimeout(timeout time.Duration) *HelmUserRepositoryCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) WithContext(ctx context.Context) *HelmUserRepositoryCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) WithHTTPClient(client *http.Client) *HelmUserRepositoryCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) WithID(id int64) *HelmUserRepositoryCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) SetID(id int64) {
	o.ID = id
}

// WithPayload adds the payload to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) WithPayload(payload *models.HelmAddHelmRepoURLPayload) *HelmUserRepositoryCreateParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the helm user repository create params
func (o *HelmUserRepositoryCreateParams) SetPayload(payload *models.HelmAddHelmRepoURLPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *HelmUserRepositoryCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
