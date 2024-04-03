// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// LDAPUsersReader is a Reader for the LDAPUsers structure.
type LDAPUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLDAPUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLDAPUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ldap/users] LDAPUsers", response, response.Code())
	}
}

// NewLDAPUsersOK creates a LDAPUsersOK with default headers values
func NewLDAPUsersOK() *LDAPUsersOK {
	return &LDAPUsersOK{}
}

/*
LDAPUsersOK describes a response with status code 200, with default header values.

Success
*/
type LDAPUsersOK struct {
	Payload []*models.PortainereeLDAPUser
}

// IsSuccess returns true when this l d a p users o k response has a 2xx status code
func (o *LDAPUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this l d a p users o k response has a 3xx status code
func (o *LDAPUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p users o k response has a 4xx status code
func (o *LDAPUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this l d a p users o k response has a 5xx status code
func (o *LDAPUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this l d a p users o k response a status code equal to that given
func (o *LDAPUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the l d a p users o k response
func (o *LDAPUsersOK) Code() int {
	return 200
}

func (o *LDAPUsersOK) Error() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersOK  %+v", 200, o.Payload)
}

func (o *LDAPUsersOK) String() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersOK  %+v", 200, o.Payload)
}

func (o *LDAPUsersOK) GetPayload() []*models.PortainereeLDAPUser {
	return o.Payload
}

func (o *LDAPUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLDAPUsersBadRequest creates a LDAPUsersBadRequest with default headers values
func NewLDAPUsersBadRequest() *LDAPUsersBadRequest {
	return &LDAPUsersBadRequest{}
}

/*
LDAPUsersBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type LDAPUsersBadRequest struct {
}

// IsSuccess returns true when this l d a p users bad request response has a 2xx status code
func (o *LDAPUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this l d a p users bad request response has a 3xx status code
func (o *LDAPUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p users bad request response has a 4xx status code
func (o *LDAPUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this l d a p users bad request response has a 5xx status code
func (o *LDAPUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this l d a p users bad request response a status code equal to that given
func (o *LDAPUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the l d a p users bad request response
func (o *LDAPUsersBadRequest) Code() int {
	return 400
}

func (o *LDAPUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersBadRequest ", 400)
}

func (o *LDAPUsersBadRequest) String() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersBadRequest ", 400)
}

func (o *LDAPUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLDAPUsersInternalServerError creates a LDAPUsersInternalServerError with default headers values
func NewLDAPUsersInternalServerError() *LDAPUsersInternalServerError {
	return &LDAPUsersInternalServerError{}
}

/*
LDAPUsersInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type LDAPUsersInternalServerError struct {
}

// IsSuccess returns true when this l d a p users internal server error response has a 2xx status code
func (o *LDAPUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this l d a p users internal server error response has a 3xx status code
func (o *LDAPUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p users internal server error response has a 4xx status code
func (o *LDAPUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this l d a p users internal server error response has a 5xx status code
func (o *LDAPUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this l d a p users internal server error response a status code equal to that given
func (o *LDAPUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the l d a p users internal server error response
func (o *LDAPUsersInternalServerError) Code() int {
	return 500
}

func (o *LDAPUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersInternalServerError ", 500)
}

func (o *LDAPUsersInternalServerError) String() string {
	return fmt.Sprintf("[POST /ldap/users][%d] lDAPUsersInternalServerError ", 500)
}

func (o *LDAPUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
