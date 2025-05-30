// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EcrDeleteRepositoryReader is a Reader for the EcrDeleteRepository structure.
type EcrDeleteRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EcrDeleteRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEcrDeleteRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEcrDeleteRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEcrDeleteRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEcrDeleteRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEcrDeleteRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /registries/{id}/ecr/repositories/{repositoryName}] ecrDeleteRepository", response, response.Code())
	}
}

// NewEcrDeleteRepositoryOK creates a EcrDeleteRepositoryOK with default headers values
func NewEcrDeleteRepositoryOK() *EcrDeleteRepositoryOK {
	return &EcrDeleteRepositoryOK{}
}

/*
EcrDeleteRepositoryOK describes a response with status code 200, with default header values.

Success
*/
type EcrDeleteRepositoryOK struct {
}

// IsSuccess returns true when this ecr delete repository o k response has a 2xx status code
func (o *EcrDeleteRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ecr delete repository o k response has a 3xx status code
func (o *EcrDeleteRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ecr delete repository o k response has a 4xx status code
func (o *EcrDeleteRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ecr delete repository o k response has a 5xx status code
func (o *EcrDeleteRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ecr delete repository o k response a status code equal to that given
func (o *EcrDeleteRepositoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ecr delete repository o k response
func (o *EcrDeleteRepositoryOK) Code() int {
	return 200
}

func (o *EcrDeleteRepositoryOK) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryOK", 200)
}

func (o *EcrDeleteRepositoryOK) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryOK", 200)
}

func (o *EcrDeleteRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEcrDeleteRepositoryBadRequest creates a EcrDeleteRepositoryBadRequest with default headers values
func NewEcrDeleteRepositoryBadRequest() *EcrDeleteRepositoryBadRequest {
	return &EcrDeleteRepositoryBadRequest{}
}

/*
EcrDeleteRepositoryBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EcrDeleteRepositoryBadRequest struct {
}

// IsSuccess returns true when this ecr delete repository bad request response has a 2xx status code
func (o *EcrDeleteRepositoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ecr delete repository bad request response has a 3xx status code
func (o *EcrDeleteRepositoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ecr delete repository bad request response has a 4xx status code
func (o *EcrDeleteRepositoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ecr delete repository bad request response has a 5xx status code
func (o *EcrDeleteRepositoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ecr delete repository bad request response a status code equal to that given
func (o *EcrDeleteRepositoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ecr delete repository bad request response
func (o *EcrDeleteRepositoryBadRequest) Code() int {
	return 400
}

func (o *EcrDeleteRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryBadRequest", 400)
}

func (o *EcrDeleteRepositoryBadRequest) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryBadRequest", 400)
}

func (o *EcrDeleteRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEcrDeleteRepositoryForbidden creates a EcrDeleteRepositoryForbidden with default headers values
func NewEcrDeleteRepositoryForbidden() *EcrDeleteRepositoryForbidden {
	return &EcrDeleteRepositoryForbidden{}
}

/*
EcrDeleteRepositoryForbidden describes a response with status code 403, with default header values.

Permission denied to access registry
*/
type EcrDeleteRepositoryForbidden struct {
}

// IsSuccess returns true when this ecr delete repository forbidden response has a 2xx status code
func (o *EcrDeleteRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ecr delete repository forbidden response has a 3xx status code
func (o *EcrDeleteRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ecr delete repository forbidden response has a 4xx status code
func (o *EcrDeleteRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ecr delete repository forbidden response has a 5xx status code
func (o *EcrDeleteRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ecr delete repository forbidden response a status code equal to that given
func (o *EcrDeleteRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ecr delete repository forbidden response
func (o *EcrDeleteRepositoryForbidden) Code() int {
	return 403
}

func (o *EcrDeleteRepositoryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryForbidden", 403)
}

func (o *EcrDeleteRepositoryForbidden) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryForbidden", 403)
}

func (o *EcrDeleteRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEcrDeleteRepositoryNotFound creates a EcrDeleteRepositoryNotFound with default headers values
func NewEcrDeleteRepositoryNotFound() *EcrDeleteRepositoryNotFound {
	return &EcrDeleteRepositoryNotFound{}
}

/*
EcrDeleteRepositoryNotFound describes a response with status code 404, with default header values.

Registry not found
*/
type EcrDeleteRepositoryNotFound struct {
}

// IsSuccess returns true when this ecr delete repository not found response has a 2xx status code
func (o *EcrDeleteRepositoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ecr delete repository not found response has a 3xx status code
func (o *EcrDeleteRepositoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ecr delete repository not found response has a 4xx status code
func (o *EcrDeleteRepositoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ecr delete repository not found response has a 5xx status code
func (o *EcrDeleteRepositoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ecr delete repository not found response a status code equal to that given
func (o *EcrDeleteRepositoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ecr delete repository not found response
func (o *EcrDeleteRepositoryNotFound) Code() int {
	return 404
}

func (o *EcrDeleteRepositoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryNotFound", 404)
}

func (o *EcrDeleteRepositoryNotFound) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryNotFound", 404)
}

func (o *EcrDeleteRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEcrDeleteRepositoryInternalServerError creates a EcrDeleteRepositoryInternalServerError with default headers values
func NewEcrDeleteRepositoryInternalServerError() *EcrDeleteRepositoryInternalServerError {
	return &EcrDeleteRepositoryInternalServerError{}
}

/*
EcrDeleteRepositoryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EcrDeleteRepositoryInternalServerError struct {
}

// IsSuccess returns true when this ecr delete repository internal server error response has a 2xx status code
func (o *EcrDeleteRepositoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ecr delete repository internal server error response has a 3xx status code
func (o *EcrDeleteRepositoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ecr delete repository internal server error response has a 4xx status code
func (o *EcrDeleteRepositoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ecr delete repository internal server error response has a 5xx status code
func (o *EcrDeleteRepositoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ecr delete repository internal server error response a status code equal to that given
func (o *EcrDeleteRepositoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ecr delete repository internal server error response
func (o *EcrDeleteRepositoryInternalServerError) Code() int {
	return 500
}

func (o *EcrDeleteRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryInternalServerError", 500)
}

func (o *EcrDeleteRepositoryInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}/ecr/repositories/{repositoryName}][%d] ecrDeleteRepositoryInternalServerError", 500)
}

func (o *EcrDeleteRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
