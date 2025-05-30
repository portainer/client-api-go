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

// UserUpdateReader is a Reader for the UserUpdate structure.
type UserUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUserUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /users/{id}] UserUpdate", response, response.Code())
	}
}

// NewUserUpdateOK creates a UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {
	return &UserUpdateOK{}
}

/*
UserUpdateOK describes a response with status code 200, with default header values.

Success
*/
type UserUpdateOK struct {
	Payload *models.PortainereeUser
}

// IsSuccess returns true when this user update o k response has a 2xx status code
func (o *UserUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user update o k response has a 3xx status code
func (o *UserUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update o k response has a 4xx status code
func (o *UserUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user update o k response has a 5xx status code
func (o *UserUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user update o k response a status code equal to that given
func (o *UserUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user update o k response
func (o *UserUpdateOK) Code() int {
	return 200
}

func (o *UserUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateOK %s", 200, payload)
}

func (o *UserUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateOK %s", 200, payload)
}

func (o *UserUpdateOK) GetPayload() *models.PortainereeUser {
	return o.Payload
}

func (o *UserUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateBadRequest creates a UserUpdateBadRequest with default headers values
func NewUserUpdateBadRequest() *UserUpdateBadRequest {
	return &UserUpdateBadRequest{}
}

/*
UserUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserUpdateBadRequest struct {
}

// IsSuccess returns true when this user update bad request response has a 2xx status code
func (o *UserUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update bad request response has a 3xx status code
func (o *UserUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update bad request response has a 4xx status code
func (o *UserUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update bad request response has a 5xx status code
func (o *UserUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user update bad request response a status code equal to that given
func (o *UserUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user update bad request response
func (o *UserUpdateBadRequest) Code() int {
	return 400
}

func (o *UserUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateBadRequest", 400)
}

func (o *UserUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateBadRequest", 400)
}

func (o *UserUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateForbidden creates a UserUpdateForbidden with default headers values
func NewUserUpdateForbidden() *UserUpdateForbidden {
	return &UserUpdateForbidden{}
}

/*
UserUpdateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserUpdateForbidden struct {
}

// IsSuccess returns true when this user update forbidden response has a 2xx status code
func (o *UserUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update forbidden response has a 3xx status code
func (o *UserUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update forbidden response has a 4xx status code
func (o *UserUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update forbidden response has a 5xx status code
func (o *UserUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user update forbidden response a status code equal to that given
func (o *UserUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user update forbidden response
func (o *UserUpdateForbidden) Code() int {
	return 403
}

func (o *UserUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateForbidden", 403)
}

func (o *UserUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateForbidden", 403)
}

func (o *UserUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateNotFound creates a UserUpdateNotFound with default headers values
func NewUserUpdateNotFound() *UserUpdateNotFound {
	return &UserUpdateNotFound{}
}

/*
UserUpdateNotFound describes a response with status code 404, with default header values.

User not found
*/
type UserUpdateNotFound struct {
}

// IsSuccess returns true when this user update not found response has a 2xx status code
func (o *UserUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update not found response has a 3xx status code
func (o *UserUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update not found response has a 4xx status code
func (o *UserUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update not found response has a 5xx status code
func (o *UserUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user update not found response a status code equal to that given
func (o *UserUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user update not found response
func (o *UserUpdateNotFound) Code() int {
	return 404
}

func (o *UserUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateNotFound", 404)
}

func (o *UserUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateNotFound", 404)
}

func (o *UserUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateConflict creates a UserUpdateConflict with default headers values
func NewUserUpdateConflict() *UserUpdateConflict {
	return &UserUpdateConflict{}
}

/*
UserUpdateConflict describes a response with status code 409, with default header values.

Username already exist or Edge Compute features are not enabled
*/
type UserUpdateConflict struct {
}

// IsSuccess returns true when this user update conflict response has a 2xx status code
func (o *UserUpdateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update conflict response has a 3xx status code
func (o *UserUpdateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update conflict response has a 4xx status code
func (o *UserUpdateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update conflict response has a 5xx status code
func (o *UserUpdateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this user update conflict response a status code equal to that given
func (o *UserUpdateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the user update conflict response
func (o *UserUpdateConflict) Code() int {
	return 409
}

func (o *UserUpdateConflict) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateConflict", 409)
}

func (o *UserUpdateConflict) String() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateConflict", 409)
}

func (o *UserUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateInternalServerError creates a UserUpdateInternalServerError with default headers values
func NewUserUpdateInternalServerError() *UserUpdateInternalServerError {
	return &UserUpdateInternalServerError{}
}

/*
UserUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserUpdateInternalServerError struct {
}

// IsSuccess returns true when this user update internal server error response has a 2xx status code
func (o *UserUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update internal server error response has a 3xx status code
func (o *UserUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update internal server error response has a 4xx status code
func (o *UserUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user update internal server error response has a 5xx status code
func (o *UserUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user update internal server error response a status code equal to that given
func (o *UserUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user update internal server error response
func (o *UserUpdateInternalServerError) Code() int {
	return 500
}

func (o *UserUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateInternalServerError", 500)
}

func (o *UserUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /users/{id}][%d] userUpdateInternalServerError", 500)
}

func (o *UserUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
