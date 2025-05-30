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

// NewCreateKubernetesServiceParams creates a new CreateKubernetesServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKubernetesServiceParams() *CreateKubernetesServiceParams {
	return &CreateKubernetesServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKubernetesServiceParamsWithTimeout creates a new CreateKubernetesServiceParams object
// with the ability to set a timeout on a request.
func NewCreateKubernetesServiceParamsWithTimeout(timeout time.Duration) *CreateKubernetesServiceParams {
	return &CreateKubernetesServiceParams{
		timeout: timeout,
	}
}

// NewCreateKubernetesServiceParamsWithContext creates a new CreateKubernetesServiceParams object
// with the ability to set a context for a request.
func NewCreateKubernetesServiceParamsWithContext(ctx context.Context) *CreateKubernetesServiceParams {
	return &CreateKubernetesServiceParams{
		Context: ctx,
	}
}

// NewCreateKubernetesServiceParamsWithHTTPClient creates a new CreateKubernetesServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKubernetesServiceParamsWithHTTPClient(client *http.Client) *CreateKubernetesServiceParams {
	return &CreateKubernetesServiceParams{
		HTTPClient: client,
	}
}

/*
CreateKubernetesServiceParams contains all the parameters to send to the API endpoint

	for the create kubernetes service operation.

	Typically these are written to a http.Request.
*/
type CreateKubernetesServiceParams struct {

	/* Body.

	   Service definition
	*/
	Body *models.KubernetesK8sServiceInfo

	/* ID.

	   Environment identifier
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

// WithDefaults hydrates default values in the create kubernetes service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKubernetesServiceParams) WithDefaults() *CreateKubernetesServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create kubernetes service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKubernetesServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithTimeout(timeout time.Duration) *CreateKubernetesServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithContext(ctx context.Context) *CreateKubernetesServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithHTTPClient(client *http.Client) *CreateKubernetesServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithBody(body *models.KubernetesK8sServiceInfo) *CreateKubernetesServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetBody(body *models.KubernetesK8sServiceInfo) {
	o.Body = body
}

// WithID adds the id to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithID(id int64) *CreateKubernetesServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetID(id int64) {
	o.ID = id
}

// WithNamespace adds the namespace to the create kubernetes service params
func (o *CreateKubernetesServiceParams) WithNamespace(namespace string) *CreateKubernetesServiceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create kubernetes service params
func (o *CreateKubernetesServiceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKubernetesServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
