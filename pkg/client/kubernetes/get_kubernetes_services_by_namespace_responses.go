// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// GetKubernetesServicesByNamespaceReader is a Reader for the GetKubernetesServicesByNamespace structure.
type GetKubernetesServicesByNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesServicesByNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesServicesByNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesServicesByNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesServicesByNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesServicesByNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesServicesByNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesServicesByNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces/{namespace}/services] GetKubernetesServicesByNamespace", response, response.Code())
	}
}

// NewGetKubernetesServicesByNamespaceOK creates a GetKubernetesServicesByNamespaceOK with default headers values
func NewGetKubernetesServicesByNamespaceOK() *GetKubernetesServicesByNamespaceOK {
	return &GetKubernetesServicesByNamespaceOK{}
}

/*
GetKubernetesServicesByNamespaceOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesServicesByNamespaceOK struct {
	Payload []*models.KubernetesK8sServiceInfo
}

// IsSuccess returns true when this get kubernetes services by namespace o k response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes services by namespace o k response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace o k response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes services by namespace o k response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes services by namespace o k response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes services by namespace o k response
func (o *GetKubernetesServicesByNamespaceOK) Code() int {
	return 200
}

func (o *GetKubernetesServicesByNamespaceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceOK %s", 200, payload)
}

func (o *GetKubernetesServicesByNamespaceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceOK %s", 200, payload)
}

func (o *GetKubernetesServicesByNamespaceOK) GetPayload() []*models.KubernetesK8sServiceInfo {
	return o.Payload
}

func (o *GetKubernetesServicesByNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesServicesByNamespaceBadRequest creates a GetKubernetesServicesByNamespaceBadRequest with default headers values
func NewGetKubernetesServicesByNamespaceBadRequest() *GetKubernetesServicesByNamespaceBadRequest {
	return &GetKubernetesServicesByNamespaceBadRequest{}
}

/*
GetKubernetesServicesByNamespaceBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetKubernetesServicesByNamespaceBadRequest struct {
}

// IsSuccess returns true when this get kubernetes services by namespace bad request response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes services by namespace bad request response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace bad request response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes services by namespace bad request response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes services by namespace bad request response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes services by namespace bad request response
func (o *GetKubernetesServicesByNamespaceBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesServicesByNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceBadRequest", 400)
}

func (o *GetKubernetesServicesByNamespaceBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceBadRequest", 400)
}

func (o *GetKubernetesServicesByNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesServicesByNamespaceUnauthorized creates a GetKubernetesServicesByNamespaceUnauthorized with default headers values
func NewGetKubernetesServicesByNamespaceUnauthorized() *GetKubernetesServicesByNamespaceUnauthorized {
	return &GetKubernetesServicesByNamespaceUnauthorized{}
}

/*
GetKubernetesServicesByNamespaceUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetKubernetesServicesByNamespaceUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes services by namespace unauthorized response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes services by namespace unauthorized response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace unauthorized response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes services by namespace unauthorized response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes services by namespace unauthorized response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes services by namespace unauthorized response
func (o *GetKubernetesServicesByNamespaceUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesServicesByNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceUnauthorized", 401)
}

func (o *GetKubernetesServicesByNamespaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceUnauthorized", 401)
}

func (o *GetKubernetesServicesByNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesServicesByNamespaceForbidden creates a GetKubernetesServicesByNamespaceForbidden with default headers values
func NewGetKubernetesServicesByNamespaceForbidden() *GetKubernetesServicesByNamespaceForbidden {
	return &GetKubernetesServicesByNamespaceForbidden{}
}

/*
GetKubernetesServicesByNamespaceForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetKubernetesServicesByNamespaceForbidden struct {
}

// IsSuccess returns true when this get kubernetes services by namespace forbidden response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes services by namespace forbidden response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace forbidden response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes services by namespace forbidden response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes services by namespace forbidden response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes services by namespace forbidden response
func (o *GetKubernetesServicesByNamespaceForbidden) Code() int {
	return 403
}

func (o *GetKubernetesServicesByNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceForbidden", 403)
}

func (o *GetKubernetesServicesByNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceForbidden", 403)
}

func (o *GetKubernetesServicesByNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesServicesByNamespaceNotFound creates a GetKubernetesServicesByNamespaceNotFound with default headers values
func NewGetKubernetesServicesByNamespaceNotFound() *GetKubernetesServicesByNamespaceNotFound {
	return &GetKubernetesServicesByNamespaceNotFound{}
}

/*
GetKubernetesServicesByNamespaceNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier.
*/
type GetKubernetesServicesByNamespaceNotFound struct {
}

// IsSuccess returns true when this get kubernetes services by namespace not found response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes services by namespace not found response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace not found response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes services by namespace not found response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes services by namespace not found response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes services by namespace not found response
func (o *GetKubernetesServicesByNamespaceNotFound) Code() int {
	return 404
}

func (o *GetKubernetesServicesByNamespaceNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceNotFound", 404)
}

func (o *GetKubernetesServicesByNamespaceNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceNotFound", 404)
}

func (o *GetKubernetesServicesByNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesServicesByNamespaceInternalServerError creates a GetKubernetesServicesByNamespaceInternalServerError with default headers values
func NewGetKubernetesServicesByNamespaceInternalServerError() *GetKubernetesServicesByNamespaceInternalServerError {
	return &GetKubernetesServicesByNamespaceInternalServerError{}
}

/*
GetKubernetesServicesByNamespaceInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve all services for a namespace.
*/
type GetKubernetesServicesByNamespaceInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes services by namespace internal server error response has a 2xx status code
func (o *GetKubernetesServicesByNamespaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes services by namespace internal server error response has a 3xx status code
func (o *GetKubernetesServicesByNamespaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes services by namespace internal server error response has a 4xx status code
func (o *GetKubernetesServicesByNamespaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes services by namespace internal server error response has a 5xx status code
func (o *GetKubernetesServicesByNamespaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes services by namespace internal server error response a status code equal to that given
func (o *GetKubernetesServicesByNamespaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes services by namespace internal server error response
func (o *GetKubernetesServicesByNamespaceInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesServicesByNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceInternalServerError", 500)
}

func (o *GetKubernetesServicesByNamespaceInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/services][%d] getKubernetesServicesByNamespaceInternalServerError", 500)
}

func (o *GetKubernetesServicesByNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
