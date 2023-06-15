// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// CustomTemplateCreateReader is a Reader for the CustomTemplateCreate structure.
type CustomTemplateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomTemplateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomTemplateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomTemplateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCustomTemplateCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomTemplateCreateOK creates a CustomTemplateCreateOK with default headers values
func NewCustomTemplateCreateOK() *CustomTemplateCreateOK {
	return &CustomTemplateCreateOK{}
}

/*
CustomTemplateCreateOK describes a response with status code 200, with default header values.

OK
*/
type CustomTemplateCreateOK struct {
	Payload *models.PortainereeCustomTemplate
}

// IsSuccess returns true when this custom template create o k response has a 2xx status code
func (o *CustomTemplateCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom template create o k response has a 3xx status code
func (o *CustomTemplateCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template create o k response has a 4xx status code
func (o *CustomTemplateCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template create o k response has a 5xx status code
func (o *CustomTemplateCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template create o k response a status code equal to that given
func (o *CustomTemplateCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CustomTemplateCreateOK) Error() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateOK  %+v", 200, o.Payload)
}

func (o *CustomTemplateCreateOK) String() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateOK  %+v", 200, o.Payload)
}

func (o *CustomTemplateCreateOK) GetPayload() *models.PortainereeCustomTemplate {
	return o.Payload
}

func (o *CustomTemplateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeCustomTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomTemplateCreateBadRequest creates a CustomTemplateCreateBadRequest with default headers values
func NewCustomTemplateCreateBadRequest() *CustomTemplateCreateBadRequest {
	return &CustomTemplateCreateBadRequest{}
}

/*
CustomTemplateCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CustomTemplateCreateBadRequest struct {
}

// IsSuccess returns true when this custom template create bad request response has a 2xx status code
func (o *CustomTemplateCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template create bad request response has a 3xx status code
func (o *CustomTemplateCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template create bad request response has a 4xx status code
func (o *CustomTemplateCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this custom template create bad request response has a 5xx status code
func (o *CustomTemplateCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template create bad request response a status code equal to that given
func (o *CustomTemplateCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CustomTemplateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateBadRequest ", 400)
}

func (o *CustomTemplateCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateBadRequest ", 400)
}

func (o *CustomTemplateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomTemplateCreateInternalServerError creates a CustomTemplateCreateInternalServerError with default headers values
func NewCustomTemplateCreateInternalServerError() *CustomTemplateCreateInternalServerError {
	return &CustomTemplateCreateInternalServerError{}
}

/*
CustomTemplateCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CustomTemplateCreateInternalServerError struct {
}

// IsSuccess returns true when this custom template create internal server error response has a 2xx status code
func (o *CustomTemplateCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template create internal server error response has a 3xx status code
func (o *CustomTemplateCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template create internal server error response has a 4xx status code
func (o *CustomTemplateCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template create internal server error response has a 5xx status code
func (o *CustomTemplateCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this custom template create internal server error response a status code equal to that given
func (o *CustomTemplateCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CustomTemplateCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateInternalServerError ", 500)
}

func (o *CustomTemplateCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /custom_templates][%d] customTemplateCreateInternalServerError ", 500)
}

func (o *CustomTemplateCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
