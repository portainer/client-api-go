// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// RegistryCreateReader is a Reader for the RegistryCreate structure.
type RegistryCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegistryCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegistryCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegistryCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegistryCreateOK creates a RegistryCreateOK with default headers values
func NewRegistryCreateOK() *RegistryCreateOK {
	return &RegistryCreateOK{}
}

/*
RegistryCreateOK describes a response with status code 200, with default header values.

Success
*/
type RegistryCreateOK struct {
	Payload *models_ce.PortainerRegistry
}

// IsSuccess returns true when this registry create o k response has a 2xx status code
func (o *RegistryCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this registry create o k response has a 3xx status code
func (o *RegistryCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry create o k response has a 4xx status code
func (o *RegistryCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry create o k response has a 5xx status code
func (o *RegistryCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this registry create o k response a status code equal to that given
func (o *RegistryCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *RegistryCreateOK) Error() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateOK  %+v", 200, o.Payload)
}

func (o *RegistryCreateOK) String() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateOK  %+v", 200, o.Payload)
}

func (o *RegistryCreateOK) GetPayload() *models_ce.PortainerRegistry {
	return o.Payload
}

func (o *RegistryCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryCreateBadRequest creates a RegistryCreateBadRequest with default headers values
func NewRegistryCreateBadRequest() *RegistryCreateBadRequest {
	return &RegistryCreateBadRequest{}
}

/*
RegistryCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type RegistryCreateBadRequest struct {
}

// IsSuccess returns true when this registry create bad request response has a 2xx status code
func (o *RegistryCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry create bad request response has a 3xx status code
func (o *RegistryCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry create bad request response has a 4xx status code
func (o *RegistryCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this registry create bad request response has a 5xx status code
func (o *RegistryCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this registry create bad request response a status code equal to that given
func (o *RegistryCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RegistryCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateBadRequest ", 400)
}

func (o *RegistryCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateBadRequest ", 400)
}

func (o *RegistryCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryCreateInternalServerError creates a RegistryCreateInternalServerError with default headers values
func NewRegistryCreateInternalServerError() *RegistryCreateInternalServerError {
	return &RegistryCreateInternalServerError{}
}

/*
RegistryCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RegistryCreateInternalServerError struct {
}

// IsSuccess returns true when this registry create internal server error response has a 2xx status code
func (o *RegistryCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry create internal server error response has a 3xx status code
func (o *RegistryCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry create internal server error response has a 4xx status code
func (o *RegistryCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry create internal server error response has a 5xx status code
func (o *RegistryCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this registry create internal server error response a status code equal to that given
func (o *RegistryCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RegistryCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateInternalServerError ", 500)
}

func (o *RegistryCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /registries][%d] registryCreateInternalServerError ", 500)
}

func (o *RegistryCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
