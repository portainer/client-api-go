// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

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

	"github.com/portainer/client-api/models"
)

// NewPostEdgeStacksParams creates a new PostEdgeStacksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEdgeStacksParams() *PostEdgeStacksParams {
	return &PostEdgeStacksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEdgeStacksParamsWithTimeout creates a new PostEdgeStacksParams object
// with the ability to set a timeout on a request.
func NewPostEdgeStacksParamsWithTimeout(timeout time.Duration) *PostEdgeStacksParams {
	return &PostEdgeStacksParams{
		timeout: timeout,
	}
}

// NewPostEdgeStacksParamsWithContext creates a new PostEdgeStacksParams object
// with the ability to set a context for a request.
func NewPostEdgeStacksParamsWithContext(ctx context.Context) *PostEdgeStacksParams {
	return &PostEdgeStacksParams{
		Context: ctx,
	}
}

// NewPostEdgeStacksParamsWithHTTPClient creates a new PostEdgeStacksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEdgeStacksParamsWithHTTPClient(client *http.Client) *PostEdgeStacksParams {
	return &PostEdgeStacksParams{
		HTTPClient: client,
	}
}

/* PostEdgeStacksParams contains all the parameters to send to the API endpoint
   for the post edge stacks operation.

   Typically these are written to a http.Request.
*/
type PostEdgeStacksParams struct {

	/* BodyFile.

	   Required when using method=file
	*/
	BodyFile *models.EdgestacksSwarmStackFromFileUploadPayload

	/* BodyRepository.

	   Required when using method=repository
	*/
	BodyRepository *models.EdgestacksSwarmStackFromGitRepositoryPayload

	/* BodyString.

	   Required when using method=string
	*/
	BodyString *models.EdgestacksSwarmStackFromFileContentPayload

	/* Method.

	   Creation Method
	*/
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post edge stacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEdgeStacksParams) WithDefaults() *PostEdgeStacksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post edge stacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEdgeStacksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post edge stacks params
func (o *PostEdgeStacksParams) WithTimeout(timeout time.Duration) *PostEdgeStacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post edge stacks params
func (o *PostEdgeStacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post edge stacks params
func (o *PostEdgeStacksParams) WithContext(ctx context.Context) *PostEdgeStacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post edge stacks params
func (o *PostEdgeStacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post edge stacks params
func (o *PostEdgeStacksParams) WithHTTPClient(client *http.Client) *PostEdgeStacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post edge stacks params
func (o *PostEdgeStacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBodyFile adds the bodyFile to the post edge stacks params
func (o *PostEdgeStacksParams) WithBodyFile(bodyFile *models.EdgestacksSwarmStackFromFileUploadPayload) *PostEdgeStacksParams {
	o.SetBodyFile(bodyFile)
	return o
}

// SetBodyFile adds the bodyFile to the post edge stacks params
func (o *PostEdgeStacksParams) SetBodyFile(bodyFile *models.EdgestacksSwarmStackFromFileUploadPayload) {
	o.BodyFile = bodyFile
}

// WithBodyRepository adds the bodyRepository to the post edge stacks params
func (o *PostEdgeStacksParams) WithBodyRepository(bodyRepository *models.EdgestacksSwarmStackFromGitRepositoryPayload) *PostEdgeStacksParams {
	o.SetBodyRepository(bodyRepository)
	return o
}

// SetBodyRepository adds the bodyRepository to the post edge stacks params
func (o *PostEdgeStacksParams) SetBodyRepository(bodyRepository *models.EdgestacksSwarmStackFromGitRepositoryPayload) {
	o.BodyRepository = bodyRepository
}

// WithBodyString adds the bodyString to the post edge stacks params
func (o *PostEdgeStacksParams) WithBodyString(bodyString *models.EdgestacksSwarmStackFromFileContentPayload) *PostEdgeStacksParams {
	o.SetBodyString(bodyString)
	return o
}

// SetBodyString adds the bodyString to the post edge stacks params
func (o *PostEdgeStacksParams) SetBodyString(bodyString *models.EdgestacksSwarmStackFromFileContentPayload) {
	o.BodyString = bodyString
}

// WithMethod adds the method to the post edge stacks params
func (o *PostEdgeStacksParams) WithMethod(method string) *PostEdgeStacksParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the post edge stacks params
func (o *PostEdgeStacksParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *PostEdgeStacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BodyFile != nil {
		if err := r.SetBodyParam(o.BodyFile); err != nil {
			return err
		}
	}
	if o.BodyRepository != nil {
		if err := r.SetBodyParam(o.BodyRepository); err != nil {
			return err
		}
	}
	if o.BodyString != nil {
		if err := r.SetBodyParam(o.BodyString); err != nil {
			return err
		}
	}

	// query param method
	qrMethod := o.Method
	qMethod := qrMethod
	if qMethod != "" {

		if err := r.SetQueryParam("method", qMethod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
