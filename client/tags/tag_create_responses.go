// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/models"
)

// TagCreateReader is a Reader for the TagCreate structure.
type TagCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTagCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewTagCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTagCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTagCreateOK creates a TagCreateOK with default headers values
func NewTagCreateOK() *TagCreateOK {
	return &TagCreateOK{}
}

/*
TagCreateOK describes a response with status code 200, with default header values.

Success
*/
type TagCreateOK struct {
	Payload *models.PortainereeTag
}

// IsSuccess returns true when this tag create o k response has a 2xx status code
func (o *TagCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag create o k response has a 3xx status code
func (o *TagCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag create o k response has a 4xx status code
func (o *TagCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag create o k response has a 5xx status code
func (o *TagCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tag create o k response a status code equal to that given
func (o *TagCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TagCreateOK) Error() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateOK  %+v", 200, o.Payload)
}

func (o *TagCreateOK) String() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateOK  %+v", 200, o.Payload)
}

func (o *TagCreateOK) GetPayload() *models.PortainereeTag {
	return o.Payload
}

func (o *TagCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagCreateConflict creates a TagCreateConflict with default headers values
func NewTagCreateConflict() *TagCreateConflict {
	return &TagCreateConflict{}
}

/*
TagCreateConflict describes a response with status code 409, with default header values.

Tag name exists
*/
type TagCreateConflict struct {
}

// IsSuccess returns true when this tag create conflict response has a 2xx status code
func (o *TagCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag create conflict response has a 3xx status code
func (o *TagCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag create conflict response has a 4xx status code
func (o *TagCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this tag create conflict response has a 5xx status code
func (o *TagCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this tag create conflict response a status code equal to that given
func (o *TagCreateConflict) IsCode(code int) bool {
	return code == 409
}

func (o *TagCreateConflict) Error() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateConflict ", 409)
}

func (o *TagCreateConflict) String() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateConflict ", 409)
}

func (o *TagCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTagCreateInternalServerError creates a TagCreateInternalServerError with default headers values
func NewTagCreateInternalServerError() *TagCreateInternalServerError {
	return &TagCreateInternalServerError{}
}

/*
TagCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TagCreateInternalServerError struct {
}

// IsSuccess returns true when this tag create internal server error response has a 2xx status code
func (o *TagCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag create internal server error response has a 3xx status code
func (o *TagCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag create internal server error response has a 4xx status code
func (o *TagCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag create internal server error response has a 5xx status code
func (o *TagCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tag create internal server error response a status code equal to that given
func (o *TagCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TagCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateInternalServerError ", 500)
}

func (o *TagCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /tags][%d] tagCreateInternalServerError ", 500)
}

func (o *TagCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
