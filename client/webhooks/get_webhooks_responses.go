// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/models"
)

// GetWebhooksReader is a Reader for the GetWebhooks structure.
type GetWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebhooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebhooksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebhooksOK creates a GetWebhooksOK with default headers values
func NewGetWebhooksOK() *GetWebhooksOK {
	return &GetWebhooksOK{}
}

/*
GetWebhooksOK describes a response with status code 200, with default header values.

OK
*/
type GetWebhooksOK struct {
	Payload []*models.PortainereeWebhook
}

// IsSuccess returns true when this get webhooks o k response has a 2xx status code
func (o *GetWebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get webhooks o k response has a 3xx status code
func (o *GetWebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhooks o k response has a 4xx status code
func (o *GetWebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webhooks o k response has a 5xx status code
func (o *GetWebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhooks o k response a status code equal to that given
func (o *GetWebhooksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebhooksOK) Error() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksOK  %+v", 200, o.Payload)
}

func (o *GetWebhooksOK) String() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksOK  %+v", 200, o.Payload)
}

func (o *GetWebhooksOK) GetPayload() []*models.PortainereeWebhook {
	return o.Payload
}

func (o *GetWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhooksBadRequest creates a GetWebhooksBadRequest with default headers values
func NewGetWebhooksBadRequest() *GetWebhooksBadRequest {
	return &GetWebhooksBadRequest{}
}

/*
GetWebhooksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetWebhooksBadRequest struct {
}

// IsSuccess returns true when this get webhooks bad request response has a 2xx status code
func (o *GetWebhooksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhooks bad request response has a 3xx status code
func (o *GetWebhooksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhooks bad request response has a 4xx status code
func (o *GetWebhooksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webhooks bad request response has a 5xx status code
func (o *GetWebhooksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhooks bad request response a status code equal to that given
func (o *GetWebhooksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWebhooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksBadRequest ", 400)
}

func (o *GetWebhooksBadRequest) String() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksBadRequest ", 400)
}

func (o *GetWebhooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebhooksInternalServerError creates a GetWebhooksInternalServerError with default headers values
func NewGetWebhooksInternalServerError() *GetWebhooksInternalServerError {
	return &GetWebhooksInternalServerError{}
}

/*
GetWebhooksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetWebhooksInternalServerError struct {
}

// IsSuccess returns true when this get webhooks internal server error response has a 2xx status code
func (o *GetWebhooksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhooks internal server error response has a 3xx status code
func (o *GetWebhooksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhooks internal server error response has a 4xx status code
func (o *GetWebhooksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webhooks internal server error response has a 5xx status code
func (o *GetWebhooksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get webhooks internal server error response a status code equal to that given
func (o *GetWebhooksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWebhooksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksInternalServerError ", 500)
}

func (o *GetWebhooksInternalServerError) String() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksInternalServerError ", 500)
}

func (o *GetWebhooksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
