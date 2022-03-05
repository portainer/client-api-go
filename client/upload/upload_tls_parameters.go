// Code generated by go-swagger; DO NOT EDIT.

package upload

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

// NewUploadTLSParams creates a new UploadTLSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadTLSParams() *UploadTLSParams {
	return &UploadTLSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadTLSParamsWithTimeout creates a new UploadTLSParams object
// with the ability to set a timeout on a request.
func NewUploadTLSParamsWithTimeout(timeout time.Duration) *UploadTLSParams {
	return &UploadTLSParams{
		timeout: timeout,
	}
}

// NewUploadTLSParamsWithContext creates a new UploadTLSParams object
// with the ability to set a context for a request.
func NewUploadTLSParamsWithContext(ctx context.Context) *UploadTLSParams {
	return &UploadTLSParams{
		Context: ctx,
	}
}

// NewUploadTLSParamsWithHTTPClient creates a new UploadTLSParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadTLSParamsWithHTTPClient(client *http.Client) *UploadTLSParams {
	return &UploadTLSParams{
		HTTPClient: client,
	}
}

/* UploadTLSParams contains all the parameters to send to the API endpoint
   for the upload TLS operation.

   Typically these are written to a http.Request.
*/
type UploadTLSParams struct {

	/* Certificate.

	   TLS file type. Valid values are 'ca', 'cert' or 'key'.
	*/
	Certificate string

	/* File.

	   The file to upload
	*/
	File runtime.NamedReadCloser

	/* Folder.

	   Folder where the TLS file will be stored. Will be created if not existing
	*/
	Folder string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload TLS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadTLSParams) WithDefaults() *UploadTLSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload TLS params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadTLSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload TLS params
func (o *UploadTLSParams) WithTimeout(timeout time.Duration) *UploadTLSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload TLS params
func (o *UploadTLSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload TLS params
func (o *UploadTLSParams) WithContext(ctx context.Context) *UploadTLSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload TLS params
func (o *UploadTLSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload TLS params
func (o *UploadTLSParams) WithHTTPClient(client *http.Client) *UploadTLSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload TLS params
func (o *UploadTLSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificate adds the certificate to the upload TLS params
func (o *UploadTLSParams) WithCertificate(certificate string) *UploadTLSParams {
	o.SetCertificate(certificate)
	return o
}

// SetCertificate adds the certificate to the upload TLS params
func (o *UploadTLSParams) SetCertificate(certificate string) {
	o.Certificate = certificate
}

// WithFile adds the file to the upload TLS params
func (o *UploadTLSParams) WithFile(file runtime.NamedReadCloser) *UploadTLSParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload TLS params
func (o *UploadTLSParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithFolder adds the folder to the upload TLS params
func (o *UploadTLSParams) WithFolder(folder string) *UploadTLSParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the upload TLS params
func (o *UploadTLSParams) SetFolder(folder string) {
	o.Folder = folder
}

// WriteToRequest writes these params to a swagger request
func (o *UploadTLSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificate
	if err := r.SetPathParam("certificate", o.Certificate); err != nil {
		return err
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param folder
	frFolder := o.Folder
	fFolder := frFolder
	if fFolder != "" {
		if err := r.SetFormParam("folder", fFolder); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
