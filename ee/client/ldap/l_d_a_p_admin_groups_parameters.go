// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewLDAPAdminGroupsParams creates a new LDAPAdminGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLDAPAdminGroupsParams() *LDAPAdminGroupsParams {
	return &LDAPAdminGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLDAPAdminGroupsParamsWithTimeout creates a new LDAPAdminGroupsParams object
// with the ability to set a timeout on a request.
func NewLDAPAdminGroupsParamsWithTimeout(timeout time.Duration) *LDAPAdminGroupsParams {
	return &LDAPAdminGroupsParams{
		timeout: timeout,
	}
}

// NewLDAPAdminGroupsParamsWithContext creates a new LDAPAdminGroupsParams object
// with the ability to set a context for a request.
func NewLDAPAdminGroupsParamsWithContext(ctx context.Context) *LDAPAdminGroupsParams {
	return &LDAPAdminGroupsParams{
		Context: ctx,
	}
}

// NewLDAPAdminGroupsParamsWithHTTPClient creates a new LDAPAdminGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLDAPAdminGroupsParamsWithHTTPClient(client *http.Client) *LDAPAdminGroupsParams {
	return &LDAPAdminGroupsParams{
		HTTPClient: client,
	}
}

/*
LDAPAdminGroupsParams contains all the parameters to send to the API endpoint

	for the l d a p admin groups operation.

	Typically these are written to a http.Request.
*/
type LDAPAdminGroupsParams struct {

	/* Body.

	   LDAPSettings
	*/
	Body *models.LdapAdminGroupsPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the l d a p admin groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPAdminGroupsParams) WithDefaults() *LDAPAdminGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the l d a p admin groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPAdminGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) WithTimeout(timeout time.Duration) *LDAPAdminGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) WithContext(ctx context.Context) *LDAPAdminGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) WithHTTPClient(client *http.Client) *LDAPAdminGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) WithBody(body *models.LdapAdminGroupsPayload) *LDAPAdminGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the l d a p admin groups params
func (o *LDAPAdminGroupsParams) SetBody(body *models.LdapAdminGroupsPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LDAPAdminGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
