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

// EdgeStackStatusDeleteReader is a Reader for the EdgeStackStatusDelete structure.
type EdgeStackStatusDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackStatusDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackStatusDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackStatusDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeStackStatusDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeStackStatusDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackStatusDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackStatusDeleteOK creates a EdgeStackStatusDeleteOK with default headers values
func NewEdgeStackStatusDeleteOK() *EdgeStackStatusDeleteOK {
	return &EdgeStackStatusDeleteOK{}
}

/*
EdgeStackStatusDeleteOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackStatusDeleteOK struct {
	Payload *models.PortainerEdgeStack
}

// IsSuccess returns true when this edge stack status delete o k response has a 2xx status code
func (o *EdgeStackStatusDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack status delete o k response has a 3xx status code
func (o *EdgeStackStatusDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status delete o k response has a 4xx status code
func (o *EdgeStackStatusDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack status delete o k response has a 5xx status code
func (o *EdgeStackStatusDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status delete o k response a status code equal to that given
func (o *EdgeStackStatusDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackStatusDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStatusDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteOK  %+v", 200, o.Payload)
}

func (o *EdgeStackStatusDeleteOK) GetPayload() *models.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackStatusDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackStatusDeleteBadRequest creates a EdgeStackStatusDeleteBadRequest with default headers values
func NewEdgeStackStatusDeleteBadRequest() *EdgeStackStatusDeleteBadRequest {
	return &EdgeStackStatusDeleteBadRequest{}
}

/*
EdgeStackStatusDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeStackStatusDeleteBadRequest struct {
}

// IsSuccess returns true when this edge stack status delete bad request response has a 2xx status code
func (o *EdgeStackStatusDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status delete bad request response has a 3xx status code
func (o *EdgeStackStatusDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status delete bad request response has a 4xx status code
func (o *EdgeStackStatusDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status delete bad request response has a 5xx status code
func (o *EdgeStackStatusDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status delete bad request response a status code equal to that given
func (o *EdgeStackStatusDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeStackStatusDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteBadRequest ", 400)
}

func (o *EdgeStackStatusDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteBadRequest ", 400)
}

func (o *EdgeStackStatusDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusDeleteForbidden creates a EdgeStackStatusDeleteForbidden with default headers values
func NewEdgeStackStatusDeleteForbidden() *EdgeStackStatusDeleteForbidden {
	return &EdgeStackStatusDeleteForbidden{}
}

/*
EdgeStackStatusDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EdgeStackStatusDeleteForbidden struct {
}

// IsSuccess returns true when this edge stack status delete forbidden response has a 2xx status code
func (o *EdgeStackStatusDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status delete forbidden response has a 3xx status code
func (o *EdgeStackStatusDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status delete forbidden response has a 4xx status code
func (o *EdgeStackStatusDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status delete forbidden response has a 5xx status code
func (o *EdgeStackStatusDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status delete forbidden response a status code equal to that given
func (o *EdgeStackStatusDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeStackStatusDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteForbidden ", 403)
}

func (o *EdgeStackStatusDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteForbidden ", 403)
}

func (o *EdgeStackStatusDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusDeleteNotFound creates a EdgeStackStatusDeleteNotFound with default headers values
func NewEdgeStackStatusDeleteNotFound() *EdgeStackStatusDeleteNotFound {
	return &EdgeStackStatusDeleteNotFound{}
}

/*
EdgeStackStatusDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type EdgeStackStatusDeleteNotFound struct {
}

// IsSuccess returns true when this edge stack status delete not found response has a 2xx status code
func (o *EdgeStackStatusDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status delete not found response has a 3xx status code
func (o *EdgeStackStatusDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status delete not found response has a 4xx status code
func (o *EdgeStackStatusDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack status delete not found response has a 5xx status code
func (o *EdgeStackStatusDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack status delete not found response a status code equal to that given
func (o *EdgeStackStatusDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeStackStatusDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteNotFound ", 404)
}

func (o *EdgeStackStatusDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteNotFound ", 404)
}

func (o *EdgeStackStatusDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackStatusDeleteInternalServerError creates a EdgeStackStatusDeleteInternalServerError with default headers values
func NewEdgeStackStatusDeleteInternalServerError() *EdgeStackStatusDeleteInternalServerError {
	return &EdgeStackStatusDeleteInternalServerError{}
}

/*
EdgeStackStatusDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeStackStatusDeleteInternalServerError struct {
}

// IsSuccess returns true when this edge stack status delete internal server error response has a 2xx status code
func (o *EdgeStackStatusDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack status delete internal server error response has a 3xx status code
func (o *EdgeStackStatusDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack status delete internal server error response has a 4xx status code
func (o *EdgeStackStatusDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack status delete internal server error response has a 5xx status code
func (o *EdgeStackStatusDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack status delete internal server error response a status code equal to that given
func (o *EdgeStackStatusDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackStatusDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteInternalServerError ", 500)
}

func (o *EdgeStackStatusDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /edge_stacks/{id}/status/{endpoint_id}][%d] edgeStackStatusDeleteInternalServerError ", 500)
}

func (o *EdgeStackStatusDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
