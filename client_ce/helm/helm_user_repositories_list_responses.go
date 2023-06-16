// Code generated by go-swagger; DO NOT EDIT.

package helm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// HelmUserRepositoriesListReader is a Reader for the HelmUserRepositoriesList structure.
type HelmUserRepositoriesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelmUserRepositoriesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelmUserRepositoriesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelmUserRepositoriesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHelmUserRepositoriesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelmUserRepositoriesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHelmUserRepositoriesListOK creates a HelmUserRepositoriesListOK with default headers values
func NewHelmUserRepositoriesListOK() *HelmUserRepositoriesListOK {
	return &HelmUserRepositoriesListOK{}
}

/*
HelmUserRepositoriesListOK describes a response with status code 200, with default header values.

Success
*/
type HelmUserRepositoriesListOK struct {
	Payload *models_ce.HelmHelmUserRepositoryResponse
}

// IsSuccess returns true when this helm user repositories list o k response has a 2xx status code
func (o *HelmUserRepositoriesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this helm user repositories list o k response has a 3xx status code
func (o *HelmUserRepositoriesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repositories list o k response has a 4xx status code
func (o *HelmUserRepositoriesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm user repositories list o k response has a 5xx status code
func (o *HelmUserRepositoriesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repositories list o k response a status code equal to that given
func (o *HelmUserRepositoriesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *HelmUserRepositoriesListOK) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListOK  %+v", 200, o.Payload)
}

func (o *HelmUserRepositoriesListOK) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListOK  %+v", 200, o.Payload)
}

func (o *HelmUserRepositoriesListOK) GetPayload() *models_ce.HelmHelmUserRepositoryResponse {
	return o.Payload
}

func (o *HelmUserRepositoriesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.HelmHelmUserRepositoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelmUserRepositoriesListBadRequest creates a HelmUserRepositoriesListBadRequest with default headers values
func NewHelmUserRepositoriesListBadRequest() *HelmUserRepositoriesListBadRequest {
	return &HelmUserRepositoriesListBadRequest{}
}

/*
HelmUserRepositoriesListBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type HelmUserRepositoriesListBadRequest struct {
}

// IsSuccess returns true when this helm user repositories list bad request response has a 2xx status code
func (o *HelmUserRepositoriesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repositories list bad request response has a 3xx status code
func (o *HelmUserRepositoriesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repositories list bad request response has a 4xx status code
func (o *HelmUserRepositoriesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm user repositories list bad request response has a 5xx status code
func (o *HelmUserRepositoriesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repositories list bad request response a status code equal to that given
func (o *HelmUserRepositoriesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *HelmUserRepositoriesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListBadRequest ", 400)
}

func (o *HelmUserRepositoriesListBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListBadRequest ", 400)
}

func (o *HelmUserRepositoriesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmUserRepositoriesListForbidden creates a HelmUserRepositoriesListForbidden with default headers values
func NewHelmUserRepositoriesListForbidden() *HelmUserRepositoriesListForbidden {
	return &HelmUserRepositoriesListForbidden{}
}

/*
HelmUserRepositoriesListForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type HelmUserRepositoriesListForbidden struct {
}

// IsSuccess returns true when this helm user repositories list forbidden response has a 2xx status code
func (o *HelmUserRepositoriesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repositories list forbidden response has a 3xx status code
func (o *HelmUserRepositoriesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repositories list forbidden response has a 4xx status code
func (o *HelmUserRepositoriesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm user repositories list forbidden response has a 5xx status code
func (o *HelmUserRepositoriesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this helm user repositories list forbidden response a status code equal to that given
func (o *HelmUserRepositoriesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HelmUserRepositoriesListForbidden) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListForbidden ", 403)
}

func (o *HelmUserRepositoriesListForbidden) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListForbidden ", 403)
}

func (o *HelmUserRepositoriesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmUserRepositoriesListInternalServerError creates a HelmUserRepositoriesListInternalServerError with default headers values
func NewHelmUserRepositoriesListInternalServerError() *HelmUserRepositoriesListInternalServerError {
	return &HelmUserRepositoriesListInternalServerError{}
}

/*
HelmUserRepositoriesListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type HelmUserRepositoriesListInternalServerError struct {
}

// IsSuccess returns true when this helm user repositories list internal server error response has a 2xx status code
func (o *HelmUserRepositoriesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm user repositories list internal server error response has a 3xx status code
func (o *HelmUserRepositoriesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm user repositories list internal server error response has a 4xx status code
func (o *HelmUserRepositoriesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm user repositories list internal server error response has a 5xx status code
func (o *HelmUserRepositoriesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this helm user repositories list internal server error response a status code equal to that given
func (o *HelmUserRepositoriesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HelmUserRepositoriesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListInternalServerError ", 500)
}

func (o *HelmUserRepositoriesListInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm/repositories][%d] helmUserRepositoriesListInternalServerError ", 500)
}

func (o *HelmUserRepositoriesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
