// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateKubernetesNamespaceReader is a Reader for the CreateKubernetesNamespace structure.
type CreateKubernetesNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateKubernetesNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateKubernetesNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateKubernetesNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateKubernetesNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kubernetes/{id}/namespaces/{namespace}] createKubernetesNamespace", response, response.Code())
	}
}

// NewCreateKubernetesNamespaceOK creates a CreateKubernetesNamespaceOK with default headers values
func NewCreateKubernetesNamespaceOK() *CreateKubernetesNamespaceOK {
	return &CreateKubernetesNamespaceOK{}
}

/*
CreateKubernetesNamespaceOK describes a response with status code 200, with default header values.

Success
*/
type CreateKubernetesNamespaceOK struct {
	Payload string
}

// IsSuccess returns true when this create kubernetes namespace o k response has a 2xx status code
func (o *CreateKubernetesNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create kubernetes namespace o k response has a 3xx status code
func (o *CreateKubernetesNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes namespace o k response has a 4xx status code
func (o *CreateKubernetesNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create kubernetes namespace o k response has a 5xx status code
func (o *CreateKubernetesNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create kubernetes namespace o k response a status code equal to that given
func (o *CreateKubernetesNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create kubernetes namespace o k response
func (o *CreateKubernetesNamespaceOK) Code() int {
	return 200
}

func (o *CreateKubernetesNamespaceOK) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceOK  %+v", 200, o.Payload)
}

func (o *CreateKubernetesNamespaceOK) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceOK  %+v", 200, o.Payload)
}

func (o *CreateKubernetesNamespaceOK) GetPayload() string {
	return o.Payload
}

func (o *CreateKubernetesNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKubernetesNamespaceBadRequest creates a CreateKubernetesNamespaceBadRequest with default headers values
func NewCreateKubernetesNamespaceBadRequest() *CreateKubernetesNamespaceBadRequest {
	return &CreateKubernetesNamespaceBadRequest{}
}

/*
CreateKubernetesNamespaceBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CreateKubernetesNamespaceBadRequest struct {
}

// IsSuccess returns true when this create kubernetes namespace bad request response has a 2xx status code
func (o *CreateKubernetesNamespaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kubernetes namespace bad request response has a 3xx status code
func (o *CreateKubernetesNamespaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes namespace bad request response has a 4xx status code
func (o *CreateKubernetesNamespaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create kubernetes namespace bad request response has a 5xx status code
func (o *CreateKubernetesNamespaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create kubernetes namespace bad request response a status code equal to that given
func (o *CreateKubernetesNamespaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create kubernetes namespace bad request response
func (o *CreateKubernetesNamespaceBadRequest) Code() int {
	return 400
}

func (o *CreateKubernetesNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceBadRequest ", 400)
}

func (o *CreateKubernetesNamespaceBadRequest) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceBadRequest ", 400)
}

func (o *CreateKubernetesNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateKubernetesNamespaceInternalServerError creates a CreateKubernetesNamespaceInternalServerError with default headers values
func NewCreateKubernetesNamespaceInternalServerError() *CreateKubernetesNamespaceInternalServerError {
	return &CreateKubernetesNamespaceInternalServerError{}
}

/*
CreateKubernetesNamespaceInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateKubernetesNamespaceInternalServerError struct {
}

// IsSuccess returns true when this create kubernetes namespace internal server error response has a 2xx status code
func (o *CreateKubernetesNamespaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kubernetes namespace internal server error response has a 3xx status code
func (o *CreateKubernetesNamespaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kubernetes namespace internal server error response has a 4xx status code
func (o *CreateKubernetesNamespaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create kubernetes namespace internal server error response has a 5xx status code
func (o *CreateKubernetesNamespaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create kubernetes namespace internal server error response a status code equal to that given
func (o *CreateKubernetesNamespaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create kubernetes namespace internal server error response
func (o *CreateKubernetesNamespaceInternalServerError) Code() int {
	return 500
}

func (o *CreateKubernetesNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceInternalServerError ", 500)
}

func (o *CreateKubernetesNamespaceInternalServerError) String() string {
	return fmt.Sprintf("[POST /kubernetes/{id}/namespaces/{namespace}][%d] createKubernetesNamespaceInternalServerError ", 500)
}

func (o *CreateKubernetesNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
