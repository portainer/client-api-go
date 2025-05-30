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
)

// NewRestartKubernetesApplicationParams creates a new RestartKubernetesApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartKubernetesApplicationParams() *RestartKubernetesApplicationParams {
	return &RestartKubernetesApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartKubernetesApplicationParamsWithTimeout creates a new RestartKubernetesApplicationParams object
// with the ability to set a timeout on a request.
func NewRestartKubernetesApplicationParamsWithTimeout(timeout time.Duration) *RestartKubernetesApplicationParams {
	return &RestartKubernetesApplicationParams{
		timeout: timeout,
	}
}

// NewRestartKubernetesApplicationParamsWithContext creates a new RestartKubernetesApplicationParams object
// with the ability to set a context for a request.
func NewRestartKubernetesApplicationParamsWithContext(ctx context.Context) *RestartKubernetesApplicationParams {
	return &RestartKubernetesApplicationParams{
		Context: ctx,
	}
}

// NewRestartKubernetesApplicationParamsWithHTTPClient creates a new RestartKubernetesApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartKubernetesApplicationParamsWithHTTPClient(client *http.Client) *RestartKubernetesApplicationParams {
	return &RestartKubernetesApplicationParams{
		HTTPClient: client,
	}
}

/*
RestartKubernetesApplicationParams contains all the parameters to send to the API endpoint

	for the restart kubernetes application operation.

	Typically these are written to a http.Request.
*/
type RestartKubernetesApplicationParams struct {

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	/* Kind.

	   deployment, statefulset or daemonset
	*/
	Kind string

	/* Name.

	   name of the application
	*/
	Name string

	/* Namespace.

	   The namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart kubernetes application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartKubernetesApplicationParams) WithDefaults() *RestartKubernetesApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart kubernetes application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartKubernetesApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithTimeout(timeout time.Duration) *RestartKubernetesApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithContext(ctx context.Context) *RestartKubernetesApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithHTTPClient(client *http.Client) *RestartKubernetesApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithID(id int64) *RestartKubernetesApplicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetID(id int64) {
	o.ID = id
}

// WithKind adds the kind to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithKind(kind string) *RestartKubernetesApplicationParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetKind(kind string) {
	o.Kind = kind
}

// WithName adds the name to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithName(name string) *RestartKubernetesApplicationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) WithNamespace(namespace string) *RestartKubernetesApplicationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the restart kubernetes application params
func (o *RestartKubernetesApplicationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *RestartKubernetesApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
