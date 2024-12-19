// Code generated by go-swagger; DO NOT EDIT.

package status

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

// StatusNodesCountReader is a Reader for the StatusNodesCount structure.
type StatusNodesCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusNodesCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusNodesCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewStatusNodesCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /status/nodes] statusNodesCount", response, response.Code())
	}
}

// NewStatusNodesCountOK creates a StatusNodesCountOK with default headers values
func NewStatusNodesCountOK() *StatusNodesCountOK {
	return &StatusNodesCountOK{}
}

/*
StatusNodesCountOK describes a response with status code 200, with default header values.

Success
*/
type StatusNodesCountOK struct {
	Payload *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemNodesCountResponse
}

// IsSuccess returns true when this status nodes count o k response has a 2xx status code
func (o *StatusNodesCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status nodes count o k response has a 3xx status code
func (o *StatusNodesCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status nodes count o k response has a 4xx status code
func (o *StatusNodesCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status nodes count o k response has a 5xx status code
func (o *StatusNodesCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status nodes count o k response a status code equal to that given
func (o *StatusNodesCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status nodes count o k response
func (o *StatusNodesCountOK) Code() int {
	return 200
}

func (o *StatusNodesCountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/nodes][%d] statusNodesCountOK %s", 200, payload)
}

func (o *StatusNodesCountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/nodes][%d] statusNodesCountOK %s", 200, payload)
}

func (o *StatusNodesCountOK) GetPayload() *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemNodesCountResponse {
	return o.Payload
}

func (o *StatusNodesCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemNodesCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusNodesCountInternalServerError creates a StatusNodesCountInternalServerError with default headers values
func NewStatusNodesCountInternalServerError() *StatusNodesCountInternalServerError {
	return &StatusNodesCountInternalServerError{}
}

/*
StatusNodesCountInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StatusNodesCountInternalServerError struct {
}

// IsSuccess returns true when this status nodes count internal server error response has a 2xx status code
func (o *StatusNodesCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status nodes count internal server error response has a 3xx status code
func (o *StatusNodesCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status nodes count internal server error response has a 4xx status code
func (o *StatusNodesCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status nodes count internal server error response has a 5xx status code
func (o *StatusNodesCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status nodes count internal server error response a status code equal to that given
func (o *StatusNodesCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status nodes count internal server error response
func (o *StatusNodesCountInternalServerError) Code() int {
	return 500
}

func (o *StatusNodesCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status/nodes][%d] statusNodesCountInternalServerError", 500)
}

func (o *StatusNodesCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /status/nodes][%d] statusNodesCountInternalServerError", 500)
}

func (o *StatusNodesCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
