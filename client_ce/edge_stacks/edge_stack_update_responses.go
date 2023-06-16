// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// EdgeStackUpdateReader is a Reader for the EdgeStackUpdate structure.
type EdgeStackUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackUpdateOK creates a EdgeStackUpdateOK with default headers values
func NewEdgeStackUpdateOK() *EdgeStackUpdateOK {
	return &EdgeStackUpdateOK{}
}

/*
EdgeStackUpdateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackUpdateOK struct {
	Payload *models_ce.PortainerEdgeStack
}

// IsSuccess returns true when this edge stack update o k response has a 2xx status code
func (o *EdgeStackUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack update o k response has a 3xx status code
func (o *EdgeStackUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack update o k response has a 4xx status code
func (o *EdgeStackUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack update o k response has a 5xx status code
func (o *EdgeStackUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack update o k response a status code equal to that given
func (o *EdgeStackUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackUpdateOK) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackUpdateOK) GetPayload() *models_ce.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackUpdateBadRequest creates a EdgeStackUpdateBadRequest with default headers values
func NewEdgeStackUpdateBadRequest() *EdgeStackUpdateBadRequest {
	return &EdgeStackUpdateBadRequest{}
}

/*
EdgeStackUpdateBadRequest describes a response with status code 400, with default header values.

EdgeStackUpdateBadRequest edge stack update bad request
*/
type EdgeStackUpdateBadRequest struct {
}

// IsSuccess returns true when this edge stack update bad request response has a 2xx status code
func (o *EdgeStackUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack update bad request response has a 3xx status code
func (o *EdgeStackUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack update bad request response has a 4xx status code
func (o *EdgeStackUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack update bad request response has a 5xx status code
func (o *EdgeStackUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack update bad request response a status code equal to that given
func (o *EdgeStackUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeStackUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateBadRequest ", 400)
}

func (o *EdgeStackUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateBadRequest ", 400)
}

func (o *EdgeStackUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackUpdateInternalServerError creates a EdgeStackUpdateInternalServerError with default headers values
func NewEdgeStackUpdateInternalServerError() *EdgeStackUpdateInternalServerError {
	return &EdgeStackUpdateInternalServerError{}
}

/*
EdgeStackUpdateInternalServerError describes a response with status code 500, with default header values.

EdgeStackUpdateInternalServerError edge stack update internal server error
*/
type EdgeStackUpdateInternalServerError struct {
}

// IsSuccess returns true when this edge stack update internal server error response has a 2xx status code
func (o *EdgeStackUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack update internal server error response has a 3xx status code
func (o *EdgeStackUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack update internal server error response has a 4xx status code
func (o *EdgeStackUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack update internal server error response has a 5xx status code
func (o *EdgeStackUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack update internal server error response a status code equal to that given
func (o *EdgeStackUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateInternalServerError ", 500)
}

func (o *EdgeStackUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateInternalServerError ", 500)
}

func (o *EdgeStackUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackUpdateServiceUnavailable creates a EdgeStackUpdateServiceUnavailable with default headers values
func NewEdgeStackUpdateServiceUnavailable() *EdgeStackUpdateServiceUnavailable {
	return &EdgeStackUpdateServiceUnavailable{}
}

/*
EdgeStackUpdateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackUpdateServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack update service unavailable response has a 2xx status code
func (o *EdgeStackUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack update service unavailable response has a 3xx status code
func (o *EdgeStackUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack update service unavailable response has a 4xx status code
func (o *EdgeStackUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack update service unavailable response has a 5xx status code
func (o *EdgeStackUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack update service unavailable response a status code equal to that given
func (o *EdgeStackUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeStackUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateServiceUnavailable ", 503)
}

func (o *EdgeStackUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}][%d] edgeStackUpdateServiceUnavailable ", 503)
}

func (o *EdgeStackUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
