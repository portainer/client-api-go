// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// CurrentUserInspectReader is a Reader for the CurrentUserInspect structure.
type CurrentUserInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CurrentUserInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCurrentUserInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCurrentUserInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCurrentUserInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCurrentUserInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCurrentUserInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/me] CurrentUserInspect", response, response.Code())
	}
}

// NewCurrentUserInspectOK creates a CurrentUserInspectOK with default headers values
func NewCurrentUserInspectOK() *CurrentUserInspectOK {
	return &CurrentUserInspectOK{}
}

/*
CurrentUserInspectOK describes a response with status code 200, with default header values.

Success
*/
type CurrentUserInspectOK struct {
	Payload *models.PortainereeUser
}

// IsSuccess returns true when this current user inspect o k response has a 2xx status code
func (o *CurrentUserInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this current user inspect o k response has a 3xx status code
func (o *CurrentUserInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user inspect o k response has a 4xx status code
func (o *CurrentUserInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this current user inspect o k response has a 5xx status code
func (o *CurrentUserInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this current user inspect o k response a status code equal to that given
func (o *CurrentUserInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the current user inspect o k response
func (o *CurrentUserInspectOK) Code() int {
	return 200
}

func (o *CurrentUserInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectOK %s", 200, payload)
}

func (o *CurrentUserInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectOK %s", 200, payload)
}

func (o *CurrentUserInspectOK) GetPayload() *models.PortainereeUser {
	return o.Payload
}

func (o *CurrentUserInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCurrentUserInspectBadRequest creates a CurrentUserInspectBadRequest with default headers values
func NewCurrentUserInspectBadRequest() *CurrentUserInspectBadRequest {
	return &CurrentUserInspectBadRequest{}
}

/*
CurrentUserInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CurrentUserInspectBadRequest struct {
}

// IsSuccess returns true when this current user inspect bad request response has a 2xx status code
func (o *CurrentUserInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user inspect bad request response has a 3xx status code
func (o *CurrentUserInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user inspect bad request response has a 4xx status code
func (o *CurrentUserInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user inspect bad request response has a 5xx status code
func (o *CurrentUserInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this current user inspect bad request response a status code equal to that given
func (o *CurrentUserInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the current user inspect bad request response
func (o *CurrentUserInspectBadRequest) Code() int {
	return 400
}

func (o *CurrentUserInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectBadRequest", 400)
}

func (o *CurrentUserInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectBadRequest", 400)
}

func (o *CurrentUserInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserInspectForbidden creates a CurrentUserInspectForbidden with default headers values
func NewCurrentUserInspectForbidden() *CurrentUserInspectForbidden {
	return &CurrentUserInspectForbidden{}
}

/*
CurrentUserInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type CurrentUserInspectForbidden struct {
}

// IsSuccess returns true when this current user inspect forbidden response has a 2xx status code
func (o *CurrentUserInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user inspect forbidden response has a 3xx status code
func (o *CurrentUserInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user inspect forbidden response has a 4xx status code
func (o *CurrentUserInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user inspect forbidden response has a 5xx status code
func (o *CurrentUserInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this current user inspect forbidden response a status code equal to that given
func (o *CurrentUserInspectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the current user inspect forbidden response
func (o *CurrentUserInspectForbidden) Code() int {
	return 403
}

func (o *CurrentUserInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectForbidden", 403)
}

func (o *CurrentUserInspectForbidden) String() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectForbidden", 403)
}

func (o *CurrentUserInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserInspectNotFound creates a CurrentUserInspectNotFound with default headers values
func NewCurrentUserInspectNotFound() *CurrentUserInspectNotFound {
	return &CurrentUserInspectNotFound{}
}

/*
CurrentUserInspectNotFound describes a response with status code 404, with default header values.

User not found
*/
type CurrentUserInspectNotFound struct {
}

// IsSuccess returns true when this current user inspect not found response has a 2xx status code
func (o *CurrentUserInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user inspect not found response has a 3xx status code
func (o *CurrentUserInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user inspect not found response has a 4xx status code
func (o *CurrentUserInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user inspect not found response has a 5xx status code
func (o *CurrentUserInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this current user inspect not found response a status code equal to that given
func (o *CurrentUserInspectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the current user inspect not found response
func (o *CurrentUserInspectNotFound) Code() int {
	return 404
}

func (o *CurrentUserInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectNotFound", 404)
}

func (o *CurrentUserInspectNotFound) String() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectNotFound", 404)
}

func (o *CurrentUserInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserInspectInternalServerError creates a CurrentUserInspectInternalServerError with default headers values
func NewCurrentUserInspectInternalServerError() *CurrentUserInspectInternalServerError {
	return &CurrentUserInspectInternalServerError{}
}

/*
CurrentUserInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CurrentUserInspectInternalServerError struct {
}

// IsSuccess returns true when this current user inspect internal server error response has a 2xx status code
func (o *CurrentUserInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user inspect internal server error response has a 3xx status code
func (o *CurrentUserInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user inspect internal server error response has a 4xx status code
func (o *CurrentUserInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this current user inspect internal server error response has a 5xx status code
func (o *CurrentUserInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this current user inspect internal server error response a status code equal to that given
func (o *CurrentUserInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the current user inspect internal server error response
func (o *CurrentUserInspectInternalServerError) Code() int {
	return 500
}

func (o *CurrentUserInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectInternalServerError", 500)
}

func (o *CurrentUserInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/me][%d] currentUserInspectInternalServerError", 500)
}

func (o *CurrentUserInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
