// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

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

	"github.com/portainer/client-api-go/v2/models"
)

// NewCustomTemplateCreateParams creates a new CustomTemplateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomTemplateCreateParams() *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomTemplateCreateParamsWithTimeout creates a new CustomTemplateCreateParams object
// with the ability to set a timeout on a request.
func NewCustomTemplateCreateParamsWithTimeout(timeout time.Duration) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		timeout: timeout,
	}
}

// NewCustomTemplateCreateParamsWithContext creates a new CustomTemplateCreateParams object
// with the ability to set a context for a request.
func NewCustomTemplateCreateParamsWithContext(ctx context.Context) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		Context: ctx,
	}
}

// NewCustomTemplateCreateParamsWithHTTPClient creates a new CustomTemplateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomTemplateCreateParamsWithHTTPClient(client *http.Client) *CustomTemplateCreateParams {
	return &CustomTemplateCreateParams{
		HTTPClient: client,
	}
}

/*
CustomTemplateCreateParams contains all the parameters to send to the API endpoint

	for the custom template create operation.

	Typically these are written to a http.Request.
*/
type CustomTemplateCreateParams struct {

	/* Description.

	   Description of the template. required when method is file
	*/
	Description *string

	/* Note.

	   A note that will be displayed in the UI. Supports HTML content
	*/
	Note *string

	/* Platform.

	   Platform associated to the template (1 - 'linux', 2 - 'windows'). required when method is file
	*/
	Platform *int64

	/* Title.

	   Title of the template. required when method is file
	*/
	Title *string

	/* Type.

	   Type of created stack (1 - swarm, 2 - compose), required when method is file
	*/
	Type *int64

	/* BodyRepository.

	   Required when using method=repository
	*/
	BodyRepository *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload

	/* BodyString.

	   Required when using method=string
	*/
	BodyString *models.CustomtemplatesCustomTemplateFromFileContentPayload

	/* File.

	   required when method is file
	*/
	File runtime.NamedReadCloser

	/* Method.

	   method for creating template
	*/
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateParams) WithDefaults() *CustomTemplateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomTemplateCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom template create params
func (o *CustomTemplateCreateParams) WithTimeout(timeout time.Duration) *CustomTemplateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom template create params
func (o *CustomTemplateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom template create params
func (o *CustomTemplateCreateParams) WithContext(ctx context.Context) *CustomTemplateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom template create params
func (o *CustomTemplateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom template create params
func (o *CustomTemplateCreateParams) WithHTTPClient(client *http.Client) *CustomTemplateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom template create params
func (o *CustomTemplateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the custom template create params
func (o *CustomTemplateCreateParams) WithDescription(description *string) *CustomTemplateCreateParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the custom template create params
func (o *CustomTemplateCreateParams) SetDescription(description *string) {
	o.Description = description
}

// WithNote adds the note to the custom template create params
func (o *CustomTemplateCreateParams) WithNote(note *string) *CustomTemplateCreateParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the custom template create params
func (o *CustomTemplateCreateParams) SetNote(note *string) {
	o.Note = note
}

// WithPlatform adds the platform to the custom template create params
func (o *CustomTemplateCreateParams) WithPlatform(platform *int64) *CustomTemplateCreateParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the custom template create params
func (o *CustomTemplateCreateParams) SetPlatform(platform *int64) {
	o.Platform = platform
}

// WithTitle adds the title to the custom template create params
func (o *CustomTemplateCreateParams) WithTitle(title *string) *CustomTemplateCreateParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the custom template create params
func (o *CustomTemplateCreateParams) SetTitle(title *string) {
	o.Title = title
}

// WithType adds the typeVar to the custom template create params
func (o *CustomTemplateCreateParams) WithType(typeVar *int64) *CustomTemplateCreateParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the custom template create params
func (o *CustomTemplateCreateParams) SetType(typeVar *int64) {
	o.Type = typeVar
}

// WithBodyRepository adds the bodyRepository to the custom template create params
func (o *CustomTemplateCreateParams) WithBodyRepository(bodyRepository *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload) *CustomTemplateCreateParams {
	o.SetBodyRepository(bodyRepository)
	return o
}

// SetBodyRepository adds the bodyRepository to the custom template create params
func (o *CustomTemplateCreateParams) SetBodyRepository(bodyRepository *models.CustomtemplatesCustomTemplateFromGitRepositoryPayload) {
	o.BodyRepository = bodyRepository
}

// WithBodyString adds the bodyString to the custom template create params
func (o *CustomTemplateCreateParams) WithBodyString(bodyString *models.CustomtemplatesCustomTemplateFromFileContentPayload) *CustomTemplateCreateParams {
	o.SetBodyString(bodyString)
	return o
}

// SetBodyString adds the bodyString to the custom template create params
func (o *CustomTemplateCreateParams) SetBodyString(bodyString *models.CustomtemplatesCustomTemplateFromFileContentPayload) {
	o.BodyString = bodyString
}

// WithFile adds the file to the custom template create params
func (o *CustomTemplateCreateParams) WithFile(file runtime.NamedReadCloser) *CustomTemplateCreateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the custom template create params
func (o *CustomTemplateCreateParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithMethod adds the method to the custom template create params
func (o *CustomTemplateCreateParams) WithMethod(method string) *CustomTemplateCreateParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the custom template create params
func (o *CustomTemplateCreateParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *CustomTemplateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// form param Description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("Description", fDescription); err != nil {
				return err
			}
		}
	}

	if o.Note != nil {

		// form param Note
		var frNote string
		if o.Note != nil {
			frNote = *o.Note
		}
		fNote := frNote
		if fNote != "" {
			if err := r.SetFormParam("Note", fNote); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// form param Platform
		var frPlatform int64
		if o.Platform != nil {
			frPlatform = *o.Platform
		}
		fPlatform := swag.FormatInt64(frPlatform)
		if fPlatform != "" {
			if err := r.SetFormParam("Platform", fPlatform); err != nil {
				return err
			}
		}
	}

	if o.Title != nil {

		// form param Title
		var frTitle string
		if o.Title != nil {
			frTitle = *o.Title
		}
		fTitle := frTitle
		if fTitle != "" {
			if err := r.SetFormParam("Title", fTitle); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// form param Type
		var frType int64
		if o.Type != nil {
			frType = *o.Type
		}
		fType := swag.FormatInt64(frType)
		if fType != "" {
			if err := r.SetFormParam("Type", fType); err != nil {
				return err
			}
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

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
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
