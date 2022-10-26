// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BackupToS3Reader is a Reader for the BackupToS3 structure.
type BackupToS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupToS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBackupToS3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupToS3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupToS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupToS3NoContent creates a BackupToS3NoContent with default headers values
func NewBackupToS3NoContent() *BackupToS3NoContent {
	return &BackupToS3NoContent{}
}

/*
BackupToS3NoContent describes a response with status code 204, with default header values.

Success
*/
type BackupToS3NoContent struct {
}

// IsSuccess returns true when this backup to s3 no content response has a 2xx status code
func (o *BackupToS3NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup to s3 no content response has a 3xx status code
func (o *BackupToS3NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup to s3 no content response has a 4xx status code
func (o *BackupToS3NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup to s3 no content response has a 5xx status code
func (o *BackupToS3NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this backup to s3 no content response a status code equal to that given
func (o *BackupToS3NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *BackupToS3NoContent) Error() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3NoContent ", 204)
}

func (o *BackupToS3NoContent) String() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3NoContent ", 204)
}

func (o *BackupToS3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupToS3BadRequest creates a BackupToS3BadRequest with default headers values
func NewBackupToS3BadRequest() *BackupToS3BadRequest {
	return &BackupToS3BadRequest{}
}

/*
BackupToS3BadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type BackupToS3BadRequest struct {
}

// IsSuccess returns true when this backup to s3 bad request response has a 2xx status code
func (o *BackupToS3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup to s3 bad request response has a 3xx status code
func (o *BackupToS3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup to s3 bad request response has a 4xx status code
func (o *BackupToS3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup to s3 bad request response has a 5xx status code
func (o *BackupToS3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup to s3 bad request response a status code equal to that given
func (o *BackupToS3BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupToS3BadRequest) Error() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3BadRequest ", 400)
}

func (o *BackupToS3BadRequest) String() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3BadRequest ", 400)
}

func (o *BackupToS3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupToS3InternalServerError creates a BackupToS3InternalServerError with default headers values
func NewBackupToS3InternalServerError() *BackupToS3InternalServerError {
	return &BackupToS3InternalServerError{}
}

/*
BackupToS3InternalServerError describes a response with status code 500, with default header values.

Server error
*/
type BackupToS3InternalServerError struct {
}

// IsSuccess returns true when this backup to s3 internal server error response has a 2xx status code
func (o *BackupToS3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup to s3 internal server error response has a 3xx status code
func (o *BackupToS3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup to s3 internal server error response has a 4xx status code
func (o *BackupToS3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup to s3 internal server error response has a 5xx status code
func (o *BackupToS3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup to s3 internal server error response a status code equal to that given
func (o *BackupToS3InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupToS3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3InternalServerError ", 500)
}

func (o *BackupToS3InternalServerError) String() string {
	return fmt.Sprintf("[POST /backup/s3/execute][%d] backupToS3InternalServerError ", 500)
}

func (o *BackupToS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
