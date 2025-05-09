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

// NewTeamListParams creates a new TeamListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamListParams() *TeamListParams {
	return &TeamListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamListParamsWithTimeout creates a new TeamListParams object
// with the ability to set a timeout on a request.
func NewTeamListParamsWithTimeout(timeout time.Duration) *TeamListParams {
	return &TeamListParams{
		timeout: timeout,
	}
}

// NewTeamListParamsWithContext creates a new TeamListParams object
// with the ability to set a context for a request.
func NewTeamListParamsWithContext(ctx context.Context) *TeamListParams {
	return &TeamListParams{
		Context: ctx,
	}
}

// NewTeamListParamsWithHTTPClient creates a new TeamListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamListParamsWithHTTPClient(client *http.Client) *TeamListParams {
	return &TeamListParams{
		HTTPClient: client,
	}
}

/*
TeamListParams contains all the parameters to send to the API endpoint

	for the team list operation.

	Typically these are written to a http.Request.
*/
type TeamListParams struct {

	/* EnvironmentID.

	   Identifier of the environment(endpoint) that will be used to filter the authorized teams
	*/
	EnvironmentID *int64

	/* OnlyLedTeams.

	   Only list teams that the user is leader of
	*/
	OnlyLedTeams *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the team list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamListParams) WithDefaults() *TeamListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the team list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the team list params
func (o *TeamListParams) WithTimeout(timeout time.Duration) *TeamListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team list params
func (o *TeamListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team list params
func (o *TeamListParams) WithContext(ctx context.Context) *TeamListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team list params
func (o *TeamListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team list params
func (o *TeamListParams) WithHTTPClient(client *http.Client) *TeamListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team list params
func (o *TeamListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the team list params
func (o *TeamListParams) WithEnvironmentID(environmentID *int64) *TeamListParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the team list params
func (o *TeamListParams) SetEnvironmentID(environmentID *int64) {
	o.EnvironmentID = environmentID
}

// WithOnlyLedTeams adds the onlyLedTeams to the team list params
func (o *TeamListParams) WithOnlyLedTeams(onlyLedTeams *bool) *TeamListParams {
	o.SetOnlyLedTeams(onlyLedTeams)
	return o
}

// SetOnlyLedTeams adds the onlyLedTeams to the team list params
func (o *TeamListParams) SetOnlyLedTeams(onlyLedTeams *bool) {
	o.OnlyLedTeams = onlyLedTeams
}

// WriteToRequest writes these params to a swagger request
func (o *TeamListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentID != nil {

		// query param environmentId
		var qrEnvironmentID int64

		if o.EnvironmentID != nil {
			qrEnvironmentID = *o.EnvironmentID
		}
		qEnvironmentID := swag.FormatInt64(qrEnvironmentID)
		if qEnvironmentID != "" {

			if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
				return err
			}
		}
	}

	if o.OnlyLedTeams != nil {

		// query param onlyLedTeams
		var qrOnlyLedTeams bool

		if o.OnlyLedTeams != nil {
			qrOnlyLedTeams = *o.OnlyLedTeams
		}
		qOnlyLedTeams := swag.FormatBool(qrOnlyLedTeams)
		if qOnlyLedTeams != "" {

			if err := r.SetQueryParam("onlyLedTeams", qOnlyLedTeams); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
