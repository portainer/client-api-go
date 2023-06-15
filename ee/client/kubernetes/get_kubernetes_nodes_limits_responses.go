// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// GetKubernetesNodesLimitsReader is a Reader for the GetKubernetesNodesLimits structure.
type GetKubernetesNodesLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesNodesLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesNodesLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesNodesLimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKubernetesNodesLimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesNodesLimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesNodesLimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesNodesLimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKubernetesNodesLimitsOK creates a GetKubernetesNodesLimitsOK with default headers values
func NewGetKubernetesNodesLimitsOK() *GetKubernetesNodesLimitsOK {
	return &GetKubernetesNodesLimitsOK{}
}

/*
GetKubernetesNodesLimitsOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesNodesLimitsOK struct {
	Payload models.PortainereeK8sNodesLimits
}

// IsSuccess returns true when this get kubernetes nodes limits o k response has a 2xx status code
func (o *GetKubernetesNodesLimitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes nodes limits o k response has a 3xx status code
func (o *GetKubernetesNodesLimitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits o k response has a 4xx status code
func (o *GetKubernetesNodesLimitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes nodes limits o k response has a 5xx status code
func (o *GetKubernetesNodesLimitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes nodes limits o k response a status code equal to that given
func (o *GetKubernetesNodesLimitsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKubernetesNodesLimitsOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesNodesLimitsOK) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesNodesLimitsOK) GetPayload() models.PortainereeK8sNodesLimits {
	return o.Payload
}

func (o *GetKubernetesNodesLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesNodesLimitsBadRequest creates a GetKubernetesNodesLimitsBadRequest with default headers values
func NewGetKubernetesNodesLimitsBadRequest() *GetKubernetesNodesLimitsBadRequest {
	return &GetKubernetesNodesLimitsBadRequest{}
}

/*
GetKubernetesNodesLimitsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetKubernetesNodesLimitsBadRequest struct {
}

// IsSuccess returns true when this get kubernetes nodes limits bad request response has a 2xx status code
func (o *GetKubernetesNodesLimitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes nodes limits bad request response has a 3xx status code
func (o *GetKubernetesNodesLimitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits bad request response has a 4xx status code
func (o *GetKubernetesNodesLimitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes nodes limits bad request response has a 5xx status code
func (o *GetKubernetesNodesLimitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes nodes limits bad request response a status code equal to that given
func (o *GetKubernetesNodesLimitsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKubernetesNodesLimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsBadRequest ", 400)
}

func (o *GetKubernetesNodesLimitsBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsBadRequest ", 400)
}

func (o *GetKubernetesNodesLimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesNodesLimitsUnauthorized creates a GetKubernetesNodesLimitsUnauthorized with default headers values
func NewGetKubernetesNodesLimitsUnauthorized() *GetKubernetesNodesLimitsUnauthorized {
	return &GetKubernetesNodesLimitsUnauthorized{}
}

/*
GetKubernetesNodesLimitsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetKubernetesNodesLimitsUnauthorized struct {
}

// IsSuccess returns true when this get kubernetes nodes limits unauthorized response has a 2xx status code
func (o *GetKubernetesNodesLimitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes nodes limits unauthorized response has a 3xx status code
func (o *GetKubernetesNodesLimitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits unauthorized response has a 4xx status code
func (o *GetKubernetesNodesLimitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes nodes limits unauthorized response has a 5xx status code
func (o *GetKubernetesNodesLimitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes nodes limits unauthorized response a status code equal to that given
func (o *GetKubernetesNodesLimitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKubernetesNodesLimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsUnauthorized ", 401)
}

func (o *GetKubernetesNodesLimitsUnauthorized) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsUnauthorized ", 401)
}

func (o *GetKubernetesNodesLimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesNodesLimitsForbidden creates a GetKubernetesNodesLimitsForbidden with default headers values
func NewGetKubernetesNodesLimitsForbidden() *GetKubernetesNodesLimitsForbidden {
	return &GetKubernetesNodesLimitsForbidden{}
}

/*
GetKubernetesNodesLimitsForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type GetKubernetesNodesLimitsForbidden struct {
}

// IsSuccess returns true when this get kubernetes nodes limits forbidden response has a 2xx status code
func (o *GetKubernetesNodesLimitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes nodes limits forbidden response has a 3xx status code
func (o *GetKubernetesNodesLimitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits forbidden response has a 4xx status code
func (o *GetKubernetesNodesLimitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes nodes limits forbidden response has a 5xx status code
func (o *GetKubernetesNodesLimitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes nodes limits forbidden response a status code equal to that given
func (o *GetKubernetesNodesLimitsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKubernetesNodesLimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsForbidden ", 403)
}

func (o *GetKubernetesNodesLimitsForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsForbidden ", 403)
}

func (o *GetKubernetesNodesLimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesNodesLimitsNotFound creates a GetKubernetesNodesLimitsNotFound with default headers values
func NewGetKubernetesNodesLimitsNotFound() *GetKubernetesNodesLimitsNotFound {
	return &GetKubernetesNodesLimitsNotFound{}
}

/*
GetKubernetesNodesLimitsNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) not found
*/
type GetKubernetesNodesLimitsNotFound struct {
}

// IsSuccess returns true when this get kubernetes nodes limits not found response has a 2xx status code
func (o *GetKubernetesNodesLimitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes nodes limits not found response has a 3xx status code
func (o *GetKubernetesNodesLimitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits not found response has a 4xx status code
func (o *GetKubernetesNodesLimitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes nodes limits not found response has a 5xx status code
func (o *GetKubernetesNodesLimitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes nodes limits not found response a status code equal to that given
func (o *GetKubernetesNodesLimitsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKubernetesNodesLimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsNotFound ", 404)
}

func (o *GetKubernetesNodesLimitsNotFound) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsNotFound ", 404)
}

func (o *GetKubernetesNodesLimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesNodesLimitsInternalServerError creates a GetKubernetesNodesLimitsInternalServerError with default headers values
func NewGetKubernetesNodesLimitsInternalServerError() *GetKubernetesNodesLimitsInternalServerError {
	return &GetKubernetesNodesLimitsInternalServerError{}
}

/*
GetKubernetesNodesLimitsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetKubernetesNodesLimitsInternalServerError struct {
}

// IsSuccess returns true when this get kubernetes nodes limits internal server error response has a 2xx status code
func (o *GetKubernetesNodesLimitsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes nodes limits internal server error response has a 3xx status code
func (o *GetKubernetesNodesLimitsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes nodes limits internal server error response has a 4xx status code
func (o *GetKubernetesNodesLimitsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes nodes limits internal server error response has a 5xx status code
func (o *GetKubernetesNodesLimitsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kubernetes nodes limits internal server error response a status code equal to that given
func (o *GetKubernetesNodesLimitsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKubernetesNodesLimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsInternalServerError ", 500)
}

func (o *GetKubernetesNodesLimitsInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/nodes_limits][%d] getKubernetesNodesLimitsInternalServerError ", 500)
}

func (o *GetKubernetesNodesLimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
