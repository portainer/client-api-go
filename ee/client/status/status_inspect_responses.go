// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// StatusInspectReader is a Reader for the StatusInspect structure.
type StatusInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStatusInspectOK creates a StatusInspectOK with default headers values
func NewStatusInspectOK() *StatusInspectOK {
	return &StatusInspectOK{}
}

/*
StatusInspectOK describes a response with status code 200, with default header values.

Success
*/
type StatusInspectOK struct {
	Payload *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus
}

// IsSuccess returns true when this status inspect o k response has a 2xx status code
func (o *StatusInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status inspect o k response has a 3xx status code
func (o *StatusInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status inspect o k response has a 4xx status code
func (o *StatusInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status inspect o k response has a 5xx status code
func (o *StatusInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status inspect o k response a status code equal to that given
func (o *StatusInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *StatusInspectOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] statusInspectOK  %+v", 200, o.Payload)
}

func (o *StatusInspectOK) String() string {
	return fmt.Sprintf("[GET /status][%d] statusInspectOK  %+v", 200, o.Payload)
}

func (o *StatusInspectOK) GetPayload() *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus {
	return o.Payload
}

func (o *StatusInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
