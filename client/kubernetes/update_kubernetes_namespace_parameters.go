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

// NewUpdateKubernetesNamespaceParams creates a new UpdateKubernetesNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateKubernetesNamespaceParams() *UpdateKubernetesNamespaceParams {
	return &UpdateKubernetesNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKubernetesNamespaceParamsWithTimeout creates a new UpdateKubernetesNamespaceParams object
// with the ability to set a timeout on a request.
func NewUpdateKubernetesNamespaceParamsWithTimeout(timeout time.Duration) *UpdateKubernetesNamespaceParams {
	return &UpdateKubernetesNamespaceParams{
		timeout: timeout,
	}
}

// NewUpdateKubernetesNamespaceParamsWithContext creates a new UpdateKubernetesNamespaceParams object
// with the ability to set a context for a request.
func NewUpdateKubernetesNamespaceParamsWithContext(ctx context.Context) *UpdateKubernetesNamespaceParams {
	return &UpdateKubernetesNamespaceParams{
		Context: ctx,
	}
}

// NewUpdateKubernetesNamespaceParamsWithHTTPClient creates a new UpdateKubernetesNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateKubernetesNamespaceParamsWithHTTPClient(client *http.Client) *UpdateKubernetesNamespaceParams {
	return &UpdateKubernetesNamespaceParams{
		HTTPClient: client,
	}
}

/*
UpdateKubernetesNamespaceParams contains all the parameters to send to the API endpoint

	for the update kubernetes namespace operation.

	Typically these are written to a http.Request.
*/
type UpdateKubernetesNamespaceParams struct {

	/* Body.

	   Namespace details
	*/
	Body *models.ModelsK8sNamespaceDetails

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* Namespace.

	   Namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update kubernetes namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesNamespaceParams) WithDefaults() *UpdateKubernetesNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update kubernetes namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesNamespaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithTimeout(timeout time.Duration) *UpdateKubernetesNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithContext(ctx context.Context) *UpdateKubernetesNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithHTTPClient(client *http.Client) *UpdateKubernetesNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithBody(body *models.ModelsK8sNamespaceDetails) *UpdateKubernetesNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetBody(body *models.ModelsK8sNamespaceDetails) {
	o.Body = body
}

// WithID adds the id to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithID(id int64) *UpdateKubernetesNamespaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetID(id int64) {
	o.ID = id
}

// WithNamespace adds the namespace to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) WithNamespace(namespace string) *UpdateKubernetesNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update kubernetes namespace params
func (o *UpdateKubernetesNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKubernetesNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
