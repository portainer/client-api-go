// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOK creates a DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*
DeleteOK describes a response with status code 200, with default header values.

OK
*/
type DeleteOK struct {
	Payload *models.ModelsCloudCredential
}

// IsSuccess returns true when this delete o k response has a 2xx status code
func (o *DeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete o k response has a 3xx status code
func (o *DeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete o k response has a 4xx status code
func (o *DeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete o k response has a 5xx status code
func (o *DeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete o k response a status code equal to that given
func (o *DeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOK) Error() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) String() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) GetPayload() *models.ModelsCloudCredential {
	return o.Payload
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCloudCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBadRequest creates a DeleteBadRequest with default headers values
func NewDeleteBadRequest() *DeleteBadRequest {
	return &DeleteBadRequest{}
}

/*
DeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type DeleteBadRequest struct {
}

// IsSuccess returns true when this delete bad request response has a 2xx status code
func (o *DeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete bad request response has a 3xx status code
func (o *DeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete bad request response has a 4xx status code
func (o *DeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete bad request response has a 5xx status code
func (o *DeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete bad request response a status code equal to that given
func (o *DeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteBadRequest ", 400)
}

func (o *DeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteBadRequest ", 400)
}

func (o *DeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInternalServerError creates a DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {
	return &DeleteInternalServerError{}
}

/*
DeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteInternalServerError struct {
}

// IsSuccess returns true when this delete internal server error response has a 2xx status code
func (o *DeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete internal server error response has a 3xx status code
func (o *DeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete internal server error response has a 4xx status code
func (o *DeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete internal server error response has a 5xx status code
func (o *DeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete internal server error response a status code equal to that given
func (o *DeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteInternalServerError ", 500)
}

func (o *DeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloudcredentials][%d] deleteInternalServerError ", 500)
}

func (o *DeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
