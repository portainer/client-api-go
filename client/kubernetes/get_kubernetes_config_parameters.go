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

// NewGetKubernetesConfigParams creates a new GetKubernetesConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubernetesConfigParams() *GetKubernetesConfigParams {
	return &GetKubernetesConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesConfigParamsWithTimeout creates a new GetKubernetesConfigParams object
// with the ability to set a timeout on a request.
func NewGetKubernetesConfigParamsWithTimeout(timeout time.Duration) *GetKubernetesConfigParams {
	return &GetKubernetesConfigParams{
		timeout: timeout,
	}
}

// NewGetKubernetesConfigParamsWithContext creates a new GetKubernetesConfigParams object
// with the ability to set a context for a request.
func NewGetKubernetesConfigParamsWithContext(ctx context.Context) *GetKubernetesConfigParams {
	return &GetKubernetesConfigParams{
		Context: ctx,
	}
}

// NewGetKubernetesConfigParamsWithHTTPClient creates a new GetKubernetesConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubernetesConfigParamsWithHTTPClient(client *http.Client) *GetKubernetesConfigParams {
	return &GetKubernetesConfigParams{
		HTTPClient: client,
	}
}

/*
GetKubernetesConfigParams contains all the parameters to send to the API endpoint

	for the get kubernetes config operation.

	Typically these are written to a http.Request.
*/
type GetKubernetesConfigParams struct {

	/* ExcludeIds.

	   will exclude these environments(endpoints)
	*/
	ExcludeIds []int64

	/* Ids.

	   will include only these environments(endpoints)
	*/
	Ids []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubernetes config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesConfigParams) WithDefaults() *GetKubernetesConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubernetes config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubernetesConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubernetes config params
func (o *GetKubernetesConfigParams) WithTimeout(timeout time.Duration) *GetKubernetesConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes config params
func (o *GetKubernetesConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes config params
func (o *GetKubernetesConfigParams) WithContext(ctx context.Context) *GetKubernetesConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes config params
func (o *GetKubernetesConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes config params
func (o *GetKubernetesConfigParams) WithHTTPClient(client *http.Client) *GetKubernetesConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes config params
func (o *GetKubernetesConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeIds adds the excludeIds to the get kubernetes config params
func (o *GetKubernetesConfigParams) WithExcludeIds(excludeIds []int64) *GetKubernetesConfigParams {
	o.SetExcludeIds(excludeIds)
	return o
}

// SetExcludeIds adds the excludeIds to the get kubernetes config params
func (o *GetKubernetesConfigParams) SetExcludeIds(excludeIds []int64) {
	o.ExcludeIds = excludeIds
}

// WithIds adds the ids to the get kubernetes config params
func (o *GetKubernetesConfigParams) WithIds(ids []int64) *GetKubernetesConfigParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get kubernetes config params
func (o *GetKubernetesConfigParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeIds != nil {

		// binding items for excludeIds
		joinedExcludeIds := o.bindParamExcludeIds(reg)

		// query array param excludeIds
		if err := r.SetQueryParam("excludeIds", joinedExcludeIds...); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetKubernetesConfig binds the parameter excludeIds
func (o *GetKubernetesConfigParams) bindParamExcludeIds(formats strfmt.Registry) []string {
	excludeIdsIR := o.ExcludeIds

	var excludeIdsIC []string
	for _, excludeIdsIIR := range excludeIdsIR { // explode []int64

		excludeIdsIIV := swag.FormatInt64(excludeIdsIIR) // int64 as string
		excludeIdsIC = append(excludeIdsIC, excludeIdsIIV)
	}

	// items.CollectionFormat: "csv"
	excludeIdsIS := swag.JoinByFormat(excludeIdsIC, "csv")

	return excludeIdsIS
}

// bindParamGetKubernetesConfig binds the parameter ids
func (o *GetKubernetesConfigParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []int64

		idsIIV := swag.FormatInt64(idsIIR) // int64 as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
