// Code generated by go-swagger; DO NOT EDIT.

package kaas

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

// NewMicrok8sAddonsParams creates a new Microk8sAddonsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMicrok8sAddonsParams() *Microk8sAddonsParams {
	return &Microk8sAddonsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMicrok8sAddonsParamsWithTimeout creates a new Microk8sAddonsParams object
// with the ability to set a timeout on a request.
func NewMicrok8sAddonsParamsWithTimeout(timeout time.Duration) *Microk8sAddonsParams {
	return &Microk8sAddonsParams{
		timeout: timeout,
	}
}

// NewMicrok8sAddonsParamsWithContext creates a new Microk8sAddonsParams object
// with the ability to set a context for a request.
func NewMicrok8sAddonsParamsWithContext(ctx context.Context) *Microk8sAddonsParams {
	return &Microk8sAddonsParams{
		Context: ctx,
	}
}

// NewMicrok8sAddonsParamsWithHTTPClient creates a new Microk8sAddonsParams object
// with the ability to set a custom HTTPClient for a request.
func NewMicrok8sAddonsParamsWithHTTPClient(client *http.Client) *Microk8sAddonsParams {
	return &Microk8sAddonsParams{
		HTTPClient: client,
	}
}

/*
Microk8sAddonsParams contains all the parameters to send to the API endpoint

	for the microk8s addons operation.

	Typically these are written to a http.Request.
*/
type Microk8sAddonsParams struct {

	/* CredentialID.

	   The credential ID to use to connect to a node in the MicroK8s cluster.
	*/
	CredentialID int64

	/* EnvironmentID.

	   The environment ID of the cluster within Portainer.
	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the microk8s addons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Microk8sAddonsParams) WithDefaults() *Microk8sAddonsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the microk8s addons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Microk8sAddonsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the microk8s addons params
func (o *Microk8sAddonsParams) WithTimeout(timeout time.Duration) *Microk8sAddonsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the microk8s addons params
func (o *Microk8sAddonsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the microk8s addons params
func (o *Microk8sAddonsParams) WithContext(ctx context.Context) *Microk8sAddonsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the microk8s addons params
func (o *Microk8sAddonsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the microk8s addons params
func (o *Microk8sAddonsParams) WithHTTPClient(client *http.Client) *Microk8sAddonsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the microk8s addons params
func (o *Microk8sAddonsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the microk8s addons params
func (o *Microk8sAddonsParams) WithCredentialID(credentialID int64) *Microk8sAddonsParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the microk8s addons params
func (o *Microk8sAddonsParams) SetCredentialID(credentialID int64) {
	o.CredentialID = credentialID
}

// WithEnvironmentID adds the environmentID to the microk8s addons params
func (o *Microk8sAddonsParams) WithEnvironmentID(environmentID int64) *Microk8sAddonsParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the microk8s addons params
func (o *Microk8sAddonsParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *Microk8sAddonsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param credentialID
	qrCredentialID := o.CredentialID
	qCredentialID := swag.FormatInt64(qrCredentialID)
	if qCredentialID != "" {

		if err := r.SetQueryParam("credentialID", qCredentialID); err != nil {
			return err
		}
	}

	// query param environmentID
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := swag.FormatInt64(qrEnvironmentID)
	if qEnvironmentID != "" {

		if err := r.SetQueryParam("environmentID", qEnvironmentID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
