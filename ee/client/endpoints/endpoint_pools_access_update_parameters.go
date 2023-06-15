// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

	"github.com/portainer/client-api-go/ee/v2/models"
)

// NewEndpointPoolsAccessUpdateParams creates a new EndpointPoolsAccessUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointPoolsAccessUpdateParams() *EndpointPoolsAccessUpdateParams {
	return &EndpointPoolsAccessUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointPoolsAccessUpdateParamsWithTimeout creates a new EndpointPoolsAccessUpdateParams object
// with the ability to set a timeout on a request.
func NewEndpointPoolsAccessUpdateParamsWithTimeout(timeout time.Duration) *EndpointPoolsAccessUpdateParams {
	return &EndpointPoolsAccessUpdateParams{
		timeout: timeout,
	}
}

// NewEndpointPoolsAccessUpdateParamsWithContext creates a new EndpointPoolsAccessUpdateParams object
// with the ability to set a context for a request.
func NewEndpointPoolsAccessUpdateParamsWithContext(ctx context.Context) *EndpointPoolsAccessUpdateParams {
	return &EndpointPoolsAccessUpdateParams{
		Context: ctx,
	}
}

// NewEndpointPoolsAccessUpdateParamsWithHTTPClient creates a new EndpointPoolsAccessUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointPoolsAccessUpdateParamsWithHTTPClient(client *http.Client) *EndpointPoolsAccessUpdateParams {
	return &EndpointPoolsAccessUpdateParams{
		HTTPClient: client,
	}
}

/*
EndpointPoolsAccessUpdateParams contains all the parameters to send to the API endpoint

	for the endpoint pools access update operation.

	Typically these are written to a http.Request.
*/
type EndpointPoolsAccessUpdateParams struct {

	/* Body.

	   details
	*/
	Body *models.EndpointsResourcePoolUpdatePayload

	/* ID.

	   endpoint id
	*/
	ID int64

	/* Rpn.

	   namespace
	*/
	Rpn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint pools access update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointPoolsAccessUpdateParams) WithDefaults() *EndpointPoolsAccessUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint pools access update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointPoolsAccessUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithTimeout(timeout time.Duration) *EndpointPoolsAccessUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithContext(ctx context.Context) *EndpointPoolsAccessUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithHTTPClient(client *http.Client) *EndpointPoolsAccessUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithBody(body *models.EndpointsResourcePoolUpdatePayload) *EndpointPoolsAccessUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetBody(body *models.EndpointsResourcePoolUpdatePayload) {
	o.Body = body
}

// WithID adds the id to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithID(id int64) *EndpointPoolsAccessUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetID(id int64) {
	o.ID = id
}

// WithRpn adds the rpn to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) WithRpn(rpn int64) *EndpointPoolsAccessUpdateParams {
	o.SetRpn(rpn)
	return o
}

// SetRpn adds the rpn to the endpoint pools access update params
func (o *EndpointPoolsAccessUpdateParams) SetRpn(rpn int64) {
	o.Rpn = rpn
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointPoolsAccessUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param rpn
	if err := r.SetPathParam("rpn", swag.FormatInt64(o.Rpn)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
