// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateKubernetesServiceReader is a Reader for the UpdateKubernetesService structure.
type UpdateKubernetesServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKubernetesServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateKubernetesServiceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateKubernetesServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateKubernetesServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateKubernetesServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateKubernetesServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateKubernetesServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /kubernetes/{id}/namespaces/{namespace}/services] UpdateKubernetesService", response, response.Code())
	}
}

// NewUpdateKubernetesServiceNoContent creates a UpdateKubernetesServiceNoContent with default headers values
func NewUpdateKubernetesServiceNoContent() *UpdateKubernetesServiceNoContent {
	return &UpdateKubernetesServiceNoContent{}
}

/*
UpdateKubernetesServiceNoContent describes a response with status code 204, with default header values.

Success
*/
type UpdateKubernetesServiceNoContent struct {
}

// IsSuccess returns true when this update kubernetes service no content response has a 2xx status code
func (o *UpdateKubernetesServiceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update kubernetes service no content response has a 3xx status code
func (o *UpdateKubernetesServiceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service no content response has a 4xx status code
func (o *UpdateKubernetesServiceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update kubernetes service no content response has a 5xx status code
func (o *UpdateKubernetesServiceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update kubernetes service no content response a status code equal to that given
func (o *UpdateKubernetesServiceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update kubernetes service no content response
func (o *UpdateKubernetesServiceNoContent) Code() int {
	return 204
}

func (o *UpdateKubernetesServiceNoContent) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceNoContent", 204)
}

func (o *UpdateKubernetesServiceNoContent) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceNoContent", 204)
}

func (o *UpdateKubernetesServiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKubernetesServiceBadRequest creates a UpdateKubernetesServiceBadRequest with default headers values
func NewUpdateKubernetesServiceBadRequest() *UpdateKubernetesServiceBadRequest {
	return &UpdateKubernetesServiceBadRequest{}
}

/*
UpdateKubernetesServiceBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type UpdateKubernetesServiceBadRequest struct {
}

// IsSuccess returns true when this update kubernetes service bad request response has a 2xx status code
func (o *UpdateKubernetesServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kubernetes service bad request response has a 3xx status code
func (o *UpdateKubernetesServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service bad request response has a 4xx status code
func (o *UpdateKubernetesServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kubernetes service bad request response has a 5xx status code
func (o *UpdateKubernetesServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update kubernetes service bad request response a status code equal to that given
func (o *UpdateKubernetesServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update kubernetes service bad request response
func (o *UpdateKubernetesServiceBadRequest) Code() int {
	return 400
}

func (o *UpdateKubernetesServiceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceBadRequest", 400)
}

func (o *UpdateKubernetesServiceBadRequest) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceBadRequest", 400)
}

func (o *UpdateKubernetesServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKubernetesServiceUnauthorized creates a UpdateKubernetesServiceUnauthorized with default headers values
func NewUpdateKubernetesServiceUnauthorized() *UpdateKubernetesServiceUnauthorized {
	return &UpdateKubernetesServiceUnauthorized{}
}

/*
UpdateKubernetesServiceUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type UpdateKubernetesServiceUnauthorized struct {
}

// IsSuccess returns true when this update kubernetes service unauthorized response has a 2xx status code
func (o *UpdateKubernetesServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kubernetes service unauthorized response has a 3xx status code
func (o *UpdateKubernetesServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service unauthorized response has a 4xx status code
func (o *UpdateKubernetesServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kubernetes service unauthorized response has a 5xx status code
func (o *UpdateKubernetesServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update kubernetes service unauthorized response a status code equal to that given
func (o *UpdateKubernetesServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update kubernetes service unauthorized response
func (o *UpdateKubernetesServiceUnauthorized) Code() int {
	return 401
}

func (o *UpdateKubernetesServiceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceUnauthorized", 401)
}

func (o *UpdateKubernetesServiceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceUnauthorized", 401)
}

func (o *UpdateKubernetesServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKubernetesServiceForbidden creates a UpdateKubernetesServiceForbidden with default headers values
func NewUpdateKubernetesServiceForbidden() *UpdateKubernetesServiceForbidden {
	return &UpdateKubernetesServiceForbidden{}
}

/*
UpdateKubernetesServiceForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type UpdateKubernetesServiceForbidden struct {
}

// IsSuccess returns true when this update kubernetes service forbidden response has a 2xx status code
func (o *UpdateKubernetesServiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kubernetes service forbidden response has a 3xx status code
func (o *UpdateKubernetesServiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service forbidden response has a 4xx status code
func (o *UpdateKubernetesServiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kubernetes service forbidden response has a 5xx status code
func (o *UpdateKubernetesServiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update kubernetes service forbidden response a status code equal to that given
func (o *UpdateKubernetesServiceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update kubernetes service forbidden response
func (o *UpdateKubernetesServiceForbidden) Code() int {
	return 403
}

func (o *UpdateKubernetesServiceForbidden) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceForbidden", 403)
}

func (o *UpdateKubernetesServiceForbidden) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceForbidden", 403)
}

func (o *UpdateKubernetesServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKubernetesServiceNotFound creates a UpdateKubernetesServiceNotFound with default headers values
func NewUpdateKubernetesServiceNotFound() *UpdateKubernetesServiceNotFound {
	return &UpdateKubernetesServiceNotFound{}
}

/*
UpdateKubernetesServiceNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier or unable to find the service to update.
*/
type UpdateKubernetesServiceNotFound struct {
}

// IsSuccess returns true when this update kubernetes service not found response has a 2xx status code
func (o *UpdateKubernetesServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kubernetes service not found response has a 3xx status code
func (o *UpdateKubernetesServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service not found response has a 4xx status code
func (o *UpdateKubernetesServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kubernetes service not found response has a 5xx status code
func (o *UpdateKubernetesServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update kubernetes service not found response a status code equal to that given
func (o *UpdateKubernetesServiceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update kubernetes service not found response
func (o *UpdateKubernetesServiceNotFound) Code() int {
	return 404
}

func (o *UpdateKubernetesServiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceNotFound", 404)
}

func (o *UpdateKubernetesServiceNotFound) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceNotFound", 404)
}

func (o *UpdateKubernetesServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKubernetesServiceInternalServerError creates a UpdateKubernetesServiceInternalServerError with default headers values
func NewUpdateKubernetesServiceInternalServerError() *UpdateKubernetesServiceInternalServerError {
	return &UpdateKubernetesServiceInternalServerError{}
}

/*
UpdateKubernetesServiceInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to update a service.
*/
type UpdateKubernetesServiceInternalServerError struct {
}

// IsSuccess returns true when this update kubernetes service internal server error response has a 2xx status code
func (o *UpdateKubernetesServiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kubernetes service internal server error response has a 3xx status code
func (o *UpdateKubernetesServiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kubernetes service internal server error response has a 4xx status code
func (o *UpdateKubernetesServiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update kubernetes service internal server error response has a 5xx status code
func (o *UpdateKubernetesServiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update kubernetes service internal server error response a status code equal to that given
func (o *UpdateKubernetesServiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update kubernetes service internal server error response
func (o *UpdateKubernetesServiceInternalServerError) Code() int {
	return 500
}

func (o *UpdateKubernetesServiceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceInternalServerError", 500)
}

func (o *UpdateKubernetesServiceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{id}/namespaces/{namespace}/services][%d] updateKubernetesServiceInternalServerError", 500)
}

func (o *UpdateKubernetesServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
