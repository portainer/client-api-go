// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// TeamUpdateReader is a Reader for the TeamUpdate structure.
type TeamUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTeamUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTeamUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamUpdateOK creates a TeamUpdateOK with default headers values
func NewTeamUpdateOK() *TeamUpdateOK {
	return &TeamUpdateOK{}
}

/*
TeamUpdateOK describes a response with status code 200, with default header values.

Success
*/
type TeamUpdateOK struct {
	Payload *models.PortainereeTeam
}

// IsSuccess returns true when this team update o k response has a 2xx status code
func (o *TeamUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team update o k response has a 3xx status code
func (o *TeamUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update o k response has a 4xx status code
func (o *TeamUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this team update o k response has a 5xx status code
func (o *TeamUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this team update o k response a status code equal to that given
func (o *TeamUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TeamUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateOK  %+v", 200, o.Payload)
}

func (o *TeamUpdateOK) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateOK  %+v", 200, o.Payload)
}

func (o *TeamUpdateOK) GetPayload() *models.PortainereeTeam {
	return o.Payload
}

func (o *TeamUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeTeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamUpdateNoContent creates a TeamUpdateNoContent with default headers values
func NewTeamUpdateNoContent() *TeamUpdateNoContent {
	return &TeamUpdateNoContent{}
}

/*
TeamUpdateNoContent describes a response with status code 204, with default header values.

Success
*/
type TeamUpdateNoContent struct {
}

// IsSuccess returns true when this team update no content response has a 2xx status code
func (o *TeamUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team update no content response has a 3xx status code
func (o *TeamUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update no content response has a 4xx status code
func (o *TeamUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this team update no content response has a 5xx status code
func (o *TeamUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this team update no content response a status code equal to that given
func (o *TeamUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TeamUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateNoContent ", 204)
}

func (o *TeamUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateNoContent ", 204)
}

func (o *TeamUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamUpdateBadRequest creates a TeamUpdateBadRequest with default headers values
func NewTeamUpdateBadRequest() *TeamUpdateBadRequest {
	return &TeamUpdateBadRequest{}
}

/*
TeamUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamUpdateBadRequest struct {
}

// IsSuccess returns true when this team update bad request response has a 2xx status code
func (o *TeamUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team update bad request response has a 3xx status code
func (o *TeamUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update bad request response has a 4xx status code
func (o *TeamUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team update bad request response has a 5xx status code
func (o *TeamUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team update bad request response a status code equal to that given
func (o *TeamUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TeamUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateBadRequest ", 400)
}

func (o *TeamUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateBadRequest ", 400)
}

func (o *TeamUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamUpdateForbidden creates a TeamUpdateForbidden with default headers values
func NewTeamUpdateForbidden() *TeamUpdateForbidden {
	return &TeamUpdateForbidden{}
}

/*
TeamUpdateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type TeamUpdateForbidden struct {
}

// IsSuccess returns true when this team update forbidden response has a 2xx status code
func (o *TeamUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team update forbidden response has a 3xx status code
func (o *TeamUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update forbidden response has a 4xx status code
func (o *TeamUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team update forbidden response has a 5xx status code
func (o *TeamUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team update forbidden response a status code equal to that given
func (o *TeamUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TeamUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateForbidden ", 403)
}

func (o *TeamUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateForbidden ", 403)
}

func (o *TeamUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamUpdateNotFound creates a TeamUpdateNotFound with default headers values
func NewTeamUpdateNotFound() *TeamUpdateNotFound {
	return &TeamUpdateNotFound{}
}

/*
TeamUpdateNotFound describes a response with status code 404, with default header values.

Team not found
*/
type TeamUpdateNotFound struct {
}

// IsSuccess returns true when this team update not found response has a 2xx status code
func (o *TeamUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team update not found response has a 3xx status code
func (o *TeamUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update not found response has a 4xx status code
func (o *TeamUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this team update not found response has a 5xx status code
func (o *TeamUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this team update not found response a status code equal to that given
func (o *TeamUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TeamUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateNotFound ", 404)
}

func (o *TeamUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateNotFound ", 404)
}

func (o *TeamUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamUpdateInternalServerError creates a TeamUpdateInternalServerError with default headers values
func NewTeamUpdateInternalServerError() *TeamUpdateInternalServerError {
	return &TeamUpdateInternalServerError{}
}

/*
TeamUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamUpdateInternalServerError struct {
}

// IsSuccess returns true when this team update internal server error response has a 2xx status code
func (o *TeamUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team update internal server error response has a 3xx status code
func (o *TeamUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team update internal server error response has a 4xx status code
func (o *TeamUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team update internal server error response has a 5xx status code
func (o *TeamUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team update internal server error response a status code equal to that given
func (o *TeamUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TeamUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateInternalServerError ", 500)
}

func (o *TeamUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /teams/{id}][%d] teamUpdateInternalServerError ", 500)
}

func (o *TeamUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
