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

	"github.com/portainer/client-api-go/models"
)

// NewEdgeStackUpdateParams creates a new EdgeStackUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeStackUpdateParams() *EdgeStackUpdateParams {
	return &EdgeStackUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeStackUpdateParamsWithTimeout creates a new EdgeStackUpdateParams object
// with the ability to set a timeout on a request.
func NewEdgeStackUpdateParamsWithTimeout(timeout time.Duration) *EdgeStackUpdateParams {
	return &EdgeStackUpdateParams{
		timeout: timeout,
	}
}

// NewEdgeStackUpdateParamsWithContext creates a new EdgeStackUpdateParams object
// with the ability to set a context for a request.
func NewEdgeStackUpdateParamsWithContext(ctx context.Context) *EdgeStackUpdateParams {
	return &EdgeStackUpdateParams{
		Context: ctx,
	}
}

// NewEdgeStackUpdateParamsWithHTTPClient creates a new EdgeStackUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeStackUpdateParamsWithHTTPClient(client *http.Client) *EdgeStackUpdateParams {
	return &EdgeStackUpdateParams{
		HTTPClient: client,
	}
}

/*
EdgeStackUpdateParams contains all the parameters to send to the API endpoint

	for the edge stack update operation.

	Typically these are written to a http.Request.
*/
type EdgeStackUpdateParams struct {

	/* Body.

	   EdgeStack data
	*/
	Body *models.EdgestacksUpdateEdgeStackPayload

	/* ID.

	   EdgeStack Id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge stack update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackUpdateParams) WithDefaults() *EdgeStackUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge stack update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeStackUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge stack update params
func (o *EdgeStackUpdateParams) WithTimeout(timeout time.Duration) *EdgeStackUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge stack update params
func (o *EdgeStackUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge stack update params
func (o *EdgeStackUpdateParams) WithContext(ctx context.Context) *EdgeStackUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge stack update params
func (o *EdgeStackUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge stack update params
func (o *EdgeStackUpdateParams) WithHTTPClient(client *http.Client) *EdgeStackUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge stack update params
func (o *EdgeStackUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edge stack update params
func (o *EdgeStackUpdateParams) WithBody(body *models.EdgestacksUpdateEdgeStackPayload) *EdgeStackUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge stack update params
func (o *EdgeStackUpdateParams) SetBody(body *models.EdgestacksUpdateEdgeStackPayload) {
	o.Body = body
}

// WithID adds the id to the edge stack update params
func (o *EdgeStackUpdateParams) WithID(id string) *EdgeStackUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge stack update params
func (o *EdgeStackUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeStackUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
