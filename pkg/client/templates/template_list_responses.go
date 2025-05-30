// Code generated by go-swagger; DO NOT EDIT.

package templates

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

// TemplateListReader is a Reader for the TemplateList structure.
type TemplateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTemplateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /templates] TemplateList", response, response.Code())
	}
}

// NewTemplateListOK creates a TemplateListOK with default headers values
func NewTemplateListOK() *TemplateListOK {
	return &TemplateListOK{}
}

/*
TemplateListOK describes a response with status code 200, with default header values.

Success
*/
type TemplateListOK struct {
	Payload *models.TemplatesListResponse
}

// IsSuccess returns true when this template list o k response has a 2xx status code
func (o *TemplateListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template list o k response has a 3xx status code
func (o *TemplateListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template list o k response has a 4xx status code
func (o *TemplateListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template list o k response has a 5xx status code
func (o *TemplateListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template list o k response a status code equal to that given
func (o *TemplateListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template list o k response
func (o *TemplateListOK) Code() int {
	return 200
}

func (o *TemplateListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /templates][%d] templateListOK %s", 200, payload)
}

func (o *TemplateListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /templates][%d] templateListOK %s", 200, payload)
}

func (o *TemplateListOK) GetPayload() *models.TemplatesListResponse {
	return o.Payload
}

func (o *TemplateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplatesListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateListInternalServerError creates a TemplateListInternalServerError with default headers values
func NewTemplateListInternalServerError() *TemplateListInternalServerError {
	return &TemplateListInternalServerError{}
}

/*
TemplateListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TemplateListInternalServerError struct {
}

// IsSuccess returns true when this template list internal server error response has a 2xx status code
func (o *TemplateListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this template list internal server error response has a 3xx status code
func (o *TemplateListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template list internal server error response has a 4xx status code
func (o *TemplateListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this template list internal server error response has a 5xx status code
func (o *TemplateListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this template list internal server error response a status code equal to that given
func (o *TemplateListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the template list internal server error response
func (o *TemplateListInternalServerError) Code() int {
	return 500
}

func (o *TemplateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /templates][%d] templateListInternalServerError", 500)
}

func (o *TemplateListInternalServerError) String() string {
	return fmt.Sprintf("[GET /templates][%d] templateListInternalServerError", 500)
}

func (o *TemplateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
