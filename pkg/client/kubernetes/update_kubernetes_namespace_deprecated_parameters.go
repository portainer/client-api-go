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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewUpdateKubernetesNamespaceDeprecatedParams creates a new UpdateKubernetesNamespaceDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateKubernetesNamespaceDeprecatedParams() *UpdateKubernetesNamespaceDeprecatedParams {
	return &UpdateKubernetesNamespaceDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKubernetesNamespaceDeprecatedParamsWithTimeout creates a new UpdateKubernetesNamespaceDeprecatedParams object
// with the ability to set a timeout on a request.
func NewUpdateKubernetesNamespaceDeprecatedParamsWithTimeout(timeout time.Duration) *UpdateKubernetesNamespaceDeprecatedParams {
	return &UpdateKubernetesNamespaceDeprecatedParams{
		timeout: timeout,
	}
}

// NewUpdateKubernetesNamespaceDeprecatedParamsWithContext creates a new UpdateKubernetesNamespaceDeprecatedParams object
// with the ability to set a context for a request.
func NewUpdateKubernetesNamespaceDeprecatedParamsWithContext(ctx context.Context) *UpdateKubernetesNamespaceDeprecatedParams {
	return &UpdateKubernetesNamespaceDeprecatedParams{
		Context: ctx,
	}
}

// NewUpdateKubernetesNamespaceDeprecatedParamsWithHTTPClient creates a new UpdateKubernetesNamespaceDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateKubernetesNamespaceDeprecatedParamsWithHTTPClient(client *http.Client) *UpdateKubernetesNamespaceDeprecatedParams {
	return &UpdateKubernetesNamespaceDeprecatedParams{
		HTTPClient: client,
	}
}

/*
UpdateKubernetesNamespaceDeprecatedParams contains all the parameters to send to the API endpoint

	for the update kubernetes namespace deprecated operation.

	Typically these are written to a http.Request.
*/
type UpdateKubernetesNamespaceDeprecatedParams struct {

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

// WithDefaults hydrates default values in the update kubernetes namespace deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithDefaults() *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update kubernetes namespace deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithTimeout(timeout time.Duration) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithContext(ctx context.Context) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithHTTPClient(client *http.Client) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithBody(body *models.ModelsK8sNamespaceDetails) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetBody(body *models.ModelsK8sNamespaceDetails) {
	o.Body = body
}

// WithID adds the id to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithID(id int64) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetID(id int64) {
	o.ID = id
}

// WithNamespace adds the namespace to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) WithNamespace(namespace string) *UpdateKubernetesNamespaceDeprecatedParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update kubernetes namespace deprecated params
func (o *UpdateKubernetesNamespaceDeprecatedParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKubernetesNamespaceDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
