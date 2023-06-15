// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// UserInspectReader is a Reader for the UserInspect structure.
type UserInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserInspectOK creates a UserInspectOK with default headers values
func NewUserInspectOK() *UserInspectOK {
	return &UserInspectOK{}
}

/*
UserInspectOK describes a response with status code 200, with default header values.

Success
*/
type UserInspectOK struct {
	Payload *models.PortainerUser
}

// IsSuccess returns true when this user inspect o k response has a 2xx status code
func (o *UserInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user inspect o k response has a 3xx status code
func (o *UserInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user inspect o k response has a 4xx status code
func (o *UserInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user inspect o k response has a 5xx status code
func (o *UserInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user inspect o k response a status code equal to that given
func (o *UserInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserInspectOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectOK  %+v", 200, o.Payload)
}

func (o *UserInspectOK) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectOK  %+v", 200, o.Payload)
}

func (o *UserInspectOK) GetPayload() *models.PortainerUser {
	return o.Payload
}

func (o *UserInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserInspectBadRequest creates a UserInspectBadRequest with default headers values
func NewUserInspectBadRequest() *UserInspectBadRequest {
	return &UserInspectBadRequest{}
}

/*
UserInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserInspectBadRequest struct {
}

// IsSuccess returns true when this user inspect bad request response has a 2xx status code
func (o *UserInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user inspect bad request response has a 3xx status code
func (o *UserInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user inspect bad request response has a 4xx status code
func (o *UserInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user inspect bad request response has a 5xx status code
func (o *UserInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user inspect bad request response a status code equal to that given
func (o *UserInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectBadRequest ", 400)
}

func (o *UserInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectBadRequest ", 400)
}

func (o *UserInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserInspectForbidden creates a UserInspectForbidden with default headers values
func NewUserInspectForbidden() *UserInspectForbidden {
	return &UserInspectForbidden{}
}

/*
UserInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserInspectForbidden struct {
}

// IsSuccess returns true when this user inspect forbidden response has a 2xx status code
func (o *UserInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user inspect forbidden response has a 3xx status code
func (o *UserInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user inspect forbidden response has a 4xx status code
func (o *UserInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user inspect forbidden response has a 5xx status code
func (o *UserInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user inspect forbidden response a status code equal to that given
func (o *UserInspectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectForbidden ", 403)
}

func (o *UserInspectForbidden) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectForbidden ", 403)
}

func (o *UserInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserInspectNotFound creates a UserInspectNotFound with default headers values
func NewUserInspectNotFound() *UserInspectNotFound {
	return &UserInspectNotFound{}
}

/*
UserInspectNotFound describes a response with status code 404, with default header values.

User not found
*/
type UserInspectNotFound struct {
}

// IsSuccess returns true when this user inspect not found response has a 2xx status code
func (o *UserInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user inspect not found response has a 3xx status code
func (o *UserInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user inspect not found response has a 4xx status code
func (o *UserInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user inspect not found response has a 5xx status code
func (o *UserInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user inspect not found response a status code equal to that given
func (o *UserInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectNotFound ", 404)
}

func (o *UserInspectNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectNotFound ", 404)
}

func (o *UserInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserInspectInternalServerError creates a UserInspectInternalServerError with default headers values
func NewUserInspectInternalServerError() *UserInspectInternalServerError {
	return &UserInspectInternalServerError{}
}

/*
UserInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserInspectInternalServerError struct {
}

// IsSuccess returns true when this user inspect internal server error response has a 2xx status code
func (o *UserInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user inspect internal server error response has a 3xx status code
func (o *UserInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user inspect internal server error response has a 4xx status code
func (o *UserInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user inspect internal server error response has a 5xx status code
func (o *UserInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user inspect internal server error response a status code equal to that given
func (o *UserInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectInternalServerError ", 500)
}

func (o *UserInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] userInspectInternalServerError ", 500)
}

func (o *UserInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
