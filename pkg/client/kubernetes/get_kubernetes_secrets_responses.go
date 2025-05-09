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

// GetKubernetesSecretsReader is a Reader for the GetKubernetesSecrets structure.
type GetKubernetesSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesSecretsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesSecretsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesSecretsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/secrets] GetKubernetesSecrets", response, response.Code())
	}
}

// NewGetKubernetesSecretsOK creates a GetKubernetesSecretsOK with default headers values
func NewGetKubernetesSecretsOK() *GetKubernetesSecretsOK {
	return &GetKubernetesSecretsOK{}
}

/*
GetKubernetesSecretsOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesSecretsOK struct {
	Payload []*models.KubernetesK8sSecret
}

// IsSuccess returns true when this get kubernetes secrets o k response has a 2xx status code
func (o *GetKubernetesSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes secrets o k response has a 3xx status code
func (o *GetKubernetesSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets o k response has a 4xx status code
func (o *GetKubernetesSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes secrets o k response has a 5xx status code
func (o *GetKubernetesSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secrets o k response a status code equal to that given
func (o *GetKubernetesSecretsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes secrets o k response
func (o *GetKubernetesSecretsOK) Code() int {
	return 200
}

func (o *GetKubernetesSecretsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsOK %s", 200, payload)
}

func (o *GetKubernetesSecretsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsOK %s", 200, payload)
}

func (o *GetKubernetesSecretsOK) GetPayload() []*models.KubernetesK8sSecret {
	return o.Payload
}

func (o *GetKubernetesSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesSecretsBadRequest creates a GetKubernetesSecretsBadRequest with default headers values
func NewGetKubernetesSecretsBadRequest() *GetKubernetesSecretsBadRequest {
	return &GetKubernetesSecretsBadRequest{}
}

/*
GetKubernetesSecretsBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetKubernetesSecretsBadRequest struct {
}

// IsSuccess returns true when this get kubernetes secrets bad request response has a 2xx status code
func (o *GetKubernetesSecretsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secrets bad request response has a 3xx status code
func (o *GetKubernetesSecretsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets bad request response has a 4xx status code
func (o *GetKubernetesSecretsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secrets bad request response has a 5xx status code
func (o *GetKubernetesSecretsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secrets bad request response a status code equal to that given
func (o *GetKubernetesSecretsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes secrets bad request response
func (o *GetKubernetesSecretsBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesSecretsBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsBadRequest", 400)
}

func (o *GetKubernetesSecretsBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsBadRequest", 400)
}

func (o *GetKubernetesSecretsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretsUnauthorized creates a GetKubernetesSecretsUnauthorized with default headers values
func NewGetKubernetesSecretsUnauthorized() *GetKubernetesSecretsUnauthorized {
	return &GetKubernetesSecretsUnauthorized{}
}

/*
GetKubernetesSecretsUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetKubernetesSecretsUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes secrets unauthorized response has a 2xx status code
func (o *GetKubernetesSecretsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secrets unauthorized response has a 3xx status code
func (o *GetKubernetesSecretsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets unauthorized response has a 4xx status code
func (o *GetKubernetesSecretsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secrets unauthorized response has a 5xx status code
func (o *GetKubernetesSecretsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secrets unauthorized response a status code equal to that given
func (o *GetKubernetesSecretsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes secrets unauthorized response
func (o *GetKubernetesSecretsUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesSecretsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsUnauthorized", 401)
}

func (o *GetKubernetesSecretsUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsUnauthorized", 401)
}

func (o *GetKubernetesSecretsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretsForbidden creates a GetKubernetesSecretsForbidden with default headers values
func NewGetKubernetesSecretsForbidden() *GetKubernetesSecretsForbidden {
	return &GetKubernetesSecretsForbidden{}
}

/*
GetKubernetesSecretsForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetKubernetesSecretsForbidden struct {
}

// IsSuccess returns true when this get kubernetes secrets forbidden response has a 2xx status code
func (o *GetKubernetesSecretsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secrets forbidden response has a 3xx status code
func (o *GetKubernetesSecretsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets forbidden response has a 4xx status code
func (o *GetKubernetesSecretsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secrets forbidden response has a 5xx status code
func (o *GetKubernetesSecretsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secrets forbidden response a status code equal to that given
func (o *GetKubernetesSecretsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes secrets forbidden response
func (o *GetKubernetesSecretsForbidden) Code() int {
	return 403
}

func (o *GetKubernetesSecretsForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsForbidden", 403)
}

func (o *GetKubernetesSecretsForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsForbidden", 403)
}

func (o *GetKubernetesSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretsNotFound creates a GetKubernetesSecretsNotFound with default headers values
func NewGetKubernetesSecretsNotFound() *GetKubernetesSecretsNotFound {
	return &GetKubernetesSecretsNotFound{}
}

/*
GetKubernetesSecretsNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier.
*/
type GetKubernetesSecretsNotFound struct {
}

// IsSuccess returns true when this get kubernetes secrets not found response has a 2xx status code
func (o *GetKubernetesSecretsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secrets not found response has a 3xx status code
func (o *GetKubernetesSecretsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets not found response has a 4xx status code
func (o *GetKubernetesSecretsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes secrets not found response has a 5xx status code
func (o *GetKubernetesSecretsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes secrets not found response a status code equal to that given
func (o *GetKubernetesSecretsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes secrets not found response
func (o *GetKubernetesSecretsNotFound) Code() int {
	return 404
}

func (o *GetKubernetesSecretsNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsNotFound", 404)
}

func (o *GetKubernetesSecretsNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsNotFound", 404)
}

func (o *GetKubernetesSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesSecretsInternalServerError creates a GetKubernetesSecretsInternalServerError with default headers values
func NewGetKubernetesSecretsInternalServerError() *GetKubernetesSecretsInternalServerError {
	return &GetKubernetesSecretsInternalServerError{}
}

/*
GetKubernetesSecretsInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve all secrets from the cluster.
*/
type GetKubernetesSecretsInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes secrets internal server error response has a 2xx status code
func (o *GetKubernetesSecretsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes secrets internal server error response has a 3xx status code
func (o *GetKubernetesSecretsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes secrets internal server error response has a 4xx status code
func (o *GetKubernetesSecretsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes secrets internal server error response has a 5xx status code
func (o *GetKubernetesSecretsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes secrets internal server error response a status code equal to that given
func (o *GetKubernetesSecretsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes secrets internal server error response
func (o *GetKubernetesSecretsInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesSecretsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsInternalServerError", 500)
}

func (o *GetKubernetesSecretsInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/secrets][%d] getKubernetesSecretsInternalServerError", 500)
}

func (o *GetKubernetesSecretsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
