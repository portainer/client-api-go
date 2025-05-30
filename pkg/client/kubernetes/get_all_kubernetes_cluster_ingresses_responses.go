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

// GetAllKubernetesClusterIngressesReader is a Reader for the GetAllKubernetesClusterIngresses structure.
type GetAllKubernetesClusterIngressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllKubernetesClusterIngressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllKubernetesClusterIngressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllKubernetesClusterIngressesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllKubernetesClusterIngressesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllKubernetesClusterIngressesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllKubernetesClusterIngressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllKubernetesClusterIngressesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/ingresses] GetAllKubernetesClusterIngresses", response, response.Code())
	}
}

// NewGetAllKubernetesClusterIngressesOK creates a GetAllKubernetesClusterIngressesOK with default headers values
func NewGetAllKubernetesClusterIngressesOK() *GetAllKubernetesClusterIngressesOK {
	return &GetAllKubernetesClusterIngressesOK{}
}

/*
GetAllKubernetesClusterIngressesOK describes a response with status code 200, with default header values.

Success
*/
type GetAllKubernetesClusterIngressesOK struct {
	Payload []*models.KubernetesK8sIngressInfo
}

// IsSuccess returns true when this get all kubernetes cluster ingresses o k response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all kubernetes cluster ingresses o k response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses o k response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all kubernetes cluster ingresses o k response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes cluster ingresses o k response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all kubernetes cluster ingresses o k response
func (o *GetAllKubernetesClusterIngressesOK) Code() int {
	return 200
}

func (o *GetAllKubernetesClusterIngressesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesOK %s", 200, payload)
}

func (o *GetAllKubernetesClusterIngressesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesOK %s", 200, payload)
}

func (o *GetAllKubernetesClusterIngressesOK) GetPayload() []*models.KubernetesK8sIngressInfo {
	return o.Payload
}

func (o *GetAllKubernetesClusterIngressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllKubernetesClusterIngressesBadRequest creates a GetAllKubernetesClusterIngressesBadRequest with default headers values
func NewGetAllKubernetesClusterIngressesBadRequest() *GetAllKubernetesClusterIngressesBadRequest {
	return &GetAllKubernetesClusterIngressesBadRequest{}
}

/*
GetAllKubernetesClusterIngressesBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetAllKubernetesClusterIngressesBadRequest struct {
}

// IsSuccess returns true when this get all kubernetes cluster ingresses bad request response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes cluster ingresses bad request response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses bad request response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes cluster ingresses bad request response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes cluster ingresses bad request response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get all kubernetes cluster ingresses bad request response
func (o *GetAllKubernetesClusterIngressesBadRequest) Code() int {
	return 400
}

func (o *GetAllKubernetesClusterIngressesBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesBadRequest", 400)
}

func (o *GetAllKubernetesClusterIngressesBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesBadRequest", 400)
}

func (o *GetAllKubernetesClusterIngressesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesClusterIngressesUnauthorized creates a GetAllKubernetesClusterIngressesUnauthorized with default headers values
func NewGetAllKubernetesClusterIngressesUnauthorized() *GetAllKubernetesClusterIngressesUnauthorized {
	return &GetAllKubernetesClusterIngressesUnauthorized{}
}

/*
GetAllKubernetesClusterIngressesUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetAllKubernetesClusterIngressesUnauthorized struct {
}

// IsSuccess returns true when this get all kubernetes cluster ingresses unauthorized response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes cluster ingresses unauthorized response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses unauthorized response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes cluster ingresses unauthorized response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes cluster ingresses unauthorized response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all kubernetes cluster ingresses unauthorized response
func (o *GetAllKubernetesClusterIngressesUnauthorized) Code() int {
	return 401
}

func (o *GetAllKubernetesClusterIngressesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesUnauthorized", 401)
}

func (o *GetAllKubernetesClusterIngressesUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesUnauthorized", 401)
}

func (o *GetAllKubernetesClusterIngressesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesClusterIngressesForbidden creates a GetAllKubernetesClusterIngressesForbidden with default headers values
func NewGetAllKubernetesClusterIngressesForbidden() *GetAllKubernetesClusterIngressesForbidden {
	return &GetAllKubernetesClusterIngressesForbidden{}
}

/*
GetAllKubernetesClusterIngressesForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetAllKubernetesClusterIngressesForbidden struct {
}

// IsSuccess returns true when this get all kubernetes cluster ingresses forbidden response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes cluster ingresses forbidden response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses forbidden response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes cluster ingresses forbidden response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes cluster ingresses forbidden response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all kubernetes cluster ingresses forbidden response
func (o *GetAllKubernetesClusterIngressesForbidden) Code() int {
	return 403
}

func (o *GetAllKubernetesClusterIngressesForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesForbidden", 403)
}

func (o *GetAllKubernetesClusterIngressesForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesForbidden", 403)
}

func (o *GetAllKubernetesClusterIngressesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesClusterIngressesNotFound creates a GetAllKubernetesClusterIngressesNotFound with default headers values
func NewGetAllKubernetesClusterIngressesNotFound() *GetAllKubernetesClusterIngressesNotFound {
	return &GetAllKubernetesClusterIngressesNotFound{}
}

/*
GetAllKubernetesClusterIngressesNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier.
*/
type GetAllKubernetesClusterIngressesNotFound struct {
}

// IsSuccess returns true when this get all kubernetes cluster ingresses not found response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes cluster ingresses not found response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses not found response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes cluster ingresses not found response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes cluster ingresses not found response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all kubernetes cluster ingresses not found response
func (o *GetAllKubernetesClusterIngressesNotFound) Code() int {
	return 404
}

func (o *GetAllKubernetesClusterIngressesNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesNotFound", 404)
}

func (o *GetAllKubernetesClusterIngressesNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesNotFound", 404)
}

func (o *GetAllKubernetesClusterIngressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesClusterIngressesInternalServerError creates a GetAllKubernetesClusterIngressesInternalServerError with default headers values
func NewGetAllKubernetesClusterIngressesInternalServerError() *GetAllKubernetesClusterIngressesInternalServerError {
	return &GetAllKubernetesClusterIngressesInternalServerError{}
}

/*
GetAllKubernetesClusterIngressesInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve ingresses.
*/
type GetAllKubernetesClusterIngressesInternalServerError struct {
}

// IsSuccess returns true when this get all kubernetes cluster ingresses internal server error response has a 2xx status code
func (o *GetAllKubernetesClusterIngressesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes cluster ingresses internal server error response has a 3xx status code
func (o *GetAllKubernetesClusterIngressesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes cluster ingresses internal server error response has a 4xx status code
func (o *GetAllKubernetesClusterIngressesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all kubernetes cluster ingresses internal server error response has a 5xx status code
func (o *GetAllKubernetesClusterIngressesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all kubernetes cluster ingresses internal server error response a status code equal to that given
func (o *GetAllKubernetesClusterIngressesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all kubernetes cluster ingresses internal server error response
func (o *GetAllKubernetesClusterIngressesInternalServerError) Code() int {
	return 500
}

func (o *GetAllKubernetesClusterIngressesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesInternalServerError", 500)
}

func (o *GetAllKubernetesClusterIngressesInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/ingresses][%d] getAllKubernetesClusterIngressesInternalServerError", 500)
}

func (o *GetAllKubernetesClusterIngressesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
