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

// GetKubernetesEventsForNamespaceReader is a Reader for the GetKubernetesEventsForNamespace structure.
type GetKubernetesEventsForNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesEventsForNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesEventsForNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesEventsForNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesEventsForNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesEventsForNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesEventsForNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces/{namespace}/events] getKubernetesEventsForNamespace", response, response.Code())
	}
}

// NewGetKubernetesEventsForNamespaceOK creates a GetKubernetesEventsForNamespaceOK with default headers values
func NewGetKubernetesEventsForNamespaceOK() *GetKubernetesEventsForNamespaceOK {
	return &GetKubernetesEventsForNamespaceOK{}
}

/*
GetKubernetesEventsForNamespaceOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesEventsForNamespaceOK struct {
	Payload []*models.KubernetesK8sEvent
}

// IsSuccess returns true when this get kubernetes events for namespace o k response has a 2xx status code
func (o *GetKubernetesEventsForNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes events for namespace o k response has a 3xx status code
func (o *GetKubernetesEventsForNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes events for namespace o k response has a 4xx status code
func (o *GetKubernetesEventsForNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes events for namespace o k response has a 5xx status code
func (o *GetKubernetesEventsForNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes events for namespace o k response a status code equal to that given
func (o *GetKubernetesEventsForNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes events for namespace o k response
func (o *GetKubernetesEventsForNamespaceOK) Code() int {
	return 200
}

func (o *GetKubernetesEventsForNamespaceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceOK %s", 200, payload)
}

func (o *GetKubernetesEventsForNamespaceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceOK %s", 200, payload)
}

func (o *GetKubernetesEventsForNamespaceOK) GetPayload() []*models.KubernetesK8sEvent {
	return o.Payload
}

func (o *GetKubernetesEventsForNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesEventsForNamespaceBadRequest creates a GetKubernetesEventsForNamespaceBadRequest with default headers values
func NewGetKubernetesEventsForNamespaceBadRequest() *GetKubernetesEventsForNamespaceBadRequest {
	return &GetKubernetesEventsForNamespaceBadRequest{}
}

/*
GetKubernetesEventsForNamespaceBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetKubernetesEventsForNamespaceBadRequest struct {
}

// IsSuccess returns true when this get kubernetes events for namespace bad request response has a 2xx status code
func (o *GetKubernetesEventsForNamespaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes events for namespace bad request response has a 3xx status code
func (o *GetKubernetesEventsForNamespaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes events for namespace bad request response has a 4xx status code
func (o *GetKubernetesEventsForNamespaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes events for namespace bad request response has a 5xx status code
func (o *GetKubernetesEventsForNamespaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes events for namespace bad request response a status code equal to that given
func (o *GetKubernetesEventsForNamespaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes events for namespace bad request response
func (o *GetKubernetesEventsForNamespaceBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesEventsForNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceBadRequest", 400)
}

func (o *GetKubernetesEventsForNamespaceBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceBadRequest", 400)
}

func (o *GetKubernetesEventsForNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesEventsForNamespaceUnauthorized creates a GetKubernetesEventsForNamespaceUnauthorized with default headers values
func NewGetKubernetesEventsForNamespaceUnauthorized() *GetKubernetesEventsForNamespaceUnauthorized {
	return &GetKubernetesEventsForNamespaceUnauthorized{}
}

/*
GetKubernetesEventsForNamespaceUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetKubernetesEventsForNamespaceUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes events for namespace unauthorized response has a 2xx status code
func (o *GetKubernetesEventsForNamespaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes events for namespace unauthorized response has a 3xx status code
func (o *GetKubernetesEventsForNamespaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes events for namespace unauthorized response has a 4xx status code
func (o *GetKubernetesEventsForNamespaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes events for namespace unauthorized response has a 5xx status code
func (o *GetKubernetesEventsForNamespaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes events for namespace unauthorized response a status code equal to that given
func (o *GetKubernetesEventsForNamespaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes events for namespace unauthorized response
func (o *GetKubernetesEventsForNamespaceUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesEventsForNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceUnauthorized", 401)
}

func (o *GetKubernetesEventsForNamespaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceUnauthorized", 401)
}

func (o *GetKubernetesEventsForNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesEventsForNamespaceForbidden creates a GetKubernetesEventsForNamespaceForbidden with default headers values
func NewGetKubernetesEventsForNamespaceForbidden() *GetKubernetesEventsForNamespaceForbidden {
	return &GetKubernetesEventsForNamespaceForbidden{}
}

/*
GetKubernetesEventsForNamespaceForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetKubernetesEventsForNamespaceForbidden struct {
}

// IsSuccess returns true when this get kubernetes events for namespace forbidden response has a 2xx status code
func (o *GetKubernetesEventsForNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes events for namespace forbidden response has a 3xx status code
func (o *GetKubernetesEventsForNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes events for namespace forbidden response has a 4xx status code
func (o *GetKubernetesEventsForNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes events for namespace forbidden response has a 5xx status code
func (o *GetKubernetesEventsForNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes events for namespace forbidden response a status code equal to that given
func (o *GetKubernetesEventsForNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes events for namespace forbidden response
func (o *GetKubernetesEventsForNamespaceForbidden) Code() int {
	return 403
}

func (o *GetKubernetesEventsForNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceForbidden", 403)
}

func (o *GetKubernetesEventsForNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceForbidden", 403)
}

func (o *GetKubernetesEventsForNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesEventsForNamespaceInternalServerError creates a GetKubernetesEventsForNamespaceInternalServerError with default headers values
func NewGetKubernetesEventsForNamespaceInternalServerError() *GetKubernetesEventsForNamespaceInternalServerError {
	return &GetKubernetesEventsForNamespaceInternalServerError{}
}

/*
GetKubernetesEventsForNamespaceInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve the events within the specified namespace.
*/
type GetKubernetesEventsForNamespaceInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes events for namespace internal server error response has a 2xx status code
func (o *GetKubernetesEventsForNamespaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes events for namespace internal server error response has a 3xx status code
func (o *GetKubernetesEventsForNamespaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes events for namespace internal server error response has a 4xx status code
func (o *GetKubernetesEventsForNamespaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes events for namespace internal server error response has a 5xx status code
func (o *GetKubernetesEventsForNamespaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes events for namespace internal server error response a status code equal to that given
func (o *GetKubernetesEventsForNamespaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes events for namespace internal server error response
func (o *GetKubernetesEventsForNamespaceInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesEventsForNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceInternalServerError", 500)
}

func (o *GetKubernetesEventsForNamespaceInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/events][%d] getKubernetesEventsForNamespaceInternalServerError", 500)
}

func (o *GetKubernetesEventsForNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
