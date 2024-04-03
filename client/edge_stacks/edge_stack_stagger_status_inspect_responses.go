// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// EdgeStackStaggerStatusInspectReader is a Reader for the EdgeStackStaggerStatusInspect structure.
type EdgeStackStaggerStatusInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackStaggerStatusInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackStaggerStatusInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackStaggerStatusInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackStaggerStatusInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackStaggerStatusInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /edge_stacks/{id}/stagger/status] EdgeStackStaggerStatusInspect", response, response.Code())
	}
}

// NewEdgeStackStaggerStatusInspectOK creates a EdgeStackStaggerStatusInspectOK with default headers values
func NewEdgeStackStaggerStatusInspectOK() *EdgeStackStaggerStatusInspectOK {
	return &EdgeStackStaggerStatusInspectOK{}
}

/*
EdgeStackStaggerStatusInspectOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackStaggerStatusInspectOK struct {
	Payload *models.EdgestacksEdgeStackStaggerStatusResponse
}

// IsSuccess returns true when this edge stack stagger status inspect o k response has a 2xx status code
func (o *EdgeStackStaggerStatusInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack stagger status inspect o k response has a 3xx status code
func (o *EdgeStackStaggerStatusInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack stagger status inspect o k response has a 4xx status code
func (o *EdgeStackStaggerStatusInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack stagger status inspect o k response has a 5xx status code
func (o *EdgeStackStaggerStatusInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack stagger status inspect o k response a status code equal to that given
func (o *EdgeStackStaggerStatusInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge stack stagger status inspect o k response
func (o *EdgeStackStaggerStatusInspectOK) Code() int {
	return 200
}

func (o *EdgeStackStaggerStatusInspectOK) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStaggerStatusInspectOK) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStaggerStatusInspectOK) GetPayload() *models.EdgestacksEdgeStackStaggerStatusResponse {
	return o.Payload
}

func (o *EdgeStackStaggerStatusInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgestacksEdgeStackStaggerStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackStaggerStatusInspectBadRequest creates a EdgeStackStaggerStatusInspectBadRequest with default headers values
func NewEdgeStackStaggerStatusInspectBadRequest() *EdgeStackStaggerStatusInspectBadRequest {
	return &EdgeStackStaggerStatusInspectBadRequest{}
}

/*
EdgeStackStaggerStatusInspectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeStackStaggerStatusInspectBadRequest struct {
}

// IsSuccess returns true when this edge stack stagger status inspect bad request response has a 2xx status code
func (o *EdgeStackStaggerStatusInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack stagger status inspect bad request response has a 3xx status code
func (o *EdgeStackStaggerStatusInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack stagger status inspect bad request response has a 4xx status code
func (o *EdgeStackStaggerStatusInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack stagger status inspect bad request response has a 5xx status code
func (o *EdgeStackStaggerStatusInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack stagger status inspect bad request response a status code equal to that given
func (o *EdgeStackStaggerStatusInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge stack stagger status inspect bad request response
func (o *EdgeStackStaggerStatusInspectBadRequest) Code() int {
	return 400
}

func (o *EdgeStackStaggerStatusInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectBadRequest ", 400)
}

func (o *EdgeStackStaggerStatusInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectBadRequest ", 400)
}

func (o *EdgeStackStaggerStatusInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStaggerStatusInspectInternalServerError creates a EdgeStackStaggerStatusInspectInternalServerError with default headers values
func NewEdgeStackStaggerStatusInspectInternalServerError() *EdgeStackStaggerStatusInspectInternalServerError {
	return &EdgeStackStaggerStatusInspectInternalServerError{}
}

/*
EdgeStackStaggerStatusInspectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeStackStaggerStatusInspectInternalServerError struct {
}

// IsSuccess returns true when this edge stack stagger status inspect internal server error response has a 2xx status code
func (o *EdgeStackStaggerStatusInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack stagger status inspect internal server error response has a 3xx status code
func (o *EdgeStackStaggerStatusInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack stagger status inspect internal server error response has a 4xx status code
func (o *EdgeStackStaggerStatusInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack stagger status inspect internal server error response has a 5xx status code
func (o *EdgeStackStaggerStatusInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack stagger status inspect internal server error response a status code equal to that given
func (o *EdgeStackStaggerStatusInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge stack stagger status inspect internal server error response
func (o *EdgeStackStaggerStatusInspectInternalServerError) Code() int {
	return 500
}

func (o *EdgeStackStaggerStatusInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectInternalServerError ", 500)
}

func (o *EdgeStackStaggerStatusInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectInternalServerError ", 500)
}

func (o *EdgeStackStaggerStatusInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStaggerStatusInspectServiceUnavailable creates a EdgeStackStaggerStatusInspectServiceUnavailable with default headers values
func NewEdgeStackStaggerStatusInspectServiceUnavailable() *EdgeStackStaggerStatusInspectServiceUnavailable {
	return &EdgeStackStaggerStatusInspectServiceUnavailable{}
}

/*
EdgeStackStaggerStatusInspectServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackStaggerStatusInspectServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack stagger status inspect service unavailable response has a 2xx status code
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack stagger status inspect service unavailable response has a 3xx status code
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack stagger status inspect service unavailable response has a 4xx status code
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack stagger status inspect service unavailable response has a 5xx status code
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack stagger status inspect service unavailable response a status code equal to that given
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge stack stagger status inspect service unavailable response
func (o *EdgeStackStaggerStatusInspectServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeStackStaggerStatusInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectServiceUnavailable ", 503)
}

func (o *EdgeStackStaggerStatusInspectServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/stagger/status][%d] edgeStackStaggerStatusInspectServiceUnavailable ", 503)
}

func (o *EdgeStackStaggerStatusInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
