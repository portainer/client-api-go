// Code generated by go-swagger; DO NOT EDIT.

package useractivity

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

// NewAuthLogsListParams creates a new AuthLogsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthLogsListParams() *AuthLogsListParams {
	return &AuthLogsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthLogsListParamsWithTimeout creates a new AuthLogsListParams object
// with the ability to set a timeout on a request.
func NewAuthLogsListParamsWithTimeout(timeout time.Duration) *AuthLogsListParams {
	return &AuthLogsListParams{
		timeout: timeout,
	}
}

// NewAuthLogsListParamsWithContext creates a new AuthLogsListParams object
// with the ability to set a context for a request.
func NewAuthLogsListParamsWithContext(ctx context.Context) *AuthLogsListParams {
	return &AuthLogsListParams{
		Context: ctx,
	}
}

// NewAuthLogsListParamsWithHTTPClient creates a new AuthLogsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthLogsListParamsWithHTTPClient(client *http.Client) *AuthLogsListParams {
	return &AuthLogsListParams{
		HTTPClient: client,
	}
}

/*
AuthLogsListParams contains all the parameters to send to the API endpoint

	for the auth logs list operation.

	Typically these are written to a http.Request.
*/
type AuthLogsListParams struct {

	/* After.

	   Results after timestamp (unix)
	*/
	After *int64

	/* Before.

	   Results before timestamp (unix)
	*/
	Before *int64

	/* Keyword.

	   Query logs by this keyword
	*/
	Keyword *string

	/* Limit.

	   Limit results
	*/
	Limit *int64

	/* Offset.

	   Pagination offset
	*/
	Offset *int64

	/* SortBy.

	   Sort by this column
	*/
	SortBy *string

	/* SortDesc.

	   Sort order, if true will return results by descending order
	*/
	SortDesc *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth logs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthLogsListParams) WithDefaults() *AuthLogsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth logs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthLogsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth logs list params
func (o *AuthLogsListParams) WithTimeout(timeout time.Duration) *AuthLogsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth logs list params
func (o *AuthLogsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth logs list params
func (o *AuthLogsListParams) WithContext(ctx context.Context) *AuthLogsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth logs list params
func (o *AuthLogsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth logs list params
func (o *AuthLogsListParams) WithHTTPClient(client *http.Client) *AuthLogsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth logs list params
func (o *AuthLogsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the auth logs list params
func (o *AuthLogsListParams) WithAfter(after *int64) *AuthLogsListParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the auth logs list params
func (o *AuthLogsListParams) SetAfter(after *int64) {
	o.After = after
}

// WithBefore adds the before to the auth logs list params
func (o *AuthLogsListParams) WithBefore(before *int64) *AuthLogsListParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the auth logs list params
func (o *AuthLogsListParams) SetBefore(before *int64) {
	o.Before = before
}

// WithKeyword adds the keyword to the auth logs list params
func (o *AuthLogsListParams) WithKeyword(keyword *string) *AuthLogsListParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the auth logs list params
func (o *AuthLogsListParams) SetKeyword(keyword *string) {
	o.Keyword = keyword
}

// WithLimit adds the limit to the auth logs list params
func (o *AuthLogsListParams) WithLimit(limit *int64) *AuthLogsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the auth logs list params
func (o *AuthLogsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the auth logs list params
func (o *AuthLogsListParams) WithOffset(offset *int64) *AuthLogsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the auth logs list params
func (o *AuthLogsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSortBy adds the sortBy to the auth logs list params
func (o *AuthLogsListParams) WithSortBy(sortBy *string) *AuthLogsListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the auth logs list params
func (o *AuthLogsListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDesc adds the sortDesc to the auth logs list params
func (o *AuthLogsListParams) WithSortDesc(sortDesc *bool) *AuthLogsListParams {
	o.SetSortDesc(sortDesc)
	return o
}

// SetSortDesc adds the sortDesc to the auth logs list params
func (o *AuthLogsListParams) SetSortDesc(sortDesc *bool) {
	o.SortDesc = sortDesc
}

// WriteToRequest writes these params to a swagger request
func (o *AuthLogsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter int64

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := swag.FormatInt64(qrAfter)
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore int64

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := swag.FormatInt64(qrBefore)
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Keyword != nil {

		// query param keyword
		var qrKeyword string

		if o.Keyword != nil {
			qrKeyword = *o.Keyword
		}
		qKeyword := qrKeyword
		if qKeyword != "" {

			if err := r.SetQueryParam("keyword", qKeyword); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortDesc != nil {

		// query param sortDesc
		var qrSortDesc bool

		if o.SortDesc != nil {
			qrSortDesc = *o.SortDesc
		}
		qSortDesc := swag.FormatBool(qrSortDesc)
		if qSortDesc != "" {

			if err := r.SetQueryParam("sortDesc", qSortDesc); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
