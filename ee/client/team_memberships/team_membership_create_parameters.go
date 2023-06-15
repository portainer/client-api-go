// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

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

	"github.com/portainer/client-api-go/ee/v2/models"
)

// NewTeamMembershipCreateParams creates a new TeamMembershipCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamMembershipCreateParams() *TeamMembershipCreateParams {
	return &TeamMembershipCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamMembershipCreateParamsWithTimeout creates a new TeamMembershipCreateParams object
// with the ability to set a timeout on a request.
func NewTeamMembershipCreateParamsWithTimeout(timeout time.Duration) *TeamMembershipCreateParams {
	return &TeamMembershipCreateParams{
		timeout: timeout,
	}
}

// NewTeamMembershipCreateParamsWithContext creates a new TeamMembershipCreateParams object
// with the ability to set a context for a request.
func NewTeamMembershipCreateParamsWithContext(ctx context.Context) *TeamMembershipCreateParams {
	return &TeamMembershipCreateParams{
		Context: ctx,
	}
}

// NewTeamMembershipCreateParamsWithHTTPClient creates a new TeamMembershipCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamMembershipCreateParamsWithHTTPClient(client *http.Client) *TeamMembershipCreateParams {
	return &TeamMembershipCreateParams{
		HTTPClient: client,
	}
}

/*
TeamMembershipCreateParams contains all the parameters to send to the API endpoint

	for the team membership create operation.

	Typically these are written to a http.Request.
*/
type TeamMembershipCreateParams struct {

	/* Body.

	   Team membership details
	*/
	Body *models.TeammembershipsTeamMembershipCreatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the team membership create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamMembershipCreateParams) WithDefaults() *TeamMembershipCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the team membership create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamMembershipCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the team membership create params
func (o *TeamMembershipCreateParams) WithTimeout(timeout time.Duration) *TeamMembershipCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team membership create params
func (o *TeamMembershipCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team membership create params
func (o *TeamMembershipCreateParams) WithContext(ctx context.Context) *TeamMembershipCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team membership create params
func (o *TeamMembershipCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team membership create params
func (o *TeamMembershipCreateParams) WithHTTPClient(client *http.Client) *TeamMembershipCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team membership create params
func (o *TeamMembershipCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the team membership create params
func (o *TeamMembershipCreateParams) WithBody(body *models.TeammembershipsTeamMembershipCreatePayload) *TeamMembershipCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the team membership create params
func (o *TeamMembershipCreateParams) SetBody(body *models.TeammembershipsTeamMembershipCreatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TeamMembershipCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
