// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamInspectParams creates a new TeamInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamInspectParams() *TeamInspectParams {
	return &TeamInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamInspectParamsWithTimeout creates a new TeamInspectParams object
// with the ability to set a timeout on a request.
func NewTeamInspectParamsWithTimeout(timeout time.Duration) *TeamInspectParams {
	return &TeamInspectParams{
		timeout: timeout,
	}
}

// NewTeamInspectParamsWithContext creates a new TeamInspectParams object
// with the ability to set a context for a request.
func NewTeamInspectParamsWithContext(ctx context.Context) *TeamInspectParams {
	return &TeamInspectParams{
		Context: ctx,
	}
}

// NewTeamInspectParamsWithHTTPClient creates a new TeamInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamInspectParamsWithHTTPClient(client *http.Client) *TeamInspectParams {
	return &TeamInspectParams{
		HTTPClient: client,
	}
}

/*
TeamInspectParams contains all the parameters to send to the API endpoint

	for the team inspect operation.

	Typically these are written to a http.Request.
*/
type TeamInspectParams struct {

	/* ID.

	   Team identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the team inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamInspectParams) WithDefaults() *TeamInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the team inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the team inspect params
func (o *TeamInspectParams) WithTimeout(timeout time.Duration) *TeamInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team inspect params
func (o *TeamInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team inspect params
func (o *TeamInspectParams) WithContext(ctx context.Context) *TeamInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team inspect params
func (o *TeamInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team inspect params
func (o *TeamInspectParams) WithHTTPClient(client *http.Client) *TeamInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team inspect params
func (o *TeamInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the team inspect params
func (o *TeamInspectParams) WithID(id int64) *TeamInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the team inspect params
func (o *TeamInspectParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TeamInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
