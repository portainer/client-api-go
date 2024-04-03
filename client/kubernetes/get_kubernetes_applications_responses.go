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

// GetKubernetesApplicationsReader is a Reader for the GetKubernetesApplications structure.
type GetKubernetesApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces/{namespace}/applications] getKubernetesApplications", response, response.Code())
	}
}

// NewGetKubernetesApplicationsOK creates a GetKubernetesApplicationsOK with default headers values
func NewGetKubernetesApplicationsOK() *GetKubernetesApplicationsOK {
	return &GetKubernetesApplicationsOK{}
}

/*
GetKubernetesApplicationsOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesApplicationsOK struct {
	Payload []*models.ModelsK8sApplication
}

// IsSuccess returns true when this get kubernetes applications o k response has a 2xx status code
func (o *GetKubernetesApplicationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes applications o k response has a 3xx status code
func (o *GetKubernetesApplicationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications o k response has a 4xx status code
func (o *GetKubernetesApplicationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes applications o k response has a 5xx status code
func (o *GetKubernetesApplicationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes applications o k response a status code equal to that given
func (o *GetKubernetesApplicationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes applications o k response
func (o *GetKubernetesApplicationsOK) Code() int {
	return 200
}

func (o *GetKubernetesApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesApplicationsOK) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesApplicationsOK) GetPayload() []*models.ModelsK8sApplication {
	return o.Payload
}

func (o *GetKubernetesApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesApplicationsBadRequest creates a GetKubernetesApplicationsBadRequest with default headers values
func NewGetKubernetesApplicationsBadRequest() *GetKubernetesApplicationsBadRequest {
	return &GetKubernetesApplicationsBadRequest{}
}

/*
GetKubernetesApplicationsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetKubernetesApplicationsBadRequest struct {
}

// IsSuccess returns true when this get kubernetes applications bad request response has a 2xx status code
func (o *GetKubernetesApplicationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes applications bad request response has a 3xx status code
func (o *GetKubernetesApplicationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications bad request response has a 4xx status code
func (o *GetKubernetesApplicationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes applications bad request response has a 5xx status code
func (o *GetKubernetesApplicationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes applications bad request response a status code equal to that given
func (o *GetKubernetesApplicationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes applications bad request response
func (o *GetKubernetesApplicationsBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsBadRequest ", 400)
}

func (o *GetKubernetesApplicationsBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsBadRequest ", 400)
}

func (o *GetKubernetesApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationsUnauthorized creates a GetKubernetesApplicationsUnauthorized with default headers values
func NewGetKubernetesApplicationsUnauthorized() *GetKubernetesApplicationsUnauthorized {
	return &GetKubernetesApplicationsUnauthorized{}
}

/*
GetKubernetesApplicationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetKubernetesApplicationsUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes applications unauthorized response has a 2xx status code
func (o *GetKubernetesApplicationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes applications unauthorized response has a 3xx status code
func (o *GetKubernetesApplicationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications unauthorized response has a 4xx status code
func (o *GetKubernetesApplicationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes applications unauthorized response has a 5xx status code
func (o *GetKubernetesApplicationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes applications unauthorized response a status code equal to that given
func (o *GetKubernetesApplicationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes applications unauthorized response
func (o *GetKubernetesApplicationsUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsUnauthorized ", 401)
}

func (o *GetKubernetesApplicationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsUnauthorized ", 401)
}

func (o *GetKubernetesApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationsForbidden creates a GetKubernetesApplicationsForbidden with default headers values
func NewGetKubernetesApplicationsForbidden() *GetKubernetesApplicationsForbidden {
	return &GetKubernetesApplicationsForbidden{}
}

/*
GetKubernetesApplicationsForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type GetKubernetesApplicationsForbidden struct {
}

// IsSuccess returns true when this get kubernetes applications forbidden response has a 2xx status code
func (o *GetKubernetesApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes applications forbidden response has a 3xx status code
func (o *GetKubernetesApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications forbidden response has a 4xx status code
func (o *GetKubernetesApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes applications forbidden response has a 5xx status code
func (o *GetKubernetesApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes applications forbidden response a status code equal to that given
func (o *GetKubernetesApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes applications forbidden response
func (o *GetKubernetesApplicationsForbidden) Code() int {
	return 403
}

func (o *GetKubernetesApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsForbidden ", 403)
}

func (o *GetKubernetesApplicationsForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsForbidden ", 403)
}

func (o *GetKubernetesApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationsNotFound creates a GetKubernetesApplicationsNotFound with default headers values
func NewGetKubernetesApplicationsNotFound() *GetKubernetesApplicationsNotFound {
	return &GetKubernetesApplicationsNotFound{}
}

/*
GetKubernetesApplicationsNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) or ServiceAccount not found
*/
type GetKubernetesApplicationsNotFound struct {
}

// IsSuccess returns true when this get kubernetes applications not found response has a 2xx status code
func (o *GetKubernetesApplicationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes applications not found response has a 3xx status code
func (o *GetKubernetesApplicationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications not found response has a 4xx status code
func (o *GetKubernetesApplicationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes applications not found response has a 5xx status code
func (o *GetKubernetesApplicationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes applications not found response a status code equal to that given
func (o *GetKubernetesApplicationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes applications not found response
func (o *GetKubernetesApplicationsNotFound) Code() int {
	return 404
}

func (o *GetKubernetesApplicationsNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsNotFound ", 404)
}

func (o *GetKubernetesApplicationsNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsNotFound ", 404)
}

func (o *GetKubernetesApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationsInternalServerError creates a GetKubernetesApplicationsInternalServerError with default headers values
func NewGetKubernetesApplicationsInternalServerError() *GetKubernetesApplicationsInternalServerError {
	return &GetKubernetesApplicationsInternalServerError{}
}

/*
GetKubernetesApplicationsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetKubernetesApplicationsInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes applications internal server error response has a 2xx status code
func (o *GetKubernetesApplicationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes applications internal server error response has a 3xx status code
func (o *GetKubernetesApplicationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes applications internal server error response has a 4xx status code
func (o *GetKubernetesApplicationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes applications internal server error response has a 5xx status code
func (o *GetKubernetesApplicationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes applications internal server error response a status code equal to that given
func (o *GetKubernetesApplicationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes applications internal server error response
func (o *GetKubernetesApplicationsInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsInternalServerError ", 500)
}

func (o *GetKubernetesApplicationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications][%d] getKubernetesApplicationsInternalServerError ", 500)
}

func (o *GetKubernetesApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
