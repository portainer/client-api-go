// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamMembershipDeleteReader is a Reader for the TeamMembershipDelete structure.
type TeamMembershipDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTeamMembershipDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamMembershipDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamMembershipDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTeamMembershipDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamMembershipDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /team_memberships/{id}] TeamMembershipDelete", response, response.Code())
	}
}

// NewTeamMembershipDeleteNoContent creates a TeamMembershipDeleteNoContent with default headers values
func NewTeamMembershipDeleteNoContent() *TeamMembershipDeleteNoContent {
	return &TeamMembershipDeleteNoContent{}
}

/*
TeamMembershipDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type TeamMembershipDeleteNoContent struct {
}

// IsSuccess returns true when this team membership delete no content response has a 2xx status code
func (o *TeamMembershipDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team membership delete no content response has a 3xx status code
func (o *TeamMembershipDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership delete no content response has a 4xx status code
func (o *TeamMembershipDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership delete no content response has a 5xx status code
func (o *TeamMembershipDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership delete no content response a status code equal to that given
func (o *TeamMembershipDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the team membership delete no content response
func (o *TeamMembershipDeleteNoContent) Code() int {
	return 204
}

func (o *TeamMembershipDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteNoContent", 204)
}

func (o *TeamMembershipDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteNoContent", 204)
}

func (o *TeamMembershipDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipDeleteBadRequest creates a TeamMembershipDeleteBadRequest with default headers values
func NewTeamMembershipDeleteBadRequest() *TeamMembershipDeleteBadRequest {
	return &TeamMembershipDeleteBadRequest{}
}

/*
TeamMembershipDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamMembershipDeleteBadRequest struct {
}

// IsSuccess returns true when this team membership delete bad request response has a 2xx status code
func (o *TeamMembershipDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership delete bad request response has a 3xx status code
func (o *TeamMembershipDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership delete bad request response has a 4xx status code
func (o *TeamMembershipDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership delete bad request response has a 5xx status code
func (o *TeamMembershipDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership delete bad request response a status code equal to that given
func (o *TeamMembershipDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the team membership delete bad request response
func (o *TeamMembershipDeleteBadRequest) Code() int {
	return 400
}

func (o *TeamMembershipDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteBadRequest", 400)
}

func (o *TeamMembershipDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteBadRequest", 400)
}

func (o *TeamMembershipDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipDeleteForbidden creates a TeamMembershipDeleteForbidden with default headers values
func NewTeamMembershipDeleteForbidden() *TeamMembershipDeleteForbidden {
	return &TeamMembershipDeleteForbidden{}
}

/*
TeamMembershipDeleteForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type TeamMembershipDeleteForbidden struct {
}

// IsSuccess returns true when this team membership delete forbidden response has a 2xx status code
func (o *TeamMembershipDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership delete forbidden response has a 3xx status code
func (o *TeamMembershipDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership delete forbidden response has a 4xx status code
func (o *TeamMembershipDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership delete forbidden response has a 5xx status code
func (o *TeamMembershipDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership delete forbidden response a status code equal to that given
func (o *TeamMembershipDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the team membership delete forbidden response
func (o *TeamMembershipDeleteForbidden) Code() int {
	return 403
}

func (o *TeamMembershipDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteForbidden", 403)
}

func (o *TeamMembershipDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteForbidden", 403)
}

func (o *TeamMembershipDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipDeleteNotFound creates a TeamMembershipDeleteNotFound with default headers values
func NewTeamMembershipDeleteNotFound() *TeamMembershipDeleteNotFound {
	return &TeamMembershipDeleteNotFound{}
}

/*
TeamMembershipDeleteNotFound describes a response with status code 404, with default header values.

TeamMembership not found
*/
type TeamMembershipDeleteNotFound struct {
}

// IsSuccess returns true when this team membership delete not found response has a 2xx status code
func (o *TeamMembershipDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership delete not found response has a 3xx status code
func (o *TeamMembershipDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership delete not found response has a 4xx status code
func (o *TeamMembershipDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this team membership delete not found response has a 5xx status code
func (o *TeamMembershipDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this team membership delete not found response a status code equal to that given
func (o *TeamMembershipDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the team membership delete not found response
func (o *TeamMembershipDeleteNotFound) Code() int {
	return 404
}

func (o *TeamMembershipDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteNotFound", 404)
}

func (o *TeamMembershipDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteNotFound", 404)
}

func (o *TeamMembershipDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipDeleteInternalServerError creates a TeamMembershipDeleteInternalServerError with default headers values
func NewTeamMembershipDeleteInternalServerError() *TeamMembershipDeleteInternalServerError {
	return &TeamMembershipDeleteInternalServerError{}
}

/*
TeamMembershipDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamMembershipDeleteInternalServerError struct {
}

// IsSuccess returns true when this team membership delete internal server error response has a 2xx status code
func (o *TeamMembershipDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team membership delete internal server error response has a 3xx status code
func (o *TeamMembershipDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team membership delete internal server error response has a 4xx status code
func (o *TeamMembershipDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team membership delete internal server error response has a 5xx status code
func (o *TeamMembershipDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team membership delete internal server error response a status code equal to that given
func (o *TeamMembershipDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the team membership delete internal server error response
func (o *TeamMembershipDeleteInternalServerError) Code() int {
	return 500
}

func (o *TeamMembershipDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteInternalServerError", 500)
}

func (o *TeamMembershipDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /team_memberships/{id}][%d] teamMembershipDeleteInternalServerError", 500)
}

func (o *TeamMembershipDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
