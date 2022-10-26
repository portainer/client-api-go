// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/models"
)

// BackupStatusFetchReader is a Reader for the BackupStatusFetch structure.
type BackupStatusFetchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupStatusFetchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupStatusFetchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewBackupStatusFetchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupStatusFetchOK creates a BackupStatusFetchOK with default headers values
func NewBackupStatusFetchOK() *BackupStatusFetchOK {
	return &BackupStatusFetchOK{}
}

/*
BackupStatusFetchOK describes a response with status code 200, with default header values.

Success
*/
type BackupStatusFetchOK struct {
	Payload *models.BackupBackupStatus
}

// IsSuccess returns true when this backup status fetch o k response has a 2xx status code
func (o *BackupStatusFetchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup status fetch o k response has a 3xx status code
func (o *BackupStatusFetchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup status fetch o k response has a 4xx status code
func (o *BackupStatusFetchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup status fetch o k response has a 5xx status code
func (o *BackupStatusFetchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup status fetch o k response a status code equal to that given
func (o *BackupStatusFetchOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupStatusFetchOK) Error() string {
	return fmt.Sprintf("[GET /backup/s3/status][%d] backupStatusFetchOK  %+v", 200, o.Payload)
}

func (o *BackupStatusFetchOK) String() string {
	return fmt.Sprintf("[GET /backup/s3/status][%d] backupStatusFetchOK  %+v", 200, o.Payload)
}

func (o *BackupStatusFetchOK) GetPayload() *models.BackupBackupStatus {
	return o.Payload
}

func (o *BackupStatusFetchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupBackupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupStatusFetchInternalServerError creates a BackupStatusFetchInternalServerError with default headers values
func NewBackupStatusFetchInternalServerError() *BackupStatusFetchInternalServerError {
	return &BackupStatusFetchInternalServerError{}
}

/*
BackupStatusFetchInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type BackupStatusFetchInternalServerError struct {
}

// IsSuccess returns true when this backup status fetch internal server error response has a 2xx status code
func (o *BackupStatusFetchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup status fetch internal server error response has a 3xx status code
func (o *BackupStatusFetchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup status fetch internal server error response has a 4xx status code
func (o *BackupStatusFetchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup status fetch internal server error response has a 5xx status code
func (o *BackupStatusFetchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup status fetch internal server error response a status code equal to that given
func (o *BackupStatusFetchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupStatusFetchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /backup/s3/status][%d] backupStatusFetchInternalServerError ", 500)
}

func (o *BackupStatusFetchInternalServerError) String() string {
	return fmt.Sprintf("[GET /backup/s3/status][%d] backupStatusFetchInternalServerError ", 500)
}

func (o *BackupStatusFetchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
