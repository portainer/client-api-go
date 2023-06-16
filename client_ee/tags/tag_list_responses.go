// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ee"
)

// TagListReader is a Reader for the TagList structure.
type TagListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTagListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTagListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTagListOK creates a TagListOK with default headers values
func NewTagListOK() *TagListOK {
	return &TagListOK{}
}

/*
TagListOK describes a response with status code 200, with default header values.

Success
*/
type TagListOK struct {
	Payload []*models_ee.PortainereeTag
}

// IsSuccess returns true when this tag list o k response has a 2xx status code
func (o *TagListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tag list o k response has a 3xx status code
func (o *TagListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag list o k response has a 4xx status code
func (o *TagListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag list o k response has a 5xx status code
func (o *TagListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tag list o k response a status code equal to that given
func (o *TagListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TagListOK) Error() string {
	return fmt.Sprintf("[GET /tags][%d] tagListOK  %+v", 200, o.Payload)
}

func (o *TagListOK) String() string {
	return fmt.Sprintf("[GET /tags][%d] tagListOK  %+v", 200, o.Payload)
}

func (o *TagListOK) GetPayload() []*models_ee.PortainereeTag {
	return o.Payload
}

func (o *TagListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagListInternalServerError creates a TagListInternalServerError with default headers values
func NewTagListInternalServerError() *TagListInternalServerError {
	return &TagListInternalServerError{}
}

/*
TagListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TagListInternalServerError struct {
}

// IsSuccess returns true when this tag list internal server error response has a 2xx status code
func (o *TagListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tag list internal server error response has a 3xx status code
func (o *TagListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tag list internal server error response has a 4xx status code
func (o *TagListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tag list internal server error response has a 5xx status code
func (o *TagListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tag list internal server error response a status code equal to that given
func (o *TagListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TagListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tags][%d] tagListInternalServerError ", 500)
}

func (o *TagListInternalServerError) String() string {
	return fmt.Sprintf("[GET /tags][%d] tagListInternalServerError ", 500)
}

func (o *TagListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
