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

// GetKubernetesSecretReader is a Reader for the GetKubernetesSecret structure.
type GetKubernetesSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}] GetKubernetesSecret", response, response.Code())
	}
}

// NewGetKubernetesSecretOK creates a GetKubernetesSecretOK with default headers values
func NewGetKubernetesSecretOK() *GetKubernetesSecretOK {
	return &GetKubernetesSecretOK{}
}

/*
GetKubernetesSecretOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesSecretOK struct {
	Payload *models.KubernetesK8sSecret
}

// IsSuccess returns true when this get kubernetes secret o k response has a 2xx status code
func (o *GetKubernetesSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes secret o k response has a 3xx status code
func (o *GetKubernetesSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret o k response has a 4xx status code
func (o *GetKubernetesSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes secret o k response has a 5xx status code
func (o *GetKubernetesSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secret o k response a status code equal to that given
func (o *GetKubernetesSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes secret o k response
func (o *GetKubernetesSecretOK) Code() int {
	return 200
}

func (o *GetKubernetesSecretOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretOK %s", 200, payload)
}

func (o *GetKubernetesSecretOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretOK %s", 200, payload)
}

func (o *GetKubernetesSecretOK) GetPayload() *models.KubernetesK8sSecret {
	return o.Payload
}

func (o *GetKubernetesSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesK8sSecret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesSecretBadRequest creates a GetKubernetesSecretBadRequest with default headers values
func NewGetKubernetesSecretBadRequest() *GetKubernetesSecretBadRequest {
	return &GetKubernetesSecretBadRequest{}
}

/*
GetKubernetesSecretBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetKubernetesSecretBadRequest struct {
}

// IsSuccess returns true when this get kubernetes secret bad request response has a 2xx status code
func (o *GetKubernetesSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secret bad request response has a 3xx status code
func (o *GetKubernetesSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret bad request response has a 4xx status code
func (o *GetKubernetesSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secret bad request response has a 5xx status code
func (o *GetKubernetesSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secret bad request response a status code equal to that given
func (o *GetKubernetesSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes secret bad request response
func (o *GetKubernetesSecretBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesSecretBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretBadRequest", 400)
}

func (o *GetKubernetesSecretBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretBadRequest", 400)
}

func (o *GetKubernetesSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretUnauthorized creates a GetKubernetesSecretUnauthorized with default headers values
func NewGetKubernetesSecretUnauthorized() *GetKubernetesSecretUnauthorized {
	return &GetKubernetesSecretUnauthorized{}
}

/*
GetKubernetesSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetKubernetesSecretUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes secret unauthorized response has a 2xx status code
func (o *GetKubernetesSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secret unauthorized response has a 3xx status code
func (o *GetKubernetesSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret unauthorized response has a 4xx status code
func (o *GetKubernetesSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secret unauthorized response has a 5xx status code
func (o *GetKubernetesSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secret unauthorized response a status code equal to that given
func (o *GetKubernetesSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes secret unauthorized response
func (o *GetKubernetesSecretUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretUnauthorized", 401)
}

func (o *GetKubernetesSecretUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretUnauthorized", 401)
}

func (o *GetKubernetesSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretForbidden creates a GetKubernetesSecretForbidden with default headers values
func NewGetKubernetesSecretForbidden() *GetKubernetesSecretForbidden {
	return &GetKubernetesSecretForbidden{}
}

/*
GetKubernetesSecretForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetKubernetesSecretForbidden struct {
}

// IsSuccess returns true when this get kubernetes secret forbidden response has a 2xx status code
func (o *GetKubernetesSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secret forbidden response has a 3xx status code
func (o *GetKubernetesSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret forbidden response has a 4xx status code
func (o *GetKubernetesSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secret forbidden response has a 5xx status code
func (o *GetKubernetesSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secret forbidden response a status code equal to that given
func (o *GetKubernetesSecretForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes secret forbidden response
func (o *GetKubernetesSecretForbidden) Code() int {
	return 403
}

func (o *GetKubernetesSecretForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretForbidden", 403)
}

func (o *GetKubernetesSecretForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretForbidden", 403)
}

func (o *GetKubernetesSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretNotFound creates a GetKubernetesSecretNotFound with default headers values
func NewGetKubernetesSecretNotFound() *GetKubernetesSecretNotFound {
	return &GetKubernetesSecretNotFound{}
}

/*
GetKubernetesSecretNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier.
*/
type GetKubernetesSecretNotFound struct {
}

// IsSuccess returns true when this get kubernetes secret not found response has a 2xx status code
func (o *GetKubernetesSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secret not found response has a 3xx status code
func (o *GetKubernetesSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret not found response has a 4xx status code
func (o *GetKubernetesSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secret not found response has a 5xx status code
func (o *GetKubernetesSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secret not found response a status code equal to that given
func (o *GetKubernetesSecretNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes secret not found response
func (o *GetKubernetesSecretNotFound) Code() int {
	return 404
}

func (o *GetKubernetesSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretNotFound", 404)
}

func (o *GetKubernetesSecretNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretNotFound", 404)
}

func (o *GetKubernetesSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretInternalServerError creates a GetKubernetesSecretInternalServerError with default headers values
func NewGetKubernetesSecretInternalServerError() *GetKubernetesSecretInternalServerError {
	return &GetKubernetesSecretInternalServerError{}
}

/*
GetKubernetesSecretInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve a secret by name belong in a namespace.
*/
type GetKubernetesSecretInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes secret internal server error response has a 2xx status code
func (o *GetKubernetesSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secret internal server error response has a 3xx status code
func (o *GetKubernetesSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secret internal server error response has a 4xx status code
func (o *GetKubernetesSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes secret internal server error response has a 5xx status code
func (o *GetKubernetesSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes secret internal server error response a status code equal to that given
func (o *GetKubernetesSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes secret internal server error response
func (o *GetKubernetesSecretInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesSecretInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretInternalServerError", 500)
}

func (o *GetKubernetesSecretInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/secrets/{secret}][%d] getKubernetesSecretInternalServerError", 500)
}

func (o *GetKubernetesSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
