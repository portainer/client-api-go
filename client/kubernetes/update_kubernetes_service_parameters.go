// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewUpdateKubernetesServiceParams creates a new UpdateKubernetesServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateKubernetesServiceParams() *UpdateKubernetesServiceParams {
	return &UpdateKubernetesServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKubernetesServiceParamsWithTimeout creates a new UpdateKubernetesServiceParams object
// with the ability to set a timeout on a request.
func NewUpdateKubernetesServiceParamsWithTimeout(timeout time.Duration) *UpdateKubernetesServiceParams {
	return &UpdateKubernetesServiceParams{
		timeout: timeout,
	}
}

// NewUpdateKubernetesServiceParamsWithContext creates a new UpdateKubernetesServiceParams object
// with the ability to set a context for a request.
func NewUpdateKubernetesServiceParamsWithContext(ctx context.Context) *UpdateKubernetesServiceParams {
	return &UpdateKubernetesServiceParams{
		Context: ctx,
	}
}

// NewUpdateKubernetesServiceParamsWithHTTPClient creates a new UpdateKubernetesServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateKubernetesServiceParamsWithHTTPClient(client *http.Client) *UpdateKubernetesServiceParams {
	return &UpdateKubernetesServiceParams{
		HTTPClient: client,
	}
}

/*
UpdateKubernetesServiceParams contains all the parameters to send to the API endpoint

	for the update kubernetes service operation.

	Typically these are written to a http.Request.
*/
type UpdateKubernetesServiceParams struct {

	/* Body.

	   Service definition
	*/
	Body *models.ModelsK8sServiceInfo

	/* ID.

	   Environment (Endpoint) identifier
	*/
	ID int64

	/* Namespace.

	   Namespace name
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update kubernetes service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesServiceParams) WithDefaults() *UpdateKubernetesServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update kubernetes service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithTimeout(timeout time.Duration) *UpdateKubernetesServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithContext(ctx context.Context) *UpdateKubernetesServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithHTTPClient(client *http.Client) *UpdateKubernetesServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithBody(body *models.ModelsK8sServiceInfo) *UpdateKubernetesServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetBody(body *models.ModelsK8sServiceInfo) {
	o.Body = body
}

// WithID adds the id to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithID(id int64) *UpdateKubernetesServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetID(id int64) {
	o.ID = id
}

// WithNamespace adds the namespace to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) WithNamespace(namespace string) *UpdateKubernetesServiceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update kubernetes service params
func (o *UpdateKubernetesServiceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKubernetesServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
