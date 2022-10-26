// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// NewSnapshotContainersListParams creates a new SnapshotContainersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotContainersListParams() *SnapshotContainersListParams {
	return &SnapshotContainersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotContainersListParamsWithTimeout creates a new SnapshotContainersListParams object
// with the ability to set a timeout on a request.
func NewSnapshotContainersListParamsWithTimeout(timeout time.Duration) *SnapshotContainersListParams {
	return &SnapshotContainersListParams{
		timeout: timeout,
	}
}

// NewSnapshotContainersListParamsWithContext creates a new SnapshotContainersListParams object
// with the ability to set a context for a request.
func NewSnapshotContainersListParamsWithContext(ctx context.Context) *SnapshotContainersListParams {
	return &SnapshotContainersListParams{
		Context: ctx,
	}
}

// NewSnapshotContainersListParamsWithHTTPClient creates a new SnapshotContainersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotContainersListParamsWithHTTPClient(client *http.Client) *SnapshotContainersListParams {
	return &SnapshotContainersListParams{
		HTTPClient: client,
	}
}

/*
SnapshotContainersListParams contains all the parameters to send to the API endpoint

	for the snapshot containers list operation.

	Typically these are written to a http.Request.
*/
type SnapshotContainersListParams struct {

	/* EdgeStackID.

	   Edge stack identifier, will return only containers for this edge stack
	*/
	EdgeStackID *int64

	/* EnvironmentID.

	   Environment identifier
	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot containers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotContainersListParams) WithDefaults() *SnapshotContainersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot containers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotContainersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot containers list params
func (o *SnapshotContainersListParams) WithTimeout(timeout time.Duration) *SnapshotContainersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot containers list params
func (o *SnapshotContainersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot containers list params
func (o *SnapshotContainersListParams) WithContext(ctx context.Context) *SnapshotContainersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot containers list params
func (o *SnapshotContainersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot containers list params
func (o *SnapshotContainersListParams) WithHTTPClient(client *http.Client) *SnapshotContainersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot containers list params
func (o *SnapshotContainersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeStackID adds the edgeStackID to the snapshot containers list params
func (o *SnapshotContainersListParams) WithEdgeStackID(edgeStackID *int64) *SnapshotContainersListParams {
	o.SetEdgeStackID(edgeStackID)
	return o
}

// SetEdgeStackID adds the edgeStackId to the snapshot containers list params
func (o *SnapshotContainersListParams) SetEdgeStackID(edgeStackID *int64) {
	o.EdgeStackID = edgeStackID
}

// WithEnvironmentID adds the environmentID to the snapshot containers list params
func (o *SnapshotContainersListParams) WithEnvironmentID(environmentID int64) *SnapshotContainersListParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the snapshot containers list params
func (o *SnapshotContainersListParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotContainersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EdgeStackID != nil {

		// query param edgeStackId
		var qrEdgeStackID int64

		if o.EdgeStackID != nil {
			qrEdgeStackID = *o.EdgeStackID
		}
		qEdgeStackID := swag.FormatInt64(qrEdgeStackID)
		if qEdgeStackID != "" {

			if err := r.SetQueryParam("edgeStackId", qEdgeStackID); err != nil {
				return err
			}
		}
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
