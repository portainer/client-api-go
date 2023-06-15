// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// EdgeStackStatusUpdateReader is a Reader for the EdgeStackStatusUpdate structure.
type EdgeStackStatusUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackStatusUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackStatusUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackStatusUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeStackStatusUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeStackStatusUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackStatusUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackStatusUpdateOK creates a EdgeStackStatusUpdateOK with default headers values
func NewEdgeStackStatusUpdateOK() *EdgeStackStatusUpdateOK {
	return &EdgeStackStatusUpdateOK{}
}

/*
EdgeStackStatusUpdateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackStatusUpdateOK struct {
	Payload *models.PortainerEdgeStack
}

// IsSuccess returns true when this edge stack status update o k response has a 2xx status code
func (o *EdgeStackStatusUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack status update o k response has a 3xx status code
func (o *EdgeStackStatusUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status update o k response has a 4xx status code
func (o *EdgeStackStatusUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack status update o k response has a 5xx status code
func (o *EdgeStackStatusUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status update o k response a status code equal to that given
func (o *EdgeStackStatusUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackStatusUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStatusUpdateOK) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStatusUpdateOK) GetPayload() *models.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackStatusUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackStatusUpdateBadRequest creates a EdgeStackStatusUpdateBadRequest with default headers values
func NewEdgeStackStatusUpdateBadRequest() *EdgeStackStatusUpdateBadRequest {
	return &EdgeStackStatusUpdateBadRequest{}
}

/*
EdgeStackStatusUpdateBadRequest describes a response with status code 400, with default header values.

EdgeStackStatusUpdateBadRequest edge stack status update bad request
*/
type EdgeStackStatusUpdateBadRequest struct {
}

// IsSuccess returns true when this edge stack status update bad request response has a 2xx status code
func (o *EdgeStackStatusUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status update bad request response has a 3xx status code
func (o *EdgeStackStatusUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status update bad request response has a 4xx status code
func (o *EdgeStackStatusUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status update bad request response has a 5xx status code
func (o *EdgeStackStatusUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status update bad request response a status code equal to that given
func (o *EdgeStackStatusUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeStackStatusUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateBadRequest ", 400)
}

func (o *EdgeStackStatusUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateBadRequest ", 400)
}

func (o *EdgeStackStatusUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusUpdateForbidden creates a EdgeStackStatusUpdateForbidden with default headers values
func NewEdgeStackStatusUpdateForbidden() *EdgeStackStatusUpdateForbidden {
	return &EdgeStackStatusUpdateForbidden{}
}

/*
EdgeStackStatusUpdateForbidden describes a response with status code 403, with default header values.

EdgeStackStatusUpdateForbidden edge stack status update forbidden
*/
type EdgeStackStatusUpdateForbidden struct {
}

// IsSuccess returns true when this edge stack status update forbidden response has a 2xx status code
func (o *EdgeStackStatusUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status update forbidden response has a 3xx status code
func (o *EdgeStackStatusUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status update forbidden response has a 4xx status code
func (o *EdgeStackStatusUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status update forbidden response has a 5xx status code
func (o *EdgeStackStatusUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status update forbidden response a status code equal to that given
func (o *EdgeStackStatusUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeStackStatusUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateForbidden ", 403)
}

func (o *EdgeStackStatusUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateForbidden ", 403)
}

func (o *EdgeStackStatusUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusUpdateNotFound creates a EdgeStackStatusUpdateNotFound with default headers values
func NewEdgeStackStatusUpdateNotFound() *EdgeStackStatusUpdateNotFound {
	return &EdgeStackStatusUpdateNotFound{}
}

/*
EdgeStackStatusUpdateNotFound describes a response with status code 404, with default header values.

EdgeStackStatusUpdateNotFound edge stack status update not found
*/
type EdgeStackStatusUpdateNotFound struct {
}

// IsSuccess returns true when this edge stack status update not found response has a 2xx status code
func (o *EdgeStackStatusUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status update not found response has a 3xx status code
func (o *EdgeStackStatusUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status update not found response has a 4xx status code
func (o *EdgeStackStatusUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status update not found response has a 5xx status code
func (o *EdgeStackStatusUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status update not found response a status code equal to that given
func (o *EdgeStackStatusUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeStackStatusUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateNotFound ", 404)
}

func (o *EdgeStackStatusUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateNotFound ", 404)
}

func (o *EdgeStackStatusUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusUpdateInternalServerError creates a EdgeStackStatusUpdateInternalServerError with default headers values
func NewEdgeStackStatusUpdateInternalServerError() *EdgeStackStatusUpdateInternalServerError {
	return &EdgeStackStatusUpdateInternalServerError{}
}

/*
EdgeStackStatusUpdateInternalServerError describes a response with status code 500, with default header values.

EdgeStackStatusUpdateInternalServerError edge stack status update internal server error
*/
type EdgeStackStatusUpdateInternalServerError struct {
}

// IsSuccess returns true when this edge stack status update internal server error response has a 2xx status code
func (o *EdgeStackStatusUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status update internal server error response has a 3xx status code
func (o *EdgeStackStatusUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status update internal server error response has a 4xx status code
func (o *EdgeStackStatusUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack status update internal server error response has a 5xx status code
func (o *EdgeStackStatusUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack status update internal server error response a status code equal to that given
func (o *EdgeStackStatusUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackStatusUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateInternalServerError ", 500)
}

func (o *EdgeStackStatusUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /edge_stacks/{id}/status][%d] edgeStackStatusUpdateInternalServerError ", 500)
}

func (o *EdgeStackStatusUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
