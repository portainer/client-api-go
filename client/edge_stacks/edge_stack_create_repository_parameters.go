// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

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

	"github.com/portainer/client-api-go/v2/models"
)

// NewEdgeStackCreateRepositoryParams creates a new EdgeStackCreateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeStackCreateRepositoryParams() *EdgeStackCreateRepositoryParams {
	return &EdgeStackCreateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeStackCreateRepositoryParamsWithTimeout creates a new EdgeStackCreateRepositoryParams object
// with the ability to set a timeout on a request.
func NewEdgeStackCreateRepositoryParamsWithTimeout(timeout time.Duration) *EdgeStackCreateRepositoryParams {
	return &EdgeStackCreateRepositoryParams{
		timeout: timeout,
	}
}

// NewEdgeStackCreateRepositoryParamsWithContext creates a new EdgeStackCreateRepositoryParams object
// with the ability to set a context for a request.
func NewEdgeStackCreateRepositoryParamsWithContext(ctx context.Context) *EdgeStackCreateRepositoryParams {
	return &EdgeStackCreateRepositoryParams{
		Context: ctx,
	}
}

// NewEdgeStackCreateRepositoryParamsWithHTTPClient creates a new EdgeStackCreateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeStackCreateRepositoryParamsWithHTTPClient(client *http.Client) *EdgeStackCreateRepositoryParams {
	return &EdgeStackCreateRepositoryParams{
		HTTPClient: client,
	}
}

/*
EdgeStackCreateRepositoryParams contains all the parameters to send to the API endpoint

	for the edge stack create repository operation.

	Typically these are written to a http.Request.
*/
type EdgeStackCreateRepositoryParams struct {

	/* Body.

	   stack config
	*/
	Body *models.EdgestacksEdgeStackFromGitRepositoryPayload

	/* Dryrun.

	   if true, will not create an edge stack, but just will check the settings and return a non-persisted edge stack object
	*/
	Dryrun *string

	/* Method.

	   Creation Method
	*/
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge stack create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackCreateRepositoryParams) WithDefaults() *EdgeStackCreateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge stack create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackCreateRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithTimeout(timeout time.Duration) *EdgeStackCreateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithContext(ctx context.Context) *EdgeStackCreateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithHTTPClient(client *http.Client) *EdgeStackCreateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithBody(body *models.EdgestacksEdgeStackFromGitRepositoryPayload) *EdgeStackCreateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetBody(body *models.EdgestacksEdgeStackFromGitRepositoryPayload) {
	o.Body = body
}

// WithDryrun adds the dryrun to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithDryrun(dryrun *string) *EdgeStackCreateRepositoryParams {
	o.SetDryrun(dryrun)
	return o
}

// SetDryrun adds the dryrun to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetDryrun(dryrun *string) {
	o.Dryrun = dryrun
}

// WithMethod adds the method to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) WithMethod(method string) *EdgeStackCreateRepositoryParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the edge stack create repository params
func (o *EdgeStackCreateRepositoryParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeStackCreateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Dryrun != nil {

		// query param dryrun
		var qrDryrun string

		if o.Dryrun != nil {
			qrDryrun = *o.Dryrun
		}
		qDryrun := qrDryrun
		if qDryrun != "" {

			if err := r.SetQueryParam("dryrun", qDryrun); err != nil {
				return err
			}
		}
	}

	// query param method
	qrMethod := o.Method
	qMethod := qrMethod
	if qMethod != "" {

		if err := r.SetQueryParam("method", qMethod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
