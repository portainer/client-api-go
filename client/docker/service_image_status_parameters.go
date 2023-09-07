// Code generated by go-swagger; DO NOT EDIT.

package docker

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

// NewServiceImageStatusParams creates a new ServiceImageStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceImageStatusParams() *ServiceImageStatusParams {
	return &ServiceImageStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceImageStatusParamsWithTimeout creates a new ServiceImageStatusParams object
// with the ability to set a timeout on a request.
func NewServiceImageStatusParamsWithTimeout(timeout time.Duration) *ServiceImageStatusParams {
	return &ServiceImageStatusParams{
		timeout: timeout,
	}
}

// NewServiceImageStatusParamsWithContext creates a new ServiceImageStatusParams object
// with the ability to set a context for a request.
func NewServiceImageStatusParamsWithContext(ctx context.Context) *ServiceImageStatusParams {
	return &ServiceImageStatusParams{
		Context: ctx,
	}
}

// NewServiceImageStatusParamsWithHTTPClient creates a new ServiceImageStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceImageStatusParamsWithHTTPClient(client *http.Client) *ServiceImageStatusParams {
	return &ServiceImageStatusParams{
		HTTPClient: client,
	}
}

/*
ServiceImageStatusParams contains all the parameters to send to the API endpoint

	for the service image status operation.

	Typically these are written to a http.Request.
*/
type ServiceImageStatusParams struct {

	/* EnvironmentID.

	   Environment identifier
	*/
	EnvironmentID int64

	/* ServiceID.

	   Service identifier
	*/
	ServiceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service image status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceImageStatusParams) WithDefaults() *ServiceImageStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service image status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceImageStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service image status params
func (o *ServiceImageStatusParams) WithTimeout(timeout time.Duration) *ServiceImageStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service image status params
func (o *ServiceImageStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service image status params
func (o *ServiceImageStatusParams) WithContext(ctx context.Context) *ServiceImageStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service image status params
func (o *ServiceImageStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service image status params
func (o *ServiceImageStatusParams) WithHTTPClient(client *http.Client) *ServiceImageStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service image status params
func (o *ServiceImageStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the service image status params
func (o *ServiceImageStatusParams) WithEnvironmentID(environmentID int64) *ServiceImageStatusParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the service image status params
func (o *ServiceImageStatusParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WithServiceID adds the serviceID to the service image status params
func (o *ServiceImageStatusParams) WithServiceID(serviceID int64) *ServiceImageStatusParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service image status params
func (o *ServiceImageStatusParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceImageStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	// path param serviceId
	if err := r.SetPathParam("serviceId", swag.FormatInt64(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
