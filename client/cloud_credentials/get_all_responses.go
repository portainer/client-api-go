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

// GetAllReader is a Reader for the GetAll structure.
type GetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/credentials] getAll", response, response.Code())
	}
}

// NewGetAllOK creates a GetAllOK with default headers values
func NewGetAllOK() *GetAllOK {
	return &GetAllOK{}
}

/*
GetAllOK describes a response with status code 200, with default header values.

OK
*/
type GetAllOK struct {
	Payload *models.ModelsCloudCredential
}

// IsSuccess returns true when this get all o k response has a 2xx status code
func (o *GetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all o k response has a 3xx status code
func (o *GetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all o k response has a 4xx status code
func (o *GetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all o k response has a 5xx status code
func (o *GetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all o k response a status code equal to that given
func (o *GetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all o k response
func (o *GetAllOK) Code() int {
	return 200
}

func (o *GetAllOK) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllOK  %+v", 200, o.Payload)
}

func (o *GetAllOK) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllOK  %+v", 200, o.Payload)
}

func (o *GetAllOK) GetPayload() *models.ModelsCloudCredential {
	return o.Payload
}

func (o *GetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCloudCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllBadRequest creates a GetAllBadRequest with default headers values
func NewGetAllBadRequest() *GetAllBadRequest {
	return &GetAllBadRequest{}
}

/*
GetAllBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetAllBadRequest struct {
}

// IsSuccess returns true when this get all bad request response has a 2xx status code
func (o *GetAllBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all bad request response has a 3xx status code
func (o *GetAllBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all bad request response has a 4xx status code
func (o *GetAllBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all bad request response has a 5xx status code
func (o *GetAllBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all bad request response a status code equal to that given
func (o *GetAllBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get all bad request response
func (o *GetAllBadRequest) Code() int {
	return 400
}

func (o *GetAllBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllBadRequest ", 400)
}

func (o *GetAllBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllBadRequest ", 400)
}

func (o *GetAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllInternalServerError creates a GetAllInternalServerError with default headers values
func NewGetAllInternalServerError() *GetAllInternalServerError {
	return &GetAllInternalServerError{}
}

/*
GetAllInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetAllInternalServerError struct {
}

// IsSuccess returns true when this get all internal server error response has a 2xx status code
func (o *GetAllInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all internal server error response has a 3xx status code
func (o *GetAllInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all internal server error response has a 4xx status code
func (o *GetAllInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all internal server error response has a 5xx status code
func (o *GetAllInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all internal server error response a status code equal to that given
func (o *GetAllInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all internal server error response
func (o *GetAllInternalServerError) Code() int {
	return 500
}

func (o *GetAllInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllInternalServerError ", 500)
}

func (o *GetAllInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/credentials][%d] getAllInternalServerError ", 500)
}

func (o *GetAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
