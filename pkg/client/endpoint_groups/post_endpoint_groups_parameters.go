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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewPostEndpointGroupsParams creates a new PostEndpointGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEndpointGroupsParams() *PostEndpointGroupsParams {
	return &PostEndpointGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEndpointGroupsParamsWithTimeout creates a new PostEndpointGroupsParams object
// with the ability to set a timeout on a request.
func NewPostEndpointGroupsParamsWithTimeout(timeout time.Duration) *PostEndpointGroupsParams {
	return &PostEndpointGroupsParams{
		timeout: timeout,
	}
}

// NewPostEndpointGroupsParamsWithContext creates a new PostEndpointGroupsParams object
// with the ability to set a context for a request.
func NewPostEndpointGroupsParamsWithContext(ctx context.Context) *PostEndpointGroupsParams {
	return &PostEndpointGroupsParams{
		Context: ctx,
	}
}

// NewPostEndpointGroupsParamsWithHTTPClient creates a new PostEndpointGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEndpointGroupsParamsWithHTTPClient(client *http.Client) *PostEndpointGroupsParams {
	return &PostEndpointGroupsParams{
		HTTPClient: client,
	}
}

/*
PostEndpointGroupsParams contains all the parameters to send to the API endpoint

	for the post endpoint groups operation.

	Typically these are written to a http.Request.
*/
type PostEndpointGroupsParams struct {

	/* Body.

	   Environment(Endpoint) Group details
	*/
	Body *models.EndpointgroupsEndpointGroupCreatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post endpoint groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEndpointGroupsParams) WithDefaults() *PostEndpointGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post endpoint groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEndpointGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post endpoint groups params
func (o *PostEndpointGroupsParams) WithTimeout(timeout time.Duration) *PostEndpointGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post endpoint groups params
func (o *PostEndpointGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post endpoint groups params
func (o *PostEndpointGroupsParams) WithContext(ctx context.Context) *PostEndpointGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post endpoint groups params
func (o *PostEndpointGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post endpoint groups params
func (o *PostEndpointGroupsParams) WithHTTPClient(client *http.Client) *PostEndpointGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post endpoint groups params
func (o *PostEndpointGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post endpoint groups params
func (o *PostEndpointGroupsParams) WithBody(body *models.EndpointgroupsEndpointGroupCreatePayload) *PostEndpointGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post endpoint groups params
func (o *PostEndpointGroupsParams) SetBody(body *models.EndpointgroupsEndpointGroupCreatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEndpointGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
