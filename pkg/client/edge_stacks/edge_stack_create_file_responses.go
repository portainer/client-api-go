// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

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

// EdgeStackCreateFileReader is a Reader for the EdgeStackCreateFile structure.
type EdgeStackCreateFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackCreateFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackCreateFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackCreateFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackCreateFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackCreateFileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /edge_stacks/create/file] EdgeStackCreateFile", response, response.Code())
	}
}

// NewEdgeStackCreateFileOK creates a EdgeStackCreateFileOK with default headers values
func NewEdgeStackCreateFileOK() *EdgeStackCreateFileOK {
	return &EdgeStackCreateFileOK{}
}

/*
EdgeStackCreateFileOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackCreateFileOK struct {
	Payload *models.PortainereeEdgeStack
}

// IsSuccess returns true when this edge stack create file o k response has a 2xx status code
func (o *EdgeStackCreateFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack create file o k response has a 3xx status code
func (o *EdgeStackCreateFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create file o k response has a 4xx status code
func (o *EdgeStackCreateFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create file o k response has a 5xx status code
func (o *EdgeStackCreateFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create file o k response a status code equal to that given
func (o *EdgeStackCreateFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge stack create file o k response
func (o *EdgeStackCreateFileOK) Code() int {
	return 200
}

func (o *EdgeStackCreateFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileOK %s", 200, payload)
}

func (o *EdgeStackCreateFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileOK %s", 200, payload)
}

func (o *EdgeStackCreateFileOK) GetPayload() *models.PortainereeEdgeStack {
	return o.Payload
}

func (o *EdgeStackCreateFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackCreateFileBadRequest creates a EdgeStackCreateFileBadRequest with default headers values
func NewEdgeStackCreateFileBadRequest() *EdgeStackCreateFileBadRequest {
	return &EdgeStackCreateFileBadRequest{}
}

/*
EdgeStackCreateFileBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type EdgeStackCreateFileBadRequest struct {
}

// IsSuccess returns true when this edge stack create file bad request response has a 2xx status code
func (o *EdgeStackCreateFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create file bad request response has a 3xx status code
func (o *EdgeStackCreateFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create file bad request response has a 4xx status code
func (o *EdgeStackCreateFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack create file bad request response has a 5xx status code
func (o *EdgeStackCreateFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create file bad request response a status code equal to that given
func (o *EdgeStackCreateFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge stack create file bad request response
func (o *EdgeStackCreateFileBadRequest) Code() int {
	return 400
}

func (o *EdgeStackCreateFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileBadRequest", 400)
}

func (o *EdgeStackCreateFileBadRequest) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileBadRequest", 400)
}

func (o *EdgeStackCreateFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateFileInternalServerError creates a EdgeStackCreateFileInternalServerError with default headers values
func NewEdgeStackCreateFileInternalServerError() *EdgeStackCreateFileInternalServerError {
	return &EdgeStackCreateFileInternalServerError{}
}

/*
EdgeStackCreateFileInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type EdgeStackCreateFileInternalServerError struct {
}

// IsSuccess returns true when this edge stack create file internal server error response has a 2xx status code
func (o *EdgeStackCreateFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create file internal server error response has a 3xx status code
func (o *EdgeStackCreateFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create file internal server error response has a 4xx status code
func (o *EdgeStackCreateFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create file internal server error response has a 5xx status code
func (o *EdgeStackCreateFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create file internal server error response a status code equal to that given
func (o *EdgeStackCreateFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge stack create file internal server error response
func (o *EdgeStackCreateFileInternalServerError) Code() int {
	return 500
}

func (o *EdgeStackCreateFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileInternalServerError", 500)
}

func (o *EdgeStackCreateFileInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileInternalServerError", 500)
}

func (o *EdgeStackCreateFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateFileServiceUnavailable creates a EdgeStackCreateFileServiceUnavailable with default headers values
func NewEdgeStackCreateFileServiceUnavailable() *EdgeStackCreateFileServiceUnavailable {
	return &EdgeStackCreateFileServiceUnavailable{}
}

/*
EdgeStackCreateFileServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackCreateFileServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack create file service unavailable response has a 2xx status code
func (o *EdgeStackCreateFileServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create file service unavailable response has a 3xx status code
func (o *EdgeStackCreateFileServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create file service unavailable response has a 4xx status code
func (o *EdgeStackCreateFileServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create file service unavailable response has a 5xx status code
func (o *EdgeStackCreateFileServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create file service unavailable response a status code equal to that given
func (o *EdgeStackCreateFileServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge stack create file service unavailable response
func (o *EdgeStackCreateFileServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeStackCreateFileServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileServiceUnavailable", 503)
}

func (o *EdgeStackCreateFileServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/file][%d] edgeStackCreateFileServiceUnavailable", 503)
}

func (o *EdgeStackCreateFileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
