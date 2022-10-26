// Code generated by go-swagger; DO NOT EDIT.

package edge_groups

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

// NewEgeGroupUpdateParams creates a new EgeGroupUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEgeGroupUpdateParams() *EgeGroupUpdateParams {
	return &EgeGroupUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEgeGroupUpdateParamsWithTimeout creates a new EgeGroupUpdateParams object
// with the ability to set a timeout on a request.
func NewEgeGroupUpdateParamsWithTimeout(timeout time.Duration) *EgeGroupUpdateParams {
	return &EgeGroupUpdateParams{
		timeout: timeout,
	}
}

// NewEgeGroupUpdateParamsWithContext creates a new EgeGroupUpdateParams object
// with the ability to set a context for a request.
func NewEgeGroupUpdateParamsWithContext(ctx context.Context) *EgeGroupUpdateParams {
	return &EgeGroupUpdateParams{
		Context: ctx,
	}
}

// NewEgeGroupUpdateParamsWithHTTPClient creates a new EgeGroupUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEgeGroupUpdateParamsWithHTTPClient(client *http.Client) *EgeGroupUpdateParams {
	return &EgeGroupUpdateParams{
		HTTPClient: client,
	}
}

/*
EgeGroupUpdateParams contains all the parameters to send to the API endpoint

	for the ege group update operation.

	Typically these are written to a http.Request.
*/
type EgeGroupUpdateParams struct {

	/* Body.

	   EdgeGroup data
	*/
	Body *models.EdgegroupsEdgeGroupUpdatePayload

	/* ID.

	   EdgeGroup Id
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ege group update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EgeGroupUpdateParams) WithDefaults() *EgeGroupUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ege group update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EgeGroupUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ege group update params
func (o *EgeGroupUpdateParams) WithTimeout(timeout time.Duration) *EgeGroupUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ege group update params
func (o *EgeGroupUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ege group update params
func (o *EgeGroupUpdateParams) WithContext(ctx context.Context) *EgeGroupUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ege group update params
func (o *EgeGroupUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ege group update params
func (o *EgeGroupUpdateParams) WithHTTPClient(client *http.Client) *EgeGroupUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ege group update params
func (o *EgeGroupUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ege group update params
func (o *EgeGroupUpdateParams) WithBody(body *models.EdgegroupsEdgeGroupUpdatePayload) *EgeGroupUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ege group update params
func (o *EgeGroupUpdateParams) SetBody(body *models.EdgegroupsEdgeGroupUpdatePayload) {
	o.Body = body
}

// WithID adds the id to the ege group update params
func (o *EgeGroupUpdateParams) WithID(id int64) *EgeGroupUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ege group update params
func (o *EgeGroupUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EgeGroupUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
