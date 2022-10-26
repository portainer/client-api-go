// Code generated by go-swagger; DO NOT EDIT.

package backup

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

	"github.com/portainer/client-api-go/v2/models"
)

// NewBackupParams creates a new BackupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupParams() *BackupParams {
	return &BackupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupParamsWithTimeout creates a new BackupParams object
// with the ability to set a timeout on a request.
func NewBackupParamsWithTimeout(timeout time.Duration) *BackupParams {
	return &BackupParams{
		timeout: timeout,
	}
}

// NewBackupParamsWithContext creates a new BackupParams object
// with the ability to set a context for a request.
func NewBackupParamsWithContext(ctx context.Context) *BackupParams {
	return &BackupParams{
		Context: ctx,
	}
}

// NewBackupParamsWithHTTPClient creates a new BackupParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupParamsWithHTTPClient(client *http.Client) *BackupParams {
	return &BackupParams{
		HTTPClient: client,
	}
}

/*
BackupParams contains all the parameters to send to the API endpoint

	for the backup operation.

	Typically these are written to a http.Request.
*/
type BackupParams struct {

	/* Body.

	   An object contains the password to encrypt the backup with
	*/
	Body *models.BackupBackupPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupParams) WithDefaults() *BackupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup params
func (o *BackupParams) WithTimeout(timeout time.Duration) *BackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup params
func (o *BackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup params
func (o *BackupParams) WithContext(ctx context.Context) *BackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup params
func (o *BackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup params
func (o *BackupParams) WithHTTPClient(client *http.Client) *BackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup params
func (o *BackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the backup params
func (o *BackupParams) WithBody(body *models.BackupBackupPayload) *BackupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup params
func (o *BackupParams) SetBody(body *models.BackupBackupPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
