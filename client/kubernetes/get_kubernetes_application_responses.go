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

// GetKubernetesApplicationReader is a Reader for the GetKubernetesApplication structure.
type GetKubernetesApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesApplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKubernetesApplicationOK creates a GetKubernetesApplicationOK with default headers values
func NewGetKubernetesApplicationOK() *GetKubernetesApplicationOK {
	return &GetKubernetesApplicationOK{}
}

/*
GetKubernetesApplicationOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesApplicationOK struct {
	Payload *models.ModelsK8sApplication
}

// IsSuccess returns true when this get kubernetes application o k response has a 2xx status code
func (o *GetKubernetesApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes application o k response has a 3xx status code
func (o *GetKubernetesApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application o k response has a 4xx status code
func (o *GetKubernetesApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes application o k response has a 5xx status code
func (o *GetKubernetesApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes application o k response a status code equal to that given
func (o *GetKubernetesApplicationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKubernetesApplicationOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesApplicationOK) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesApplicationOK) GetPayload() *models.ModelsK8sApplication {
	return o.Payload
}

func (o *GetKubernetesApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsK8sApplication)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesApplicationBadRequest creates a GetKubernetesApplicationBadRequest with default headers values
func NewGetKubernetesApplicationBadRequest() *GetKubernetesApplicationBadRequest {
	return &GetKubernetesApplicationBadRequest{}
}

/*
GetKubernetesApplicationBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetKubernetesApplicationBadRequest struct {
}

// IsSuccess returns true when this get kubernetes application bad request response has a 2xx status code
func (o *GetKubernetesApplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes application bad request response has a 3xx status code
func (o *GetKubernetesApplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application bad request response has a 4xx status code
func (o *GetKubernetesApplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes application bad request response has a 5xx status code
func (o *GetKubernetesApplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes application bad request response a status code equal to that given
func (o *GetKubernetesApplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKubernetesApplicationBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationBadRequest ", 400)
}

func (o *GetKubernetesApplicationBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationBadRequest ", 400)
}

func (o *GetKubernetesApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationUnauthorized creates a GetKubernetesApplicationUnauthorized with default headers values
func NewGetKubernetesApplicationUnauthorized() *GetKubernetesApplicationUnauthorized {
	return &GetKubernetesApplicationUnauthorized{}
}

/*
GetKubernetesApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetKubernetesApplicationUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes application unauthorized response has a 2xx status code
func (o *GetKubernetesApplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes application unauthorized response has a 3xx status code
func (o *GetKubernetesApplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application unauthorized response has a 4xx status code
func (o *GetKubernetesApplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes application unauthorized response has a 5xx status code
func (o *GetKubernetesApplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes application unauthorized response a status code equal to that given
func (o *GetKubernetesApplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKubernetesApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationUnauthorized ", 401)
}

func (o *GetKubernetesApplicationUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationUnauthorized ", 401)
}

func (o *GetKubernetesApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationForbidden creates a GetKubernetesApplicationForbidden with default headers values
func NewGetKubernetesApplicationForbidden() *GetKubernetesApplicationForbidden {
	return &GetKubernetesApplicationForbidden{}
}

/*
GetKubernetesApplicationForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type GetKubernetesApplicationForbidden struct {
}

// IsSuccess returns true when this get kubernetes application forbidden response has a 2xx status code
func (o *GetKubernetesApplicationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes application forbidden response has a 3xx status code
func (o *GetKubernetesApplicationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application forbidden response has a 4xx status code
func (o *GetKubernetesApplicationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes application forbidden response has a 5xx status code
func (o *GetKubernetesApplicationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes application forbidden response a status code equal to that given
func (o *GetKubernetesApplicationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKubernetesApplicationForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationForbidden ", 403)
}

func (o *GetKubernetesApplicationForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationForbidden ", 403)
}

func (o *GetKubernetesApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationNotFound creates a GetKubernetesApplicationNotFound with default headers values
func NewGetKubernetesApplicationNotFound() *GetKubernetesApplicationNotFound {
	return &GetKubernetesApplicationNotFound{}
}

/*
GetKubernetesApplicationNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) or ServiceAccount not found
*/
type GetKubernetesApplicationNotFound struct {
}

// IsSuccess returns true when this get kubernetes application not found response has a 2xx status code
func (o *GetKubernetesApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes application not found response has a 3xx status code
func (o *GetKubernetesApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application not found response has a 4xx status code
func (o *GetKubernetesApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes application not found response has a 5xx status code
func (o *GetKubernetesApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes application not found response a status code equal to that given
func (o *GetKubernetesApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKubernetesApplicationNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationNotFound ", 404)
}

func (o *GetKubernetesApplicationNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationNotFound ", 404)
}

func (o *GetKubernetesApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesApplicationInternalServerError creates a GetKubernetesApplicationInternalServerError with default headers values
func NewGetKubernetesApplicationInternalServerError() *GetKubernetesApplicationInternalServerError {
	return &GetKubernetesApplicationInternalServerError{}
}

/*
GetKubernetesApplicationInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetKubernetesApplicationInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes application internal server error response has a 2xx status code
func (o *GetKubernetesApplicationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes application internal server error response has a 3xx status code
func (o *GetKubernetesApplicationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes application internal server error response has a 4xx status code
func (o *GetKubernetesApplicationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes application internal server error response has a 5xx status code
func (o *GetKubernetesApplicationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes application internal server error response a status code equal to that given
func (o *GetKubernetesApplicationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKubernetesApplicationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationInternalServerError ", 500)
}

func (o *GetKubernetesApplicationInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/applications/{kind}/{name}][%d] getKubernetesApplicationInternalServerError ", 500)
}

func (o *GetKubernetesApplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
