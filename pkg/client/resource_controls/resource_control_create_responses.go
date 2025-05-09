// Code generated by go-swagger; DO NOT EDIT.

package resource_controls

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

// ResourceControlCreateReader is a Reader for the ResourceControlCreate structure.
type ResourceControlCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceControlCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceControlCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResourceControlCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewResourceControlCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceControlCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /resource_controls] ResourceControlCreate", response, response.Code())
	}
}

// NewResourceControlCreateOK creates a ResourceControlCreateOK with default headers values
func NewResourceControlCreateOK() *ResourceControlCreateOK {
	return &ResourceControlCreateOK{}
}

/*
ResourceControlCreateOK describes a response with status code 200, with default header values.

Success
*/
type ResourceControlCreateOK struct {
	Payload *models.PortainerResourceControl
}

// IsSuccess returns true when this resource control create o k response has a 2xx status code
func (o *ResourceControlCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource control create o k response has a 3xx status code
func (o *ResourceControlCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control create o k response has a 4xx status code
func (o *ResourceControlCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource control create o k response has a 5xx status code
func (o *ResourceControlCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control create o k response a status code equal to that given
func (o *ResourceControlCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource control create o k response
func (o *ResourceControlCreateOK) Code() int {
	return 200
}

func (o *ResourceControlCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateOK %s", 200, payload)
}

func (o *ResourceControlCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateOK %s", 200, payload)
}

func (o *ResourceControlCreateOK) GetPayload() *models.PortainerResourceControl {
	return o.Payload
}

func (o *ResourceControlCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerResourceControl)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceControlCreateBadRequest creates a ResourceControlCreateBadRequest with default headers values
func NewResourceControlCreateBadRequest() *ResourceControlCreateBadRequest {
	return &ResourceControlCreateBadRequest{}
}

/*
ResourceControlCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type ResourceControlCreateBadRequest struct {
}

// IsSuccess returns true when this resource control create bad request response has a 2xx status code
func (o *ResourceControlCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control create bad request response has a 3xx status code
func (o *ResourceControlCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control create bad request response has a 4xx status code
func (o *ResourceControlCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource control create bad request response has a 5xx status code
func (o *ResourceControlCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control create bad request response a status code equal to that given
func (o *ResourceControlCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the resource control create bad request response
func (o *ResourceControlCreateBadRequest) Code() int {
	return 400
}

func (o *ResourceControlCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateBadRequest", 400)
}

func (o *ResourceControlCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateBadRequest", 400)
}

func (o *ResourceControlCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourceControlCreateConflict creates a ResourceControlCreateConflict with default headers values
func NewResourceControlCreateConflict() *ResourceControlCreateConflict {
	return &ResourceControlCreateConflict{}
}

/*
ResourceControlCreateConflict describes a response with status code 409, with default header values.

A resource control is already associated to this resource
*/
type ResourceControlCreateConflict struct {
}

// IsSuccess returns true when this resource control create conflict response has a 2xx status code
func (o *ResourceControlCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control create conflict response has a 3xx status code
func (o *ResourceControlCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control create conflict response has a 4xx status code
func (o *ResourceControlCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource control create conflict response has a 5xx status code
func (o *ResourceControlCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control create conflict response a status code equal to that given
func (o *ResourceControlCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the resource control create conflict response
func (o *ResourceControlCreateConflict) Code() int {
	return 409
}

func (o *ResourceControlCreateConflict) Error() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateConflict", 409)
}

func (o *ResourceControlCreateConflict) String() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateConflict", 409)
}

func (o *ResourceControlCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourceControlCreateInternalServerError creates a ResourceControlCreateInternalServerError with default headers values
func NewResourceControlCreateInternalServerError() *ResourceControlCreateInternalServerError {
	return &ResourceControlCreateInternalServerError{}
}

/*
ResourceControlCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ResourceControlCreateInternalServerError struct {
}

// IsSuccess returns true when this resource control create internal server error response has a 2xx status code
func (o *ResourceControlCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control create internal server error response has a 3xx status code
func (o *ResourceControlCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control create internal server error response has a 4xx status code
func (o *ResourceControlCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource control create internal server error response has a 5xx status code
func (o *ResourceControlCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource control create internal server error response a status code equal to that given
func (o *ResourceControlCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource control create internal server error response
func (o *ResourceControlCreateInternalServerError) Code() int {
	return 500
}

func (o *ResourceControlCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateInternalServerError", 500)
}

func (o *ResourceControlCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /resource_controls][%d] resourceControlCreateInternalServerError", 500)
}

func (o *ResourceControlCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
