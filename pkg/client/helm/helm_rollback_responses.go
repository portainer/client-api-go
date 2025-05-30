// Code generated by go-swagger; DO NOT EDIT.

package helm

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

// HelmRollbackReader is a Reader for the HelmRollback structure.
type HelmRollbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelmRollbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelmRollbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelmRollbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHelmRollbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHelmRollbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHelmRollbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelmRollbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback] HelmRollback", response, response.Code())
	}
}

// NewHelmRollbackOK creates a HelmRollbackOK with default headers values
func NewHelmRollbackOK() *HelmRollbackOK {
	return &HelmRollbackOK{}
}

/*
HelmRollbackOK describes a response with status code 200, with default header values.

Success
*/
type HelmRollbackOK struct {
	Payload *models.ReleaseRelease
}

// IsSuccess returns true when this helm rollback o k response has a 2xx status code
func (o *HelmRollbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this helm rollback o k response has a 3xx status code
func (o *HelmRollbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback o k response has a 4xx status code
func (o *HelmRollbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm rollback o k response has a 5xx status code
func (o *HelmRollbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this helm rollback o k response a status code equal to that given
func (o *HelmRollbackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the helm rollback o k response
func (o *HelmRollbackOK) Code() int {
	return 200
}

func (o *HelmRollbackOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackOK %s", 200, payload)
}

func (o *HelmRollbackOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackOK %s", 200, payload)
}

func (o *HelmRollbackOK) GetPayload() *models.ReleaseRelease {
	return o.Payload
}

func (o *HelmRollbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelmRollbackBadRequest creates a HelmRollbackBadRequest with default headers values
func NewHelmRollbackBadRequest() *HelmRollbackBadRequest {
	return &HelmRollbackBadRequest{}
}

/*
HelmRollbackBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type HelmRollbackBadRequest struct {
}

// IsSuccess returns true when this helm rollback bad request response has a 2xx status code
func (o *HelmRollbackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm rollback bad request response has a 3xx status code
func (o *HelmRollbackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback bad request response has a 4xx status code
func (o *HelmRollbackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm rollback bad request response has a 5xx status code
func (o *HelmRollbackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this helm rollback bad request response a status code equal to that given
func (o *HelmRollbackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the helm rollback bad request response
func (o *HelmRollbackBadRequest) Code() int {
	return 400
}

func (o *HelmRollbackBadRequest) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackBadRequest", 400)
}

func (o *HelmRollbackBadRequest) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackBadRequest", 400)
}

func (o *HelmRollbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmRollbackUnauthorized creates a HelmRollbackUnauthorized with default headers values
func NewHelmRollbackUnauthorized() *HelmRollbackUnauthorized {
	return &HelmRollbackUnauthorized{}
}

/*
HelmRollbackUnauthorized describes a response with status code 401, with default header values.

Unauthorized access - the user is not authenticated or does not have the necessary permissions. Ensure that you have provided a valid API key or JWT token, and that you have the required permissions.
*/
type HelmRollbackUnauthorized struct {
}

// IsSuccess returns true when this helm rollback unauthorized response has a 2xx status code
func (o *HelmRollbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm rollback unauthorized response has a 3xx status code
func (o *HelmRollbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback unauthorized response has a 4xx status code
func (o *HelmRollbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm rollback unauthorized response has a 5xx status code
func (o *HelmRollbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this helm rollback unauthorized response a status code equal to that given
func (o *HelmRollbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the helm rollback unauthorized response
func (o *HelmRollbackUnauthorized) Code() int {
	return 401
}

func (o *HelmRollbackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackUnauthorized", 401)
}

func (o *HelmRollbackUnauthorized) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackUnauthorized", 401)
}

func (o *HelmRollbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmRollbackForbidden creates a HelmRollbackForbidden with default headers values
func NewHelmRollbackForbidden() *HelmRollbackForbidden {
	return &HelmRollbackForbidden{}
}

/*
HelmRollbackForbidden describes a response with status code 403, with default header values.

Permission denied - the user is authenticated but does not have the necessary permissions to access the requested resource or perform the specified operation. Check your user roles and permissions.
*/
type HelmRollbackForbidden struct {
}

// IsSuccess returns true when this helm rollback forbidden response has a 2xx status code
func (o *HelmRollbackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm rollback forbidden response has a 3xx status code
func (o *HelmRollbackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback forbidden response has a 4xx status code
func (o *HelmRollbackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm rollback forbidden response has a 5xx status code
func (o *HelmRollbackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this helm rollback forbidden response a status code equal to that given
func (o *HelmRollbackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the helm rollback forbidden response
func (o *HelmRollbackForbidden) Code() int {
	return 403
}

func (o *HelmRollbackForbidden) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackForbidden", 403)
}

func (o *HelmRollbackForbidden) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackForbidden", 403)
}

func (o *HelmRollbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmRollbackNotFound creates a HelmRollbackNotFound with default headers values
func NewHelmRollbackNotFound() *HelmRollbackNotFound {
	return &HelmRollbackNotFound{}
}

/*
HelmRollbackNotFound describes a response with status code 404, with default header values.

Unable to find an environment with the specified identifier or release name.
*/
type HelmRollbackNotFound struct {
}

// IsSuccess returns true when this helm rollback not found response has a 2xx status code
func (o *HelmRollbackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm rollback not found response has a 3xx status code
func (o *HelmRollbackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback not found response has a 4xx status code
func (o *HelmRollbackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm rollback not found response has a 5xx status code
func (o *HelmRollbackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this helm rollback not found response a status code equal to that given
func (o *HelmRollbackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the helm rollback not found response
func (o *HelmRollbackNotFound) Code() int {
	return 404
}

func (o *HelmRollbackNotFound) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackNotFound", 404)
}

func (o *HelmRollbackNotFound) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackNotFound", 404)
}

func (o *HelmRollbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmRollbackInternalServerError creates a HelmRollbackInternalServerError with default headers values
func NewHelmRollbackInternalServerError() *HelmRollbackInternalServerError {
	return &HelmRollbackInternalServerError{}
}

/*
HelmRollbackInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to rollback the release.
*/
type HelmRollbackInternalServerError struct {
}

// IsSuccess returns true when this helm rollback internal server error response has a 2xx status code
func (o *HelmRollbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm rollback internal server error response has a 3xx status code
func (o *HelmRollbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm rollback internal server error response has a 4xx status code
func (o *HelmRollbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm rollback internal server error response has a 5xx status code
func (o *HelmRollbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this helm rollback internal server error response a status code equal to that given
func (o *HelmRollbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the helm rollback internal server error response
func (o *HelmRollbackInternalServerError) Code() int {
	return 500
}

func (o *HelmRollbackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackInternalServerError", 500)
}

func (o *HelmRollbackInternalServerError) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/{release}/rollback][%d] helmRollbackInternalServerError", 500)
}

func (o *HelmRollbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
