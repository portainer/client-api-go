// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

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

// EdgeJobListReader is a Reader for the EdgeJobList structure.
type EdgeJobListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /edge_jobs] EdgeJobList", response, response.Code())
	}
}

// NewEdgeJobListOK creates a EdgeJobListOK with default headers values
func NewEdgeJobListOK() *EdgeJobListOK {
	return &EdgeJobListOK{}
}

/*
EdgeJobListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobListOK struct {
	Payload []*models.PortainerEdgeJob
}

// IsSuccess returns true when this edge job list o k response has a 2xx status code
func (o *EdgeJobListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job list o k response has a 3xx status code
func (o *EdgeJobListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job list o k response has a 4xx status code
func (o *EdgeJobListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job list o k response has a 5xx status code
func (o *EdgeJobListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job list o k response a status code equal to that given
func (o *EdgeJobListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge job list o k response
func (o *EdgeJobListOK) Code() int {
	return 200
}

func (o *EdgeJobListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListOK %s", 200, payload)
}

func (o *EdgeJobListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListOK %s", 200, payload)
}

func (o *EdgeJobListOK) GetPayload() []*models.PortainerEdgeJob {
	return o.Payload
}

func (o *EdgeJobListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobListBadRequest creates a EdgeJobListBadRequest with default headers values
func NewEdgeJobListBadRequest() *EdgeJobListBadRequest {
	return &EdgeJobListBadRequest{}
}

/*
EdgeJobListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeJobListBadRequest struct {
}

// IsSuccess returns true when this edge job list bad request response has a 2xx status code
func (o *EdgeJobListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job list bad request response has a 3xx status code
func (o *EdgeJobListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job list bad request response has a 4xx status code
func (o *EdgeJobListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge job list bad request response has a 5xx status code
func (o *EdgeJobListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job list bad request response a status code equal to that given
func (o *EdgeJobListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge job list bad request response
func (o *EdgeJobListBadRequest) Code() int {
	return 400
}

func (o *EdgeJobListBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListBadRequest", 400)
}

func (o *EdgeJobListBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListBadRequest", 400)
}

func (o *EdgeJobListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobListInternalServerError creates a EdgeJobListInternalServerError with default headers values
func NewEdgeJobListInternalServerError() *EdgeJobListInternalServerError {
	return &EdgeJobListInternalServerError{}
}

/*
EdgeJobListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeJobListInternalServerError struct {
}

// IsSuccess returns true when this edge job list internal server error response has a 2xx status code
func (o *EdgeJobListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job list internal server error response has a 3xx status code
func (o *EdgeJobListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job list internal server error response has a 4xx status code
func (o *EdgeJobListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job list internal server error response has a 5xx status code
func (o *EdgeJobListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job list internal server error response a status code equal to that given
func (o *EdgeJobListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge job list internal server error response
func (o *EdgeJobListInternalServerError) Code() int {
	return 500
}

func (o *EdgeJobListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListInternalServerError", 500)
}

func (o *EdgeJobListInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListInternalServerError", 500)
}

func (o *EdgeJobListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobListServiceUnavailable creates a EdgeJobListServiceUnavailable with default headers values
func NewEdgeJobListServiceUnavailable() *EdgeJobListServiceUnavailable {
	return &EdgeJobListServiceUnavailable{}
}

/*
EdgeJobListServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobListServiceUnavailable struct {
}

// IsSuccess returns true when this edge job list service unavailable response has a 2xx status code
func (o *EdgeJobListServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job list service unavailable response has a 3xx status code
func (o *EdgeJobListServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job list service unavailable response has a 4xx status code
func (o *EdgeJobListServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job list service unavailable response has a 5xx status code
func (o *EdgeJobListServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job list service unavailable response a status code equal to that given
func (o *EdgeJobListServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge job list service unavailable response
func (o *EdgeJobListServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeJobListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListServiceUnavailable", 503)
}

func (o *EdgeJobListServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListServiceUnavailable", 503)
}

func (o *EdgeJobListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
