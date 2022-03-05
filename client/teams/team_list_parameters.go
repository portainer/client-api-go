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

/* TeamListParams contains all the parameters to send to the API endpoint
   for the team list operation.

   Typically these are written to a http.Request.
*/
type TeamListParams struct {
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

// WriteToRequest writes these params to a swagger request
func (o *TeamListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
