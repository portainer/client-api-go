// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ee"
)

// StackListReader is a Reader for the StackList structure.
type StackListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewStackListNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackListOK creates a StackListOK with default headers values
func NewStackListOK() *StackListOK {
	return &StackListOK{}
}

/*
StackListOK describes a response with status code 200, with default header values.

Success
*/
type StackListOK struct {
	Payload []*models_ee.PortainereeStack
}

// IsSuccess returns true when this stack list o k response has a 2xx status code
func (o *StackListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack list o k response has a 3xx status code
func (o *StackListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack list o k response has a 4xx status code
func (o *StackListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack list o k response has a 5xx status code
func (o *StackListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack list o k response a status code equal to that given
func (o *StackListOK) IsCode(code int) bool {
	return code == 200
}

func (o *StackListOK) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListOK  %+v", 200, o.Payload)
}

func (o *StackListOK) String() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListOK  %+v", 200, o.Payload)
}

func (o *StackListOK) GetPayload() []*models_ee.PortainereeStack {
	return o.Payload
}

func (o *StackListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackListNoContent creates a StackListNoContent with default headers values
func NewStackListNoContent() *StackListNoContent {
	return &StackListNoContent{}
}

/*
StackListNoContent describes a response with status code 204, with default header values.

Success
*/
type StackListNoContent struct {
}

// IsSuccess returns true when this stack list no content response has a 2xx status code
func (o *StackListNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack list no content response has a 3xx status code
func (o *StackListNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack list no content response has a 4xx status code
func (o *StackListNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack list no content response has a 5xx status code
func (o *StackListNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stack list no content response a status code equal to that given
func (o *StackListNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *StackListNoContent) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListNoContent ", 204)
}

func (o *StackListNoContent) String() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListNoContent ", 204)
}

func (o *StackListNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackListBadRequest creates a StackListBadRequest with default headers values
func NewStackListBadRequest() *StackListBadRequest {
	return &StackListBadRequest{}
}

/*
StackListBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackListBadRequest struct {
}

// IsSuccess returns true when this stack list bad request response has a 2xx status code
func (o *StackListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack list bad request response has a 3xx status code
func (o *StackListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack list bad request response has a 4xx status code
func (o *StackListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack list bad request response has a 5xx status code
func (o *StackListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack list bad request response a status code equal to that given
func (o *StackListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StackListBadRequest) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListBadRequest ", 400)
}

func (o *StackListBadRequest) String() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListBadRequest ", 400)
}

func (o *StackListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackListInternalServerError creates a StackListInternalServerError with default headers values
func NewStackListInternalServerError() *StackListInternalServerError {
	return &StackListInternalServerError{}
}

/*
StackListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackListInternalServerError struct {
}

// IsSuccess returns true when this stack list internal server error response has a 2xx status code
func (o *StackListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack list internal server error response has a 3xx status code
func (o *StackListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack list internal server error response has a 4xx status code
func (o *StackListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack list internal server error response has a 5xx status code
func (o *StackListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack list internal server error response a status code equal to that given
func (o *StackListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StackListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListInternalServerError ", 500)
}

func (o *StackListInternalServerError) String() string {
	return fmt.Sprintf("[GET /stacks][%d] stackListInternalServerError ", 500)
}

func (o *StackListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
