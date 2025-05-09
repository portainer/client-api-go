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

// NewGetKubernetesMaxResourceLimitsParams creates a new GetKubernetesMaxResourceLimitsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesMaxResourceLimitsParams() *GetKubernetesMaxResourceLimitsParams {
	return &GetKubernetesMaxResourceLimitsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesMaxResourceLimitsParamsWithTimeout creates a new GetKubernetesMaxResourceLimitsParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesMaxResourceLimitsParamsWithTimeout(timeout time.Duration) *GetKubernetesMaxResourceLimitsParams {
	return &GetKubernetesMaxResourceLimitsParams{
		timeout: timeout,
	}
}

// NewGetKubernetesMaxResourceLimitsParamsWithContext creates a new GetKubernetesMaxResourceLimitsParams object
// with the ability to set a context for a request.
func NewGetKubernetesMaxResourceLimitsParamsWithContext(ctx context.Context) *GetKubernetesMaxResourceLimitsParams {
	return &GetKubernetesMaxResourceLimitsParams{
		Context: ctx,
	}
}

// NewGetKubernetesMaxResourceLimitsParamsWithHTTPClient creates a new GetKubernetesMaxResourceLimitsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesMaxResourceLimitsParamsWithHTTPClient(client *http.Client) *GetKubernetesMaxResourceLimitsParams {
	return &GetKubernetesMaxResourceLimitsParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesMaxResourceLimitsParams contains all the parameters to send to the API endpoint

	for the get kubernetes max resource limits operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesMaxResourceLimitsParams struct {

	/* ID.

	   Environment(Endpoint) identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes max resource limits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesMaxResourceLimitsParams) WithDefaults() *GetKubernetesMaxResourceLimitsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes max resource limits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesMaxResourceLimitsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) WithTimeout(timeout time.Duration) *GetKubernetesMaxResourceLimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) WithContext(ctx context.Context) *GetKubernetesMaxResourceLimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) WithHTTPClient(client *http.Client) *GetKubernetesMaxResourceLimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) WithID(id int64) *GetKubernetesMaxResourceLimitsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes max resource limits params
func (o *GetKubernetesMaxResourceLimitsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesMaxResourceLimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
