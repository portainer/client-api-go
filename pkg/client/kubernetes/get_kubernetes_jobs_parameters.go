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

// NewGetKubernetesJobsParams creates a new GetKubernetesJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesJobsParams() *GetKubernetesJobsParams {
	return &GetKubernetesJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesJobsParamsWithTimeout creates a new GetKubernetesJobsParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesJobsParamsWithTimeout(timeout time.Duration) *GetKubernetesJobsParams {
	return &GetKubernetesJobsParams{
		timeout: timeout,
	}
}

// NewGetKubernetesJobsParamsWithContext creates a new GetKubernetesJobsParams object
// with the ability to set a context for a request.
func NewGetKubernetesJobsParamsWithContext(ctx context.Context) *GetKubernetesJobsParams {
	return &GetKubernetesJobsParams{
		Context: ctx,
	}
}

// NewGetKubernetesJobsParamsWithHTTPClient creates a new GetKubernetesJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesJobsParamsWithHTTPClient(client *http.Client) *GetKubernetesJobsParams {
	return &GetKubernetesJobsParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesJobsParams contains all the parameters to send to the API endpoint

	for the get kubernetes jobs operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesJobsParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* IncludeCronJobChildren.

	   Whether to include Jobs that have a cronjob owner
	*/
	IncludeCronJobChildren *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesJobsParams) WithDefaults() *GetKubernetesJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) WithTimeout(timeout time.Duration) *GetKubernetesJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) WithContext(ctx context.Context) *GetKubernetesJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) WithHTTPClient(client *http.Client) *GetKubernetesJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) WithID(id int64) *GetKubernetesJobsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) SetID(id int64) {
	o.ID = id
}

// WithIncludeCronJobChildren adds the includeCronJobChildren to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) WithIncludeCronJobChildren(includeCronJobChildren *bool) *GetKubernetesJobsParams {
	o.SetIncludeCronJobChildren(includeCronJobChildren)
	return o
}

// SetIncludeCronJobChildren adds the includeCronJobChildren to the get kubernetes jobs params
func (o *GetKubernetesJobsParams) SetIncludeCronJobChildren(includeCronJobChildren *bool) {
	o.IncludeCronJobChildren = includeCronJobChildren
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.IncludeCronJobChildren != nil {

		// query param includeCronJobChildren
		var qrIncludeCronJobChildren bool

		if o.IncludeCronJobChildren != nil {
			qrIncludeCronJobChildren = *o.IncludeCronJobChildren
		}
		qIncludeCronJobChildren := swag.FormatBool(qrIncludeCronJobChildren)
		if qIncludeCronJobChildren != "" {

			if err := r.SetQueryParam("includeCronJobChildren", qIncludeCronJobChildren); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
