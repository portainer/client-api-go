// Code generated by go-swagger; DO NOT EDIT.

package edge_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EdgeConfigCreateReader is a Reader for the EdgeConfigCreate structure.
type EdgeConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEdgeConfigCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeConfigCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /edge_configurations] EdgeConfigCreate", response, response.Code())
	}
}

// NewEdgeConfigCreateNoContent creates a EdgeConfigCreateNoContent with default headers values
func NewEdgeConfigCreateNoContent() *EdgeConfigCreateNoContent {
	return &EdgeConfigCreateNoContent{}
}

/*
EdgeConfigCreateNoContent describes a response with status code 204, with default header values.

No Content
*/
type EdgeConfigCreateNoContent struct {
}

// IsSuccess returns true when this edge config create no content response has a 2xx status code
func (o *EdgeConfigCreateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge config create no content response has a 3xx status code
func (o *EdgeConfigCreateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge config create no content response has a 4xx status code
func (o *EdgeConfigCreateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge config create no content response has a 5xx status code
func (o *EdgeConfigCreateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this edge config create no content response a status code equal to that given
func (o *EdgeConfigCreateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the edge config create no content response
func (o *EdgeConfigCreateNoContent) Code() int {
	return 204
}

func (o *EdgeConfigCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /edge_configurations][%d] edgeConfigCreateNoContent ", 204)
}

func (o *EdgeConfigCreateNoContent) String() string {
	return fmt.Sprintf("[POST /edge_configurations][%d] edgeConfigCreateNoContent ", 204)
}

func (o *EdgeConfigCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeConfigCreateBadRequest creates a EdgeConfigCreateBadRequest with default headers values
func NewEdgeConfigCreateBadRequest() *EdgeConfigCreateBadRequest {
	return &EdgeConfigCreateBadRequest{}
}

/*
EdgeConfigCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EdgeConfigCreateBadRequest struct {
}

// IsSuccess returns true when this edge config create bad request response has a 2xx status code
func (o *EdgeConfigCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge config create bad request response has a 3xx status code
func (o *EdgeConfigCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge config create bad request response has a 4xx status code
func (o *EdgeConfigCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge config create bad request response has a 5xx status code
func (o *EdgeConfigCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge config create bad request response a status code equal to that given
func (o *EdgeConfigCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge config create bad request response
func (o *EdgeConfigCreateBadRequest) Code() int {
	return 400
}

func (o *EdgeConfigCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /edge_configurations][%d] edgeConfigCreateBadRequest ", 400)
}

func (o *EdgeConfigCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /edge_configurations][%d] edgeConfigCreateBadRequest ", 400)
}

func (o *EdgeConfigCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
