// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

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

// TeamMembershipUpdateReader is a Reader for the TeamMembershipUpdate structure.
type TeamMembershipUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamMembershipUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamMembershipUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamMembershipUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTeamMembershipUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamMembershipUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /team_memberships/{id}] TeamMembershipUpdate", response, response.Code())
	}
}

// NewTeamMembershipUpdateOK creates a TeamMembershipUpdateOK with default headers values
func NewTeamMembershipUpdateOK() *TeamMembershipUpdateOK {
	return &TeamMembershipUpdateOK{}
}

/*
TeamMembershipUpdateOK describes a response with status code 200, with default header values.

Success
*/
type TeamMembershipUpdateOK struct {
	Payload *models.PortainerTeamMembership
}

// IsSuccess returns true when this team membership update o k response has a 2xx status code
func (o *TeamMembershipUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team membership update o k response has a 3xx status code
func (o *TeamMembershipUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership update o k response has a 4xx status code
func (o *TeamMembershipUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership update o k response has a 5xx status code
func (o *TeamMembershipUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership update o k response a status code equal to that given
func (o *TeamMembershipUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the team membership update o k response
func (o *TeamMembershipUpdateOK) Code() int {
	return 200
}

func (o *TeamMembershipUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateOK %s", 200, payload)
}

func (o *TeamMembershipUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateOK %s", 200, payload)
}

func (o *TeamMembershipUpdateOK) GetPayload() *models.PortainerTeamMembership {
	return o.Payload
}

func (o *TeamMembershipUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerTeamMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipUpdateBadRequest creates a TeamMembershipUpdateBadRequest with default headers values
func NewTeamMembershipUpdateBadRequest() *TeamMembershipUpdateBadRequest {
	return &TeamMembershipUpdateBadRequest{}
}

/*
TeamMembershipUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamMembershipUpdateBadRequest struct {
}

// IsSuccess returns true when this team membership update bad request response has a 2xx status code
func (o *TeamMembershipUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership update bad request response has a 3xx status code
func (o *TeamMembershipUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership update bad request response has a 4xx status code
func (o *TeamMembershipUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership update bad request response has a 5xx status code
func (o *TeamMembershipUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership update bad request response a status code equal to that given
func (o *TeamMembershipUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the team membership update bad request response
func (o *TeamMembershipUpdateBadRequest) Code() int {
	return 400
}

func (o *TeamMembershipUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateBadRequest", 400)
}

func (o *TeamMembershipUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateBadRequest", 400)
}

func (o *TeamMembershipUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipUpdateForbidden creates a TeamMembershipUpdateForbidden with default headers values
func NewTeamMembershipUpdateForbidden() *TeamMembershipUpdateForbidden {
	return &TeamMembershipUpdateForbidden{}
}

/*
TeamMembershipUpdateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type TeamMembershipUpdateForbidden struct {
}

// IsSuccess returns true when this team membership update forbidden response has a 2xx status code
func (o *TeamMembershipUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership update forbidden response has a 3xx status code
func (o *TeamMembershipUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership update forbidden response has a 4xx status code
func (o *TeamMembershipUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership update forbidden response has a 5xx status code
func (o *TeamMembershipUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership update forbidden response a status code equal to that given
func (o *TeamMembershipUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the team membership update forbidden response
func (o *TeamMembershipUpdateForbidden) Code() int {
	return 403
}

func (o *TeamMembershipUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateForbidden", 403)
}

func (o *TeamMembershipUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateForbidden", 403)
}

func (o *TeamMembershipUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipUpdateNotFound creates a TeamMembershipUpdateNotFound with default headers values
func NewTeamMembershipUpdateNotFound() *TeamMembershipUpdateNotFound {
	return &TeamMembershipUpdateNotFound{}
}

/*
TeamMembershipUpdateNotFound describes a response with status code 404, with default header values.

TeamMembership not found
*/
type TeamMembershipUpdateNotFound struct {
}

// IsSuccess returns true when this team membership update not found response has a 2xx status code
func (o *TeamMembershipUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership update not found response has a 3xx status code
func (o *TeamMembershipUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership update not found response has a 4xx status code
func (o *TeamMembershipUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership update not found response has a 5xx status code
func (o *TeamMembershipUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership update not found response a status code equal to that given
func (o *TeamMembershipUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the team membership update not found response
func (o *TeamMembershipUpdateNotFound) Code() int {
	return 404
}

func (o *TeamMembershipUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateNotFound", 404)
}

func (o *TeamMembershipUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateNotFound", 404)
}

func (o *TeamMembershipUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipUpdateInternalServerError creates a TeamMembershipUpdateInternalServerError with default headers values
func NewTeamMembershipUpdateInternalServerError() *TeamMembershipUpdateInternalServerError {
	return &TeamMembershipUpdateInternalServerError{}
}

/*
TeamMembershipUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamMembershipUpdateInternalServerError struct {
}

// IsSuccess returns true when this team membership update internal server error response has a 2xx status code
func (o *TeamMembershipUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership update internal server error response has a 3xx status code
func (o *TeamMembershipUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership update internal server error response has a 4xx status code
func (o *TeamMembershipUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership update internal server error response has a 5xx status code
func (o *TeamMembershipUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team membership update internal server error response a status code equal to that given
func (o *TeamMembershipUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the team membership update internal server error response
func (o *TeamMembershipUpdateInternalServerError) Code() int {
	return 500
}

func (o *TeamMembershipUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateInternalServerError", 500)
}

func (o *TeamMembershipUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /team_memberships/{id}][%d] teamMembershipUpdateInternalServerError", 500)
}

func (o *TeamMembershipUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
