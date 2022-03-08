// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OpenAMTConfigureReader is a Reader for the OpenAMTConfigure structure.
type OpenAMTConfigureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenAMTConfigureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOpenAMTConfigureNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenAMTConfigureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenAMTConfigureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenAMTConfigureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenAMTConfigureNoContent creates a OpenAMTConfigureNoContent with default headers values
func NewOpenAMTConfigureNoContent() *OpenAMTConfigureNoContent {
	return &OpenAMTConfigureNoContent{}
}

/* OpenAMTConfigureNoContent describes a response with status code 204, with default header values.

Success
*/
type OpenAMTConfigureNoContent struct {
}

func (o *OpenAMTConfigureNoContent) Error() string {
	return fmt.Sprintf("[POST /open_amt][%d] openAMTConfigureNoContent ", 204)
}

func (o *OpenAMTConfigureNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTConfigureBadRequest creates a OpenAMTConfigureBadRequest with default headers values
func NewOpenAMTConfigureBadRequest() *OpenAMTConfigureBadRequest {
	return &OpenAMTConfigureBadRequest{}
}

/* OpenAMTConfigureBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type OpenAMTConfigureBadRequest struct {
}

func (o *OpenAMTConfigureBadRequest) Error() string {
	return fmt.Sprintf("[POST /open_amt][%d] openAMTConfigureBadRequest ", 400)
}

func (o *OpenAMTConfigureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTConfigureForbidden creates a OpenAMTConfigureForbidden with default headers values
func NewOpenAMTConfigureForbidden() *OpenAMTConfigureForbidden {
	return &OpenAMTConfigureForbidden{}
}

/* OpenAMTConfigureForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type OpenAMTConfigureForbidden struct {
}

func (o *OpenAMTConfigureForbidden) Error() string {
	return fmt.Sprintf("[POST /open_amt][%d] openAMTConfigureForbidden ", 403)
}

func (o *OpenAMTConfigureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAMTConfigureInternalServerError creates a OpenAMTConfigureInternalServerError with default headers values
func NewOpenAMTConfigureInternalServerError() *OpenAMTConfigureInternalServerError {
	return &OpenAMTConfigureInternalServerError{}
}

/* OpenAMTConfigureInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type OpenAMTConfigureInternalServerError struct {
}

func (o *OpenAMTConfigureInternalServerError) Error() string {
	return fmt.Sprintf("[POST /open_amt][%d] openAMTConfigureInternalServerError ", 500)
}

func (o *OpenAMTConfigureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
