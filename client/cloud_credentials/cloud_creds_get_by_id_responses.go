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

// CloudCredsGetByIDReader is a Reader for the CloudCredsGetByID structure.
type CloudCredsGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredsGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredsGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredsGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredsGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/credentials/{id}] cloudCredsGetByID", response, response.Code())
	}
}

// NewCloudCredsGetByIDOK creates a CloudCredsGetByIDOK with default headers values
func NewCloudCredsGetByIDOK() *CloudCredsGetByIDOK {
	return &CloudCredsGetByIDOK{}
}

/*
CloudCredsGetByIDOK describes a response with status code 200, with default header values.

OK
*/
type CloudCredsGetByIDOK struct {
	Payload *models.ModelsCloudCredential
}

// IsSuccess returns true when this cloud creds get by Id o k response has a 2xx status code
func (o *CloudCredsGetByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud creds get by Id o k response has a 3xx status code
func (o *CloudCredsGetByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud creds get by Id o k response has a 4xx status code
func (o *CloudCredsGetByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud creds get by Id o k response has a 5xx status code
func (o *CloudCredsGetByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud creds get by Id o k response a status code equal to that given
func (o *CloudCredsGetByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud creds get by Id o k response
func (o *CloudCredsGetByIDOK) Code() int {
	return 200
}

func (o *CloudCredsGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdOK  %+v", 200, o.Payload)
}

func (o *CloudCredsGetByIDOK) String() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdOK  %+v", 200, o.Payload)
}

func (o *CloudCredsGetByIDOK) GetPayload() *models.ModelsCloudCredential {
	return o.Payload
}

func (o *CloudCredsGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCloudCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredsGetByIDBadRequest creates a CloudCredsGetByIDBadRequest with default headers values
func NewCloudCredsGetByIDBadRequest() *CloudCredsGetByIDBadRequest {
	return &CloudCredsGetByIDBadRequest{}
}

/*
CloudCredsGetByIDBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CloudCredsGetByIDBadRequest struct {
}

// IsSuccess returns true when this cloud creds get by Id bad request response has a 2xx status code
func (o *CloudCredsGetByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud creds get by Id bad request response has a 3xx status code
func (o *CloudCredsGetByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud creds get by Id bad request response has a 4xx status code
func (o *CloudCredsGetByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud creds get by Id bad request response has a 5xx status code
func (o *CloudCredsGetByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud creds get by Id bad request response a status code equal to that given
func (o *CloudCredsGetByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cloud creds get by Id bad request response
func (o *CloudCredsGetByIDBadRequest) Code() int {
	return 400
}

func (o *CloudCredsGetByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdBadRequest ", 400)
}

func (o *CloudCredsGetByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdBadRequest ", 400)
}

func (o *CloudCredsGetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloudCredsGetByIDInternalServerError creates a CloudCredsGetByIDInternalServerError with default headers values
func NewCloudCredsGetByIDInternalServerError() *CloudCredsGetByIDInternalServerError {
	return &CloudCredsGetByIDInternalServerError{}
}

/*
CloudCredsGetByIDInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CloudCredsGetByIDInternalServerError struct {
}

// IsSuccess returns true when this cloud creds get by Id internal server error response has a 2xx status code
func (o *CloudCredsGetByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud creds get by Id internal server error response has a 3xx status code
func (o *CloudCredsGetByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud creds get by Id internal server error response has a 4xx status code
func (o *CloudCredsGetByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud creds get by Id internal server error response has a 5xx status code
func (o *CloudCredsGetByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud creds get by Id internal server error response a status code equal to that given
func (o *CloudCredsGetByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cloud creds get by Id internal server error response
func (o *CloudCredsGetByIDInternalServerError) Code() int {
	return 500
}

func (o *CloudCredsGetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdInternalServerError ", 500)
}

func (o *CloudCredsGetByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/credentials/{id}][%d] cloudCredsGetByIdInternalServerError ", 500)
}

func (o *CloudCredsGetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
