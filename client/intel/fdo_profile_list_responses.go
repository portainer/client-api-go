// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FdoProfileListReader is a Reader for the FdoProfileList structure.
type FdoProfileListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FdoProfileListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFdoProfileListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFdoProfileListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFdoProfileListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFdoProfileListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFdoProfileListOK creates a FdoProfileListOK with default headers values
func NewFdoProfileListOK() *FdoProfileListOK {
	return &FdoProfileListOK{}
}

/*
FdoProfileListOK describes a response with status code 200, with default header values.

Success
*/
type FdoProfileListOK struct {
}

// IsSuccess returns true when this fdo profile list o k response has a 2xx status code
func (o *FdoProfileListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fdo profile list o k response has a 3xx status code
func (o *FdoProfileListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo profile list o k response has a 4xx status code
func (o *FdoProfileListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fdo profile list o k response has a 5xx status code
func (o *FdoProfileListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo profile list o k response a status code equal to that given
func (o *FdoProfileListOK) IsCode(code int) bool {
	return code == 200
}

func (o *FdoProfileListOK) Error() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListOK ", 200)
}

func (o *FdoProfileListOK) String() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListOK ", 200)
}

func (o *FdoProfileListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoProfileListBadRequest creates a FdoProfileListBadRequest with default headers values
func NewFdoProfileListBadRequest() *FdoProfileListBadRequest {
	return &FdoProfileListBadRequest{}
}

/*
FdoProfileListBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type FdoProfileListBadRequest struct {
}

// IsSuccess returns true when this fdo profile list bad request response has a 2xx status code
func (o *FdoProfileListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo profile list bad request response has a 3xx status code
func (o *FdoProfileListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo profile list bad request response has a 4xx status code
func (o *FdoProfileListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this fdo profile list bad request response has a 5xx status code
func (o *FdoProfileListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo profile list bad request response a status code equal to that given
func (o *FdoProfileListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FdoProfileListBadRequest) Error() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListBadRequest ", 400)
}

func (o *FdoProfileListBadRequest) String() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListBadRequest ", 400)
}

func (o *FdoProfileListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoProfileListForbidden creates a FdoProfileListForbidden with default headers values
func NewFdoProfileListForbidden() *FdoProfileListForbidden {
	return &FdoProfileListForbidden{}
}

/*
FdoProfileListForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type FdoProfileListForbidden struct {
}

// IsSuccess returns true when this fdo profile list forbidden response has a 2xx status code
func (o *FdoProfileListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo profile list forbidden response has a 3xx status code
func (o *FdoProfileListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo profile list forbidden response has a 4xx status code
func (o *FdoProfileListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this fdo profile list forbidden response has a 5xx status code
func (o *FdoProfileListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo profile list forbidden response a status code equal to that given
func (o *FdoProfileListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FdoProfileListForbidden) Error() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListForbidden ", 403)
}

func (o *FdoProfileListForbidden) String() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListForbidden ", 403)
}

func (o *FdoProfileListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoProfileListInternalServerError creates a FdoProfileListInternalServerError with default headers values
func NewFdoProfileListInternalServerError() *FdoProfileListInternalServerError {
	return &FdoProfileListInternalServerError{}
}

/*
FdoProfileListInternalServerError describes a response with status code 500, with default header values.

Bad gateway
*/
type FdoProfileListInternalServerError struct {
}

// IsSuccess returns true when this fdo profile list internal server error response has a 2xx status code
func (o *FdoProfileListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo profile list internal server error response has a 3xx status code
func (o *FdoProfileListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo profile list internal server error response has a 4xx status code
func (o *FdoProfileListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this fdo profile list internal server error response has a 5xx status code
func (o *FdoProfileListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this fdo profile list internal server error response a status code equal to that given
func (o *FdoProfileListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FdoProfileListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListInternalServerError ", 500)
}

func (o *FdoProfileListInternalServerError) String() string {
	return fmt.Sprintf("[GET /fdo/profiles][%d] fdoProfileListInternalServerError ", 500)
}

func (o *FdoProfileListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
