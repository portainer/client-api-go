// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud/credentials] Create", response, response.Code())
	}
}

// NewCreateOK creates a CreateOK with default headers values
func NewCreateOK() *CreateOK {
	return &CreateOK{}
}

/*
CreateOK describes a response with status code 200, with default header values.

OK
*/
type CreateOK struct {
	Payload *models.ModelsCloudCredential
}

// IsSuccess returns true when this create o k response has a 2xx status code
func (o *CreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o k response has a 3xx status code
func (o *CreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o k response has a 4xx status code
func (o *CreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o k response has a 5xx status code
func (o *CreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create o k response a status code equal to that given
func (o *CreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create o k response
func (o *CreateOK) Code() int {
	return 200
}

func (o *CreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/credentials][%d] createOK %s", 200, payload)
}

func (o *CreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/credentials][%d] createOK %s", 200, payload)
}

func (o *CreateOK) GetPayload() *models.ModelsCloudCredential {
	return o.Payload
}

func (o *CreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCloudCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBadRequest creates a CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {
	return &CreateBadRequest{}
}

/*
CreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CreateBadRequest struct {
}

// IsSuccess returns true when this create bad request response has a 2xx status code
func (o *CreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bad request response has a 3xx status code
func (o *CreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bad request response has a 4xx status code
func (o *CreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bad request response has a 5xx status code
func (o *CreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bad request response a status code equal to that given
func (o *CreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create bad request response
func (o *CreateBadRequest) Code() int {
	return 400
}

func (o *CreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud/credentials][%d] createBadRequest", 400)
}

func (o *CreateBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud/credentials][%d] createBadRequest", 400)
}

func (o *CreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInternalServerError creates a CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {
	return &CreateInternalServerError{}
}

/*
CreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateInternalServerError struct {
}

// IsSuccess returns true when this create internal server error response has a 2xx status code
func (o *CreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create internal server error response has a 3xx status code
func (o *CreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create internal server error response has a 4xx status code
func (o *CreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create internal server error response has a 5xx status code
func (o *CreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create internal server error response a status code equal to that given
func (o *CreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create internal server error response
func (o *CreateInternalServerError) Code() int {
	return 500
}

func (o *CreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud/credentials][%d] createInternalServerError", 500)
}

func (o *CreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud/credentials][%d] createInternalServerError", 500)
}

func (o *CreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
