// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// PostEdgeJobsReader is a Reader for the PostEdgeJobs structure.
type PostEdgeJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEdgeJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEdgeJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostEdgeJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostEdgeJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostEdgeJobsOK creates a PostEdgeJobsOK with default headers values
func NewPostEdgeJobsOK() *PostEdgeJobsOK {
	return &PostEdgeJobsOK{}
}

/* PostEdgeJobsOK describes a response with status code 200, with default header values.

OK
*/
type PostEdgeJobsOK struct {
	Payload *models.PortainerEdgeGroup
}

func (o *PostEdgeJobsOK) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] postEdgeJobsOK  %+v", 200, o.Payload)
}
func (o *PostEdgeJobsOK) GetPayload() *models.PortainerEdgeGroup {
	return o.Payload
}

func (o *PostEdgeJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEdgeJobsInternalServerError creates a PostEdgeJobsInternalServerError with default headers values
func NewPostEdgeJobsInternalServerError() *PostEdgeJobsInternalServerError {
	return &PostEdgeJobsInternalServerError{}
}

/* PostEdgeJobsInternalServerError describes a response with status code 500, with default header values.

PostEdgeJobsInternalServerError post edge jobs internal server error
*/
type PostEdgeJobsInternalServerError struct {
}

func (o *PostEdgeJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] postEdgeJobsInternalServerError ", 500)
}

func (o *PostEdgeJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEdgeJobsServiceUnavailable creates a PostEdgeJobsServiceUnavailable with default headers values
func NewPostEdgeJobsServiceUnavailable() *PostEdgeJobsServiceUnavailable {
	return &PostEdgeJobsServiceUnavailable{}
}

/* PostEdgeJobsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PostEdgeJobsServiceUnavailable struct {
}

func (o *PostEdgeJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] postEdgeJobsServiceUnavailable ", 503)
}

func (o *PostEdgeJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
