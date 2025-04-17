// Code generated by go-swagger; DO NOT EDIT.

package edge_groups

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

// EdgeGroupUpdateReader is a Reader for the EdgeGroupUpdate structure.
type EdgeGroupUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGroupUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeGroupUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeGroupUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeGroupUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /edge_groups/{id}] EdgeGroupUpdate", response, response.Code())
	}
}

// NewEdgeGroupUpdateOK creates a EdgeGroupUpdateOK with default headers values
func NewEdgeGroupUpdateOK() *EdgeGroupUpdateOK {
	return &EdgeGroupUpdateOK{}
}

/*
EdgeGroupUpdateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeGroupUpdateOK struct {
	Payload *models.PortainereeEdgeGroup
}

// IsSuccess returns true when this edge group update o k response has a 2xx status code
func (o *EdgeGroupUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge group update o k response has a 3xx status code
func (o *EdgeGroupUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group update o k response has a 4xx status code
func (o *EdgeGroupUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group update o k response has a 5xx status code
func (o *EdgeGroupUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge group update o k response a status code equal to that given
func (o *EdgeGroupUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge group update o k response
func (o *EdgeGroupUpdateOK) Code() int {
	return 200
}

func (o *EdgeGroupUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateOK %s", 200, payload)
}

func (o *EdgeGroupUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateOK %s", 200, payload)
}

func (o *EdgeGroupUpdateOK) GetPayload() *models.PortainereeEdgeGroup {
	return o.Payload
}

func (o *EdgeGroupUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGroupUpdateInternalServerError creates a EdgeGroupUpdateInternalServerError with default headers values
func NewEdgeGroupUpdateInternalServerError() *EdgeGroupUpdateInternalServerError {
	return &EdgeGroupUpdateInternalServerError{}
}

/*
EdgeGroupUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeGroupUpdateInternalServerError struct {
}

// IsSuccess returns true when this edge group update internal server error response has a 2xx status code
func (o *EdgeGroupUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group update internal server error response has a 3xx status code
func (o *EdgeGroupUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group update internal server error response has a 4xx status code
func (o *EdgeGroupUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group update internal server error response has a 5xx status code
func (o *EdgeGroupUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group update internal server error response a status code equal to that given
func (o *EdgeGroupUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge group update internal server error response
func (o *EdgeGroupUpdateInternalServerError) Code() int {
	return 500
}

func (o *EdgeGroupUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateInternalServerError", 500)
}

func (o *EdgeGroupUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateInternalServerError", 500)
}

func (o *EdgeGroupUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGroupUpdateServiceUnavailable creates a EdgeGroupUpdateServiceUnavailable with default headers values
func NewEdgeGroupUpdateServiceUnavailable() *EdgeGroupUpdateServiceUnavailable {
	return &EdgeGroupUpdateServiceUnavailable{}
}

/*
EdgeGroupUpdateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeGroupUpdateServiceUnavailable struct {
}

// IsSuccess returns true when this edge group update service unavailable response has a 2xx status code
func (o *EdgeGroupUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group update service unavailable response has a 3xx status code
func (o *EdgeGroupUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group update service unavailable response has a 4xx status code
func (o *EdgeGroupUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group update service unavailable response has a 5xx status code
func (o *EdgeGroupUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group update service unavailable response a status code equal to that given
func (o *EdgeGroupUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge group update service unavailable response
func (o *EdgeGroupUpdateServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeGroupUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateServiceUnavailable", 503)
}

func (o *EdgeGroupUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /edge_groups/{id}][%d] edgeGroupUpdateServiceUnavailable", 503)
}

func (o *EdgeGroupUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
