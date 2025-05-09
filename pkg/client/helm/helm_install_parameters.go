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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewHelmInstallParams creates a new HelmInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHelmInstallParams() *HelmInstallParams {
	return &HelmInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHelmInstallParamsWithTimeout creates a new HelmInstallParams object
// with the ability to set a timeout on a request.
func NewHelmInstallParamsWithTimeout(timeout time.Duration) *HelmInstallParams {
	return &HelmInstallParams{
		timeout: timeout,
	}
}

// NewHelmInstallParamsWithContext creates a new HelmInstallParams object
// with the ability to set a context for a request.
func NewHelmInstallParamsWithContext(ctx context.Context) *HelmInstallParams {
	return &HelmInstallParams{
		Context: ctx,
	}
}

// NewHelmInstallParamsWithHTTPClient creates a new HelmInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewHelmInstallParamsWithHTTPClient(client *http.Client) *HelmInstallParams {
	return &HelmInstallParams{
		HTTPClient: client,
	}
}

/*
HelmInstallParams contains all the parameters to send to the API endpoint

	for the helm install operation.

	Typically these are written to a http.Request.
*/
type HelmInstallParams struct {

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	/* Payload.

	   Chart details
	*/
	Payload *models.HelmInstallChartPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the helm install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmInstallParams) WithDefaults() *HelmInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the helm install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HelmInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the helm install params
func (o *HelmInstallParams) WithTimeout(timeout time.Duration) *HelmInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the helm install params
func (o *HelmInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the helm install params
func (o *HelmInstallParams) WithContext(ctx context.Context) *HelmInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the helm install params
func (o *HelmInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the helm install params
func (o *HelmInstallParams) WithHTTPClient(client *http.Client) *HelmInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the helm install params
func (o *HelmInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the helm install params
func (o *HelmInstallParams) WithID(id int64) *HelmInstallParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the helm install params
func (o *HelmInstallParams) SetID(id int64) {
	o.ID = id
}

// WithPayload adds the payload to the helm install params
func (o *HelmInstallParams) WithPayload(payload *models.HelmInstallChartPayload) *HelmInstallParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the helm install params
func (o *HelmInstallParams) SetPayload(payload *models.HelmInstallChartPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *HelmInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
