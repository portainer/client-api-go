// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteJobsReader is a Reader for the DeleteJobs structure.
type DeleteJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJobsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kubernetes/{id}/jobs/delete] DeleteJobs", response, response.Code())
	}
}

// NewDeleteJobsNoContent creates a DeleteJobsNoContent with default headers values
func NewDeleteJobsNoContent() *DeleteJobsNoContent {
	return &DeleteJobsNoContent{}
}

/*
DeleteJobsNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteJobsNoContent struct {
}

// IsSuccess returns true when this delete jobs no content response has a 2xx status code
func (o *DeleteJobsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete jobs no content response has a 3xx status code
func (o *DeleteJobsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs no content response has a 4xx status code
func (o *DeleteJobsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete jobs no content response has a 5xx status code
func (o *DeleteJobsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete jobs no content response a status code equal to that given
func (o *DeleteJobsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete jobs no content response
func (o *DeleteJobsNoContent) Code() int {
	return 204
}

func (o *DeleteJobsNoContent) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsNoContent", 204)
}

func (o *DeleteJobsNoContent) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsNoContent", 204)
}

func (o *DeleteJobsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJobsBadRequest creates a DeleteJobsBadRequest with default headers values
func NewDeleteJobsBadRequest() *DeleteJobsBadRequest {
	return &DeleteJobsBadRequest{}
}

/*
DeleteJobsBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type DeleteJobsBadRequest struct {
}

// IsSuccess returns true when this delete jobs bad request response has a 2xx status code
func (o *DeleteJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete jobs bad request response has a 3xx status code
func (o *DeleteJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs bad request response has a 4xx status code
func (o *DeleteJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete jobs bad request response has a 5xx status code
func (o *DeleteJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete jobs bad request response a status code equal to that given
func (o *DeleteJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete jobs bad request response
func (o *DeleteJobsBadRequest) Code() int {
	return 400
}

func (o *DeleteJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsBadRequest", 400)
}

func (o *DeleteJobsBadRequest) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsBadRequest", 400)
}

func (o *DeleteJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJobsUnauthorized creates a DeleteJobsUnauthorized with default headers values
func NewDeleteJobsUnauthorized() *DeleteJobsUnauthorized {
	return &DeleteJobsUnauthorized{}
}

/*
DeleteJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type DeleteJobsUnauthorized struct {
}

// IsSuccess returns true when this delete jobs unauthorized response has a 2xx status code
func (o *DeleteJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete jobs unauthorized response has a 3xx status code
func (o *DeleteJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs unauthorized response has a 4xx status code
func (o *DeleteJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete jobs unauthorized response has a 5xx status code
func (o *DeleteJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete jobs unauthorized response a status code equal to that given
func (o *DeleteJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete jobs unauthorized response
func (o *DeleteJobsUnauthorized) Code() int {
	return 401
}

func (o *DeleteJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsUnauthorized", 401)
}

func (o *DeleteJobsUnauthorized) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsUnauthorized", 401)
}

func (o *DeleteJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJobsForbidden creates a DeleteJobsForbidden with default headers values
func NewDeleteJobsForbidden() *DeleteJobsForbidden {
	return &DeleteJobsForbidden{}
}

/*
DeleteJobsForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type DeleteJobsForbidden struct {
}

// IsSuccess returns true when this delete jobs forbidden response has a 2xx status code
func (o *DeleteJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete jobs forbidden response has a 3xx status code
func (o *DeleteJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs forbidden response has a 4xx status code
func (o *DeleteJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete jobs forbidden response has a 5xx status code
func (o *DeleteJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete jobs forbidden response a status code equal to that given
func (o *DeleteJobsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete jobs forbidden response
func (o *DeleteJobsForbidden) Code() int {
	return 403
}

func (o *DeleteJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsForbidden", 403)
}

func (o *DeleteJobsForbidden) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsForbidden", 403)
}

func (o *DeleteJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJobsNotFound creates a DeleteJobsNotFound with default headers values
func NewDeleteJobsNotFound() *DeleteJobsNotFound {
	return &DeleteJobsNotFound{}
}

/*
DeleteJobsNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier or unable to find a specific service account.
*/
type DeleteJobsNotFound struct {
}

// IsSuccess returns true when this delete jobs not found response has a 2xx status code
func (o *DeleteJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete jobs not found response has a 3xx status code
func (o *DeleteJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs not found response has a 4xx status code
func (o *DeleteJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete jobs not found response has a 5xx status code
func (o *DeleteJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete jobs not found response a status code equal to that given
func (o *DeleteJobsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete jobs not found response
func (o *DeleteJobsNotFound) Code() int {
	return 404
}

func (o *DeleteJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsNotFound", 404)
}

func (o *DeleteJobsNotFound) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsNotFound", 404)
}

func (o *DeleteJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJobsInternalServerError creates a DeleteJobsInternalServerError with default headers values
func NewDeleteJobsInternalServerError() *DeleteJobsInternalServerError {
	return &DeleteJobsInternalServerError{}
}

/*
DeleteJobsInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to delete Jobs.
*/
type DeleteJobsInternalServerError struct {
}

// IsSuccess returns true when this delete jobs internal server error response has a 2xx status code
func (o *DeleteJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete jobs internal server error response has a 3xx status code
func (o *DeleteJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete jobs internal server error response has a 4xx status code
func (o *DeleteJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete jobs internal server error response has a 5xx status code
func (o *DeleteJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete jobs internal server error response a status code equal to that given
func (o *DeleteJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete jobs internal server error response
func (o *DeleteJobsInternalServerError) Code() int {
	return 500
}

func (o *DeleteJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsInternalServerError", 500)
}

func (o *DeleteJobsInternalServerError) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/jobs/delete][%d] deleteJobsInternalServerError", 500)
}

func (o *DeleteJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
