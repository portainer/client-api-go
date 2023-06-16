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

// EdgeStackListReader is a Reader for the EdgeStackList structure.
type EdgeStackListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackListOK creates a EdgeStackListOK with default headers values
func NewEdgeStackListOK() *EdgeStackListOK {
	return &EdgeStackListOK{}
}

/*
EdgeStackListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackListOK struct {
	Payload []*models_ce.PortainerEdgeStack
}

// IsSuccess returns true when this edge stack list o k response has a 2xx status code
func (o *EdgeStackListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack list o k response has a 3xx status code
func (o *EdgeStackListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack list o k response has a 4xx status code
func (o *EdgeStackListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack list o k response has a 5xx status code
func (o *EdgeStackListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack list o k response a status code equal to that given
func (o *EdgeStackListOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackListOK) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListOK  %+v", 200, o.Payload)
}

func (o *EdgeStackListOK) String() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListOK  %+v", 200, o.Payload)
}

func (o *EdgeStackListOK) GetPayload() []*models_ce.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackListBadRequest creates a EdgeStackListBadRequest with default headers values
func NewEdgeStackListBadRequest() *EdgeStackListBadRequest {
	return &EdgeStackListBadRequest{}
}

/*
EdgeStackListBadRequest describes a response with status code 400, with default header values.

EdgeStackListBadRequest edge stack list bad request
*/
type EdgeStackListBadRequest struct {
}

// IsSuccess returns true when this edge stack list bad request response has a 2xx status code
func (o *EdgeStackListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack list bad request response has a 3xx status code
func (o *EdgeStackListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack list bad request response has a 4xx status code
func (o *EdgeStackListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack list bad request response has a 5xx status code
func (o *EdgeStackListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack list bad request response a status code equal to that given
func (o *EdgeStackListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeStackListBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListBadRequest ", 400)
}

func (o *EdgeStackListBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListBadRequest ", 400)
}

func (o *EdgeStackListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackListInternalServerError creates a EdgeStackListInternalServerError with default headers values
func NewEdgeStackListInternalServerError() *EdgeStackListInternalServerError {
	return &EdgeStackListInternalServerError{}
}

/*
EdgeStackListInternalServerError describes a response with status code 500, with default header values.

EdgeStackListInternalServerError edge stack list internal server error
*/
type EdgeStackListInternalServerError struct {
}

// IsSuccess returns true when this edge stack list internal server error response has a 2xx status code
func (o *EdgeStackListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack list internal server error response has a 3xx status code
func (o *EdgeStackListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack list internal server error response has a 4xx status code
func (o *EdgeStackListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack list internal server error response has a 5xx status code
func (o *EdgeStackListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack list internal server error response a status code equal to that given
func (o *EdgeStackListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListInternalServerError ", 500)
}

func (o *EdgeStackListInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListInternalServerError ", 500)
}

func (o *EdgeStackListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackListServiceUnavailable creates a EdgeStackListServiceUnavailable with default headers values
func NewEdgeStackListServiceUnavailable() *EdgeStackListServiceUnavailable {
	return &EdgeStackListServiceUnavailable{}
}

/*
EdgeStackListServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackListServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack list service unavailable response has a 2xx status code
func (o *EdgeStackListServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack list service unavailable response has a 3xx status code
func (o *EdgeStackListServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack list service unavailable response has a 4xx status code
func (o *EdgeStackListServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack list service unavailable response has a 5xx status code
func (o *EdgeStackListServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack list service unavailable response a status code equal to that given
func (o *EdgeStackListServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeStackListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListServiceUnavailable ", 503)
}

func (o *EdgeStackListServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListServiceUnavailable ", 503)
}

func (o *EdgeStackListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
