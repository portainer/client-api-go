// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ContainerImageStatusReader is a Reader for the ContainerImageStatus structure.
type ContainerImageStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerImageStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerImageStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerImageStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerImageStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerImageStatusOK creates a ContainerImageStatusOK with default headers values
func NewContainerImageStatusOK() *ContainerImageStatusOK {
	return &ContainerImageStatusOK{}
}

/*
ContainerImageStatusOK describes a response with status code 200, with default header values.

Success
*/
type ContainerImageStatusOK struct {
}

// IsSuccess returns true when this container image status o k response has a 2xx status code
func (o *ContainerImageStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container image status o k response has a 3xx status code
func (o *ContainerImageStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container image status o k response has a 4xx status code
func (o *ContainerImageStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container image status o k response has a 5xx status code
func (o *ContainerImageStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container image status o k response a status code equal to that given
func (o *ContainerImageStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerImageStatusOK) Error() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusOK ", 200)
}

func (o *ContainerImageStatusOK) String() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusOK ", 200)
}

func (o *ContainerImageStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerImageStatusBadRequest creates a ContainerImageStatusBadRequest with default headers values
func NewContainerImageStatusBadRequest() *ContainerImageStatusBadRequest {
	return &ContainerImageStatusBadRequest{}
}

/*
ContainerImageStatusBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ContainerImageStatusBadRequest struct {
}

// IsSuccess returns true when this container image status bad request response has a 2xx status code
func (o *ContainerImageStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container image status bad request response has a 3xx status code
func (o *ContainerImageStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container image status bad request response has a 4xx status code
func (o *ContainerImageStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container image status bad request response has a 5xx status code
func (o *ContainerImageStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container image status bad request response a status code equal to that given
func (o *ContainerImageStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerImageStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusBadRequest ", 400)
}

func (o *ContainerImageStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusBadRequest ", 400)
}

func (o *ContainerImageStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerImageStatusInternalServerError creates a ContainerImageStatusInternalServerError with default headers values
func NewContainerImageStatusInternalServerError() *ContainerImageStatusInternalServerError {
	return &ContainerImageStatusInternalServerError{}
}

/*
ContainerImageStatusInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerImageStatusInternalServerError struct {
}

// IsSuccess returns true when this container image status internal server error response has a 2xx status code
func (o *ContainerImageStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container image status internal server error response has a 3xx status code
func (o *ContainerImageStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container image status internal server error response has a 4xx status code
func (o *ContainerImageStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container image status internal server error response has a 5xx status code
func (o *ContainerImageStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container image status internal server error response a status code equal to that given
func (o *ContainerImageStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerImageStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusInternalServerError ", 500)
}

func (o *ContainerImageStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/containers/{containerId}/image_status][%d] containerImageStatusInternalServerError ", 500)
}

func (o *ContainerImageStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
