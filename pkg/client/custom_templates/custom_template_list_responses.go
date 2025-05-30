// Code generated by go-swagger; DO NOT EDIT.

package custom_templates

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

// CustomTemplateListReader is a Reader for the CustomTemplateList structure.
type CustomTemplateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomTemplateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomTemplateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCustomTemplateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /custom_templates] CustomTemplateList", response, response.Code())
	}
}

// NewCustomTemplateListOK creates a CustomTemplateListOK with default headers values
func NewCustomTemplateListOK() *CustomTemplateListOK {
	return &CustomTemplateListOK{}
}

/*
CustomTemplateListOK describes a response with status code 200, with default header values.

Success
*/
type CustomTemplateListOK struct {
	Payload []*models.PortainereeCustomTemplate
}

// IsSuccess returns true when this custom template list o k response has a 2xx status code
func (o *CustomTemplateListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom template list o k response has a 3xx status code
func (o *CustomTemplateListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template list o k response has a 4xx status code
func (o *CustomTemplateListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template list o k response has a 5xx status code
func (o *CustomTemplateListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom template list o k response a status code equal to that given
func (o *CustomTemplateListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom template list o k response
func (o *CustomTemplateListOK) Code() int {
	return 200
}

func (o *CustomTemplateListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /custom_templates][%d] customTemplateListOK %s", 200, payload)
}

func (o *CustomTemplateListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /custom_templates][%d] customTemplateListOK %s", 200, payload)
}

func (o *CustomTemplateListOK) GetPayload() []*models.PortainereeCustomTemplate {
	return o.Payload
}

func (o *CustomTemplateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomTemplateListInternalServerError creates a CustomTemplateListInternalServerError with default headers values
func NewCustomTemplateListInternalServerError() *CustomTemplateListInternalServerError {
	return &CustomTemplateListInternalServerError{}
}

/*
CustomTemplateListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CustomTemplateListInternalServerError struct {
}

// IsSuccess returns true when this custom template list internal server error response has a 2xx status code
func (o *CustomTemplateListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this custom template list internal server error response has a 3xx status code
func (o *CustomTemplateListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom template list internal server error response has a 4xx status code
func (o *CustomTemplateListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom template list internal server error response has a 5xx status code
func (o *CustomTemplateListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this custom template list internal server error response a status code equal to that given
func (o *CustomTemplateListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the custom template list internal server error response
func (o *CustomTemplateListInternalServerError) Code() int {
	return 500
}

func (o *CustomTemplateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /custom_templates][%d] customTemplateListInternalServerError", 500)
}

func (o *CustomTemplateListInternalServerError) String() string {
	return fmt.Sprintf("[GET /custom_templates][%d] customTemplateListInternalServerError", 500)
}

func (o *CustomTemplateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
