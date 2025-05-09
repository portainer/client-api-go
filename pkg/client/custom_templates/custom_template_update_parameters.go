// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewCustomTemplateUpdateParams creates a new CustomTemplateUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomTemplateUpdateParams() *CustomTemplateUpdateParams {
	return &CustomTemplateUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomTemplateUpdateParamsWithTimeout creates a new CustomTemplateUpdateParams object
// with the ability to set a timeout on a request.
func NewCustomTemplateUpdateParamsWithTimeout(timeout time.Duration) *CustomTemplateUpdateParams {
	return &CustomTemplateUpdateParams{
		timeout: timeout,
	}
}

// NewCustomTemplateUpdateParamsWithContext creates a new CustomTemplateUpdateParams object
// with the ability to set a context for a request.
func NewCustomTemplateUpdateParamsWithContext(ctx context.Context) *CustomTemplateUpdateParams {
	return &CustomTemplateUpdateParams{
		Context: ctx,
	}
}

// NewCustomTemplateUpdateParamsWithHTTPClient creates a new CustomTemplateUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomTemplateUpdateParamsWithHTTPClient(client *http.Client) *CustomTemplateUpdateParams {
	return &CustomTemplateUpdateParams{
		HTTPClient: client,
	}
}

/*
CustomTemplateUpdateParams contains all the parameters to send to the API endpoint

	for the custom template update operation.

	Typically these are written to a http.Request.
*/
type CustomTemplateUpdateParams struct {

	/* Body.

	   Template details
	*/
	Body *models.CustomtemplatesCustomTemplateUpdatePayload

	/* ID.

	   Template identifier
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom template update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateUpdateParams) WithDefaults() *CustomTemplateUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom template update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom template update params
func (o *CustomTemplateUpdateParams) WithTimeout(timeout time.Duration) *CustomTemplateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom template update params
func (o *CustomTemplateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom template update params
func (o *CustomTemplateUpdateParams) WithContext(ctx context.Context) *CustomTemplateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom template update params
func (o *CustomTemplateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom template update params
func (o *CustomTemplateUpdateParams) WithHTTPClient(client *http.Client) *CustomTemplateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom template update params
func (o *CustomTemplateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the custom template update params
func (o *CustomTemplateUpdateParams) WithBody(body *models.CustomtemplatesCustomTemplateUpdatePayload) *CustomTemplateUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the custom template update params
func (o *CustomTemplateUpdateParams) SetBody(body *models.CustomtemplatesCustomTemplateUpdatePayload) {
	o.Body = body
}

// WithID adds the id to the custom template update params
func (o *CustomTemplateUpdateParams) WithID(id int64) *CustomTemplateUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom template update params
func (o *CustomTemplateUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomTemplateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
