// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// EdgeUpdateScheduleListReader is a Reader for the EdgeUpdateScheduleList structure.
type EdgeUpdateScheduleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpdateScheduleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeUpdateScheduleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeUpdateScheduleListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /edge_update_schedules] EdgeUpdateScheduleList", response, response.Code())
	}
}

// NewEdgeUpdateScheduleListOK creates a EdgeUpdateScheduleListOK with default headers values
func NewEdgeUpdateScheduleListOK() *EdgeUpdateScheduleListOK {
	return &EdgeUpdateScheduleListOK{}
}

/*
EdgeUpdateScheduleListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeUpdateScheduleListOK struct {
	Payload []*models.EdgeupdateschedulesDecoratedUpdateSchedule
}

// IsSuccess returns true when this edge update schedule list o k response has a 2xx status code
func (o *EdgeUpdateScheduleListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge update schedule list o k response has a 3xx status code
func (o *EdgeUpdateScheduleListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule list o k response has a 4xx status code
func (o *EdgeUpdateScheduleListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule list o k response has a 5xx status code
func (o *EdgeUpdateScheduleListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule list o k response a status code equal to that given
func (o *EdgeUpdateScheduleListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge update schedule list o k response
func (o *EdgeUpdateScheduleListOK) Code() int {
	return 200
}

func (o *EdgeUpdateScheduleListOK) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules][%d] edgeUpdateScheduleListOK  %+v", 200, o.Payload)
}

func (o *EdgeUpdateScheduleListOK) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules][%d] edgeUpdateScheduleListOK  %+v", 200, o.Payload)
}

func (o *EdgeUpdateScheduleListOK) GetPayload() []*models.EdgeupdateschedulesDecoratedUpdateSchedule {
	return o.Payload
}

func (o *EdgeUpdateScheduleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeUpdateScheduleListInternalServerError creates a EdgeUpdateScheduleListInternalServerError with default headers values
func NewEdgeUpdateScheduleListInternalServerError() *EdgeUpdateScheduleListInternalServerError {
	return &EdgeUpdateScheduleListInternalServerError{}
}

/*
EdgeUpdateScheduleListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeUpdateScheduleListInternalServerError struct {
}

// IsSuccess returns true when this edge update schedule list internal server error response has a 2xx status code
func (o *EdgeUpdateScheduleListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule list internal server error response has a 3xx status code
func (o *EdgeUpdateScheduleListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule list internal server error response has a 4xx status code
func (o *EdgeUpdateScheduleListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule list internal server error response has a 5xx status code
func (o *EdgeUpdateScheduleListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge update schedule list internal server error response a status code equal to that given
func (o *EdgeUpdateScheduleListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge update schedule list internal server error response
func (o *EdgeUpdateScheduleListInternalServerError) Code() int {
	return 500
}

func (o *EdgeUpdateScheduleListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules][%d] edgeUpdateScheduleListInternalServerError ", 500)
}

func (o *EdgeUpdateScheduleListInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules][%d] edgeUpdateScheduleListInternalServerError ", 500)
}

func (o *EdgeUpdateScheduleListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
