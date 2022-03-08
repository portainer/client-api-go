// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DuplicateReader is a Reader for the Duplicate structure.
type DuplicateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuplicateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDuplicateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDuplicateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDuplicateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDuplicateOK creates a DuplicateOK with default headers values
func NewDuplicateOK() *DuplicateOK {
	return &DuplicateOK{}
}

/* DuplicateOK describes a response with status code 200, with default header values.

Success
*/
type DuplicateOK struct {
}

func (o *DuplicateOK) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles/{id}/duplicate][%d] duplicateOK ", 200)
}

func (o *DuplicateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDuplicateBadRequest creates a DuplicateBadRequest with default headers values
func NewDuplicateBadRequest() *DuplicateBadRequest {
	return &DuplicateBadRequest{}
}

/* DuplicateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type DuplicateBadRequest struct {
}

func (o *DuplicateBadRequest) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles/{id}/duplicate][%d] duplicateBadRequest ", 400)
}

func (o *DuplicateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDuplicateInternalServerError creates a DuplicateInternalServerError with default headers values
func NewDuplicateInternalServerError() *DuplicateInternalServerError {
	return &DuplicateInternalServerError{}
}

/* DuplicateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DuplicateInternalServerError struct {
}

func (o *DuplicateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles/{id}/duplicate][%d] duplicateInternalServerError ", 500)
}

func (o *DuplicateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
