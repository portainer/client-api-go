// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// BackupSettingsFetchReader is a Reader for the BackupSettingsFetch structure.
type BackupSettingsFetchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupSettingsFetchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupSettingsFetchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBackupSettingsFetchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupSettingsFetchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /backup/s3/settings] BackupSettingsFetch", response, response.Code())
	}
}

// NewBackupSettingsFetchOK creates a BackupSettingsFetchOK with default headers values
func NewBackupSettingsFetchOK() *BackupSettingsFetchOK {
	return &BackupSettingsFetchOK{}
}

/*
BackupSettingsFetchOK describes a response with status code 200, with default header values.

Success
*/
type BackupSettingsFetchOK struct {
	Payload *models.PortainereeS3BackupSettings
}

// IsSuccess returns true when this backup settings fetch o k response has a 2xx status code
func (o *BackupSettingsFetchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup settings fetch o k response has a 3xx status code
func (o *BackupSettingsFetchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup settings fetch o k response has a 4xx status code
func (o *BackupSettingsFetchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup settings fetch o k response has a 5xx status code
func (o *BackupSettingsFetchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup settings fetch o k response a status code equal to that given
func (o *BackupSettingsFetchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the backup settings fetch o k response
func (o *BackupSettingsFetchOK) Code() int {
	return 200
}

func (o *BackupSettingsFetchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchOK %s", 200, payload)
}

func (o *BackupSettingsFetchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchOK %s", 200, payload)
}

func (o *BackupSettingsFetchOK) GetPayload() *models.PortainereeS3BackupSettings {
	return o.Payload
}

func (o *BackupSettingsFetchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeS3BackupSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupSettingsFetchUnauthorized creates a BackupSettingsFetchUnauthorized with default headers values
func NewBackupSettingsFetchUnauthorized() *BackupSettingsFetchUnauthorized {
	return &BackupSettingsFetchUnauthorized{}
}

/*
BackupSettingsFetchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupSettingsFetchUnauthorized struct {
}

// IsSuccess returns true when this backup settings fetch unauthorized response has a 2xx status code
func (o *BackupSettingsFetchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup settings fetch unauthorized response has a 3xx status code
func (o *BackupSettingsFetchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup settings fetch unauthorized response has a 4xx status code
func (o *BackupSettingsFetchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup settings fetch unauthorized response has a 5xx status code
func (o *BackupSettingsFetchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup settings fetch unauthorized response a status code equal to that given
func (o *BackupSettingsFetchUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the backup settings fetch unauthorized response
func (o *BackupSettingsFetchUnauthorized) Code() int {
	return 401
}

func (o *BackupSettingsFetchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchUnauthorized", 401)
}

func (o *BackupSettingsFetchUnauthorized) String() string {
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchUnauthorized", 401)
}

func (o *BackupSettingsFetchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupSettingsFetchInternalServerError creates a BackupSettingsFetchInternalServerError with default headers values
func NewBackupSettingsFetchInternalServerError() *BackupSettingsFetchInternalServerError {
	return &BackupSettingsFetchInternalServerError{}
}

/*
BackupSettingsFetchInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type BackupSettingsFetchInternalServerError struct {
}

// IsSuccess returns true when this backup settings fetch internal server error response has a 2xx status code
func (o *BackupSettingsFetchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup settings fetch internal server error response has a 3xx status code
func (o *BackupSettingsFetchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup settings fetch internal server error response has a 4xx status code
func (o *BackupSettingsFetchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup settings fetch internal server error response has a 5xx status code
func (o *BackupSettingsFetchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup settings fetch internal server error response a status code equal to that given
func (o *BackupSettingsFetchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the backup settings fetch internal server error response
func (o *BackupSettingsFetchInternalServerError) Code() int {
	return 500
}

func (o *BackupSettingsFetchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchInternalServerError", 500)
}

func (o *BackupSettingsFetchInternalServerError) String() string {
	return fmt.Sprintf("[GET /backup/s3/settings][%d] backupSettingsFetchInternalServerError", 500)
}

func (o *BackupSettingsFetchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
