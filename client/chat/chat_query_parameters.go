// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// NewChatQueryParams creates a new ChatQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChatQueryParams() *ChatQueryParams {
	return &ChatQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChatQueryParamsWithTimeout creates a new ChatQueryParams object
// with the ability to set a timeout on a request.
func NewChatQueryParamsWithTimeout(timeout time.Duration) *ChatQueryParams {
	return &ChatQueryParams{
		timeout: timeout,
	}
}

// NewChatQueryParamsWithContext creates a new ChatQueryParams object
// with the ability to set a context for a request.
func NewChatQueryParamsWithContext(ctx context.Context) *ChatQueryParams {
	return &ChatQueryParams{
		Context: ctx,
	}
}

// NewChatQueryParamsWithHTTPClient creates a new ChatQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewChatQueryParamsWithHTTPClient(client *http.Client) *ChatQueryParams {
	return &ChatQueryParams{
		HTTPClient: client,
	}
}

/*
ChatQueryParams contains all the parameters to send to the API endpoint

	for the chat query operation.

	Typically these are written to a http.Request.
*/
type ChatQueryParams struct {

	/* Body.

	   Query
	*/
	Body *models.ChatChatQueryPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chat query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChatQueryParams) WithDefaults() *ChatQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chat query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChatQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chat query params
func (o *ChatQueryParams) WithTimeout(timeout time.Duration) *ChatQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chat query params
func (o *ChatQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chat query params
func (o *ChatQueryParams) WithContext(ctx context.Context) *ChatQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chat query params
func (o *ChatQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chat query params
func (o *ChatQueryParams) WithHTTPClient(client *http.Client) *ChatQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chat query params
func (o *ChatQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the chat query params
func (o *ChatQueryParams) WithBody(body *models.ChatChatQueryPayload) *ChatQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chat query params
func (o *ChatQueryParams) SetBody(body *models.ChatChatQueryPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChatQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
