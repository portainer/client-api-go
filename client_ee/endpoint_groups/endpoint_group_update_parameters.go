// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

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

	"github.com/portainer/client-api-go/v2/models_ee"
)

// NewEndpointGroupUpdateParams creates a new EndpointGroupUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointGroupUpdateParams() *EndpointGroupUpdateParams {
	return &EndpointGroupUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointGroupUpdateParamsWithTimeout creates a new EndpointGroupUpdateParams object
// with the ability to set a timeout on a request.
func NewEndpointGroupUpdateParamsWithTimeout(timeout time.Duration) *EndpointGroupUpdateParams {
	return &EndpointGroupUpdateParams{
		timeout: timeout,
	}
}

// NewEndpointGroupUpdateParamsWithContext creates a new EndpointGroupUpdateParams object
// with the ability to set a context for a request.
func NewEndpointGroupUpdateParamsWithContext(ctx context.Context) *EndpointGroupUpdateParams {
	return &EndpointGroupUpdateParams{
		Context: ctx,
	}
}

// NewEndpointGroupUpdateParamsWithHTTPClient creates a new EndpointGroupUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointGroupUpdateParamsWithHTTPClient(client *http.Client) *EndpointGroupUpdateParams {
	return &EndpointGroupUpdateParams{
		HTTPClient: client,
	}
}

/*
EndpointGroupUpdateParams contains all the parameters to send to the API endpoint

	for the endpoint group update operation.

	Typically these are written to a http.Request.
*/
type EndpointGroupUpdateParams struct {

	/* Body.

	   EndpointGroup details
	*/
	Body *models_ee.EndpointgroupsEndpointGroupUpdatePayload

	/* ID.

	   EndpointGroup identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint group update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupUpdateParams) WithDefaults() *EndpointGroupUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint group update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointGroupUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint group update params
func (o *EndpointGroupUpdateParams) WithTimeout(timeout time.Duration) *EndpointGroupUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint group update params
func (o *EndpointGroupUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint group update params
func (o *EndpointGroupUpdateParams) WithContext(ctx context.Context) *EndpointGroupUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint group update params
func (o *EndpointGroupUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint group update params
func (o *EndpointGroupUpdateParams) WithHTTPClient(client *http.Client) *EndpointGroupUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint group update params
func (o *EndpointGroupUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the endpoint group update params
func (o *EndpointGroupUpdateParams) WithBody(body *models_ee.EndpointgroupsEndpointGroupUpdatePayload) *EndpointGroupUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the endpoint group update params
func (o *EndpointGroupUpdateParams) SetBody(body *models_ee.EndpointgroupsEndpointGroupUpdatePayload) {
	o.Body = body
}

// WithID adds the id to the endpoint group update params
func (o *EndpointGroupUpdateParams) WithID(id int64) *EndpointGroupUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the endpoint group update params
func (o *EndpointGroupUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointGroupUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
