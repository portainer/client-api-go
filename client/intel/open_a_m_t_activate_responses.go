// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OpenAMTActivateReader is a Reader for the OpenAMTActivate structure.
type OpenAMTActivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenAMTActivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenAMTActivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenAMTActivateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenAMTActivateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenAMTActivateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenAMTActivateOK creates a OpenAMTActivateOK with default headers values
func NewOpenAMTActivateOK() *OpenAMTActivateOK {
	return &OpenAMTActivateOK{}
}

/* OpenAMTActivateOK describes a response with status code 200, with default header values.

Success
*/
type OpenAMTActivateOK struct {
}

func (o *OpenAMTActivateOK) Error() string {
	return fmt.Sprintf("[POST /open_amt/{id}/activate][%d] openAMTActivateOK ", 200)
}

func (o *OpenAMTActivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTActivateBadRequest creates a OpenAMTActivateBadRequest with default headers values
func NewOpenAMTActivateBadRequest() *OpenAMTActivateBadRequest {
	return &OpenAMTActivateBadRequest{}
}

/* OpenAMTActivateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type OpenAMTActivateBadRequest struct {
}

func (o *OpenAMTActivateBadRequest) Error() string {
	return fmt.Sprintf("[POST /open_amt/{id}/activate][%d] openAMTActivateBadRequest ", 400)
}

func (o *OpenAMTActivateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTActivateForbidden creates a OpenAMTActivateForbidden with default headers values
func NewOpenAMTActivateForbidden() *OpenAMTActivateForbidden {
	return &OpenAMTActivateForbidden{}
}

/* OpenAMTActivateForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type OpenAMTActivateForbidden struct {
}

func (o *OpenAMTActivateForbidden) Error() string {
	return fmt.Sprintf("[POST /open_amt/{id}/activate][%d] openAMTActivateForbidden ", 403)
}

func (o *OpenAMTActivateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTActivateInternalServerError creates a OpenAMTActivateInternalServerError with default headers values
func NewOpenAMTActivateInternalServerError() *OpenAMTActivateInternalServerError {
	return &OpenAMTActivateInternalServerError{}
}

/* OpenAMTActivateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type OpenAMTActivateInternalServerError struct {
}

func (o *OpenAMTActivateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open_amt/{id}/activate][%d] openAMTActivateInternalServerError ", 500)
}

func (o *OpenAMTActivateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
