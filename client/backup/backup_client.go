// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Backup(params *BackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupOK, error)

	BackupSettingsFetch(params *BackupSettingsFetchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupSettingsFetchOK, error)

	BackupStatusFetch(params *BackupStatusFetchParams, opts ...ClientOption) (*BackupStatusFetchOK, error)

	BackupToS3(params *BackupToS3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupToS3NoContent, error)

	Restore(params *RestoreParams, opts ...ClientOption) (*RestoreOK, error)

	RestoreFromS3(params *RestoreFromS3Params, opts ...ClientOption) (*RestoreFromS3OK, error)

	UpdateS3Settings(params *UpdateS3SettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateS3SettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	Backup creates an archive with a system data snapshot that could be used to restore the system

	Creates an archive with a system data snapshot that could be used to restore the system.

**Access policy**: admin
*/
func (a *Client) Backup(params *BackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup",
		Method:             "POST",
		PathPattern:        "/backup",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupSettingsFetch fetches s3 backup settings configurations

**Access policy**: administrator
*/
func (a *Client) BackupSettingsFetch(params *BackupSettingsFetchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupSettingsFetchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupSettingsFetchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupSettingsFetch",
		Method:             "GET",
		PathPattern:        "/backup/s3/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupSettingsFetchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupSettingsFetchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BackupSettingsFetch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupStatusFetch fetches the status of the last scheduled backup run

**Access policy**: public
*/
func (a *Client) BackupStatusFetch(params *BackupStatusFetchParams, opts ...ClientOption) (*BackupStatusFetchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupStatusFetchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupStatusFetch",
		Method:             "GET",
		PathPattern:        "/backup/s3/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupStatusFetchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupStatusFetchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BackupStatusFetch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	BackupToS3 executes backup to a w s s3 bucket

	Creates an archive with a system data snapshot and upload it to the target S3 bucket

**Access policy**: administrator
*/
func (a *Client) BackupToS3(params *BackupToS3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupToS3NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupToS3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupToS3",
		Method:             "POST",
		PathPattern:        "/backup/s3/execute",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupToS3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupToS3NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BackupToS3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	Restore triggers a system restore using provided backup file

	Triggers a system restore using provided backup file

**Access policy**: public
*/
func (a *Client) Restore(params *RestoreParams, opts ...ClientOption) (*RestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Restore",
		Method:             "POST",
		PathPattern:        "/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Restore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RestoreFromS3 triggers a system restore using details of s3 backup

	Triggers a system restore using details of s3 backup

**Access policy**: public
*/
func (a *Client) RestoreFromS3(params *RestoreFromS3Params, opts ...ClientOption) (*RestoreFromS3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreFromS3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreFromS3",
		Method:             "POST",
		PathPattern:        "/backup/s3/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreFromS3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreFromS3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RestoreFromS3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	UpdateS3Settings updates stored s3 backup settings and updates running cron jobs as needed

	Updates stored s3 backup settings and updates running cron jobs as needed

**Access policy**: administrator
*/
func (a *Client) UpdateS3Settings(params *UpdateS3SettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateS3SettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateS3SettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateS3Settings",
		Method:             "POST",
		PathPattern:        "/backup/s3/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateS3SettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateS3SettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateS3Settings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
