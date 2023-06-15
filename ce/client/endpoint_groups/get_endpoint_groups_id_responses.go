// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// GetEndpointGroupsIDReader is a Reader for the GetEndpointGroupsID structure.
type GetEndpointGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointGroupsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointGroupsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEndpointGroupsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointGroupsIDOK creates a GetEndpointGroupsIDOK with default headers values
func NewGetEndpointGroupsIDOK() *GetEndpointGroupsIDOK {
	return &GetEndpointGroupsIDOK{}
}

/*
GetEndpointGroupsIDOK describes a response with status code 200, with default header values.

Success
*/
type GetEndpointGroupsIDOK struct {
	Payload *models.PortainerEndpointGroup
}

// IsSuccess returns true when this get endpoint groups Id o k response has a 2xx status code
func (o *GetEndpointGroupsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint groups Id o k response has a 3xx status code
func (o *GetEndpointGroupsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint groups Id o k response has a 4xx status code
func (o *GetEndpointGroupsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint groups Id o k response has a 5xx status code
func (o *GetEndpointGroupsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint groups Id o k response a status code equal to that given
func (o *GetEndpointGroupsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEndpointGroupsIDOK) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GetEndpointGroupsIDOK) String() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GetEndpointGroupsIDOK) GetPayload() *models.PortainerEndpointGroup {
	return o.Payload
}

func (o *GetEndpointGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEndpointGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointGroupsIDBadRequest creates a GetEndpointGroupsIDBadRequest with default headers values
func NewGetEndpointGroupsIDBadRequest() *GetEndpointGroupsIDBadRequest {
	return &GetEndpointGroupsIDBadRequest{}
}

/*
GetEndpointGroupsIDBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetEndpointGroupsIDBadRequest struct {
}

// IsSuccess returns true when this get endpoint groups Id bad request response has a 2xx status code
func (o *GetEndpointGroupsIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint groups Id bad request response has a 3xx status code
func (o *GetEndpointGroupsIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint groups Id bad request response has a 4xx status code
func (o *GetEndpointGroupsIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint groups Id bad request response has a 5xx status code
func (o *GetEndpointGroupsIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint groups Id bad request response a status code equal to that given
func (o *GetEndpointGroupsIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetEndpointGroupsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdBadRequest ", 400)
}

func (o *GetEndpointGroupsIDBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdBadRequest ", 400)
}

func (o *GetEndpointGroupsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointGroupsIDNotFound creates a GetEndpointGroupsIDNotFound with default headers values
func NewGetEndpointGroupsIDNotFound() *GetEndpointGroupsIDNotFound {
	return &GetEndpointGroupsIDNotFound{}
}

/*
GetEndpointGroupsIDNotFound describes a response with status code 404, with default header values.

EndpointGroup not found
*/
type GetEndpointGroupsIDNotFound struct {
}

// IsSuccess returns true when this get endpoint groups Id not found response has a 2xx status code
func (o *GetEndpointGroupsIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint groups Id not found response has a 3xx status code
func (o *GetEndpointGroupsIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint groups Id not found response has a 4xx status code
func (o *GetEndpointGroupsIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint groups Id not found response has a 5xx status code
func (o *GetEndpointGroupsIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint groups Id not found response a status code equal to that given
func (o *GetEndpointGroupsIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEndpointGroupsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdNotFound ", 404)
}

func (o *GetEndpointGroupsIDNotFound) String() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdNotFound ", 404)
}

func (o *GetEndpointGroupsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointGroupsIDInternalServerError creates a GetEndpointGroupsIDInternalServerError with default headers values
func NewGetEndpointGroupsIDInternalServerError() *GetEndpointGroupsIDInternalServerError {
	return &GetEndpointGroupsIDInternalServerError{}
}

/*
GetEndpointGroupsIDInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetEndpointGroupsIDInternalServerError struct {
}

// IsSuccess returns true when this get endpoint groups Id internal server error response has a 2xx status code
func (o *GetEndpointGroupsIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint groups Id internal server error response has a 3xx status code
func (o *GetEndpointGroupsIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint groups Id internal server error response has a 4xx status code
func (o *GetEndpointGroupsIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint groups Id internal server error response has a 5xx status code
func (o *GetEndpointGroupsIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get endpoint groups Id internal server error response a status code equal to that given
func (o *GetEndpointGroupsIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetEndpointGroupsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdInternalServerError ", 500)
}

func (o *GetEndpointGroupsIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoint_groups/{id}][%d] getEndpointGroupsIdInternalServerError ", 500)
}

func (o *GetEndpointGroupsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
