// Code generated by go-swagger; DO NOT EDIT.

package support

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

// GetDebugLogStatusReader is a Reader for the GetDebugLogStatus structure.
type GetDebugLogStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugLogStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDebugLogStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetDebugLogStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /support/debug_log] GetDebugLogStatus", response, response.Code())
	}
}

// NewGetDebugLogStatusOK creates a GetDebugLogStatusOK with default headers values
func NewGetDebugLogStatusOK() *GetDebugLogStatusOK {
	return &GetDebugLogStatusOK{}
}

/*
GetDebugLogStatusOK describes a response with status code 200, with default header values.

Success
*/
type GetDebugLogStatusOK struct {
	Payload *models.SupportDebugLogPayload
}

// IsSuccess returns true when this get debug log status o k response has a 2xx status code
func (o *GetDebugLogStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get debug log status o k response has a 3xx status code
func (o *GetDebugLogStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get debug log status o k response has a 4xx status code
func (o *GetDebugLogStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get debug log status o k response has a 5xx status code
func (o *GetDebugLogStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get debug log status o k response a status code equal to that given
func (o *GetDebugLogStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get debug log status o k response
func (o *GetDebugLogStatusOK) Code() int {
	return 200
}

func (o *GetDebugLogStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/debug_log][%d] getDebugLogStatusOK %s", 200, payload)
}

func (o *GetDebugLogStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/debug_log][%d] getDebugLogStatusOK %s", 200, payload)
}

func (o *GetDebugLogStatusOK) GetPayload() *models.SupportDebugLogPayload {
	return o.Payload
}

func (o *GetDebugLogStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportDebugLogPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugLogStatusInternalServerError creates a GetDebugLogStatusInternalServerError with default headers values
func NewGetDebugLogStatusInternalServerError() *GetDebugLogStatusInternalServerError {
	return &GetDebugLogStatusInternalServerError{}
}

/*
GetDebugLogStatusInternalServerError describes a response with status code 500, with default header values.

Server error occurred.
*/
type GetDebugLogStatusInternalServerError struct {
}

// IsSuccess returns true when this get debug log status internal server error response has a 2xx status code
func (o *GetDebugLogStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get debug log status internal server error response has a 3xx status code
func (o *GetDebugLogStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get debug log status internal server error response has a 4xx status code
func (o *GetDebugLogStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get debug log status internal server error response has a 5xx status code
func (o *GetDebugLogStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get debug log status internal server error response a status code equal to that given
func (o *GetDebugLogStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get debug log status internal server error response
func (o *GetDebugLogStatusInternalServerError) Code() int {
	return 500
}

func (o *GetDebugLogStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /support/debug_log][%d] getDebugLogStatusInternalServerError", 500)
}

func (o *GetDebugLogStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /support/debug_log][%d] getDebugLogStatusInternalServerError", 500)
}

func (o *GetDebugLogStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
