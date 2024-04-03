// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// TeamListReader is a Reader for the TeamList structure.
type TeamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTeamListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams] TeamList", response, response.Code())
	}
}

// NewTeamListOK creates a TeamListOK with default headers values
func NewTeamListOK() *TeamListOK {
	return &TeamListOK{}
}

/*
TeamListOK describes a response with status code 200, with default header values.

Success
*/
type TeamListOK struct {
	Payload []*models.PortainereeTeam
}

// IsSuccess returns true when this team list o k response has a 2xx status code
func (o *TeamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this team list o k response has a 3xx status code
func (o *TeamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team list o k response has a 4xx status code
func (o *TeamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this team list o k response has a 5xx status code
func (o *TeamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this team list o k response a status code equal to that given
func (o *TeamListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the team list o k response
func (o *TeamListOK) Code() int {
	return 200
}

func (o *TeamListOK) Error() string {
	return fmt.Sprintf("[GET /teams][%d] teamListOK  %+v", 200, o.Payload)
}

func (o *TeamListOK) String() string {
	return fmt.Sprintf("[GET /teams][%d] teamListOK  %+v", 200, o.Payload)
}

func (o *TeamListOK) GetPayload() []*models.PortainereeTeam {
	return o.Payload
}

func (o *TeamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTeamListInternalServerError creates a TeamListInternalServerError with default headers values
func NewTeamListInternalServerError() *TeamListInternalServerError {
	return &TeamListInternalServerError{}
}

/*
TeamListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TeamListInternalServerError struct {
}

// IsSuccess returns true when this team list internal server error response has a 2xx status code
func (o *TeamListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this team list internal server error response has a 3xx status code
func (o *TeamListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this team list internal server error response has a 4xx status code
func (o *TeamListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this team list internal server error response has a 5xx status code
func (o *TeamListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this team list internal server error response a status code equal to that given
func (o *TeamListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the team list internal server error response
func (o *TeamListInternalServerError) Code() int {
	return 500
}

func (o *TeamListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams][%d] teamListInternalServerError ", 500)
}

func (o *TeamListInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams][%d] teamListInternalServerError ", 500)
}

func (o *TeamListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
