// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Microk8sGetNodeStatusReader is a Reader for the Microk8sGetNodeStatus structure.
type Microk8sGetNodeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Microk8sGetNodeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMicrok8sGetNodeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMicrok8sGetNodeStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMicrok8sGetNodeStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMicrok8sGetNodeStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewMicrok8sGetNodeStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus] microk8sGetNodeStatus", response, response.Code())
	}
}

// NewMicrok8sGetNodeStatusOK creates a Microk8sGetNodeStatusOK with default headers values
func NewMicrok8sGetNodeStatusOK() *Microk8sGetNodeStatusOK {
	return &Microk8sGetNodeStatusOK{}
}

/*
Microk8sGetNodeStatusOK describes a response with status code 200, with default header values.

Success
*/
type Microk8sGetNodeStatusOK struct {
}

// IsSuccess returns true when this microk8s get node status o k response has a 2xx status code
func (o *Microk8sGetNodeStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this microk8s get node status o k response has a 3xx status code
func (o *Microk8sGetNodeStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get node status o k response has a 4xx status code
func (o *Microk8sGetNodeStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get node status o k response has a 5xx status code
func (o *Microk8sGetNodeStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get node status o k response a status code equal to that given
func (o *Microk8sGetNodeStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the microk8s get node status o k response
func (o *Microk8sGetNodeStatusOK) Code() int {
	return 200
}

func (o *Microk8sGetNodeStatusOK) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusOK", 200)
}

func (o *Microk8sGetNodeStatusOK) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusOK", 200)
}

func (o *Microk8sGetNodeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetNodeStatusBadRequest creates a Microk8sGetNodeStatusBadRequest with default headers values
func NewMicrok8sGetNodeStatusBadRequest() *Microk8sGetNodeStatusBadRequest {
	return &Microk8sGetNodeStatusBadRequest{}
}

/*
Microk8sGetNodeStatusBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type Microk8sGetNodeStatusBadRequest struct {
}

// IsSuccess returns true when this microk8s get node status bad request response has a 2xx status code
func (o *Microk8sGetNodeStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get node status bad request response has a 3xx status code
func (o *Microk8sGetNodeStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get node status bad request response has a 4xx status code
func (o *Microk8sGetNodeStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this microk8s get node status bad request response has a 5xx status code
func (o *Microk8sGetNodeStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get node status bad request response a status code equal to that given
func (o *Microk8sGetNodeStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the microk8s get node status bad request response
func (o *Microk8sGetNodeStatusBadRequest) Code() int {
	return 400
}

func (o *Microk8sGetNodeStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusBadRequest", 400)
}

func (o *Microk8sGetNodeStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusBadRequest", 400)
}

func (o *Microk8sGetNodeStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetNodeStatusForbidden creates a Microk8sGetNodeStatusForbidden with default headers values
func NewMicrok8sGetNodeStatusForbidden() *Microk8sGetNodeStatusForbidden {
	return &Microk8sGetNodeStatusForbidden{}
}

/*
Microk8sGetNodeStatusForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type Microk8sGetNodeStatusForbidden struct {
}

// IsSuccess returns true when this microk8s get node status forbidden response has a 2xx status code
func (o *Microk8sGetNodeStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get node status forbidden response has a 3xx status code
func (o *Microk8sGetNodeStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get node status forbidden response has a 4xx status code
func (o *Microk8sGetNodeStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this microk8s get node status forbidden response has a 5xx status code
func (o *Microk8sGetNodeStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get node status forbidden response a status code equal to that given
func (o *Microk8sGetNodeStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the microk8s get node status forbidden response
func (o *Microk8sGetNodeStatusForbidden) Code() int {
	return 403
}

func (o *Microk8sGetNodeStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusForbidden", 403)
}

func (o *Microk8sGetNodeStatusForbidden) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusForbidden", 403)
}

func (o *Microk8sGetNodeStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetNodeStatusInternalServerError creates a Microk8sGetNodeStatusInternalServerError with default headers values
func NewMicrok8sGetNodeStatusInternalServerError() *Microk8sGetNodeStatusInternalServerError {
	return &Microk8sGetNodeStatusInternalServerError{}
}

/*
Microk8sGetNodeStatusInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type Microk8sGetNodeStatusInternalServerError struct {
}

// IsSuccess returns true when this microk8s get node status internal server error response has a 2xx status code
func (o *Microk8sGetNodeStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get node status internal server error response has a 3xx status code
func (o *Microk8sGetNodeStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get node status internal server error response has a 4xx status code
func (o *Microk8sGetNodeStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get node status internal server error response has a 5xx status code
func (o *Microk8sGetNodeStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this microk8s get node status internal server error response a status code equal to that given
func (o *Microk8sGetNodeStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the microk8s get node status internal server error response
func (o *Microk8sGetNodeStatusInternalServerError) Code() int {
	return 500
}

func (o *Microk8sGetNodeStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusInternalServerError", 500)
}

func (o *Microk8sGetNodeStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusInternalServerError", 500)
}

func (o *Microk8sGetNodeStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetNodeStatusServiceUnavailable creates a Microk8sGetNodeStatusServiceUnavailable with default headers values
func NewMicrok8sGetNodeStatusServiceUnavailable() *Microk8sGetNodeStatusServiceUnavailable {
	return &Microk8sGetNodeStatusServiceUnavailable{}
}

/*
Microk8sGetNodeStatusServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type Microk8sGetNodeStatusServiceUnavailable struct {
}

// IsSuccess returns true when this microk8s get node status service unavailable response has a 2xx status code
func (o *Microk8sGetNodeStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get node status service unavailable response has a 3xx status code
func (o *Microk8sGetNodeStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get node status service unavailable response has a 4xx status code
func (o *Microk8sGetNodeStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get node status service unavailable response has a 5xx status code
func (o *Microk8sGetNodeStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this microk8s get node status service unavailable response a status code equal to that given
func (o *Microk8sGetNodeStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the microk8s get node status service unavailable response
func (o *Microk8sGetNodeStatusServiceUnavailable) Code() int {
	return 503
}

func (o *Microk8sGetNodeStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusServiceUnavailable", 503)
}

func (o *Microk8sGetNodeStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/nodes/nodestatus][%d] microk8sGetNodeStatusServiceUnavailable", 503)
}

func (o *Microk8sGetNodeStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
