// Code generated by go-swagger; DO NOT EDIT.

package registries

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

// RegistryUpdateReader is a Reader for the RegistryUpdate structure.
type RegistryUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegistryUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegistryUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegistryUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegistryUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegistryUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegistryUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /registries/{id}] RegistryUpdate", response, response.Code())
	}
}

// NewRegistryUpdateOK creates a RegistryUpdateOK with default headers values
func NewRegistryUpdateOK() *RegistryUpdateOK {
	return &RegistryUpdateOK{}
}

/*
RegistryUpdateOK describes a response with status code 200, with default header values.

Success
*/
type RegistryUpdateOK struct {
	Payload *models.PortainereeRegistry
}

// IsSuccess returns true when this registry update o k response has a 2xx status code
func (o *RegistryUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this registry update o k response has a 3xx status code
func (o *RegistryUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry update o k response has a 4xx status code
func (o *RegistryUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry update o k response has a 5xx status code
func (o *RegistryUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this registry update o k response a status code equal to that given
func (o *RegistryUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the registry update o k response
func (o *RegistryUpdateOK) Code() int {
	return 200
}

func (o *RegistryUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateOK %s", 200, payload)
}

func (o *RegistryUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateOK %s", 200, payload)
}

func (o *RegistryUpdateOK) GetPayload() *models.PortainereeRegistry {
	return o.Payload
}

func (o *RegistryUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegistryUpdateBadRequest creates a RegistryUpdateBadRequest with default headers values
func NewRegistryUpdateBadRequest() *RegistryUpdateBadRequest {
	return &RegistryUpdateBadRequest{}
}

/*
RegistryUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type RegistryUpdateBadRequest struct {
}

// IsSuccess returns true when this registry update bad request response has a 2xx status code
func (o *RegistryUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry update bad request response has a 3xx status code
func (o *RegistryUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry update bad request response has a 4xx status code
func (o *RegistryUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this registry update bad request response has a 5xx status code
func (o *RegistryUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this registry update bad request response a status code equal to that given
func (o *RegistryUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the registry update bad request response
func (o *RegistryUpdateBadRequest) Code() int {
	return 400
}

func (o *RegistryUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateBadRequest", 400)
}

func (o *RegistryUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateBadRequest", 400)
}

func (o *RegistryUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryUpdateNotFound creates a RegistryUpdateNotFound with default headers values
func NewRegistryUpdateNotFound() *RegistryUpdateNotFound {
	return &RegistryUpdateNotFound{}
}

/*
RegistryUpdateNotFound describes a response with status code 404, with default header values.

Registry not found
*/
type RegistryUpdateNotFound struct {
}

// IsSuccess returns true when this registry update not found response has a 2xx status code
func (o *RegistryUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry update not found response has a 3xx status code
func (o *RegistryUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry update not found response has a 4xx status code
func (o *RegistryUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this registry update not found response has a 5xx status code
func (o *RegistryUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this registry update not found response a status code equal to that given
func (o *RegistryUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the registry update not found response
func (o *RegistryUpdateNotFound) Code() int {
	return 404
}

func (o *RegistryUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateNotFound", 404)
}

func (o *RegistryUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateNotFound", 404)
}

func (o *RegistryUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryUpdateConflict creates a RegistryUpdateConflict with default headers values
func NewRegistryUpdateConflict() *RegistryUpdateConflict {
	return &RegistryUpdateConflict{}
}

/*
RegistryUpdateConflict describes a response with status code 409, with default header values.

Another registry with either the same name or same URL & credentials already exists
*/
type RegistryUpdateConflict struct {
}

// IsSuccess returns true when this registry update conflict response has a 2xx status code
func (o *RegistryUpdateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry update conflict response has a 3xx status code
func (o *RegistryUpdateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry update conflict response has a 4xx status code
func (o *RegistryUpdateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this registry update conflict response has a 5xx status code
func (o *RegistryUpdateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this registry update conflict response a status code equal to that given
func (o *RegistryUpdateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the registry update conflict response
func (o *RegistryUpdateConflict) Code() int {
	return 409
}

func (o *RegistryUpdateConflict) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateConflict", 409)
}

func (o *RegistryUpdateConflict) String() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateConflict", 409)
}

func (o *RegistryUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegistryUpdateInternalServerError creates a RegistryUpdateInternalServerError with default headers values
func NewRegistryUpdateInternalServerError() *RegistryUpdateInternalServerError {
	return &RegistryUpdateInternalServerError{}
}

/*
RegistryUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RegistryUpdateInternalServerError struct {
}

// IsSuccess returns true when this registry update internal server error response has a 2xx status code
func (o *RegistryUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this registry update internal server error response has a 3xx status code
func (o *RegistryUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this registry update internal server error response has a 4xx status code
func (o *RegistryUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this registry update internal server error response has a 5xx status code
func (o *RegistryUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this registry update internal server error response a status code equal to that given
func (o *RegistryUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the registry update internal server error response
func (o *RegistryUpdateInternalServerError) Code() int {
	return 500
}

func (o *RegistryUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateInternalServerError", 500)
}

func (o *RegistryUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /registries/{id}][%d] registryUpdateInternalServerError", 500)
}

func (o *RegistryUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
