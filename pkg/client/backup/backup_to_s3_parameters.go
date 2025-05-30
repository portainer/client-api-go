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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// NewBackupToS3Params creates a new BackupToS3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupToS3Params() *BackupToS3Params {
	return &BackupToS3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupToS3ParamsWithTimeout creates a new BackupToS3Params object
// with the ability to set a timeout on a request.
func NewBackupToS3ParamsWithTimeout(timeout time.Duration) *BackupToS3Params {
	return &BackupToS3Params{
		timeout: timeout,
	}
}

// NewBackupToS3ParamsWithContext creates a new BackupToS3Params object
// with the ability to set a context for a request.
func NewBackupToS3ParamsWithContext(ctx context.Context) *BackupToS3Params {
	return &BackupToS3Params{
		Context: ctx,
	}
}

// NewBackupToS3ParamsWithHTTPClient creates a new BackupToS3Params object
// with the ability to set a custom HTTPClient for a request.
func NewBackupToS3ParamsWithHTTPClient(client *http.Client) *BackupToS3Params {
	return &BackupToS3Params{
		HTTPClient: client,
	}
}

/*
BackupToS3Params contains all the parameters to send to the API endpoint

	for the backup to s3 operation.

	Typically these are written to a http.Request.
*/
type BackupToS3Params struct {

	/* Body.

	   S3 backup settings
	*/
	Body *models.BackupS3BackupPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup to s3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupToS3Params) WithDefaults() *BackupToS3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup to s3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupToS3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup to s3 params
func (o *BackupToS3Params) WithTimeout(timeout time.Duration) *BackupToS3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup to s3 params
func (o *BackupToS3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup to s3 params
func (o *BackupToS3Params) WithContext(ctx context.Context) *BackupToS3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup to s3 params
func (o *BackupToS3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup to s3 params
func (o *BackupToS3Params) WithHTTPClient(client *http.Client) *BackupToS3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup to s3 params
func (o *BackupToS3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the backup to s3 params
func (o *BackupToS3Params) WithBody(body *models.BackupS3BackupPayload) *BackupToS3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup to s3 params
func (o *BackupToS3Params) SetBody(body *models.BackupS3BackupPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BackupToS3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
