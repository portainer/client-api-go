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

// NewGetKubernetesRoleBindingsParams creates a new GetKubernetesRoleBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesRoleBindingsParams() *GetKubernetesRoleBindingsParams {
	return &GetKubernetesRoleBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesRoleBindingsParamsWithTimeout creates a new GetKubernetesRoleBindingsParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesRoleBindingsParamsWithTimeout(timeout time.Duration) *GetKubernetesRoleBindingsParams {
	return &GetKubernetesRoleBindingsParams{
		timeout: timeout,
	}
}

// NewGetKubernetesRoleBindingsParamsWithContext creates a new GetKubernetesRoleBindingsParams object
// with the ability to set a context for a request.
func NewGetKubernetesRoleBindingsParamsWithContext(ctx context.Context) *GetKubernetesRoleBindingsParams {
	return &GetKubernetesRoleBindingsParams{
		Context: ctx,
	}
}

// NewGetKubernetesRoleBindingsParamsWithHTTPClient creates a new GetKubernetesRoleBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesRoleBindingsParamsWithHTTPClient(client *http.Client) *GetKubernetesRoleBindingsParams {
	return &GetKubernetesRoleBindingsParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesRoleBindingsParams contains all the parameters to send to the API endpoint

	for the get kubernetes role bindings operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesRoleBindingsParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes role bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesRoleBindingsParams) WithDefaults() *GetKubernetesRoleBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes role bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesRoleBindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) WithTimeout(timeout time.Duration) *GetKubernetesRoleBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) WithContext(ctx context.Context) *GetKubernetesRoleBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) WithHTTPClient(client *http.Client) *GetKubernetesRoleBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) WithID(id int64) *GetKubernetesRoleBindingsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes role bindings params
func (o *GetKubernetesRoleBindingsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesRoleBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
