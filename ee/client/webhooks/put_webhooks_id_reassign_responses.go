// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// PutWebhooksIDReassignReader is a Reader for the PutWebhooksIDReassign structure.
type PutWebhooksIDReassignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebhooksIDReassignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWebhooksIDReassignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPutWebhooksIDReassignNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutWebhooksIDReassignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutWebhooksIDReassignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutWebhooksIDReassignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutWebhooksIDReassignOK creates a PutWebhooksIDReassignOK with default headers values
func NewPutWebhooksIDReassignOK() *PutWebhooksIDReassignOK {
	return &PutWebhooksIDReassignOK{}
}

/*
PutWebhooksIDReassignOK describes a response with status code 200, with default header values.

OK
*/
type PutWebhooksIDReassignOK struct {
	Payload *models.PortainereeWebhook
}

// IsSuccess returns true when this put webhooks Id reassign o k response has a 2xx status code
func (o *PutWebhooksIDReassignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put webhooks Id reassign o k response has a 3xx status code
func (o *PutWebhooksIDReassignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id reassign o k response has a 4xx status code
func (o *PutWebhooksIDReassignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put webhooks Id reassign o k response has a 5xx status code
func (o *PutWebhooksIDReassignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id reassign o k response a status code equal to that given
func (o *PutWebhooksIDReassignOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutWebhooksIDReassignOK) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignOK  %+v", 200, o.Payload)
}

func (o *PutWebhooksIDReassignOK) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignOK  %+v", 200, o.Payload)
}

func (o *PutWebhooksIDReassignOK) GetPayload() *models.PortainereeWebhook {
	return o.Payload
}

func (o *PutWebhooksIDReassignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeWebhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebhooksIDReassignNoContent creates a PutWebhooksIDReassignNoContent with default headers values
func NewPutWebhooksIDReassignNoContent() *PutWebhooksIDReassignNoContent {
	return &PutWebhooksIDReassignNoContent{}
}

/*
PutWebhooksIDReassignNoContent describes a response with status code 204, with default header values.

No Content
*/
type PutWebhooksIDReassignNoContent struct {
}

// IsSuccess returns true when this put webhooks Id reassign no content response has a 2xx status code
func (o *PutWebhooksIDReassignNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put webhooks Id reassign no content response has a 3xx status code
func (o *PutWebhooksIDReassignNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id reassign no content response has a 4xx status code
func (o *PutWebhooksIDReassignNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put webhooks Id reassign no content response has a 5xx status code
func (o *PutWebhooksIDReassignNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id reassign no content response a status code equal to that given
func (o *PutWebhooksIDReassignNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PutWebhooksIDReassignNoContent) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignNoContent ", 204)
}

func (o *PutWebhooksIDReassignNoContent) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignNoContent ", 204)
}

func (o *PutWebhooksIDReassignNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDReassignBadRequest creates a PutWebhooksIDReassignBadRequest with default headers values
func NewPutWebhooksIDReassignBadRequest() *PutWebhooksIDReassignBadRequest {
	return &PutWebhooksIDReassignBadRequest{}
}

/*
PutWebhooksIDReassignBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutWebhooksIDReassignBadRequest struct {
}

// IsSuccess returns true when this put webhooks Id reassign bad request response has a 2xx status code
func (o *PutWebhooksIDReassignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id reassign bad request response has a 3xx status code
func (o *PutWebhooksIDReassignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id reassign bad request response has a 4xx status code
func (o *PutWebhooksIDReassignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put webhooks Id reassign bad request response has a 5xx status code
func (o *PutWebhooksIDReassignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id reassign bad request response a status code equal to that given
func (o *PutWebhooksIDReassignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutWebhooksIDReassignBadRequest) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignBadRequest ", 400)
}

func (o *PutWebhooksIDReassignBadRequest) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignBadRequest ", 400)
}

func (o *PutWebhooksIDReassignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDReassignNotFound creates a PutWebhooksIDReassignNotFound with default headers values
func NewPutWebhooksIDReassignNotFound() *PutWebhooksIDReassignNotFound {
	return &PutWebhooksIDReassignNotFound{}
}

/*
PutWebhooksIDReassignNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutWebhooksIDReassignNotFound struct {
}

// IsSuccess returns true when this put webhooks Id reassign not found response has a 2xx status code
func (o *PutWebhooksIDReassignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id reassign not found response has a 3xx status code
func (o *PutWebhooksIDReassignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id reassign not found response has a 4xx status code
func (o *PutWebhooksIDReassignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put webhooks Id reassign not found response has a 5xx status code
func (o *PutWebhooksIDReassignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put webhooks Id reassign not found response a status code equal to that given
func (o *PutWebhooksIDReassignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutWebhooksIDReassignNotFound) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignNotFound ", 404)
}

func (o *PutWebhooksIDReassignNotFound) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignNotFound ", 404)
}

func (o *PutWebhooksIDReassignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDReassignInternalServerError creates a PutWebhooksIDReassignInternalServerError with default headers values
func NewPutWebhooksIDReassignInternalServerError() *PutWebhooksIDReassignInternalServerError {
	return &PutWebhooksIDReassignInternalServerError{}
}

/*
PutWebhooksIDReassignInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutWebhooksIDReassignInternalServerError struct {
}

// IsSuccess returns true when this put webhooks Id reassign internal server error response has a 2xx status code
func (o *PutWebhooksIDReassignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put webhooks Id reassign internal server error response has a 3xx status code
func (o *PutWebhooksIDReassignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put webhooks Id reassign internal server error response has a 4xx status code
func (o *PutWebhooksIDReassignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put webhooks Id reassign internal server error response has a 5xx status code
func (o *PutWebhooksIDReassignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put webhooks Id reassign internal server error response a status code equal to that given
func (o *PutWebhooksIDReassignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutWebhooksIDReassignInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignInternalServerError ", 500)
}

func (o *PutWebhooksIDReassignInternalServerError) String() string {
	return fmt.Sprintf("[PUT /webhooks/{id}/reassign][%d] putWebhooksIdReassignInternalServerError ", 500)
}

func (o *PutWebhooksIDReassignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
