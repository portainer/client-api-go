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

	"github.com/portainer/client-api-go/v2/models_ce"
)

// NewTeamUpdateParams creates a new TeamUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamUpdateParams() *TeamUpdateParams {
	return &TeamUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamUpdateParamsWithTimeout creates a new TeamUpdateParams object
// with the ability to set a timeout on a request.
func NewTeamUpdateParamsWithTimeout(timeout time.Duration) *TeamUpdateParams {
	return &TeamUpdateParams{
		timeout: timeout,
	}
}

// NewTeamUpdateParamsWithContext creates a new TeamUpdateParams object
// with the ability to set a context for a request.
func NewTeamUpdateParamsWithContext(ctx context.Context) *TeamUpdateParams {
	return &TeamUpdateParams{
		Context: ctx,
	}
}

// NewTeamUpdateParamsWithHTTPClient creates a new TeamUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamUpdateParamsWithHTTPClient(client *http.Client) *TeamUpdateParams {
	return &TeamUpdateParams{
		HTTPClient: client,
	}
}

/*
TeamUpdateParams contains all the parameters to send to the API endpoint

	for the team update operation.

	Typically these are written to a http.Request.
*/
type TeamUpdateParams struct {

	/* Body.

	   Team details
	*/
	Body *models_ce.TeamsTeamUpdatePayload

	/* ID.

	   Team identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the team update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamUpdateParams) WithDefaults() *TeamUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the team update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the team update params
func (o *TeamUpdateParams) WithTimeout(timeout time.Duration) *TeamUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team update params
func (o *TeamUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team update params
func (o *TeamUpdateParams) WithContext(ctx context.Context) *TeamUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team update params
func (o *TeamUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team update params
func (o *TeamUpdateParams) WithHTTPClient(client *http.Client) *TeamUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team update params
func (o *TeamUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the team update params
func (o *TeamUpdateParams) WithBody(body *models_ce.TeamsTeamUpdatePayload) *TeamUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the team update params
func (o *TeamUpdateParams) SetBody(body *models_ce.TeamsTeamUpdatePayload) {
	o.Body = body
}

// WithID adds the id to the team update params
func (o *TeamUpdateParams) WithID(id int64) *TeamUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the team update params
func (o *TeamUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TeamUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
