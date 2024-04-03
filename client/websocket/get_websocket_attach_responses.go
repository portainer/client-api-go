// Code generated by go-swagger; DO NOT EDIT.

package websocket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWebsocketAttachReader is a Reader for the GetWebsocketAttach structure.
type GetWebsocketAttachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsocketAttachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebsocketAttachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebsocketAttachBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebsocketAttachForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebsocketAttachNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebsocketAttachInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /websocket/attach] GetWebsocketAttach", response, response.Code())
	}
}

// NewGetWebsocketAttachOK creates a GetWebsocketAttachOK with default headers values
func NewGetWebsocketAttachOK() *GetWebsocketAttachOK {
	return &GetWebsocketAttachOK{}
}

/*
GetWebsocketAttachOK describes a response with status code 200, with default header values.

OK
*/
type GetWebsocketAttachOK struct {
}

// IsSuccess returns true when this get websocket attach o k response has a 2xx status code
func (o *GetWebsocketAttachOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get websocket attach o k response has a 3xx status code
func (o *GetWebsocketAttachOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket attach o k response has a 4xx status code
func (o *GetWebsocketAttachOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get websocket attach o k response has a 5xx status code
func (o *GetWebsocketAttachOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket attach o k response a status code equal to that given
func (o *GetWebsocketAttachOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get websocket attach o k response
func (o *GetWebsocketAttachOK) Code() int {
	return 200
}

func (o *GetWebsocketAttachOK) Error() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachOK ", 200)
}

func (o *GetWebsocketAttachOK) String() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachOK ", 200)
}

func (o *GetWebsocketAttachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketAttachBadRequest creates a GetWebsocketAttachBadRequest with default headers values
func NewGetWebsocketAttachBadRequest() *GetWebsocketAttachBadRequest {
	return &GetWebsocketAttachBadRequest{}
}

/*
GetWebsocketAttachBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetWebsocketAttachBadRequest struct {
}

// IsSuccess returns true when this get websocket attach bad request response has a 2xx status code
func (o *GetWebsocketAttachBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket attach bad request response has a 3xx status code
func (o *GetWebsocketAttachBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket attach bad request response has a 4xx status code
func (o *GetWebsocketAttachBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get websocket attach bad request response has a 5xx status code
func (o *GetWebsocketAttachBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket attach bad request response a status code equal to that given
func (o *GetWebsocketAttachBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get websocket attach bad request response
func (o *GetWebsocketAttachBadRequest) Code() int {
	return 400
}

func (o *GetWebsocketAttachBadRequest) Error() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachBadRequest ", 400)
}

func (o *GetWebsocketAttachBadRequest) String() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachBadRequest ", 400)
}

func (o *GetWebsocketAttachBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketAttachForbidden creates a GetWebsocketAttachForbidden with default headers values
func NewGetWebsocketAttachForbidden() *GetWebsocketAttachForbidden {
	return &GetWebsocketAttachForbidden{}
}

/*
GetWebsocketAttachForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetWebsocketAttachForbidden struct {
}

// IsSuccess returns true when this get websocket attach forbidden response has a 2xx status code
func (o *GetWebsocketAttachForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket attach forbidden response has a 3xx status code
func (o *GetWebsocketAttachForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket attach forbidden response has a 4xx status code
func (o *GetWebsocketAttachForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get websocket attach forbidden response has a 5xx status code
func (o *GetWebsocketAttachForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket attach forbidden response a status code equal to that given
func (o *GetWebsocketAttachForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get websocket attach forbidden response
func (o *GetWebsocketAttachForbidden) Code() int {
	return 403
}

func (o *GetWebsocketAttachForbidden) Error() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachForbidden ", 403)
}

func (o *GetWebsocketAttachForbidden) String() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachForbidden ", 403)
}

func (o *GetWebsocketAttachForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketAttachNotFound creates a GetWebsocketAttachNotFound with default headers values
func NewGetWebsocketAttachNotFound() *GetWebsocketAttachNotFound {
	return &GetWebsocketAttachNotFound{}
}

/*
GetWebsocketAttachNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetWebsocketAttachNotFound struct {
}

// IsSuccess returns true when this get websocket attach not found response has a 2xx status code
func (o *GetWebsocketAttachNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket attach not found response has a 3xx status code
func (o *GetWebsocketAttachNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket attach not found response has a 4xx status code
func (o *GetWebsocketAttachNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get websocket attach not found response has a 5xx status code
func (o *GetWebsocketAttachNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get websocket attach not found response a status code equal to that given
func (o *GetWebsocketAttachNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get websocket attach not found response
func (o *GetWebsocketAttachNotFound) Code() int {
	return 404
}

func (o *GetWebsocketAttachNotFound) Error() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachNotFound ", 404)
}

func (o *GetWebsocketAttachNotFound) String() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachNotFound ", 404)
}

func (o *GetWebsocketAttachNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWebsocketAttachInternalServerError creates a GetWebsocketAttachInternalServerError with default headers values
func NewGetWebsocketAttachInternalServerError() *GetWebsocketAttachInternalServerError {
	return &GetWebsocketAttachInternalServerError{}
}

/*
GetWebsocketAttachInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetWebsocketAttachInternalServerError struct {
}

// IsSuccess returns true when this get websocket attach internal server error response has a 2xx status code
func (o *GetWebsocketAttachInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get websocket attach internal server error response has a 3xx status code
func (o *GetWebsocketAttachInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get websocket attach internal server error response has a 4xx status code
func (o *GetWebsocketAttachInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get websocket attach internal server error response has a 5xx status code
func (o *GetWebsocketAttachInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get websocket attach internal server error response a status code equal to that given
func (o *GetWebsocketAttachInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get websocket attach internal server error response
func (o *GetWebsocketAttachInternalServerError) Code() int {
	return 500
}

func (o *GetWebsocketAttachInternalServerError) Error() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachInternalServerError ", 500)
}

func (o *GetWebsocketAttachInternalServerError) String() string {
	return fmt.Sprintf("[GET /websocket/attach][%d] getWebsocketAttachInternalServerError ", 500)
}

func (o *GetWebsocketAttachInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
