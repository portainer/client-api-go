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

// NewGetKubernetesServicesParams creates a new GetKubernetesServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesServicesParams() *GetKubernetesServicesParams {
	return &GetKubernetesServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesServicesParamsWithTimeout creates a new GetKubernetesServicesParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesServicesParamsWithTimeout(timeout time.Duration) *GetKubernetesServicesParams {
	return &GetKubernetesServicesParams{
		timeout: timeout,
	}
}

// NewGetKubernetesServicesParamsWithContext creates a new GetKubernetesServicesParams object
// with the ability to set a context for a request.
func NewGetKubernetesServicesParamsWithContext(ctx context.Context) *GetKubernetesServicesParams {
	return &GetKubernetesServicesParams{
		Context: ctx,
	}
}

// NewGetKubernetesServicesParamsWithHTTPClient creates a new GetKubernetesServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesServicesParamsWithHTTPClient(client *http.Client) *GetKubernetesServicesParams {
	return &GetKubernetesServicesParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesServicesParams contains all the parameters to send to the API endpoint

	for the get kubernetes services operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesServicesParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* WithApplications.

	   Lookup applications associated with each service
	*/
	WithApplications *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesServicesParams) WithDefaults() *GetKubernetesServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes services params
func (o *GetKubernetesServicesParams) WithTimeout(timeout time.Duration) *GetKubernetesServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes services params
func (o *GetKubernetesServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes services params
func (o *GetKubernetesServicesParams) WithContext(ctx context.Context) *GetKubernetesServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes services params
func (o *GetKubernetesServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes services params
func (o *GetKubernetesServicesParams) WithHTTPClient(client *http.Client) *GetKubernetesServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes services params
func (o *GetKubernetesServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes services params
func (o *GetKubernetesServicesParams) WithID(id int64) *GetKubernetesServicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes services params
func (o *GetKubernetesServicesParams) SetID(id int64) {
	o.ID = id
}

// WithWithApplications adds the withApplications to the get kubernetes services params
func (o *GetKubernetesServicesParams) WithWithApplications(withApplications *bool) *GetKubernetesServicesParams {
	o.SetWithApplications(withApplications)
	return o
}

// SetWithApplications adds the withApplications to the get kubernetes services params
func (o *GetKubernetesServicesParams) SetWithApplications(withApplications *bool) {
	o.WithApplications = withApplications
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.WithApplications != nil {

		// query param withApplications
		var qrWithApplications bool

		if o.WithApplications != nil {
			qrWithApplications = *o.WithApplications
		}
		qWithApplications := swag.FormatBool(qrWithApplications)
		if qWithApplications != "" {

			if err := r.SetQueryParam("withApplications", qWithApplications); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
