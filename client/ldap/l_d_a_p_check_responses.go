// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LDAPCheckReader is a Reader for the LDAPCheck structure.
type LDAPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLDAPCheckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLDAPCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLDAPCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ldap/check] LDAPCheck", response, response.Code())
	}
}

// NewLDAPCheckNoContent creates a LDAPCheckNoContent with default headers values
func NewLDAPCheckNoContent() *LDAPCheckNoContent {
	return &LDAPCheckNoContent{}
}

/*
LDAPCheckNoContent describes a response with status code 204, with default header values.

Success
*/
type LDAPCheckNoContent struct {
}

// IsSuccess returns true when this l d a p check no content response has a 2xx status code
func (o *LDAPCheckNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this l d a p check no content response has a 3xx status code
func (o *LDAPCheckNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p check no content response has a 4xx status code
func (o *LDAPCheckNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this l d a p check no content response has a 5xx status code
func (o *LDAPCheckNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this l d a p check no content response a status code equal to that given
func (o *LDAPCheckNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the l d a p check no content response
func (o *LDAPCheckNoContent) Code() int {
	return 204
}

func (o *LDAPCheckNoContent) Error() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckNoContent", 204)
}

func (o *LDAPCheckNoContent) String() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckNoContent", 204)
}

func (o *LDAPCheckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLDAPCheckBadRequest creates a LDAPCheckBadRequest with default headers values
func NewLDAPCheckBadRequest() *LDAPCheckBadRequest {
	return &LDAPCheckBadRequest{}
}

/*
LDAPCheckBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type LDAPCheckBadRequest struct {
}

// IsSuccess returns true when this l d a p check bad request response has a 2xx status code
func (o *LDAPCheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this l d a p check bad request response has a 3xx status code
func (o *LDAPCheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p check bad request response has a 4xx status code
func (o *LDAPCheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this l d a p check bad request response has a 5xx status code
func (o *LDAPCheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this l d a p check bad request response a status code equal to that given
func (o *LDAPCheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the l d a p check bad request response
func (o *LDAPCheckBadRequest) Code() int {
	return 400
}

func (o *LDAPCheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckBadRequest", 400)
}

func (o *LDAPCheckBadRequest) String() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckBadRequest", 400)
}

func (o *LDAPCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLDAPCheckInternalServerError creates a LDAPCheckInternalServerError with default headers values
func NewLDAPCheckInternalServerError() *LDAPCheckInternalServerError {
	return &LDAPCheckInternalServerError{}
}

/*
LDAPCheckInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type LDAPCheckInternalServerError struct {
}

// IsSuccess returns true when this l d a p check internal server error response has a 2xx status code
func (o *LDAPCheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this l d a p check internal server error response has a 3xx status code
func (o *LDAPCheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this l d a p check internal server error response has a 4xx status code
func (o *LDAPCheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this l d a p check internal server error response has a 5xx status code
func (o *LDAPCheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this l d a p check internal server error response a status code equal to that given
func (o *LDAPCheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the l d a p check internal server error response
func (o *LDAPCheckInternalServerError) Code() int {
	return 500
}

func (o *LDAPCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckInternalServerError", 500)
}

func (o *LDAPCheckInternalServerError) String() string {
	return fmt.Sprintf("[POST /ldap/check][%d] lDAPCheckInternalServerError", 500)
}

func (o *LDAPCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
