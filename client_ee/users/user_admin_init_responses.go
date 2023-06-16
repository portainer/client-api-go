// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ee"
)

// UserAdminInitReader is a Reader for the UserAdminInit structure.
type UserAdminInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAdminInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAdminInitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserAdminInitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUserAdminInitConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserAdminInitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserAdminInitOK creates a UserAdminInitOK with default headers values
func NewUserAdminInitOK() *UserAdminInitOK {
	return &UserAdminInitOK{}
}

/*
UserAdminInitOK describes a response with status code 200, with default header values.

Success
*/
type UserAdminInitOK struct {
	Payload *models_ee.PortainereeUser
}

// IsSuccess returns true when this user admin init o k response has a 2xx status code
func (o *UserAdminInitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user admin init o k response has a 3xx status code
func (o *UserAdminInitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user admin init o k response has a 4xx status code
func (o *UserAdminInitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user admin init o k response has a 5xx status code
func (o *UserAdminInitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user admin init o k response a status code equal to that given
func (o *UserAdminInitOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserAdminInitOK) Error() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitOK  %+v", 200, o.Payload)
}

func (o *UserAdminInitOK) String() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitOK  %+v", 200, o.Payload)
}

func (o *UserAdminInitOK) GetPayload() *models_ee.PortainereeUser {
	return o.Payload
}

func (o *UserAdminInitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ee.PortainereeUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAdminInitBadRequest creates a UserAdminInitBadRequest with default headers values
func NewUserAdminInitBadRequest() *UserAdminInitBadRequest {
	return &UserAdminInitBadRequest{}
}

/*
UserAdminInitBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserAdminInitBadRequest struct {
}

// IsSuccess returns true when this user admin init bad request response has a 2xx status code
func (o *UserAdminInitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user admin init bad request response has a 3xx status code
func (o *UserAdminInitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user admin init bad request response has a 4xx status code
func (o *UserAdminInitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user admin init bad request response has a 5xx status code
func (o *UserAdminInitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user admin init bad request response a status code equal to that given
func (o *UserAdminInitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserAdminInitBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitBadRequest ", 400)
}

func (o *UserAdminInitBadRequest) String() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitBadRequest ", 400)
}

func (o *UserAdminInitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserAdminInitConflict creates a UserAdminInitConflict with default headers values
func NewUserAdminInitConflict() *UserAdminInitConflict {
	return &UserAdminInitConflict{}
}

/*
UserAdminInitConflict describes a response with status code 409, with default header values.

Admin user already initialized
*/
type UserAdminInitConflict struct {
}

// IsSuccess returns true when this user admin init conflict response has a 2xx status code
func (o *UserAdminInitConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user admin init conflict response has a 3xx status code
func (o *UserAdminInitConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user admin init conflict response has a 4xx status code
func (o *UserAdminInitConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this user admin init conflict response has a 5xx status code
func (o *UserAdminInitConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this user admin init conflict response a status code equal to that given
func (o *UserAdminInitConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UserAdminInitConflict) Error() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitConflict ", 409)
}

func (o *UserAdminInitConflict) String() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitConflict ", 409)
}

func (o *UserAdminInitConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserAdminInitInternalServerError creates a UserAdminInitInternalServerError with default headers values
func NewUserAdminInitInternalServerError() *UserAdminInitInternalServerError {
	return &UserAdminInitInternalServerError{}
}

/*
UserAdminInitInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserAdminInitInternalServerError struct {
}

// IsSuccess returns true when this user admin init internal server error response has a 2xx status code
func (o *UserAdminInitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user admin init internal server error response has a 3xx status code
func (o *UserAdminInitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user admin init internal server error response has a 4xx status code
func (o *UserAdminInitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user admin init internal server error response has a 5xx status code
func (o *UserAdminInitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user admin init internal server error response a status code equal to that given
func (o *UserAdminInitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserAdminInitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitInternalServerError ", 500)
}

func (o *UserAdminInitInternalServerError) String() string {
	return fmt.Sprintf("[POST /users/admin/init][%d] userAdminInitInternalServerError ", 500)
}

func (o *UserAdminInitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
