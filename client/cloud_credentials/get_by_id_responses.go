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

// GetByIDReader is a Reader for the GetByID structure.
type GetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetByIDOK creates a GetByIDOK with default headers values
func NewGetByIDOK() *GetByIDOK {
	return &GetByIDOK{}
}

/*
GetByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetByIDOK struct {
	Payload *models.ModelsCloudCredential
}

// IsSuccess returns true when this get by Id o k response has a 2xx status code
func (o *GetByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get by Id o k response has a 3xx status code
func (o *GetByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get by Id o k response has a 4xx status code
func (o *GetByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get by Id o k response has a 5xx status code
func (o *GetByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get by Id o k response a status code equal to that given
func (o *GetByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetByIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdOK  %+v", 200, o.Payload)
}

func (o *GetByIDOK) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdOK  %+v", 200, o.Payload)
}

func (o *GetByIDOK) GetPayload() *models.ModelsCloudCredential {
	return o.Payload
}

func (o *GetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCloudCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByIDBadRequest creates a GetByIDBadRequest with default headers values
func NewGetByIDBadRequest() *GetByIDBadRequest {
	return &GetByIDBadRequest{}
}

/*
GetByIDBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetByIDBadRequest struct {
}

// IsSuccess returns true when this get by Id bad request response has a 2xx status code
func (o *GetByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get by Id bad request response has a 3xx status code
func (o *GetByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get by Id bad request response has a 4xx status code
func (o *GetByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get by Id bad request response has a 5xx status code
func (o *GetByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get by Id bad request response a status code equal to that given
func (o *GetByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdBadRequest ", 400)
}

func (o *GetByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdBadRequest ", 400)
}

func (o *GetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDInternalServerError creates a GetByIDInternalServerError with default headers values
func NewGetByIDInternalServerError() *GetByIDInternalServerError {
	return &GetByIDInternalServerError{}
}

/*
GetByIDInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetByIDInternalServerError struct {
}

// IsSuccess returns true when this get by Id internal server error response has a 2xx status code
func (o *GetByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get by Id internal server error response has a 3xx status code
func (o *GetByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get by Id internal server error response has a 4xx status code
func (o *GetByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get by Id internal server error response has a 5xx status code
func (o *GetByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get by Id internal server error response a status code equal to that given
func (o *GetByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdInternalServerError ", 500)
}

func (o *GetByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getByIdInternalServerError ", 500)
}

func (o *GetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
