// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// PutWebhooksIDReader is a Reader for the PutWebhooksID structure.
type PutWebhooksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebhooksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWebhooksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutWebhooksIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutWebhooksIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutWebhooksIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /webhooks/{id}] PutWebhooksID", response, response.Code())
	}
}

// NewPutWebhooksIDOK creates a PutWebhooksIDOK with default headers values
func NewPutWebhooksIDOK() *PutWebhooksIDOK {
	return &PutWebhooksIDOK{}
}

/*
PutWebhooksIDOK describes a response with status code 200, with default header values.

OK
*/
type PutWebhooksIDOK struct {
	Payload *models.PortainerWebhook
}

// IsSuccess returns true when this put webhooks Id o k response has a 2xx status code
func (o *PutWebhooksIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put webhooks Id o k response has a 3xx status code
func (o *PutWebhooksIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id o k response has a 4xx status code
func (o *PutWebhooksIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put webhooks Id o k response has a 5xx status code
func (o *PutWebhooksIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id o k response a status code equal to that given
func (o *PutWebhooksIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put webhooks Id o k response
func (o *PutWebhooksIDOK) Code() int {
	return 200
}

func (o *PutWebhooksIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdOK %s", 200, payload)
}

func (o *PutWebhooksIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdOK %s", 200, payload)
}

func (o *PutWebhooksIDOK) GetPayload() *models.PortainerWebhook {
	return o.Payload
}

func (o *PutWebhooksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerWebhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebhooksIDBadRequest creates a PutWebhooksIDBadRequest with default headers values
func NewPutWebhooksIDBadRequest() *PutWebhooksIDBadRequest {
	return &PutWebhooksIDBadRequest{}
}

/*
PutWebhooksIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutWebhooksIDBadRequest struct {
}

// IsSuccess returns true when this put webhooks Id bad request response has a 2xx status code
func (o *PutWebhooksIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id bad request response has a 3xx status code
func (o *PutWebhooksIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id bad request response has a 4xx status code
func (o *PutWebhooksIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put webhooks Id bad request response has a 5xx status code
func (o *PutWebhooksIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id bad request response a status code equal to that given
func (o *PutWebhooksIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put webhooks Id bad request response
func (o *PutWebhooksIDBadRequest) Code() int {
	return 400
}

func (o *PutWebhooksIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdBadRequest", 400)
}

func (o *PutWebhooksIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdBadRequest", 400)
}

func (o *PutWebhooksIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDConflict creates a PutWebhooksIDConflict with default headers values
func NewPutWebhooksIDConflict() *PutWebhooksIDConflict {
	return &PutWebhooksIDConflict{}
}

/*
PutWebhooksIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutWebhooksIDConflict struct {
}

// IsSuccess returns true when this put webhooks Id conflict response has a 2xx status code
func (o *PutWebhooksIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id conflict response has a 3xx status code
func (o *PutWebhooksIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id conflict response has a 4xx status code
func (o *PutWebhooksIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put webhooks Id conflict response has a 5xx status code
func (o *PutWebhooksIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id conflict response a status code equal to that given
func (o *PutWebhooksIDConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the put webhooks Id conflict response
func (o *PutWebhooksIDConflict) Code() int {
	return 409
}

func (o *PutWebhooksIDConflict) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdConflict", 409)
}

func (o *PutWebhooksIDConflict) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdConflict", 409)
}

func (o *PutWebhooksIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDInternalServerError creates a PutWebhooksIDInternalServerError with default headers values
func NewPutWebhooksIDInternalServerError() *PutWebhooksIDInternalServerError {
	return &PutWebhooksIDInternalServerError{}
}

/*
PutWebhooksIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutWebhooksIDInternalServerError struct {
}

// IsSuccess returns true when this put webhooks Id internal server error response has a 2xx status code
func (o *PutWebhooksIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id internal server error response has a 3xx status code
func (o *PutWebhooksIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id internal server error response has a 4xx status code
func (o *PutWebhooksIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put webhooks Id internal server error response has a 5xx status code
func (o *PutWebhooksIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put webhooks Id internal server error response a status code equal to that given
func (o *PutWebhooksIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put webhooks Id internal server error response
func (o *PutWebhooksIDInternalServerError) Code() int {
	return 500
}

func (o *PutWebhooksIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdInternalServerError", 500)
}

func (o *PutWebhooksIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdInternalServerError", 500)
}

func (o *PutWebhooksIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
