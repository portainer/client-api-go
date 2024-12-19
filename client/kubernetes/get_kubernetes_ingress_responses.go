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

	"github.com/portainer/client-api-go/v2/models"
)

// GetKubernetesIngressReader is a Reader for the GetKubernetesIngress structure.
type GetKubernetesIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesIngressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesIngressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesIngressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesIngressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}] GetKubernetesIngress", response, response.Code())
	}
}

// NewGetKubernetesIngressOK creates a GetKubernetesIngressOK with default headers values
func NewGetKubernetesIngressOK() *GetKubernetesIngressOK {
	return &GetKubernetesIngressOK{}
}

/*
GetKubernetesIngressOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesIngressOK struct {
	Payload *models.KubernetesK8sIngressInfo
}

// IsSuccess returns true when this get kubernetes ingress o k response has a 2xx status code
func (o *GetKubernetesIngressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes ingress o k response has a 3xx status code
func (o *GetKubernetesIngressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress o k response has a 4xx status code
func (o *GetKubernetesIngressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes ingress o k response has a 5xx status code
func (o *GetKubernetesIngressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes ingress o k response a status code equal to that given
func (o *GetKubernetesIngressOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kubernetes ingress o k response
func (o *GetKubernetesIngressOK) Code() int {
	return 200
}

func (o *GetKubernetesIngressOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressOK %s", 200, payload)
}

func (o *GetKubernetesIngressOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressOK %s", 200, payload)
}

func (o *GetKubernetesIngressOK) GetPayload() *models.KubernetesK8sIngressInfo {
	return o.Payload
}

func (o *GetKubernetesIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesK8sIngressInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesIngressBadRequest creates a GetKubernetesIngressBadRequest with default headers values
func NewGetKubernetesIngressBadRequest() *GetKubernetesIngressBadRequest {
	return &GetKubernetesIngressBadRequest{}
}

/*
GetKubernetesIngressBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetKubernetesIngressBadRequest struct {
}

// IsSuccess returns true when this get kubernetes ingress bad request response has a 2xx status code
func (o *GetKubernetesIngressBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes ingress bad request response has a 3xx status code
func (o *GetKubernetesIngressBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress bad request response has a 4xx status code
func (o *GetKubernetesIngressBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes ingress bad request response has a 5xx status code
func (o *GetKubernetesIngressBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes ingress bad request response a status code equal to that given
func (o *GetKubernetesIngressBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get kubernetes ingress bad request response
func (o *GetKubernetesIngressBadRequest) Code() int {
	return 400
}

func (o *GetKubernetesIngressBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressBadRequest", 400)
}

func (o *GetKubernetesIngressBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressBadRequest", 400)
}

func (o *GetKubernetesIngressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesIngressUnauthorized creates a GetKubernetesIngressUnauthorized with default headers values
func NewGetKubernetesIngressUnauthorized() *GetKubernetesIngressUnauthorized {
	return &GetKubernetesIngressUnauthorized{}
}

/*
GetKubernetesIngressUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type GetKubernetesIngressUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes ingress unauthorized response has a 2xx status code
func (o *GetKubernetesIngressUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes ingress unauthorized response has a 3xx status code
func (o *GetKubernetesIngressUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress unauthorized response has a 4xx status code
func (o *GetKubernetesIngressUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes ingress unauthorized response has a 5xx status code
func (o *GetKubernetesIngressUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes ingress unauthorized response a status code equal to that given
func (o *GetKubernetesIngressUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kubernetes ingress unauthorized response
func (o *GetKubernetesIngressUnauthorized) Code() int {
	return 401
}

func (o *GetKubernetesIngressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressUnauthorized", 401)
}

func (o *GetKubernetesIngressUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressUnauthorized", 401)
}

func (o *GetKubernetesIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesIngressForbidden creates a GetKubernetesIngressForbidden with default headers values
func NewGetKubernetesIngressForbidden() *GetKubernetesIngressForbidden {
	return &GetKubernetesIngressForbidden{}
}

/*
GetKubernetesIngressForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type GetKubernetesIngressForbidden struct {
}

// IsSuccess returns true when this get kubernetes ingress forbidden response has a 2xx status code
func (o *GetKubernetesIngressForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes ingress forbidden response has a 3xx status code
func (o *GetKubernetesIngressForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress forbidden response has a 4xx status code
func (o *GetKubernetesIngressForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes ingress forbidden response has a 5xx status code
func (o *GetKubernetesIngressForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes ingress forbidden response a status code equal to that given
func (o *GetKubernetesIngressForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kubernetes ingress forbidden response
func (o *GetKubernetesIngressForbidden) Code() int {
	return 403
}

func (o *GetKubernetesIngressForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressForbidden", 403)
}

func (o *GetKubernetesIngressForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressForbidden", 403)
}

func (o *GetKubernetesIngressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesIngressNotFound creates a GetKubernetesIngressNotFound with default headers values
func NewGetKubernetesIngressNotFound() *GetKubernetesIngressNotFound {
	return &GetKubernetesIngressNotFound{}
}

/*
GetKubernetesIngressNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier or unable to find an ingress with the specified name.
*/
type GetKubernetesIngressNotFound struct {
}

// IsSuccess returns true when this get kubernetes ingress not found response has a 2xx status code
func (o *GetKubernetesIngressNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes ingress not found response has a 3xx status code
func (o *GetKubernetesIngressNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress not found response has a 4xx status code
func (o *GetKubernetesIngressNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes ingress not found response has a 5xx status code
func (o *GetKubernetesIngressNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes ingress not found response a status code equal to that given
func (o *GetKubernetesIngressNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get kubernetes ingress not found response
func (o *GetKubernetesIngressNotFound) Code() int {
	return 404
}

func (o *GetKubernetesIngressNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressNotFound", 404)
}

func (o *GetKubernetesIngressNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressNotFound", 404)
}

func (o *GetKubernetesIngressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesIngressInternalServerError creates a GetKubernetesIngressInternalServerError with default headers values
func NewGetKubernetesIngressInternalServerError() *GetKubernetesIngressInternalServerError {
	return &GetKubernetesIngressInternalServerError{}
}

/*
GetKubernetesIngressInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve an ingress.
*/
type GetKubernetesIngressInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes ingress internal server error response has a 2xx status code
func (o *GetKubernetesIngressInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes ingress internal server error response has a 3xx status code
func (o *GetKubernetesIngressInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes ingress internal server error response has a 4xx status code
func (o *GetKubernetesIngressInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes ingress internal server error response has a 5xx status code
func (o *GetKubernetesIngressInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes ingress internal server error response a status code equal to that given
func (o *GetKubernetesIngressInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get kubernetes ingress internal server error response
func (o *GetKubernetesIngressInternalServerError) Code() int {
	return 500
}

func (o *GetKubernetesIngressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressInternalServerError", 500)
}

func (o *GetKubernetesIngressInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/namespaces/{namespace}/ingresses/{ingress}][%d] getKubernetesIngressInternalServerError", 500)
}

func (o *GetKubernetesIngressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
