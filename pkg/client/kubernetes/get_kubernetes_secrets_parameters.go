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

// NewGetKubernetesSecretsParams creates a new GetKubernetesSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesSecretsParams() *GetKubernetesSecretsParams {
	return &GetKubernetesSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesSecretsParamsWithTimeout creates a new GetKubernetesSecretsParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesSecretsParamsWithTimeout(timeout time.Duration) *GetKubernetesSecretsParams {
	return &GetKubernetesSecretsParams{
		timeout: timeout,
	}
}

// NewGetKubernetesSecretsParamsWithContext creates a new GetKubernetesSecretsParams object
// with the ability to set a context for a request.
func NewGetKubernetesSecretsParamsWithContext(ctx context.Context) *GetKubernetesSecretsParams {
	return &GetKubernetesSecretsParams{
		Context: ctx,
	}
}

// NewGetKubernetesSecretsParamsWithHTTPClient creates a new GetKubernetesSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesSecretsParamsWithHTTPClient(client *http.Client) *GetKubernetesSecretsParams {
	return &GetKubernetesSecretsParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesSecretsParams contains all the parameters to send to the API endpoint

	for the get kubernetes secrets operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesSecretsParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* IsUsed.

	   When set to true, associate the Secrets with the applications that use them
	*/
	IsUsed bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesSecretsParams) WithDefaults() *GetKubernetesSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) WithTimeout(timeout time.Duration) *GetKubernetesSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) WithContext(ctx context.Context) *GetKubernetesSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) WithHTTPClient(client *http.Client) *GetKubernetesSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) WithID(id int64) *GetKubernetesSecretsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) SetID(id int64) {
	o.ID = id
}

// WithIsUsed adds the isUsed to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) WithIsUsed(isUsed bool) *GetKubernetesSecretsParams {
	o.SetIsUsed(isUsed)
	return o
}

// SetIsUsed adds the isUsed to the get kubernetes secrets params
func (o *GetKubernetesSecretsParams) SetIsUsed(isUsed bool) {
	o.IsUsed = isUsed
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param isUsed
	qrIsUsed := o.IsUsed
	qIsUsed := swag.FormatBool(qrIsUsed)
	if qIsUsed != "" {

		if err := r.SetQueryParam("isUsed", qIsUsed); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
