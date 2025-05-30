// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateK8sPodSecurityRuleReader is a Reader for the UpdateK8sPodSecurityRule structure.
type UpdateK8sPodSecurityRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateK8sPodSecurityRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateK8sPodSecurityRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateK8sPodSecurityRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateK8sPodSecurityRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateK8sPodSecurityRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /kubernetes/{environmentId}/opa] updateK8sPodSecurityRule", response, response.Code())
	}
}

// NewUpdateK8sPodSecurityRuleOK creates a UpdateK8sPodSecurityRuleOK with default headers values
func NewUpdateK8sPodSecurityRuleOK() *UpdateK8sPodSecurityRuleOK {
	return &UpdateK8sPodSecurityRuleOK{}
}

/*
UpdateK8sPodSecurityRuleOK describes a response with status code 200, with default header values.

Success
*/
type UpdateK8sPodSecurityRuleOK struct {
}

// IsSuccess returns true when this update k8s pod security rule o k response has a 2xx status code
func (o *UpdateK8sPodSecurityRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update k8s pod security rule o k response has a 3xx status code
func (o *UpdateK8sPodSecurityRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update k8s pod security rule o k response has a 4xx status code
func (o *UpdateK8sPodSecurityRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update k8s pod security rule o k response has a 5xx status code
func (o *UpdateK8sPodSecurityRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update k8s pod security rule o k response a status code equal to that given
func (o *UpdateK8sPodSecurityRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update k8s pod security rule o k response
func (o *UpdateK8sPodSecurityRuleOK) Code() int {
	return 200
}

func (o *UpdateK8sPodSecurityRuleOK) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleOK", 200)
}

func (o *UpdateK8sPodSecurityRuleOK) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleOK", 200)
}

func (o *UpdateK8sPodSecurityRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateK8sPodSecurityRuleBadRequest creates a UpdateK8sPodSecurityRuleBadRequest with default headers values
func NewUpdateK8sPodSecurityRuleBadRequest() *UpdateK8sPodSecurityRuleBadRequest {
	return &UpdateK8sPodSecurityRuleBadRequest{}
}

/*
UpdateK8sPodSecurityRuleBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UpdateK8sPodSecurityRuleBadRequest struct {
}

// IsSuccess returns true when this update k8s pod security rule bad request response has a 2xx status code
func (o *UpdateK8sPodSecurityRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update k8s pod security rule bad request response has a 3xx status code
func (o *UpdateK8sPodSecurityRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update k8s pod security rule bad request response has a 4xx status code
func (o *UpdateK8sPodSecurityRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update k8s pod security rule bad request response has a 5xx status code
func (o *UpdateK8sPodSecurityRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update k8s pod security rule bad request response a status code equal to that given
func (o *UpdateK8sPodSecurityRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update k8s pod security rule bad request response
func (o *UpdateK8sPodSecurityRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateK8sPodSecurityRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleBadRequest", 400)
}

func (o *UpdateK8sPodSecurityRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleBadRequest", 400)
}

func (o *UpdateK8sPodSecurityRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateK8sPodSecurityRuleNotFound creates a UpdateK8sPodSecurityRuleNotFound with default headers values
func NewUpdateK8sPodSecurityRuleNotFound() *UpdateK8sPodSecurityRuleNotFound {
	return &UpdateK8sPodSecurityRuleNotFound{}
}

/*
UpdateK8sPodSecurityRuleNotFound describes a response with status code 404, with default header values.

Pod Security Rule not found
*/
type UpdateK8sPodSecurityRuleNotFound struct {
}

// IsSuccess returns true when this update k8s pod security rule not found response has a 2xx status code
func (o *UpdateK8sPodSecurityRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update k8s pod security rule not found response has a 3xx status code
func (o *UpdateK8sPodSecurityRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update k8s pod security rule not found response has a 4xx status code
func (o *UpdateK8sPodSecurityRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update k8s pod security rule not found response has a 5xx status code
func (o *UpdateK8sPodSecurityRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update k8s pod security rule not found response a status code equal to that given
func (o *UpdateK8sPodSecurityRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update k8s pod security rule not found response
func (o *UpdateK8sPodSecurityRuleNotFound) Code() int {
	return 404
}

func (o *UpdateK8sPodSecurityRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleNotFound", 404)
}

func (o *UpdateK8sPodSecurityRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleNotFound", 404)
}

func (o *UpdateK8sPodSecurityRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateK8sPodSecurityRuleInternalServerError creates a UpdateK8sPodSecurityRuleInternalServerError with default headers values
func NewUpdateK8sPodSecurityRuleInternalServerError() *UpdateK8sPodSecurityRuleInternalServerError {
	return &UpdateK8sPodSecurityRuleInternalServerError{}
}

/*
UpdateK8sPodSecurityRuleInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateK8sPodSecurityRuleInternalServerError struct {
}

// IsSuccess returns true when this update k8s pod security rule internal server error response has a 2xx status code
func (o *UpdateK8sPodSecurityRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update k8s pod security rule internal server error response has a 3xx status code
func (o *UpdateK8sPodSecurityRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update k8s pod security rule internal server error response has a 4xx status code
func (o *UpdateK8sPodSecurityRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update k8s pod security rule internal server error response has a 5xx status code
func (o *UpdateK8sPodSecurityRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update k8s pod security rule internal server error response a status code equal to that given
func (o *UpdateK8sPodSecurityRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update k8s pod security rule internal server error response
func (o *UpdateK8sPodSecurityRuleInternalServerError) Code() int {
	return 500
}

func (o *UpdateK8sPodSecurityRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleInternalServerError", 500)
}

func (o *UpdateK8sPodSecurityRuleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /kubernetes/{environmentId}/opa][%d] updateK8sPodSecurityRuleInternalServerError", 500)
}

func (o *UpdateK8sPodSecurityRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
