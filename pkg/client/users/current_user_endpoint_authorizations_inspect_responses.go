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

// CurrentUserEndpointAuthorizationsInspectReader is a Reader for the CurrentUserEndpointAuthorizationsInspect structure.
type CurrentUserEndpointAuthorizationsInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CurrentUserEndpointAuthorizationsInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCurrentUserEndpointAuthorizationsInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCurrentUserEndpointAuthorizationsInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCurrentUserEndpointAuthorizationsInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCurrentUserEndpointAuthorizationsInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCurrentUserEndpointAuthorizationsInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/me/auth/{endpointID}] CurrentUserEndpointAuthorizationsInspect", response, response.Code())
	}
}

// NewCurrentUserEndpointAuthorizationsInspectOK creates a CurrentUserEndpointAuthorizationsInspectOK with default headers values
func NewCurrentUserEndpointAuthorizationsInspectOK() *CurrentUserEndpointAuthorizationsInspectOK {
	return &CurrentUserEndpointAuthorizationsInspectOK{}
}

/*
CurrentUserEndpointAuthorizationsInspectOK describes a response with status code 200, with default header values.

Success
*/
type CurrentUserEndpointAuthorizationsInspectOK struct {
	Payload *models.UsersCurrentUserEndpointAuthResponse
}

// IsSuccess returns true when this current user endpoint authorizations inspect o k response has a 2xx status code
func (o *CurrentUserEndpointAuthorizationsInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this current user endpoint authorizations inspect o k response has a 3xx status code
func (o *CurrentUserEndpointAuthorizationsInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user endpoint authorizations inspect o k response has a 4xx status code
func (o *CurrentUserEndpointAuthorizationsInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this current user endpoint authorizations inspect o k response has a 5xx status code
func (o *CurrentUserEndpointAuthorizationsInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this current user endpoint authorizations inspect o k response a status code equal to that given
func (o *CurrentUserEndpointAuthorizationsInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the current user endpoint authorizations inspect o k response
func (o *CurrentUserEndpointAuthorizationsInspectOK) Code() int {
	return 200
}

func (o *CurrentUserEndpointAuthorizationsInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectOK %s", 200, payload)
}

func (o *CurrentUserEndpointAuthorizationsInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectOK %s", 200, payload)
}

func (o *CurrentUserEndpointAuthorizationsInspectOK) GetPayload() *models.UsersCurrentUserEndpointAuthResponse {
	return o.Payload
}

func (o *CurrentUserEndpointAuthorizationsInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersCurrentUserEndpointAuthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCurrentUserEndpointAuthorizationsInspectBadRequest creates a CurrentUserEndpointAuthorizationsInspectBadRequest with default headers values
func NewCurrentUserEndpointAuthorizationsInspectBadRequest() *CurrentUserEndpointAuthorizationsInspectBadRequest {
	return &CurrentUserEndpointAuthorizationsInspectBadRequest{}
}

/*
CurrentUserEndpointAuthorizationsInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CurrentUserEndpointAuthorizationsInspectBadRequest struct {
}

// IsSuccess returns true when this current user endpoint authorizations inspect bad request response has a 2xx status code
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user endpoint authorizations inspect bad request response has a 3xx status code
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user endpoint authorizations inspect bad request response has a 4xx status code
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user endpoint authorizations inspect bad request response has a 5xx status code
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this current user endpoint authorizations inspect bad request response a status code equal to that given
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the current user endpoint authorizations inspect bad request response
func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) Code() int {
	return 400
}

func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectBadRequest", 400)
}

func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectBadRequest", 400)
}

func (o *CurrentUserEndpointAuthorizationsInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserEndpointAuthorizationsInspectForbidden creates a CurrentUserEndpointAuthorizationsInspectForbidden with default headers values
func NewCurrentUserEndpointAuthorizationsInspectForbidden() *CurrentUserEndpointAuthorizationsInspectForbidden {
	return &CurrentUserEndpointAuthorizationsInspectForbidden{}
}

/*
CurrentUserEndpointAuthorizationsInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type CurrentUserEndpointAuthorizationsInspectForbidden struct {
}

// IsSuccess returns true when this current user endpoint authorizations inspect forbidden response has a 2xx status code
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user endpoint authorizations inspect forbidden response has a 3xx status code
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user endpoint authorizations inspect forbidden response has a 4xx status code
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user endpoint authorizations inspect forbidden response has a 5xx status code
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this current user endpoint authorizations inspect forbidden response a status code equal to that given
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the current user endpoint authorizations inspect forbidden response
func (o *CurrentUserEndpointAuthorizationsInspectForbidden) Code() int {
	return 403
}

func (o *CurrentUserEndpointAuthorizationsInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectForbidden", 403)
}

func (o *CurrentUserEndpointAuthorizationsInspectForbidden) String() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectForbidden", 403)
}

func (o *CurrentUserEndpointAuthorizationsInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserEndpointAuthorizationsInspectNotFound creates a CurrentUserEndpointAuthorizationsInspectNotFound with default headers values
func NewCurrentUserEndpointAuthorizationsInspectNotFound() *CurrentUserEndpointAuthorizationsInspectNotFound {
	return &CurrentUserEndpointAuthorizationsInspectNotFound{}
}

/*
CurrentUserEndpointAuthorizationsInspectNotFound describes a response with status code 404, with default header values.

User not found
*/
type CurrentUserEndpointAuthorizationsInspectNotFound struct {
}

// IsSuccess returns true when this current user endpoint authorizations inspect not found response has a 2xx status code
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user endpoint authorizations inspect not found response has a 3xx status code
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user endpoint authorizations inspect not found response has a 4xx status code
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this current user endpoint authorizations inspect not found response has a 5xx status code
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this current user endpoint authorizations inspect not found response a status code equal to that given
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the current user endpoint authorizations inspect not found response
func (o *CurrentUserEndpointAuthorizationsInspectNotFound) Code() int {
	return 404
}

func (o *CurrentUserEndpointAuthorizationsInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectNotFound", 404)
}

func (o *CurrentUserEndpointAuthorizationsInspectNotFound) String() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectNotFound", 404)
}

func (o *CurrentUserEndpointAuthorizationsInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrentUserEndpointAuthorizationsInspectInternalServerError creates a CurrentUserEndpointAuthorizationsInspectInternalServerError with default headers values
func NewCurrentUserEndpointAuthorizationsInspectInternalServerError() *CurrentUserEndpointAuthorizationsInspectInternalServerError {
	return &CurrentUserEndpointAuthorizationsInspectInternalServerError{}
}

/*
CurrentUserEndpointAuthorizationsInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CurrentUserEndpointAuthorizationsInspectInternalServerError struct {
}

// IsSuccess returns true when this current user endpoint authorizations inspect internal server error response has a 2xx status code
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this current user endpoint authorizations inspect internal server error response has a 3xx status code
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this current user endpoint authorizations inspect internal server error response has a 4xx status code
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this current user endpoint authorizations inspect internal server error response has a 5xx status code
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this current user endpoint authorizations inspect internal server error response a status code equal to that given
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the current user endpoint authorizations inspect internal server error response
func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) Code() int {
	return 500
}

func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectInternalServerError", 500)
}

func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/me/auth/{endpointID}][%d] currentUserEndpointAuthorizationsInspectInternalServerError", 500)
}

func (o *CurrentUserEndpointAuthorizationsInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
