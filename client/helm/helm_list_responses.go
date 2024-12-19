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

	"github.com/portainer/client-api-go/v2/models"
)

// HelmListReader is a Reader for the HelmList structure.
type HelmListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelmListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelmListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelmListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHelmListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHelmListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelmListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /endpoints/{id}/kubernetes/helm] HelmList", response, response.Code())
	}
}

// NewHelmListOK creates a HelmListOK with default headers values
func NewHelmListOK() *HelmListOK {
	return &HelmListOK{}
}

/*
HelmListOK describes a response with status code 200, with default header values.

Success
*/
type HelmListOK struct {
	Payload []*models.ReleaseReleaseElement
}

// IsSuccess returns true when this helm list o k response has a 2xx status code
func (o *HelmListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this helm list o k response has a 3xx status code
func (o *HelmListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm list o k response has a 4xx status code
func (o *HelmListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm list o k response has a 5xx status code
func (o *HelmListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this helm list o k response a status code equal to that given
func (o *HelmListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the helm list o k response
func (o *HelmListOK) Code() int {
	return 200
}

func (o *HelmListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListOK %s", 200, payload)
}

func (o *HelmListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListOK %s", 200, payload)
}

func (o *HelmListOK) GetPayload() []*models.ReleaseReleaseElement {
	return o.Payload
}

func (o *HelmListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelmListBadRequest creates a HelmListBadRequest with default headers values
func NewHelmListBadRequest() *HelmListBadRequest {
	return &HelmListBadRequest{}
}

/*
HelmListBadRequest describes a response with status code 400, with default header values.

Invalid environment(endpoint) identifier
*/
type HelmListBadRequest struct {
}

// IsSuccess returns true when this helm list bad request response has a 2xx status code
func (o *HelmListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm list bad request response has a 3xx status code
func (o *HelmListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm list bad request response has a 4xx status code
func (o *HelmListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm list bad request response has a 5xx status code
func (o *HelmListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this helm list bad request response a status code equal to that given
func (o *HelmListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the helm list bad request response
func (o *HelmListBadRequest) Code() int {
	return 400
}

func (o *HelmListBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListBadRequest", 400)
}

func (o *HelmListBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListBadRequest", 400)
}

func (o *HelmListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmListUnauthorized creates a HelmListUnauthorized with default headers values
func NewHelmListUnauthorized() *HelmListUnauthorized {
	return &HelmListUnauthorized{}
}

/*
HelmListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type HelmListUnauthorized struct {
}

// IsSuccess returns true when this helm list unauthorized response has a 2xx status code
func (o *HelmListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm list unauthorized response has a 3xx status code
func (o *HelmListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm list unauthorized response has a 4xx status code
func (o *HelmListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm list unauthorized response has a 5xx status code
func (o *HelmListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this helm list unauthorized response a status code equal to that given
func (o *HelmListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the helm list unauthorized response
func (o *HelmListUnauthorized) Code() int {
	return 401
}

func (o *HelmListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListUnauthorized", 401)
}

func (o *HelmListUnauthorized) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListUnauthorized", 401)
}

func (o *HelmListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmListNotFound creates a HelmListNotFound with default headers values
func NewHelmListNotFound() *HelmListNotFound {
	return &HelmListNotFound{}
}

/*
HelmListNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) or ServiceAccount not found
*/
type HelmListNotFound struct {
}

// IsSuccess returns true when this helm list not found response has a 2xx status code
func (o *HelmListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm list not found response has a 3xx status code
func (o *HelmListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm list not found response has a 4xx status code
func (o *HelmListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this helm list not found response has a 5xx status code
func (o *HelmListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this helm list not found response a status code equal to that given
func (o *HelmListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the helm list not found response
func (o *HelmListNotFound) Code() int {
	return 404
}

func (o *HelmListNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListNotFound", 404)
}

func (o *HelmListNotFound) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListNotFound", 404)
}

func (o *HelmListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHelmListInternalServerError creates a HelmListInternalServerError with default headers values
func NewHelmListInternalServerError() *HelmListInternalServerError {
	return &HelmListInternalServerError{}
}

/*
HelmListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type HelmListInternalServerError struct {
}

// IsSuccess returns true when this helm list internal server error response has a 2xx status code
func (o *HelmListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this helm list internal server error response has a 3xx status code
func (o *HelmListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this helm list internal server error response has a 4xx status code
func (o *HelmListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this helm list internal server error response has a 5xx status code
func (o *HelmListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this helm list internal server error response a status code equal to that given
func (o *HelmListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the helm list internal server error response
func (o *HelmListInternalServerError) Code() int {
	return 500
}

func (o *HelmListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListInternalServerError", 500)
}

func (o *HelmListInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/kubernetes/helm][%d] helmListInternalServerError", 500)
}

func (o *HelmListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
