// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// GetKubernetesNamespacesReader is a Reader for the GetKubernetesNamespaces structure.
type GetKubernetesNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesNamespacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesNamespacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces] getKubernetesNamespaces", response, response.Code())
	}
}

// NewGetKubernetesNamespacesOK creates a GetKubernetesNamespacesOK with default headers values
func NewGetKubernetesNamespacesOK() *GetKubernetesNamespacesOK {
	return &GetKubernetesNamespacesOK{}
}

/*
GetKubernetesNamespacesOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesNamespacesOK struct {
	Payload map[string]models.PortainereeK8sNamespaceInfo
}

// IsSuccess returns true when this get kubernetes namespaces o k response has a 2xx status code
func (o *GetKubernetesNamespacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes namespaces o k response has a 3xx status code
func (o *GetKubernetesNamespacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes namespaces o k response has a 4xx status code
func (o *GetKubernetesNamespacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes namespaces o k response has a 5xx status code
func (o *GetKubernetesNamespacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes namespaces o k response a status code equal to that given
func (o *GetKubernetesNamespacesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes namespaces o k response
func (o *GetKubernetesNamespacesOK) Code() int {
	return 200
}

func (o *GetKubernetesNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesNamespacesOK) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesNamespacesOK) GetPayload() map[string]models.PortainereeK8sNamespaceInfo {
	return o.Payload
}

func (o *GetKubernetesNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesNamespacesBadRequest creates a GetKubernetesNamespacesBadRequest with default headers values
func NewGetKubernetesNamespacesBadRequest() *GetKubernetesNamespacesBadRequest {
	return &GetKubernetesNamespacesBadRequest{}
}

/*
GetKubernetesNamespacesBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetKubernetesNamespacesBadRequest struct {
}

// IsSuccess returns true when this get kubernetes namespaces bad request response has a 2xx status code
func (o *GetKubernetesNamespacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes namespaces bad request response has a 3xx status code
func (o *GetKubernetesNamespacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes namespaces bad request response has a 4xx status code
func (o *GetKubernetesNamespacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes namespaces bad request response has a 5xx status code
func (o *GetKubernetesNamespacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes namespaces bad request response a status code equal to that given
func (o *GetKubernetesNamespacesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes namespaces bad request response
func (o *GetKubernetesNamespacesBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesNamespacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesBadRequest ", 400)
}

func (o *GetKubernetesNamespacesBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesBadRequest ", 400)
}

func (o *GetKubernetesNamespacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesNamespacesInternalServerError creates a GetKubernetesNamespacesInternalServerError with default headers values
func NewGetKubernetesNamespacesInternalServerError() *GetKubernetesNamespacesInternalServerError {
	return &GetKubernetesNamespacesInternalServerError{}
}

/*
GetKubernetesNamespacesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetKubernetesNamespacesInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes namespaces internal server error response has a 2xx status code
func (o *GetKubernetesNamespacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes namespaces internal server error response has a 3xx status code
func (o *GetKubernetesNamespacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes namespaces internal server error response has a 4xx status code
func (o *GetKubernetesNamespacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes namespaces internal server error response has a 5xx status code
func (o *GetKubernetesNamespacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes namespaces internal server error response a status code equal to that given
func (o *GetKubernetesNamespacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes namespaces internal server error response
func (o *GetKubernetesNamespacesInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesNamespacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesInternalServerError ", 500)
}

func (o *GetKubernetesNamespacesInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces][%d] getKubernetesNamespacesInternalServerError ", 500)
}

func (o *GetKubernetesNamespacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
