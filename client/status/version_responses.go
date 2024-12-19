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

// VersionReader is a Reader for the Version structure.
type VersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /status/version] Version", response, response.Code())
	}
}

// NewVersionOK creates a VersionOK with default headers values
func NewVersionOK() *VersionOK {
	return &VersionOK{}
}

/*
VersionOK describes a response with status code 200, with default header values.

Success
*/
type VersionOK struct {
	Payload *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemVersionResponse
}

// IsSuccess returns true when this version o k response has a 2xx status code
func (o *VersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this version o k response has a 3xx status code
func (o *VersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this version o k response has a 4xx status code
func (o *VersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this version o k response has a 5xx status code
func (o *VersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this version o k response a status code equal to that given
func (o *VersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the version o k response
func (o *VersionOK) Code() int {
	return 200
}

func (o *VersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/version][%d] versionOK %s", 200, payload)
}

func (o *VersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/version][%d] versionOK %s", 200, payload)
}

func (o *VersionOK) GetPayload() *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemVersionResponse {
	return o.Payload
}

func (o *VersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
