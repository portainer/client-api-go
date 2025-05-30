// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// NewStackCreateKubernetesGitParams creates a new StackCreateKubernetesGitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackCreateKubernetesGitParams() *StackCreateKubernetesGitParams {
	return &StackCreateKubernetesGitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackCreateKubernetesGitParamsWithTimeout creates a new StackCreateKubernetesGitParams object
// with the ability to set a timeout on a request.
func NewStackCreateKubernetesGitParamsWithTimeout(timeout time.Duration) *StackCreateKubernetesGitParams {
	return &StackCreateKubernetesGitParams{
		timeout: timeout,
	}
}

// NewStackCreateKubernetesGitParamsWithContext creates a new StackCreateKubernetesGitParams object
// with the ability to set a context for a request.
func NewStackCreateKubernetesGitParamsWithContext(ctx context.Context) *StackCreateKubernetesGitParams {
	return &StackCreateKubernetesGitParams{
		Context: ctx,
	}
}

// NewStackCreateKubernetesGitParamsWithHTTPClient creates a new StackCreateKubernetesGitParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackCreateKubernetesGitParamsWithHTTPClient(client *http.Client) *StackCreateKubernetesGitParams {
	return &StackCreateKubernetesGitParams{
		HTTPClient: client,
	}
}

/*
StackCreateKubernetesGitParams contains all the parameters to send to the API endpoint

	for the stack create kubernetes git operation.

	Typically these are written to a http.Request.
*/
type StackCreateKubernetesGitParams struct {

	/* Body.

	   stack config
	*/
	Body *models.StacksKubernetesGitDeploymentPayload

	/* EndpointID.

	   Identifier of the environment that will be used to deploy the stack
	*/
	EndpointID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack create kubernetes git params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateKubernetesGitParams) WithDefaults() *StackCreateKubernetesGitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack create kubernetes git params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateKubernetesGitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) WithTimeout(timeout time.Duration) *StackCreateKubernetesGitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) WithContext(ctx context.Context) *StackCreateKubernetesGitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) WithHTTPClient(client *http.Client) *StackCreateKubernetesGitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) WithBody(body *models.StacksKubernetesGitDeploymentPayload) *StackCreateKubernetesGitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) SetBody(body *models.StacksKubernetesGitDeploymentPayload) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) WithEndpointID(endpointID int64) *StackCreateKubernetesGitParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack create kubernetes git params
func (o *StackCreateKubernetesGitParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WriteToRequest writes these params to a swagger request
func (o *StackCreateKubernetesGitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param endpointId
	qrEndpointID := o.EndpointID
	qEndpointID := swag.FormatInt64(qrEndpointID)
	if qEndpointID != "" {

		if err := r.SetQueryParam("endpointId", qEndpointID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
