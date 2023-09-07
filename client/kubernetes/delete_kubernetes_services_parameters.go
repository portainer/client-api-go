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

// NewDeleteKubernetesServicesParams creates a new DeleteKubernetesServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKubernetesServicesParams() *DeleteKubernetesServicesParams {
	return &DeleteKubernetesServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKubernetesServicesParamsWithTimeout creates a new DeleteKubernetesServicesParams object
// with the ability to set a timeout on a request.
func NewDeleteKubernetesServicesParamsWithTimeout(timeout time.Duration) *DeleteKubernetesServicesParams {
	return &DeleteKubernetesServicesParams{
		timeout: timeout,
	}
}

// NewDeleteKubernetesServicesParamsWithContext creates a new DeleteKubernetesServicesParams object
// with the ability to set a context for a request.
func NewDeleteKubernetesServicesParamsWithContext(ctx context.Context) *DeleteKubernetesServicesParams {
	return &DeleteKubernetesServicesParams{
		Context: ctx,
	}
}

// NewDeleteKubernetesServicesParamsWithHTTPClient creates a new DeleteKubernetesServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKubernetesServicesParamsWithHTTPClient(client *http.Client) *DeleteKubernetesServicesParams {
	return &DeleteKubernetesServicesParams{
		HTTPClient: client,
	}
}

/*
DeleteKubernetesServicesParams contains all the parameters to send to the API endpoint

	for the delete kubernetes services operation.

	Typically these are written to a http.Request.
*/
type DeleteKubernetesServicesParams struct {

	/* Body.

	   A map where the key is the namespace and the value is an array of services to delete
	*/
	Body models.ModelsK8sServiceDeleteRequests

	/* ID.

	   Environment (Endpoint) identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete kubernetes services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKubernetesServicesParams) WithDefaults() *DeleteKubernetesServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete kubernetes services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKubernetesServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) WithTimeout(timeout time.Duration) *DeleteKubernetesServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) WithContext(ctx context.Context) *DeleteKubernetesServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) WithHTTPClient(client *http.Client) *DeleteKubernetesServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) WithBody(body models.ModelsK8sServiceDeleteRequests) *DeleteKubernetesServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) SetBody(body models.ModelsK8sServiceDeleteRequests) {
	o.Body = body
}

// WithID adds the id to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) WithID(id int64) *DeleteKubernetesServicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete kubernetes services params
func (o *DeleteKubernetesServicesParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKubernetesServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
