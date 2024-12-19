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

	"github.com/portainer/client-api-go/v2/models"
)

// TeamMembershipCreateReader is a Reader for the TeamMembershipCreate structure.
type TeamMembershipCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamMembershipCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewTeamMembershipCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamMembershipCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamMembershipCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTeamMembershipCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamMembershipCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /team_memberships] TeamMembershipCreate", response, response.Code())
	}
}

// NewTeamMembershipCreateOK creates a TeamMembershipCreateOK with default headers values
func NewTeamMembershipCreateOK() *TeamMembershipCreateOK {
	return &TeamMembershipCreateOK{}
}

/*
TeamMembershipCreateOK describes a response with status code 200, with default header values.

Success
*/
type TeamMembershipCreateOK struct {
	Payload *models.PortainerTeamMembership
}

// IsSuccess returns true when this team membership create o k response has a 2xx status code
func (o *TeamMembershipCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team membership create o k response has a 3xx status code
func (o *TeamMembershipCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create o k response has a 4xx status code
func (o *TeamMembershipCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership create o k response has a 5xx status code
func (o *TeamMembershipCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership create o k response a status code equal to that given
func (o *TeamMembershipCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the team membership create o k response
func (o *TeamMembershipCreateOK) Code() int {
	return 200
}

func (o *TeamMembershipCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateOK %s", 200, payload)
}

func (o *TeamMembershipCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateOK %s", 200, payload)
}

func (o *TeamMembershipCreateOK) GetPayload() *models.PortainerTeamMembership {
	return o.Payload
}

func (o *TeamMembershipCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerTeamMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipCreateNoContent creates a TeamMembershipCreateNoContent with default headers values
func NewTeamMembershipCreateNoContent() *TeamMembershipCreateNoContent {
	return &TeamMembershipCreateNoContent{}
}

/*
TeamMembershipCreateNoContent describes a response with status code 204, with default header values.

Success
*/
type TeamMembershipCreateNoContent struct {
}

// IsSuccess returns true when this team membership create no content response has a 2xx status code
func (o *TeamMembershipCreateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team membership create no content response has a 3xx status code
func (o *TeamMembershipCreateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create no content response has a 4xx status code
func (o *TeamMembershipCreateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership create no content response has a 5xx status code
func (o *TeamMembershipCreateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership create no content response a status code equal to that given
func (o *TeamMembershipCreateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the team membership create no content response
func (o *TeamMembershipCreateNoContent) Code() int {
	return 204
}

func (o *TeamMembershipCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateNoContent", 204)
}

func (o *TeamMembershipCreateNoContent) String() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateNoContent", 204)
}

func (o *TeamMembershipCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipCreateBadRequest creates a TeamMembershipCreateBadRequest with default headers values
func NewTeamMembershipCreateBadRequest() *TeamMembershipCreateBadRequest {
	return &TeamMembershipCreateBadRequest{}
}

/*
TeamMembershipCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamMembershipCreateBadRequest struct {
}

// IsSuccess returns true when this team membership create bad request response has a 2xx status code
func (o *TeamMembershipCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership create bad request response has a 3xx status code
func (o *TeamMembershipCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create bad request response has a 4xx status code
func (o *TeamMembershipCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership create bad request response has a 5xx status code
func (o *TeamMembershipCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership create bad request response a status code equal to that given
func (o *TeamMembershipCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the team membership create bad request response
func (o *TeamMembershipCreateBadRequest) Code() int {
	return 400
}

func (o *TeamMembershipCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateBadRequest", 400)
}

func (o *TeamMembershipCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateBadRequest", 400)
}

func (o *TeamMembershipCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipCreateForbidden creates a TeamMembershipCreateForbidden with default headers values
func NewTeamMembershipCreateForbidden() *TeamMembershipCreateForbidden {
	return &TeamMembershipCreateForbidden{}
}

/*
TeamMembershipCreateForbidden describes a response with status code 403, with default header values.

Permission denied to manage memberships
*/
type TeamMembershipCreateForbidden struct {
}

// IsSuccess returns true when this team membership create forbidden response has a 2xx status code
func (o *TeamMembershipCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership create forbidden response has a 3xx status code
func (o *TeamMembershipCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create forbidden response has a 4xx status code
func (o *TeamMembershipCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership create forbidden response has a 5xx status code
func (o *TeamMembershipCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership create forbidden response a status code equal to that given
func (o *TeamMembershipCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the team membership create forbidden response
func (o *TeamMembershipCreateForbidden) Code() int {
	return 403
}

func (o *TeamMembershipCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateForbidden", 403)
}

func (o *TeamMembershipCreateForbidden) String() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateForbidden", 403)
}

func (o *TeamMembershipCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipCreateConflict creates a TeamMembershipCreateConflict with default headers values
func NewTeamMembershipCreateConflict() *TeamMembershipCreateConflict {
	return &TeamMembershipCreateConflict{}
}

/*
TeamMembershipCreateConflict describes a response with status code 409, with default header values.

Team membership already registered
*/
type TeamMembershipCreateConflict struct {
}

// IsSuccess returns true when this team membership create conflict response has a 2xx status code
func (o *TeamMembershipCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership create conflict response has a 3xx status code
func (o *TeamMembershipCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create conflict response has a 4xx status code
func (o *TeamMembershipCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership create conflict response has a 5xx status code
func (o *TeamMembershipCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership create conflict response a status code equal to that given
func (o *TeamMembershipCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the team membership create conflict response
func (o *TeamMembershipCreateConflict) Code() int {
	return 409
}

func (o *TeamMembershipCreateConflict) Error() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateConflict", 409)
}

func (o *TeamMembershipCreateConflict) String() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateConflict", 409)
}

func (o *TeamMembershipCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipCreateInternalServerError creates a TeamMembershipCreateInternalServerError with default headers values
func NewTeamMembershipCreateInternalServerError() *TeamMembershipCreateInternalServerError {
	return &TeamMembershipCreateInternalServerError{}
}

/*
TeamMembershipCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamMembershipCreateInternalServerError struct {
}

// IsSuccess returns true when this team membership create internal server error response has a 2xx status code
func (o *TeamMembershipCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership create internal server error response has a 3xx status code
func (o *TeamMembershipCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership create internal server error response has a 4xx status code
func (o *TeamMembershipCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership create internal server error response has a 5xx status code
func (o *TeamMembershipCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team membership create internal server error response a status code equal to that given
func (o *TeamMembershipCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the team membership create internal server error response
func (o *TeamMembershipCreateInternalServerError) Code() int {
	return 500
}

func (o *TeamMembershipCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateInternalServerError", 500)
}

func (o *TeamMembershipCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /team_memberships][%d] teamMembershipCreateInternalServerError", 500)
}

func (o *TeamMembershipCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
