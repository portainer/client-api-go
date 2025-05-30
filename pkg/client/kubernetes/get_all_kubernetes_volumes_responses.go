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

// GetAllKubernetesVolumesReader is a Reader for the GetAllKubernetesVolumes structure.
type GetAllKubernetesVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllKubernetesVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllKubernetesVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllKubernetesVolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllKubernetesVolumesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllKubernetesVolumesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes/{id}/volumes] GetAllKubernetesVolumes", response, response.Code())
	}
}

// NewGetAllKubernetesVolumesOK creates a GetAllKubernetesVolumesOK with default headers values
func NewGetAllKubernetesVolumesOK() *GetAllKubernetesVolumesOK {
	return &GetAllKubernetesVolumesOK{}
}

/*
GetAllKubernetesVolumesOK describes a response with status code 200, with default header values.

Success
*/
type GetAllKubernetesVolumesOK struct {
	Payload map[string]models.KubernetesK8sVolumeInfo
}

// IsSuccess returns true when this get all kubernetes volumes o k response has a 2xx status code
func (o *GetAllKubernetesVolumesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all kubernetes volumes o k response has a 3xx status code
func (o *GetAllKubernetesVolumesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes volumes o k response has a 4xx status code
func (o *GetAllKubernetesVolumesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all kubernetes volumes o k response has a 5xx status code
func (o *GetAllKubernetesVolumesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes volumes o k response a status code equal to that given
func (o *GetAllKubernetesVolumesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all kubernetes volumes o k response
func (o *GetAllKubernetesVolumesOK) Code() int {
	return 200
}

func (o *GetAllKubernetesVolumesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesOK %s", 200, payload)
}

func (o *GetAllKubernetesVolumesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesOK %s", 200, payload)
}

func (o *GetAllKubernetesVolumesOK) GetPayload() map[string]models.KubernetesK8sVolumeInfo {
	return o.Payload
}

func (o *GetAllKubernetesVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllKubernetesVolumesBadRequest creates a GetAllKubernetesVolumesBadRequest with default headers values
func NewGetAllKubernetesVolumesBadRequest() *GetAllKubernetesVolumesBadRequest {
	return &GetAllKubernetesVolumesBadRequest{}
}

/*
GetAllKubernetesVolumesBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type GetAllKubernetesVolumesBadRequest struct {
}

// IsSuccess returns true when this get all kubernetes volumes bad request response has a 2xx status code
func (o *GetAllKubernetesVolumesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes volumes bad request response has a 3xx status code
func (o *GetAllKubernetesVolumesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes volumes bad request response has a 4xx status code
func (o *GetAllKubernetesVolumesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes volumes bad request response has a 5xx status code
func (o *GetAllKubernetesVolumesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes volumes bad request response a status code equal to that given
func (o *GetAllKubernetesVolumesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get all kubernetes volumes bad request response
func (o *GetAllKubernetesVolumesBadRequest) Code() int {
	return 400
}

func (o *GetAllKubernetesVolumesBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesBadRequest", 400)
}

func (o *GetAllKubernetesVolumesBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesBadRequest", 400)
}

func (o *GetAllKubernetesVolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesVolumesForbidden creates a GetAllKubernetesVolumesForbidden with default headers values
func NewGetAllKubernetesVolumesForbidden() *GetAllKubernetesVolumesForbidden {
	return &GetAllKubernetesVolumesForbidden{}
}

/*
GetAllKubernetesVolumesForbidden describes a response with status code 403, with default header values.

Unauthorized access or operation not allowed.
*/
type GetAllKubernetesVolumesForbidden struct {
}

// IsSuccess returns true when this get all kubernetes volumes forbidden response has a 2xx status code
func (o *GetAllKubernetesVolumesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes volumes forbidden response has a 3xx status code
func (o *GetAllKubernetesVolumesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes volumes forbidden response has a 4xx status code
func (o *GetAllKubernetesVolumesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all kubernetes volumes forbidden response has a 5xx status code
func (o *GetAllKubernetesVolumesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all kubernetes volumes forbidden response a status code equal to that given
func (o *GetAllKubernetesVolumesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all kubernetes volumes forbidden response
func (o *GetAllKubernetesVolumesForbidden) Code() int {
	return 403
}

func (o *GetAllKubernetesVolumesForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesForbidden", 403)
}

func (o *GetAllKubernetesVolumesForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesForbidden", 403)
}

func (o *GetAllKubernetesVolumesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllKubernetesVolumesInternalServerError creates a GetAllKubernetesVolumesInternalServerError with default headers values
func NewGetAllKubernetesVolumesInternalServerError() *GetAllKubernetesVolumesInternalServerError {
	return &GetAllKubernetesVolumesInternalServerError{}
}

/*
GetAllKubernetesVolumesInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to retrieve kubernetes volumes.
*/
type GetAllKubernetesVolumesInternalServerError struct {
}

// IsSuccess returns true when this get all kubernetes volumes internal server error response has a 2xx status code
func (o *GetAllKubernetesVolumesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all kubernetes volumes internal server error response has a 3xx status code
func (o *GetAllKubernetesVolumesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all kubernetes volumes internal server error response has a 4xx status code
func (o *GetAllKubernetesVolumesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all kubernetes volumes internal server error response has a 5xx status code
func (o *GetAllKubernetesVolumesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all kubernetes volumes internal server error response a status code equal to that given
func (o *GetAllKubernetesVolumesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all kubernetes volumes internal server error response
func (o *GetAllKubernetesVolumesInternalServerError) Code() int {
	return 500
}

func (o *GetAllKubernetesVolumesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesInternalServerError", 500)
}

func (o *GetAllKubernetesVolumesInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes/{id}/volumes][%d] getAllKubernetesVolumesInternalServerError", 500)
}

func (o *GetAllKubernetesVolumesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
