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
)

// AgentVersionsReader is a Reader for the AgentVersions structure.
type AgentVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAgentVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAgentVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /edge_update_schedules/agent_versions] AgentVersions", response, response.Code())
	}
}

// NewAgentVersionsOK creates a AgentVersionsOK with default headers values
func NewAgentVersionsOK() *AgentVersionsOK {
	return &AgentVersionsOK{}
}

/*
AgentVersionsOK describes a response with status code 200, with default header values.

OK
*/
type AgentVersionsOK struct {
	Payload []string
}

// IsSuccess returns true when this agent versions o k response has a 2xx status code
func (o *AgentVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent versions o k response has a 3xx status code
func (o *AgentVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent versions o k response has a 4xx status code
func (o *AgentVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent versions o k response has a 5xx status code
func (o *AgentVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent versions o k response a status code equal to that given
func (o *AgentVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the agent versions o k response
func (o *AgentVersionsOK) Code() int {
	return 200
}

func (o *AgentVersionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsOK %s", 200, payload)
}

func (o *AgentVersionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsOK %s", 200, payload)
}

func (o *AgentVersionsOK) GetPayload() []string {
	return o.Payload
}

func (o *AgentVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentVersionsBadRequest creates a AgentVersionsBadRequest with default headers values
func NewAgentVersionsBadRequest() *AgentVersionsBadRequest {
	return &AgentVersionsBadRequest{}
}

/*
AgentVersionsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type AgentVersionsBadRequest struct {
}

// IsSuccess returns true when this agent versions bad request response has a 2xx status code
func (o *AgentVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this agent versions bad request response has a 3xx status code
func (o *AgentVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent versions bad request response has a 4xx status code
func (o *AgentVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this agent versions bad request response has a 5xx status code
func (o *AgentVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this agent versions bad request response a status code equal to that given
func (o *AgentVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the agent versions bad request response
func (o *AgentVersionsBadRequest) Code() int {
	return 400
}

func (o *AgentVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsBadRequest", 400)
}

func (o *AgentVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsBadRequest", 400)
}

func (o *AgentVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAgentVersionsInternalServerError creates a AgentVersionsInternalServerError with default headers values
func NewAgentVersionsInternalServerError() *AgentVersionsInternalServerError {
	return &AgentVersionsInternalServerError{}
}

/*
AgentVersionsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AgentVersionsInternalServerError struct {
}

// IsSuccess returns true when this agent versions internal server error response has a 2xx status code
func (o *AgentVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this agent versions internal server error response has a 3xx status code
func (o *AgentVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent versions internal server error response has a 4xx status code
func (o *AgentVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent versions internal server error response has a 5xx status code
func (o *AgentVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this agent versions internal server error response a status code equal to that given
func (o *AgentVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the agent versions internal server error response
func (o *AgentVersionsInternalServerError) Code() int {
	return 500
}

func (o *AgentVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsInternalServerError", 500)
}

func (o *AgentVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules/agent_versions][%d] agentVersionsInternalServerError", 500)
}

func (o *AgentVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
