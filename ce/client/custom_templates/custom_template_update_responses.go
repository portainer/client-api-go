// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// CustomTemplateUpdateReader is a Reader for the CustomTemplateUpdate structure.
type CustomTemplateUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomTemplateUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomTemplateUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomTemplateUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCustomTemplateUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCustomTemplateUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCustomTemplateUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomTemplateUpdateOK creates a CustomTemplateUpdateOK with default headers values
func NewCustomTemplateUpdateOK() *CustomTemplateUpdateOK {
	return &CustomTemplateUpdateOK{}
}

/*
CustomTemplateUpdateOK describes a response with status code 200, with default header values.

Success
*/
type CustomTemplateUpdateOK struct {
	Payload *models.PortainerCustomTemplate
}

// IsSuccess returns true when this custom template update o k response has a 2xx status code
func (o *CustomTemplateUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom template update o k response has a 3xx status code
func (o *CustomTemplateUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template update o k response has a 4xx status code
func (o *CustomTemplateUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template update o k response has a 5xx status code
func (o *CustomTemplateUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template update o k response a status code equal to that given
func (o *CustomTemplateUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CustomTemplateUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateOK  %+v", 200, o.Payload)
}

func (o *CustomTemplateUpdateOK) String() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateOK  %+v", 200, o.Payload)
}

func (o *CustomTemplateUpdateOK) GetPayload() *models.PortainerCustomTemplate {
	return o.Payload
}

func (o *CustomTemplateUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerCustomTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomTemplateUpdateBadRequest creates a CustomTemplateUpdateBadRequest with default headers values
func NewCustomTemplateUpdateBadRequest() *CustomTemplateUpdateBadRequest {
	return &CustomTemplateUpdateBadRequest{}
}

/*
CustomTemplateUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CustomTemplateUpdateBadRequest struct {
}

// IsSuccess returns true when this custom template update bad request response has a 2xx status code
func (o *CustomTemplateUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template update bad request response has a 3xx status code
func (o *CustomTemplateUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template update bad request response has a 4xx status code
func (o *CustomTemplateUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this custom template update bad request response has a 5xx status code
func (o *CustomTemplateUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template update bad request response a status code equal to that given
func (o *CustomTemplateUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CustomTemplateUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateBadRequest ", 400)
}

func (o *CustomTemplateUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateBadRequest ", 400)
}

func (o *CustomTemplateUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomTemplateUpdateForbidden creates a CustomTemplateUpdateForbidden with default headers values
func NewCustomTemplateUpdateForbidden() *CustomTemplateUpdateForbidden {
	return &CustomTemplateUpdateForbidden{}
}

/*
CustomTemplateUpdateForbidden describes a response with status code 403, with default header values.

Permission denied to access template
*/
type CustomTemplateUpdateForbidden struct {
}

// IsSuccess returns true when this custom template update forbidden response has a 2xx status code
func (o *CustomTemplateUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template update forbidden response has a 3xx status code
func (o *CustomTemplateUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template update forbidden response has a 4xx status code
func (o *CustomTemplateUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this custom template update forbidden response has a 5xx status code
func (o *CustomTemplateUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template update forbidden response a status code equal to that given
func (o *CustomTemplateUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CustomTemplateUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateForbidden ", 403)
}

func (o *CustomTemplateUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateForbidden ", 403)
}

func (o *CustomTemplateUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomTemplateUpdateNotFound creates a CustomTemplateUpdateNotFound with default headers values
func NewCustomTemplateUpdateNotFound() *CustomTemplateUpdateNotFound {
	return &CustomTemplateUpdateNotFound{}
}

/*
CustomTemplateUpdateNotFound describes a response with status code 404, with default header values.

Template not found
*/
type CustomTemplateUpdateNotFound struct {
}

// IsSuccess returns true when this custom template update not found response has a 2xx status code
func (o *CustomTemplateUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template update not found response has a 3xx status code
func (o *CustomTemplateUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template update not found response has a 4xx status code
func (o *CustomTemplateUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this custom template update not found response has a 5xx status code
func (o *CustomTemplateUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template update not found response a status code equal to that given
func (o *CustomTemplateUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CustomTemplateUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateNotFound ", 404)
}

func (o *CustomTemplateUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateNotFound ", 404)
}

func (o *CustomTemplateUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomTemplateUpdateInternalServerError creates a CustomTemplateUpdateInternalServerError with default headers values
func NewCustomTemplateUpdateInternalServerError() *CustomTemplateUpdateInternalServerError {
	return &CustomTemplateUpdateInternalServerError{}
}

/*
CustomTemplateUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CustomTemplateUpdateInternalServerError struct {
}

// IsSuccess returns true when this custom template update internal server error response has a 2xx status code
func (o *CustomTemplateUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template update internal server error response has a 3xx status code
func (o *CustomTemplateUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template update internal server error response has a 4xx status code
func (o *CustomTemplateUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template update internal server error response has a 5xx status code
func (o *CustomTemplateUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this custom template update internal server error response a status code equal to that given
func (o *CustomTemplateUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CustomTemplateUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateInternalServerError ", 500)
}

func (o *CustomTemplateUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /custom_templates/{id}][%d] customTemplateUpdateInternalServerError ", 500)
}

func (o *CustomTemplateUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
