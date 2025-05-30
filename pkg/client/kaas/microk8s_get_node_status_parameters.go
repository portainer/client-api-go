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
)

// NewMicrok8sGetNodeStatusParams creates a new Microk8sGetNodeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMicrok8sGetNodeStatusParams() *Microk8sGetNodeStatusParams {
	return &Microk8sGetNodeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMicrok8sGetNodeStatusParamsWithTimeout creates a new Microk8sGetNodeStatusParams object
// with the ability to set a timeout on a request.
func NewMicrok8sGetNodeStatusParamsWithTimeout(timeout time.Duration) *Microk8sGetNodeStatusParams {
	return &Microk8sGetNodeStatusParams{
		timeout: timeout,
	}
}

// NewMicrok8sGetNodeStatusParamsWithContext creates a new Microk8sGetNodeStatusParams object
// with the ability to set a context for a request.
func NewMicrok8sGetNodeStatusParamsWithContext(ctx context.Context) *Microk8sGetNodeStatusParams {
	return &Microk8sGetNodeStatusParams{
		Context: ctx,
	}
}

// NewMicrok8sGetNodeStatusParamsWithHTTPClient creates a new Microk8sGetNodeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewMicrok8sGetNodeStatusParamsWithHTTPClient(client *http.Client) *Microk8sGetNodeStatusParams {
	return &Microk8sGetNodeStatusParams{
		HTTPClient: client,
	}
}

/*
Microk8sGetNodeStatusParams contains all the parameters to send to the API endpoint

	for the microk8s get node status operation.

	Typically these are written to a http.Request.
*/
type Microk8sGetNodeStatusParams struct {

	/* EnvironmentID.

	   Environment(Endpoint) identifier
	*/
	EnvironmentID int64

	/* NodeIP.

	   The external node ip of the control plane node.
	*/
	NodeIP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the microk8s get node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Microk8sGetNodeStatusParams) WithDefaults() *Microk8sGetNodeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the microk8s get node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Microk8sGetNodeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) WithTimeout(timeout time.Duration) *Microk8sGetNodeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) WithContext(ctx context.Context) *Microk8sGetNodeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) WithHTTPClient(client *http.Client) *Microk8sGetNodeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) WithEnvironmentID(environmentID int64) *Microk8sGetNodeStatusParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WithNodeIP adds the nodeIP to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) WithNodeIP(nodeIP string) *Microk8sGetNodeStatusParams {
	o.SetNodeIP(nodeIP)
	return o
}

// SetNodeIP adds the nodeIp to the microk8s get node status params
func (o *Microk8sGetNodeStatusParams) SetNodeIP(nodeIP string) {
	o.NodeIP = nodeIP
}

// WriteToRequest writes these params to a swagger request
func (o *Microk8sGetNodeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	// query param nodeIP
	qrNodeIP := o.NodeIP
	qNodeIP := qrNodeIP
	if qNodeIP != "" {

		if err := r.SetQueryParam("nodeIP", qNodeIP); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
