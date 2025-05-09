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

// NewDeleteServiceAccountsParams creates a new DeleteServiceAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteServiceAccountsParams() *DeleteServiceAccountsParams {
	return &DeleteServiceAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceAccountsParamsWithTimeout creates a new DeleteServiceAccountsParams object
// with the ability to set a timeout on a request.
func NewDeleteServiceAccountsParamsWithTimeout(timeout time.Duration) *DeleteServiceAccountsParams {
	return &DeleteServiceAccountsParams{
		timeout: timeout,
	}
}

// NewDeleteServiceAccountsParamsWithContext creates a new DeleteServiceAccountsParams object
// with the ability to set a context for a request.
func NewDeleteServiceAccountsParamsWithContext(ctx context.Context) *DeleteServiceAccountsParams {
	return &DeleteServiceAccountsParams{
		Context: ctx,
	}
}

// NewDeleteServiceAccountsParamsWithHTTPClient creates a new DeleteServiceAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteServiceAccountsParamsWithHTTPClient(client *http.Client) *DeleteServiceAccountsParams {
	return &DeleteServiceAccountsParams{
		HTTPClient: client,
	}
}

/*
DeleteServiceAccountsParams contains all the parameters to send to the API endpoint

	for the delete service accounts operation.

	Typically these are written to a http.Request.
*/
type DeleteServiceAccountsParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* Payload.

	   A map where the key is the namespace and the value is an array of service accounts to delete
	*/
	Payload models.KubernetesK8sServiceAccountDeleteRequests

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceAccountsParams) WithDefaults() *DeleteServiceAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete service accounts params
func (o *DeleteServiceAccountsParams) WithTimeout(timeout time.Duration) *DeleteServiceAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service accounts params
func (o *DeleteServiceAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service accounts params
func (o *DeleteServiceAccountsParams) WithContext(ctx context.Context) *DeleteServiceAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service accounts params
func (o *DeleteServiceAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service accounts params
func (o *DeleteServiceAccountsParams) WithHTTPClient(client *http.Client) *DeleteServiceAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service accounts params
func (o *DeleteServiceAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete service accounts params
func (o *DeleteServiceAccountsParams) WithID(id int64) *DeleteServiceAccountsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete service accounts params
func (o *DeleteServiceAccountsParams) SetID(id int64) {
	o.ID = id
}

// WithPayload adds the payload to the delete service accounts params
func (o *DeleteServiceAccountsParams) WithPayload(payload models.KubernetesK8sServiceAccountDeleteRequests) *DeleteServiceAccountsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the delete service accounts params
func (o *DeleteServiceAccountsParams) SetPayload(payload models.KubernetesK8sServiceAccountDeleteRequests) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
