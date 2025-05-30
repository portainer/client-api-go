// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewGetAllKubernetesConfigMapsParams creates a new GetAllKubernetesConfigMapsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllKubernetesConfigMapsParams() *GetAllKubernetesConfigMapsParams {
	return &GetAllKubernetesConfigMapsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllKubernetesConfigMapsParamsWithTimeout creates a new GetAllKubernetesConfigMapsParams object
// with the ability to set a timeout on a request.
func NewGetAllKubernetesConfigMapsParamsWithTimeout(timeout time.Duration) *GetAllKubernetesConfigMapsParams {
	return &GetAllKubernetesConfigMapsParams{
		timeout: timeout,
	}
}

// NewGetAllKubernetesConfigMapsParamsWithContext creates a new GetAllKubernetesConfigMapsParams object
// with the ability to set a context for a request.
func NewGetAllKubernetesConfigMapsParamsWithContext(ctx context.Context) *GetAllKubernetesConfigMapsParams {
	return &GetAllKubernetesConfigMapsParams{
		Context: ctx,
	}
}

// NewGetAllKubernetesConfigMapsParamsWithHTTPClient creates a new GetAllKubernetesConfigMapsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllKubernetesConfigMapsParamsWithHTTPClient(client *http.Client) *GetAllKubernetesConfigMapsParams {
	return &GetAllKubernetesConfigMapsParams{
		HTTPClient: client,
	}
}

/*
GetAllKubernetesConfigMapsParams contains all the parameters to send to the API endpoint

	for the get all kubernetes config maps operation.

	Typically these are written to a http.Request.
*/
type GetAllKubernetesConfigMapsParams struct {

	/* ID.

	   Environment identifier
	*/
	ID int64

	/* IsUsed.

	   Set to true to include information about applications that use the ConfigMaps in the response
	*/
	IsUsed bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all kubernetes config maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllKubernetesConfigMapsParams) WithDefaults() *GetAllKubernetesConfigMapsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all kubernetes config maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllKubernetesConfigMapsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) WithTimeout(timeout time.Duration) *GetAllKubernetesConfigMapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) WithContext(ctx context.Context) *GetAllKubernetesConfigMapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) WithHTTPClient(client *http.Client) *GetAllKubernetesConfigMapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) WithID(id int64) *GetAllKubernetesConfigMapsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) SetID(id int64) {
	o.ID = id
}

// WithIsUsed adds the isUsed to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) WithIsUsed(isUsed bool) *GetAllKubernetesConfigMapsParams {
	o.SetIsUsed(isUsed)
	return o
}

// SetIsUsed adds the isUsed to the get all kubernetes config maps params
func (o *GetAllKubernetesConfigMapsParams) SetIsUsed(isUsed bool) {
	o.IsUsed = isUsed
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllKubernetesConfigMapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param isUsed
	qrIsUsed := o.IsUsed
	qIsUsed := swag.FormatBool(qrIsUsed)
	if qIsUsed != "" {

		if err := r.SetQueryParam("isUsed", qIsUsed); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
