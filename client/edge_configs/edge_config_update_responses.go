// Code generated by go-swagger; DO NOT EDIT.

package edge_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EdgeConfigUpdateReader is a Reader for the EdgeConfigUpdate structure.
type EdgeConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEdgeConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeConfigUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeConfigUpdateNoContent creates a EdgeConfigUpdateNoContent with default headers values
func NewEdgeConfigUpdateNoContent() *EdgeConfigUpdateNoContent {
	return &EdgeConfigUpdateNoContent{}
}

/*
EdgeConfigUpdateNoContent describes a response with status code 204, with default header values.

No Content
*/
type EdgeConfigUpdateNoContent struct {
}

// IsSuccess returns true when this edge config update no content response has a 2xx status code
func (o *EdgeConfigUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge config update no content response has a 3xx status code
func (o *EdgeConfigUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge config update no content response has a 4xx status code
func (o *EdgeConfigUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge config update no content response has a 5xx status code
func (o *EdgeConfigUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this edge config update no content response a status code equal to that given
func (o *EdgeConfigUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *EdgeConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /edge_configurations][%d] edgeConfigUpdateNoContent ", 204)
}

func (o *EdgeConfigUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /edge_configurations][%d] edgeConfigUpdateNoContent ", 204)
}

func (o *EdgeConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeConfigUpdateBadRequest creates a EdgeConfigUpdateBadRequest with default headers values
func NewEdgeConfigUpdateBadRequest() *EdgeConfigUpdateBadRequest {
	return &EdgeConfigUpdateBadRequest{}
}

/*
EdgeConfigUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EdgeConfigUpdateBadRequest struct {
}

// IsSuccess returns true when this edge config update bad request response has a 2xx status code
func (o *EdgeConfigUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge config update bad request response has a 3xx status code
func (o *EdgeConfigUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge config update bad request response has a 4xx status code
func (o *EdgeConfigUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge config update bad request response has a 5xx status code
func (o *EdgeConfigUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge config update bad request response a status code equal to that given
func (o *EdgeConfigUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeConfigUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /edge_configurations][%d] edgeConfigUpdateBadRequest ", 400)
}

func (o *EdgeConfigUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /edge_configurations][%d] edgeConfigUpdateBadRequest ", 400)
}

func (o *EdgeConfigUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
