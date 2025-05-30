// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

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

// NewAgentVersionsParams creates a new AgentVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAgentVersionsParams() *AgentVersionsParams {
	return &AgentVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAgentVersionsParamsWithTimeout creates a new AgentVersionsParams object
// with the ability to set a timeout on a request.
func NewAgentVersionsParamsWithTimeout(timeout time.Duration) *AgentVersionsParams {
	return &AgentVersionsParams{
		timeout: timeout,
	}
}

// NewAgentVersionsParamsWithContext creates a new AgentVersionsParams object
// with the ability to set a context for a request.
func NewAgentVersionsParamsWithContext(ctx context.Context) *AgentVersionsParams {
	return &AgentVersionsParams{
		Context: ctx,
	}
}

// NewAgentVersionsParamsWithHTTPClient creates a new AgentVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAgentVersionsParamsWithHTTPClient(client *http.Client) *AgentVersionsParams {
	return &AgentVersionsParams{
		HTTPClient: client,
	}
}

/*
AgentVersionsParams contains all the parameters to send to the API endpoint

	for the agent versions operation.

	Typically these are written to a http.Request.
*/
type AgentVersionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the agent versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentVersionsParams) WithDefaults() *AgentVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the agent versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the agent versions params
func (o *AgentVersionsParams) WithTimeout(timeout time.Duration) *AgentVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agent versions params
func (o *AgentVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agent versions params
func (o *AgentVersionsParams) WithContext(ctx context.Context) *AgentVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agent versions params
func (o *AgentVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agent versions params
func (o *AgentVersionsParams) WithHTTPClient(client *http.Client) *AgentVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agent versions params
func (o *AgentVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AgentVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
