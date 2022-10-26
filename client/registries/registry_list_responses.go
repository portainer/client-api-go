// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// RegistryListReader is a Reader for the RegistryList structure.
type RegistryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegistryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRegistryListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegistryListOK creates a RegistryListOK with default headers values
func NewRegistryListOK() *RegistryListOK {
	return &RegistryListOK{}
}

/*
RegistryListOK describes a response with status code 200, with default header values.

Success
*/
type RegistryListOK struct {
	Payload []*models.PortainereeRegistry
}

// IsSuccess returns true when this registry list o k response has a 2xx status code
func (o *RegistryListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this registry list o k response has a 3xx status code
func (o *RegistryListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry list o k response has a 4xx status code
func (o *RegistryListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry list o k response has a 5xx status code
func (o *RegistryListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this registry list o k response a status code equal to that given
func (o *RegistryListOK) IsCode(code int) bool {
	return code == 200
}

func (o *RegistryListOK) Error() string {
	return fmt.Sprintf("[GET /registries][%d] registryListOK  %+v", 200, o.Payload)
}

func (o *RegistryListOK) String() string {
	return fmt.Sprintf("[GET /registries][%d] registryListOK  %+v", 200, o.Payload)
}

func (o *RegistryListOK) GetPayload() []*models.PortainereeRegistry {
	return o.Payload
}

func (o *RegistryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryListInternalServerError creates a RegistryListInternalServerError with default headers values
func NewRegistryListInternalServerError() *RegistryListInternalServerError {
	return &RegistryListInternalServerError{}
}

/*
RegistryListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RegistryListInternalServerError struct {
}

// IsSuccess returns true when this registry list internal server error response has a 2xx status code
func (o *RegistryListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry list internal server error response has a 3xx status code
func (o *RegistryListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry list internal server error response has a 4xx status code
func (o *RegistryListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry list internal server error response has a 5xx status code
func (o *RegistryListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this registry list internal server error response a status code equal to that given
func (o *RegistryListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RegistryListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries][%d] registryListInternalServerError ", 500)
}

func (o *RegistryListInternalServerError) String() string {
	return fmt.Sprintf("[GET /registries][%d] registryListInternalServerError ", 500)
}

func (o *RegistryListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
