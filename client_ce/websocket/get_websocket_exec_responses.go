// Code generated by go-swagger; DO NOT EDIT.

package websocket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWebsocketExecReader is a Reader for the GetWebsocketExec structure.
type GetWebsocketExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsocketExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsocketExecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebsocketExecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetWebsocketExecConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebsocketExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebsocketExecOK creates a GetWebsocketExecOK with default headers values
func NewGetWebsocketExecOK() *GetWebsocketExecOK {
	return &GetWebsocketExecOK{}
}

/*
GetWebsocketExecOK describes a response with status code 200, with default header values.

GetWebsocketExecOK get websocket exec o k
*/
type GetWebsocketExecOK struct {
}

// IsSuccess returns true when this get websocket exec o k response has a 2xx status code
func (o *GetWebsocketExecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get websocket exec o k response has a 3xx status code
func (o *GetWebsocketExecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket exec o k response has a 4xx status code
func (o *GetWebsocketExecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get websocket exec o k response has a 5xx status code
func (o *GetWebsocketExecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket exec o k response a status code equal to that given
func (o *GetWebsocketExecOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebsocketExecOK) Error() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecOK ", 200)
}

func (o *GetWebsocketExecOK) String() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecOK ", 200)
}

func (o *GetWebsocketExecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketExecBadRequest creates a GetWebsocketExecBadRequest with default headers values
func NewGetWebsocketExecBadRequest() *GetWebsocketExecBadRequest {
	return &GetWebsocketExecBadRequest{}
}

/*
GetWebsocketExecBadRequest describes a response with status code 400, with default header values.

GetWebsocketExecBadRequest get websocket exec bad request
*/
type GetWebsocketExecBadRequest struct {
}

// IsSuccess returns true when this get websocket exec bad request response has a 2xx status code
func (o *GetWebsocketExecBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket exec bad request response has a 3xx status code
func (o *GetWebsocketExecBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket exec bad request response has a 4xx status code
func (o *GetWebsocketExecBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get websocket exec bad request response has a 5xx status code
func (o *GetWebsocketExecBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket exec bad request response a status code equal to that given
func (o *GetWebsocketExecBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWebsocketExecBadRequest) Error() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecBadRequest ", 400)
}

func (o *GetWebsocketExecBadRequest) String() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecBadRequest ", 400)
}

func (o *GetWebsocketExecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketExecConflict creates a GetWebsocketExecConflict with default headers values
func NewGetWebsocketExecConflict() *GetWebsocketExecConflict {
	return &GetWebsocketExecConflict{}
}

/*
GetWebsocketExecConflict describes a response with status code 409, with default header values.

GetWebsocketExecConflict get websocket exec conflict
*/
type GetWebsocketExecConflict struct {
}

// IsSuccess returns true when this get websocket exec conflict response has a 2xx status code
func (o *GetWebsocketExecConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket exec conflict response has a 3xx status code
func (o *GetWebsocketExecConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket exec conflict response has a 4xx status code
func (o *GetWebsocketExecConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get websocket exec conflict response has a 5xx status code
func (o *GetWebsocketExecConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket exec conflict response a status code equal to that given
func (o *GetWebsocketExecConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetWebsocketExecConflict) Error() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecConflict ", 409)
}

func (o *GetWebsocketExecConflict) String() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecConflict ", 409)
}

func (o *GetWebsocketExecConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketExecInternalServerError creates a GetWebsocketExecInternalServerError with default headers values
func NewGetWebsocketExecInternalServerError() *GetWebsocketExecInternalServerError {
	return &GetWebsocketExecInternalServerError{}
}

/*
GetWebsocketExecInternalServerError describes a response with status code 500, with default header values.

GetWebsocketExecInternalServerError get websocket exec internal server error
*/
type GetWebsocketExecInternalServerError struct {
}

// IsSuccess returns true when this get websocket exec internal server error response has a 2xx status code
func (o *GetWebsocketExecInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket exec internal server error response has a 3xx status code
func (o *GetWebsocketExecInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket exec internal server error response has a 4xx status code
func (o *GetWebsocketExecInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get websocket exec internal server error response has a 5xx status code
func (o *GetWebsocketExecInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get websocket exec internal server error response a status code equal to that given
func (o *GetWebsocketExecInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWebsocketExecInternalServerError) Error() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecInternalServerError ", 500)
}

func (o *GetWebsocketExecInternalServerError) String() string {
	return fmt.Sprintf("[GET /websocket/exec][%d] getWebsocketExecInternalServerError ", 500)
}

func (o *GetWebsocketExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
