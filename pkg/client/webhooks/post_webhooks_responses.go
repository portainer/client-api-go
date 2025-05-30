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

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// PostWebhooksReader is a Reader for the PostWebhooks structure.
type PostWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebhooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostWebhooksConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebhooksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /webhooks] PostWebhooks", response, response.Code())
	}
}

// NewPostWebhooksOK creates a PostWebhooksOK with default headers values
func NewPostWebhooksOK() *PostWebhooksOK {
	return &PostWebhooksOK{}
}

/*
PostWebhooksOK describes a response with status code 200, with default header values.

OK
*/
type PostWebhooksOK struct {
	Payload *models.PortainerWebhook
}

// IsSuccess returns true when this post webhooks o k response has a 2xx status code
func (o *PostWebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post webhooks o k response has a 3xx status code
func (o *PostWebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webhooks o k response has a 4xx status code
func (o *PostWebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webhooks o k response has a 5xx status code
func (o *PostWebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post webhooks o k response a status code equal to that given
func (o *PostWebhooksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post webhooks o k response
func (o *PostWebhooksOK) Code() int {
	return 200
}

func (o *PostWebhooksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksOK %s", 200, payload)
}

func (o *PostWebhooksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksOK %s", 200, payload)
}

func (o *PostWebhooksOK) GetPayload() *models.PortainerWebhook {
	return o.Payload
}

func (o *PostWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerWebhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebhooksBadRequest creates a PostWebhooksBadRequest with default headers values
func NewPostWebhooksBadRequest() *PostWebhooksBadRequest {
	return &PostWebhooksBadRequest{}
}

/*
PostWebhooksBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type PostWebhooksBadRequest struct {
}

// IsSuccess returns true when this post webhooks bad request response has a 2xx status code
func (o *PostWebhooksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webhooks bad request response has a 3xx status code
func (o *PostWebhooksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webhooks bad request response has a 4xx status code
func (o *PostWebhooksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webhooks bad request response has a 5xx status code
func (o *PostWebhooksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post webhooks bad request response a status code equal to that given
func (o *PostWebhooksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post webhooks bad request response
func (o *PostWebhooksBadRequest) Code() int {
	return 400
}

func (o *PostWebhooksBadRequest) Error() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksBadRequest", 400)
}

func (o *PostWebhooksBadRequest) String() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksBadRequest", 400)
}

func (o *PostWebhooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhooksConflict creates a PostWebhooksConflict with default headers values
func NewPostWebhooksConflict() *PostWebhooksConflict {
	return &PostWebhooksConflict{}
}

/*
PostWebhooksConflict describes a response with status code 409, with default header values.

A webhook for this resource already exists
*/
type PostWebhooksConflict struct {
}

// IsSuccess returns true when this post webhooks conflict response has a 2xx status code
func (o *PostWebhooksConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webhooks conflict response has a 3xx status code
func (o *PostWebhooksConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webhooks conflict response has a 4xx status code
func (o *PostWebhooksConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webhooks conflict response has a 5xx status code
func (o *PostWebhooksConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post webhooks conflict response a status code equal to that given
func (o *PostWebhooksConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post webhooks conflict response
func (o *PostWebhooksConflict) Code() int {
	return 409
}

func (o *PostWebhooksConflict) Error() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksConflict", 409)
}

func (o *PostWebhooksConflict) String() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksConflict", 409)
}

func (o *PostWebhooksConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhooksInternalServerError creates a PostWebhooksInternalServerError with default headers values
func NewPostWebhooksInternalServerError() *PostWebhooksInternalServerError {
	return &PostWebhooksInternalServerError{}
}

/*
PostWebhooksInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type PostWebhooksInternalServerError struct {
}

// IsSuccess returns true when this post webhooks internal server error response has a 2xx status code
func (o *PostWebhooksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webhooks internal server error response has a 3xx status code
func (o *PostWebhooksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webhooks internal server error response has a 4xx status code
func (o *PostWebhooksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webhooks internal server error response has a 5xx status code
func (o *PostWebhooksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post webhooks internal server error response a status code equal to that given
func (o *PostWebhooksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post webhooks internal server error response
func (o *PostWebhooksInternalServerError) Code() int {
	return 500
}

func (o *PostWebhooksInternalServerError) Error() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksInternalServerError", 500)
}

func (o *PostWebhooksInternalServerError) String() string {
	return fmt.Sprintf("[POST /webhooks][%d] postWebhooksInternalServerError", 500)
}

func (o *PostWebhooksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
