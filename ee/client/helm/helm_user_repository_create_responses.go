// Code generated by go-swagger; DO NOT EDIT.

package helm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// HelmUserRepositoryCreateReader is a Reader for the HelmUserRepositoryCreate structure.
type HelmUserRepositoryCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelmUserRepositoryCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelmUserRepositoryCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelmUserRepositoryCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHelmUserRepositoryCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelmUserRepositoryCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHelmUserRepositoryCreateOK creates a HelmUserRepositoryCreateOK with default headers values
func NewHelmUserRepositoryCreateOK() *HelmUserRepositoryCreateOK {
	return &HelmUserRepositoryCreateOK{}
}

/*
HelmUserRepositoryCreateOK describes a response with status code 200, with default header values.

Success
*/
type HelmUserRepositoryCreateOK struct {
	Payload *models.PortainereeHelmUserRepository
}

// IsSuccess returns true when this helm user repository create o k response has a 2xx status code
func (o *HelmUserRepositoryCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this helm user repository create o k response has a 3xx status code
func (o *HelmUserRepositoryCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repository create o k response has a 4xx status code
func (o *HelmUserRepositoryCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm user repository create o k response has a 5xx status code
func (o *HelmUserRepositoryCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repository create o k response a status code equal to that given
func (o *HelmUserRepositoryCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *HelmUserRepositoryCreateOK) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateOK  %+v", 200, o.Payload)
}

func (o *HelmUserRepositoryCreateOK) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateOK  %+v", 200, o.Payload)
}

func (o *HelmUserRepositoryCreateOK) GetPayload() *models.PortainereeHelmUserRepository {
	return o.Payload
}

func (o *HelmUserRepositoryCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeHelmUserRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelmUserRepositoryCreateBadRequest creates a HelmUserRepositoryCreateBadRequest with default headers values
func NewHelmUserRepositoryCreateBadRequest() *HelmUserRepositoryCreateBadRequest {
	return &HelmUserRepositoryCreateBadRequest{}
}

/*
HelmUserRepositoryCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type HelmUserRepositoryCreateBadRequest struct {
}

// IsSuccess returns true when this helm user repository create bad request response has a 2xx status code
func (o *HelmUserRepositoryCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repository create bad request response has a 3xx status code
func (o *HelmUserRepositoryCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repository create bad request response has a 4xx status code
func (o *HelmUserRepositoryCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm user repository create bad request response has a 5xx status code
func (o *HelmUserRepositoryCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repository create bad request response a status code equal to that given
func (o *HelmUserRepositoryCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *HelmUserRepositoryCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateBadRequest ", 400)
}

func (o *HelmUserRepositoryCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateBadRequest ", 400)
}

func (o *HelmUserRepositoryCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmUserRepositoryCreateForbidden creates a HelmUserRepositoryCreateForbidden with default headers values
func NewHelmUserRepositoryCreateForbidden() *HelmUserRepositoryCreateForbidden {
	return &HelmUserRepositoryCreateForbidden{}
}

/*
HelmUserRepositoryCreateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type HelmUserRepositoryCreateForbidden struct {
}

// IsSuccess returns true when this helm user repository create forbidden response has a 2xx status code
func (o *HelmUserRepositoryCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repository create forbidden response has a 3xx status code
func (o *HelmUserRepositoryCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repository create forbidden response has a 4xx status code
func (o *HelmUserRepositoryCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm user repository create forbidden response has a 5xx status code
func (o *HelmUserRepositoryCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repository create forbidden response a status code equal to that given
func (o *HelmUserRepositoryCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HelmUserRepositoryCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateForbidden ", 403)
}

func (o *HelmUserRepositoryCreateForbidden) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateForbidden ", 403)
}

func (o *HelmUserRepositoryCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmUserRepositoryCreateInternalServerError creates a HelmUserRepositoryCreateInternalServerError with default headers values
func NewHelmUserRepositoryCreateInternalServerError() *HelmUserRepositoryCreateInternalServerError {
	return &HelmUserRepositoryCreateInternalServerError{}
}

/*
HelmUserRepositoryCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type HelmUserRepositoryCreateInternalServerError struct {
}

// IsSuccess returns true when this helm user repository create internal server error response has a 2xx status code
func (o *HelmUserRepositoryCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repository create internal server error response has a 3xx status code
func (o *HelmUserRepositoryCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repository create internal server error response has a 4xx status code
func (o *HelmUserRepositoryCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm user repository create internal server error response has a 5xx status code
func (o *HelmUserRepositoryCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this helm user repository create internal server error response a status code equal to that given
func (o *HelmUserRepositoryCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HelmUserRepositoryCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateInternalServerError ", 500)
}

func (o *HelmUserRepositoryCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoryCreateInternalServerError ", 500)
}

func (o *HelmUserRepositoryCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
