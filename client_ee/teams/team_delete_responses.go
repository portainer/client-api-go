// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamDeleteReader is a Reader for the TeamDelete structure.
type TeamDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTeamDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTeamDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamDeleteNoContent creates a TeamDeleteNoContent with default headers values
func NewTeamDeleteNoContent() *TeamDeleteNoContent {
	return &TeamDeleteNoContent{}
}

/*
TeamDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type TeamDeleteNoContent struct {
}

// IsSuccess returns true when this team delete no content response has a 2xx status code
func (o *TeamDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team delete no content response has a 3xx status code
func (o *TeamDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team delete no content response has a 4xx status code
func (o *TeamDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this team delete no content response has a 5xx status code
func (o *TeamDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this team delete no content response a status code equal to that given
func (o *TeamDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TeamDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNoContent ", 204)
}

func (o *TeamDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNoContent ", 204)
}

func (o *TeamDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamDeleteBadRequest creates a TeamDeleteBadRequest with default headers values
func NewTeamDeleteBadRequest() *TeamDeleteBadRequest {
	return &TeamDeleteBadRequest{}
}

/*
TeamDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamDeleteBadRequest struct {
}

// IsSuccess returns true when this team delete bad request response has a 2xx status code
func (o *TeamDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team delete bad request response has a 3xx status code
func (o *TeamDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team delete bad request response has a 4xx status code
func (o *TeamDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team delete bad request response has a 5xx status code
func (o *TeamDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team delete bad request response a status code equal to that given
func (o *TeamDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TeamDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteBadRequest ", 400)
}

func (o *TeamDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteBadRequest ", 400)
}

func (o *TeamDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamDeleteForbidden creates a TeamDeleteForbidden with default headers values
func NewTeamDeleteForbidden() *TeamDeleteForbidden {
	return &TeamDeleteForbidden{}
}

/*
TeamDeleteForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type TeamDeleteForbidden struct {
}

// IsSuccess returns true when this team delete forbidden response has a 2xx status code
func (o *TeamDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team delete forbidden response has a 3xx status code
func (o *TeamDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team delete forbidden response has a 4xx status code
func (o *TeamDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team delete forbidden response has a 5xx status code
func (o *TeamDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team delete forbidden response a status code equal to that given
func (o *TeamDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TeamDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteForbidden ", 403)
}

func (o *TeamDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteForbidden ", 403)
}

func (o *TeamDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamDeleteNotFound creates a TeamDeleteNotFound with default headers values
func NewTeamDeleteNotFound() *TeamDeleteNotFound {
	return &TeamDeleteNotFound{}
}

/*
TeamDeleteNotFound describes a response with status code 404, with default header values.

Team not found
*/
type TeamDeleteNotFound struct {
}

// IsSuccess returns true when this team delete not found response has a 2xx status code
func (o *TeamDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team delete not found response has a 3xx status code
func (o *TeamDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team delete not found response has a 4xx status code
func (o *TeamDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this team delete not found response has a 5xx status code
func (o *TeamDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this team delete not found response a status code equal to that given
func (o *TeamDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TeamDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNotFound ", 404)
}

func (o *TeamDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteNotFound ", 404)
}

func (o *TeamDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamDeleteInternalServerError creates a TeamDeleteInternalServerError with default headers values
func NewTeamDeleteInternalServerError() *TeamDeleteInternalServerError {
	return &TeamDeleteInternalServerError{}
}

/*
TeamDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamDeleteInternalServerError struct {
}

// IsSuccess returns true when this team delete internal server error response has a 2xx status code
func (o *TeamDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team delete internal server error response has a 3xx status code
func (o *TeamDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team delete internal server error response has a 4xx status code
func (o *TeamDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team delete internal server error response has a 5xx status code
func (o *TeamDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team delete internal server error response a status code equal to that given
func (o *TeamDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TeamDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteInternalServerError ", 500)
}

func (o *TeamDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] teamDeleteInternalServerError ", 500)
}

func (o *TeamDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
