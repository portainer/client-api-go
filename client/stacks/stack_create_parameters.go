// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

	"github.com/portainer/client-api-go/models"
)

// NewStackCreateParams creates a new StackCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStackCreateParams() *StackCreateParams {
	return &StackCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStackCreateParamsWithTimeout creates a new StackCreateParams object
// with the ability to set a timeout on a request.
func NewStackCreateParamsWithTimeout(timeout time.Duration) *StackCreateParams {
	return &StackCreateParams{
		timeout: timeout,
	}
}

// NewStackCreateParamsWithContext creates a new StackCreateParams object
// with the ability to set a context for a request.
func NewStackCreateParamsWithContext(ctx context.Context) *StackCreateParams {
	return &StackCreateParams{
		Context: ctx,
	}
}

// NewStackCreateParamsWithHTTPClient creates a new StackCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStackCreateParamsWithHTTPClient(client *http.Client) *StackCreateParams {
	return &StackCreateParams{
		HTTPClient: client,
	}
}

/*
StackCreateParams contains all the parameters to send to the API endpoint

	for the stack create operation.

	Typically these are written to a http.Request.
*/
type StackCreateParams struct {

	/* Env.

	   Environment(Endpoint) variables passed during deployment, represented as a JSON array [{'name': 'name', 'value': 'value'}]. Optional, used when method equals file and type equals 1.
	*/
	Env *string

	/* Name.

	   Name of the stack. required when method is file
	*/
	Name *string

	/* SwarmID.

	   Swarm cluster identifier. Required when method equals file and type equals 1. required when method is file
	*/
	SwarmID *string

	/* BodyComposeRepository.

	   Required when using method=repository and type=2
	*/
	BodyComposeRepository *models.StacksComposeStackFromGitRepositoryPayload

	/* BodyComposeString.

	   Required when using method=string and type=2
	*/
	BodyComposeString *models.StacksComposeStackFromFileContentPayload

	/* BodySwarmRepository.

	   Required when using method=repository and type=1
	*/
	BodySwarmRepository *models.StacksSwarmStackFromGitRepositoryPayload

	/* BodySwarmString.

	   Required when using method=string and type=1
	*/
	BodySwarmString *models.StacksSwarmStackFromFileContentPayload

	/* EndpointID.

	   Identifier of the environment(endpoint) that will be used to deploy the stack
	*/
	EndpointID int64

	/* File.

	   Stack file. required when method is file
	*/
	File runtime.NamedReadCloser

	/* Method.

	   Stack deployment method. Possible values: file, string or repository.
	*/
	Method string

	/* Type.

	   Stack deployment type. Possible values: 1 (Swarm stack) or 2 (Compose stack).
	*/
	Type int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stack create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateParams) WithDefaults() *StackCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stack create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StackCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stack create params
func (o *StackCreateParams) WithTimeout(timeout time.Duration) *StackCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack create params
func (o *StackCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack create params
func (o *StackCreateParams) WithContext(ctx context.Context) *StackCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack create params
func (o *StackCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack create params
func (o *StackCreateParams) WithHTTPClient(client *http.Client) *StackCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack create params
func (o *StackCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnv adds the env to the stack create params
func (o *StackCreateParams) WithEnv(env *string) *StackCreateParams {
	o.SetEnv(env)
	return o
}

// SetEnv adds the env to the stack create params
func (o *StackCreateParams) SetEnv(env *string) {
	o.Env = env
}

// WithName adds the name to the stack create params
func (o *StackCreateParams) WithName(name *string) *StackCreateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the stack create params
func (o *StackCreateParams) SetName(name *string) {
	o.Name = name
}

// WithSwarmID adds the swarmID to the stack create params
func (o *StackCreateParams) WithSwarmID(swarmID *string) *StackCreateParams {
	o.SetSwarmID(swarmID)
	return o
}

// SetSwarmID adds the swarmId to the stack create params
func (o *StackCreateParams) SetSwarmID(swarmID *string) {
	o.SwarmID = swarmID
}

// WithBodyComposeRepository adds the bodyComposeRepository to the stack create params
func (o *StackCreateParams) WithBodyComposeRepository(bodyComposeRepository *models.StacksComposeStackFromGitRepositoryPayload) *StackCreateParams {
	o.SetBodyComposeRepository(bodyComposeRepository)
	return o
}

// SetBodyComposeRepository adds the bodyComposeRepository to the stack create params
func (o *StackCreateParams) SetBodyComposeRepository(bodyComposeRepository *models.StacksComposeStackFromGitRepositoryPayload) {
	o.BodyComposeRepository = bodyComposeRepository
}

// WithBodyComposeString adds the bodyComposeString to the stack create params
func (o *StackCreateParams) WithBodyComposeString(bodyComposeString *models.StacksComposeStackFromFileContentPayload) *StackCreateParams {
	o.SetBodyComposeString(bodyComposeString)
	return o
}

// SetBodyComposeString adds the bodyComposeString to the stack create params
func (o *StackCreateParams) SetBodyComposeString(bodyComposeString *models.StacksComposeStackFromFileContentPayload) {
	o.BodyComposeString = bodyComposeString
}

// WithBodySwarmRepository adds the bodySwarmRepository to the stack create params
func (o *StackCreateParams) WithBodySwarmRepository(bodySwarmRepository *models.StacksSwarmStackFromGitRepositoryPayload) *StackCreateParams {
	o.SetBodySwarmRepository(bodySwarmRepository)
	return o
}

// SetBodySwarmRepository adds the bodySwarmRepository to the stack create params
func (o *StackCreateParams) SetBodySwarmRepository(bodySwarmRepository *models.StacksSwarmStackFromGitRepositoryPayload) {
	o.BodySwarmRepository = bodySwarmRepository
}

// WithBodySwarmString adds the bodySwarmString to the stack create params
func (o *StackCreateParams) WithBodySwarmString(bodySwarmString *models.StacksSwarmStackFromFileContentPayload) *StackCreateParams {
	o.SetBodySwarmString(bodySwarmString)
	return o
}

// SetBodySwarmString adds the bodySwarmString to the stack create params
func (o *StackCreateParams) SetBodySwarmString(bodySwarmString *models.StacksSwarmStackFromFileContentPayload) {
	o.BodySwarmString = bodySwarmString
}

// WithEndpointID adds the endpointID to the stack create params
func (o *StackCreateParams) WithEndpointID(endpointID int64) *StackCreateParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the stack create params
func (o *StackCreateParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WithFile adds the file to the stack create params
func (o *StackCreateParams) WithFile(file runtime.NamedReadCloser) *StackCreateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the stack create params
func (o *StackCreateParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithMethod adds the method to the stack create params
func (o *StackCreateParams) WithMethod(method string) *StackCreateParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the stack create params
func (o *StackCreateParams) SetMethod(method string) {
	o.Method = method
}

// WithType adds the typeVar to the stack create params
func (o *StackCreateParams) WithType(typeVar int64) *StackCreateParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the stack create params
func (o *StackCreateParams) SetType(typeVar int64) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *StackCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Env != nil {

		// form param Env
		var frEnv string
		if o.Env != nil {
			frEnv = *o.Env
		}
		fEnv := frEnv
		if fEnv != "" {
			if err := r.SetFormParam("Env", fEnv); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// form param Name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("Name", fName); err != nil {
				return err
			}
		}
	}

	if o.SwarmID != nil {

		// form param SwarmID
		var frSwarmID string
		if o.SwarmID != nil {
			frSwarmID = *o.SwarmID
		}
		fSwarmID := frSwarmID
		if fSwarmID != "" {
			if err := r.SetFormParam("SwarmID", fSwarmID); err != nil {
				return err
			}
		}
	}
	if o.BodyComposeRepository != nil {
		if err := r.SetBodyParam(o.BodyComposeRepository); err != nil {
			return err
		}
	}
	if o.BodyComposeString != nil {
		if err := r.SetBodyParam(o.BodyComposeString); err != nil {
			return err
		}
	}
	if o.BodySwarmRepository != nil {
		if err := r.SetBodyParam(o.BodySwarmRepository); err != nil {
			return err
		}
	}
	if o.BodySwarmString != nil {
		if err := r.SetBodyParam(o.BodySwarmString); err != nil {
			return err
		}
	}

	// query param endpointId
	qrEndpointID := o.EndpointID
	qEndpointID := swag.FormatInt64(qrEndpointID)
	if qEndpointID != "" {

		if err := r.SetQueryParam("endpointId", qEndpointID); err != nil {
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

	// query param type
	qrType := o.Type
	qType := swag.FormatInt64(qrType)
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
