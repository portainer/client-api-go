// Code generated by go-swagger; DO NOT EDIT.

package kaas

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

	"github.com/portainer/client-api-go/ee/v2/models"
)

// NewProvisionKaaSClusterParams creates a new ProvisionKaaSClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProvisionKaaSClusterParams() *ProvisionKaaSClusterParams {
	return &ProvisionKaaSClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProvisionKaaSClusterParamsWithTimeout creates a new ProvisionKaaSClusterParams object
// with the ability to set a timeout on a request.
func NewProvisionKaaSClusterParamsWithTimeout(timeout time.Duration) *ProvisionKaaSClusterParams {
	return &ProvisionKaaSClusterParams{
		timeout: timeout,
	}
}

// NewProvisionKaaSClusterParamsWithContext creates a new ProvisionKaaSClusterParams object
// with the ability to set a context for a request.
func NewProvisionKaaSClusterParamsWithContext(ctx context.Context) *ProvisionKaaSClusterParams {
	return &ProvisionKaaSClusterParams{
		Context: ctx,
	}
}

// NewProvisionKaaSClusterParamsWithHTTPClient creates a new ProvisionKaaSClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewProvisionKaaSClusterParamsWithHTTPClient(client *http.Client) *ProvisionKaaSClusterParams {
	return &ProvisionKaaSClusterParams{
		HTTPClient: client,
	}
}

/*
ProvisionKaaSClusterParams contains all the parameters to send to the API endpoint

	for the provision kaa s cluster operation.

	Typically these are written to a http.Request.
*/
type ProvisionKaaSClusterParams struct {

	/* BodyAmazon.

	   KaaS cluster provisioning details (required when provider is amazon)
	*/
	BodyAmazon *models.ProvidersAmazonProvisionPayload

	/* BodyAPI.

	   KaaS cluster provisioning details (required when provider is civo, digitalocean or linode)
	*/
	BodyAPI *models.ProvidersDefaultProvisionPayload

	/* BodyAzure.

	   KaaS cluster provisioning details (required when provider is azure)
	*/
	BodyAzure *models.ProvidersAzureProvisionPayload

	/* BodyGke.

	   KaaS cluster provisioning details (required when provider is gke)
	*/
	BodyGke *models.ProvidersGKEProvisionPayload

	/* Provider.

	   Provider
	*/
	Provider int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provision kaa s cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionKaaSClusterParams) WithDefaults() *ProvisionKaaSClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provision kaa s cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProvisionKaaSClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithTimeout(timeout time.Duration) *ProvisionKaaSClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithContext(ctx context.Context) *ProvisionKaaSClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithHTTPClient(client *http.Client) *ProvisionKaaSClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBodyAmazon adds the bodyAmazon to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithBodyAmazon(bodyAmazon *models.ProvidersAmazonProvisionPayload) *ProvisionKaaSClusterParams {
	o.SetBodyAmazon(bodyAmazon)
	return o
}

// SetBodyAmazon adds the bodyAmazon to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetBodyAmazon(bodyAmazon *models.ProvidersAmazonProvisionPayload) {
	o.BodyAmazon = bodyAmazon
}

// WithBodyAPI adds the bodyAPI to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithBodyAPI(bodyAPI *models.ProvidersDefaultProvisionPayload) *ProvisionKaaSClusterParams {
	o.SetBodyAPI(bodyAPI)
	return o
}

// SetBodyAPI adds the bodyApi to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetBodyAPI(bodyAPI *models.ProvidersDefaultProvisionPayload) {
	o.BodyAPI = bodyAPI
}

// WithBodyAzure adds the bodyAzure to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithBodyAzure(bodyAzure *models.ProvidersAzureProvisionPayload) *ProvisionKaaSClusterParams {
	o.SetBodyAzure(bodyAzure)
	return o
}

// SetBodyAzure adds the bodyAzure to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetBodyAzure(bodyAzure *models.ProvidersAzureProvisionPayload) {
	o.BodyAzure = bodyAzure
}

// WithBodyGke adds the bodyGke to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithBodyGke(bodyGke *models.ProvidersGKEProvisionPayload) *ProvisionKaaSClusterParams {
	o.SetBodyGke(bodyGke)
	return o
}

// SetBodyGke adds the bodyGke to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetBodyGke(bodyGke *models.ProvidersGKEProvisionPayload) {
	o.BodyGke = bodyGke
}

// WithProvider adds the provider to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) WithProvider(provider int64) *ProvisionKaaSClusterParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the provision kaa s cluster params
func (o *ProvisionKaaSClusterParams) SetProvider(provider int64) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisionKaaSClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BodyAmazon != nil {
		if err := r.SetBodyParam(o.BodyAmazon); err != nil {
			return err
		}
	}
	if o.BodyAPI != nil {
		if err := r.SetBodyParam(o.BodyAPI); err != nil {
			return err
		}
	}
	if o.BodyAzure != nil {
		if err := r.SetBodyParam(o.BodyAzure); err != nil {
			return err
		}
	}
	if o.BodyGke != nil {
		if err := r.SetBodyParam(o.BodyGke); err != nil {
			return err
		}
	}

	// path param provider
	if err := r.SetPathParam("provider", swag.FormatInt64(o.Provider)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
