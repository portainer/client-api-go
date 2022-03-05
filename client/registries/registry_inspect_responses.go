// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// RegistryInspectReader is a Reader for the RegistryInspect structure.
type RegistryInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegistryInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegistryInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegistryInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegistryInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegistryInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegistryInspectOK creates a RegistryInspectOK with default headers values
func NewRegistryInspectOK() *RegistryInspectOK {
	return &RegistryInspectOK{}
}

/* RegistryInspectOK describes a response with status code 200, with default header values.

Success
*/
type RegistryInspectOK struct {
	Payload *models.PortainerRegistry
}

func (o *RegistryInspectOK) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] registryInspectOK  %+v", 200, o.Payload)
}
func (o *RegistryInspectOK) GetPayload() *models.PortainerRegistry {
	return o.Payload
}

func (o *RegistryInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryInspectBadRequest creates a RegistryInspectBadRequest with default headers values
func NewRegistryInspectBadRequest() *RegistryInspectBadRequest {
	return &RegistryInspectBadRequest{}
}

/* RegistryInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type RegistryInspectBadRequest struct {
}

func (o *RegistryInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] registryInspectBadRequest ", 400)
}

func (o *RegistryInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryInspectForbidden creates a RegistryInspectForbidden with default headers values
func NewRegistryInspectForbidden() *RegistryInspectForbidden {
	return &RegistryInspectForbidden{}
}

/* RegistryInspectForbidden describes a response with status code 403, with default header values.

Permission denied to access registry
*/
type RegistryInspectForbidden struct {
}

func (o *RegistryInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] registryInspectForbidden ", 403)
}

func (o *RegistryInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryInspectNotFound creates a RegistryInspectNotFound with default headers values
func NewRegistryInspectNotFound() *RegistryInspectNotFound {
	return &RegistryInspectNotFound{}
}

/* RegistryInspectNotFound describes a response with status code 404, with default header values.

Registry not found
*/
type RegistryInspectNotFound struct {
}

func (o *RegistryInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] registryInspectNotFound ", 404)
}

func (o *RegistryInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryInspectInternalServerError creates a RegistryInspectInternalServerError with default headers values
func NewRegistryInspectInternalServerError() *RegistryInspectInternalServerError {
	return &RegistryInspectInternalServerError{}
}

/* RegistryInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RegistryInspectInternalServerError struct {
}

func (o *RegistryInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] registryInspectInternalServerError ", 500)
}

func (o *RegistryInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
