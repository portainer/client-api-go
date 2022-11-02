// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// SnapshotInspectReader is a Reader for the SnapshotInspect structure.
type SnapshotInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSnapshotInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSnapshotInspectOK creates a SnapshotInspectOK with default headers values
func NewSnapshotInspectOK() *SnapshotInspectOK {
	return &SnapshotInspectOK{}
}

/*
SnapshotInspectOK describes a response with status code 200, with default header values.

Success
*/
type SnapshotInspectOK struct {
	Payload models.PortainerDockerSnapshotRaw
}

// IsSuccess returns true when this snapshot inspect o k response has a 2xx status code
func (o *SnapshotInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot inspect o k response has a 3xx status code
func (o *SnapshotInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot inspect o k response has a 4xx status code
func (o *SnapshotInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot inspect o k response has a 5xx status code
func (o *SnapshotInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot inspect o k response a status code equal to that given
func (o *SnapshotInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnapshotInspectOK) Error() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/snapshot][%d] snapshotInspectOK  %+v", 200, o.Payload)
}

func (o *SnapshotInspectOK) String() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/snapshot][%d] snapshotInspectOK  %+v", 200, o.Payload)
}

func (o *SnapshotInspectOK) GetPayload() models.PortainerDockerSnapshotRaw {
	return o.Payload
}

func (o *SnapshotInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotInspectNotFound creates a SnapshotInspectNotFound with default headers values
func NewSnapshotInspectNotFound() *SnapshotInspectNotFound {
	return &SnapshotInspectNotFound{}
}

/*
SnapshotInspectNotFound describes a response with status code 404, with default header values.

Environment not found
*/
type SnapshotInspectNotFound struct {
}

// IsSuccess returns true when this snapshot inspect not found response has a 2xx status code
func (o *SnapshotInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this snapshot inspect not found response has a 3xx status code
func (o *SnapshotInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot inspect not found response has a 4xx status code
func (o *SnapshotInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this snapshot inspect not found response has a 5xx status code
func (o *SnapshotInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot inspect not found response a status code equal to that given
func (o *SnapshotInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SnapshotInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/snapshot][%d] snapshotInspectNotFound ", 404)
}

func (o *SnapshotInspectNotFound) String() string {
	return fmt.Sprintf("[GET /docker/{environmentId}/snapshot][%d] snapshotInspectNotFound ", 404)
}

func (o *SnapshotInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
