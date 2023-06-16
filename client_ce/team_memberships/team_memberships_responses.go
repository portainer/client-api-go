// Code generated by go-swagger; DO NOT EDIT.

package team_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// TeamMembershipsReader is a Reader for the TeamMemberships structure.
type TeamMembershipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamMembershipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTeamMembershipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTeamMembershipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTeamMembershipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamMembershipsOK creates a TeamMembershipsOK with default headers values
func NewTeamMembershipsOK() *TeamMembershipsOK {
	return &TeamMembershipsOK{}
}

/*
TeamMembershipsOK describes a response with status code 200, with default header values.

Success
*/
type TeamMembershipsOK struct {
	Payload []*models_ce.PortainerTeamMembership
}

// IsSuccess returns true when this team memberships o k response has a 2xx status code
func (o *TeamMembershipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team memberships o k response has a 3xx status code
func (o *TeamMembershipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team memberships o k response has a 4xx status code
func (o *TeamMembershipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this team memberships o k response has a 5xx status code
func (o *TeamMembershipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this team memberships o k response a status code equal to that given
func (o *TeamMembershipsOK) IsCode(code int) bool {
	return code == 200
}

func (o *TeamMembershipsOK) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsOK  %+v", 200, o.Payload)
}

func (o *TeamMembershipsOK) String() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsOK  %+v", 200, o.Payload)
}

func (o *TeamMembershipsOK) GetPayload() []*models_ce.PortainerTeamMembership {
	return o.Payload
}

func (o *TeamMembershipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamMembershipsBadRequest creates a TeamMembershipsBadRequest with default headers values
func NewTeamMembershipsBadRequest() *TeamMembershipsBadRequest {
	return &TeamMembershipsBadRequest{}
}

/*
TeamMembershipsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TeamMembershipsBadRequest struct {
}

// IsSuccess returns true when this team memberships bad request response has a 2xx status code
func (o *TeamMembershipsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team memberships bad request response has a 3xx status code
func (o *TeamMembershipsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team memberships bad request response has a 4xx status code
func (o *TeamMembershipsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this team memberships bad request response has a 5xx status code
func (o *TeamMembershipsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this team memberships bad request response a status code equal to that given
func (o *TeamMembershipsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TeamMembershipsBadRequest) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsBadRequest ", 400)
}

func (o *TeamMembershipsBadRequest) String() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsBadRequest ", 400)
}

func (o *TeamMembershipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipsForbidden creates a TeamMembershipsForbidden with default headers values
func NewTeamMembershipsForbidden() *TeamMembershipsForbidden {
	return &TeamMembershipsForbidden{}
}

/*
TeamMembershipsForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type TeamMembershipsForbidden struct {
}

// IsSuccess returns true when this team memberships forbidden response has a 2xx status code
func (o *TeamMembershipsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team memberships forbidden response has a 3xx status code
func (o *TeamMembershipsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team memberships forbidden response has a 4xx status code
func (o *TeamMembershipsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this team memberships forbidden response has a 5xx status code
func (o *TeamMembershipsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this team memberships forbidden response a status code equal to that given
func (o *TeamMembershipsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TeamMembershipsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsForbidden ", 403)
}

func (o *TeamMembershipsForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsForbidden ", 403)
}

func (o *TeamMembershipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamMembershipsInternalServerError creates a TeamMembershipsInternalServerError with default headers values
func NewTeamMembershipsInternalServerError() *TeamMembershipsInternalServerError {
	return &TeamMembershipsInternalServerError{}
}

/*
TeamMembershipsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamMembershipsInternalServerError struct {
}

// IsSuccess returns true when this team memberships internal server error response has a 2xx status code
func (o *TeamMembershipsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team memberships internal server error response has a 3xx status code
func (o *TeamMembershipsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team memberships internal server error response has a 4xx status code
func (o *TeamMembershipsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team memberships internal server error response has a 5xx status code
func (o *TeamMembershipsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team memberships internal server error response a status code equal to that given
func (o *TeamMembershipsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TeamMembershipsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsInternalServerError ", 500)
}

func (o *TeamMembershipsInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams/{id}/memberships][%d] teamMembershipsInternalServerError ", 500)
}

func (o *TeamMembershipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
