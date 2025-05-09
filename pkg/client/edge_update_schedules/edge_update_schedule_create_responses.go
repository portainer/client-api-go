// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

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

// EdgeUpdateScheduleCreateReader is a Reader for the EdgeUpdateScheduleCreate structure.
type EdgeUpdateScheduleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpdateScheduleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeUpdateScheduleCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeUpdateScheduleCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeUpdateScheduleCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeUpdateScheduleCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeUpdateScheduleCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /edge_update_schedules] EdgeUpdateScheduleCreate", response, response.Code())
	}
}

// NewEdgeUpdateScheduleCreateOK creates a EdgeUpdateScheduleCreateOK with default headers values
func NewEdgeUpdateScheduleCreateOK() *EdgeUpdateScheduleCreateOK {
	return &EdgeUpdateScheduleCreateOK{}
}

/*
EdgeUpdateScheduleCreateOK describes a response with status code 200, with default header values.

Remote update procedure successfully created
*/
type EdgeUpdateScheduleCreateOK struct {
	Payload *models.TypesUpdateSchedule
}

// IsSuccess returns true when this edge update schedule create o k response has a 2xx status code
func (o *EdgeUpdateScheduleCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge update schedule create o k response has a 3xx status code
func (o *EdgeUpdateScheduleCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule create o k response has a 4xx status code
func (o *EdgeUpdateScheduleCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule create o k response has a 5xx status code
func (o *EdgeUpdateScheduleCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule create o k response a status code equal to that given
func (o *EdgeUpdateScheduleCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge update schedule create o k response
func (o *EdgeUpdateScheduleCreateOK) Code() int {
	return 200
}

func (o *EdgeUpdateScheduleCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateOK %s", 200, payload)
}

func (o *EdgeUpdateScheduleCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateOK %s", 200, payload)
}

func (o *EdgeUpdateScheduleCreateOK) GetPayload() *models.TypesUpdateSchedule {
	return o.Payload
}

func (o *EdgeUpdateScheduleCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TypesUpdateSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeUpdateScheduleCreateBadRequest creates a EdgeUpdateScheduleCreateBadRequest with default headers values
func NewEdgeUpdateScheduleCreateBadRequest() *EdgeUpdateScheduleCreateBadRequest {
	return &EdgeUpdateScheduleCreateBadRequest{}
}

/*
EdgeUpdateScheduleCreateBadRequest describes a response with status code 400, with default header values.

Invalid request payload, such as missing required fields or fields not meeting validation criteria.
*/
type EdgeUpdateScheduleCreateBadRequest struct {
}

// IsSuccess returns true when this edge update schedule create bad request response has a 2xx status code
func (o *EdgeUpdateScheduleCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule create bad request response has a 3xx status code
func (o *EdgeUpdateScheduleCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule create bad request response has a 4xx status code
func (o *EdgeUpdateScheduleCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge update schedule create bad request response has a 5xx status code
func (o *EdgeUpdateScheduleCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule create bad request response a status code equal to that given
func (o *EdgeUpdateScheduleCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge update schedule create bad request response
func (o *EdgeUpdateScheduleCreateBadRequest) Code() int {
	return 400
}

func (o *EdgeUpdateScheduleCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateBadRequest", 400)
}

func (o *EdgeUpdateScheduleCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateBadRequest", 400)
}

func (o *EdgeUpdateScheduleCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdateScheduleCreateForbidden creates a EdgeUpdateScheduleCreateForbidden with default headers values
func NewEdgeUpdateScheduleCreateForbidden() *EdgeUpdateScheduleCreateForbidden {
	return &EdgeUpdateScheduleCreateForbidden{}
}

/*
EdgeUpdateScheduleCreateForbidden describes a response with status code 403, with default header values.

Unauthorized access or operation not allowed.
*/
type EdgeUpdateScheduleCreateForbidden struct {
}

// IsSuccess returns true when this edge update schedule create forbidden response has a 2xx status code
func (o *EdgeUpdateScheduleCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule create forbidden response has a 3xx status code
func (o *EdgeUpdateScheduleCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule create forbidden response has a 4xx status code
func (o *EdgeUpdateScheduleCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge update schedule create forbidden response has a 5xx status code
func (o *EdgeUpdateScheduleCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule create forbidden response a status code equal to that given
func (o *EdgeUpdateScheduleCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge update schedule create forbidden response
func (o *EdgeUpdateScheduleCreateForbidden) Code() int {
	return 403
}

func (o *EdgeUpdateScheduleCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateForbidden", 403)
}

func (o *EdgeUpdateScheduleCreateForbidden) String() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateForbidden", 403)
}

func (o *EdgeUpdateScheduleCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdateScheduleCreateConflict creates a EdgeUpdateScheduleCreateConflict with default headers values
func NewEdgeUpdateScheduleCreateConflict() *EdgeUpdateScheduleCreateConflict {
	return &EdgeUpdateScheduleCreateConflict{}
}

/*
EdgeUpdateScheduleCreateConflict describes a response with status code 409, with default header values.

Edge update schedule name already in use.
*/
type EdgeUpdateScheduleCreateConflict struct {
}

// IsSuccess returns true when this edge update schedule create conflict response has a 2xx status code
func (o *EdgeUpdateScheduleCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule create conflict response has a 3xx status code
func (o *EdgeUpdateScheduleCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule create conflict response has a 4xx status code
func (o *EdgeUpdateScheduleCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge update schedule create conflict response has a 5xx status code
func (o *EdgeUpdateScheduleCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule create conflict response a status code equal to that given
func (o *EdgeUpdateScheduleCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge update schedule create conflict response
func (o *EdgeUpdateScheduleCreateConflict) Code() int {
	return 409
}

func (o *EdgeUpdateScheduleCreateConflict) Error() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateConflict", 409)
}

func (o *EdgeUpdateScheduleCreateConflict) String() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateConflict", 409)
}

func (o *EdgeUpdateScheduleCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdateScheduleCreateInternalServerError creates a EdgeUpdateScheduleCreateInternalServerError with default headers values
func NewEdgeUpdateScheduleCreateInternalServerError() *EdgeUpdateScheduleCreateInternalServerError {
	return &EdgeUpdateScheduleCreateInternalServerError{}
}

/*
EdgeUpdateScheduleCreateInternalServerError describes a response with status code 500, with default header values.

Server error occurred while attempting to create the remote update procedure.
*/
type EdgeUpdateScheduleCreateInternalServerError struct {
}

// IsSuccess returns true when this edge update schedule create internal server error response has a 2xx status code
func (o *EdgeUpdateScheduleCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule create internal server error response has a 3xx status code
func (o *EdgeUpdateScheduleCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule create internal server error response has a 4xx status code
func (o *EdgeUpdateScheduleCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule create internal server error response has a 5xx status code
func (o *EdgeUpdateScheduleCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge update schedule create internal server error response a status code equal to that given
func (o *EdgeUpdateScheduleCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge update schedule create internal server error response
func (o *EdgeUpdateScheduleCreateInternalServerError) Code() int {
	return 500
}

func (o *EdgeUpdateScheduleCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateInternalServerError", 500)
}

func (o *EdgeUpdateScheduleCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_update_schedules][%d] edgeUpdateScheduleCreateInternalServerError", 500)
}

func (o *EdgeUpdateScheduleCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
