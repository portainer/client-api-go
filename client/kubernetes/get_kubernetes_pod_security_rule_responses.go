// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// GetKubernetesPodSecurityRuleReader is a Reader for the GetKubernetesPodSecurityRule structure.
type GetKubernetesPodSecurityRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesPodSecurityRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesPodSecurityRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKubernetesPodSecurityRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKubernetesPodSecurityRuleOK creates a GetKubernetesPodSecurityRuleOK with default headers values
func NewGetKubernetesPodSecurityRuleOK() *GetKubernetesPodSecurityRuleOK {
	return &GetKubernetesPodSecurityRuleOK{}
}

/*
GetKubernetesPodSecurityRuleOK describes a response with status code 200, with default header values.

Success
*/
type GetKubernetesPodSecurityRuleOK struct {
	Payload *models.PodsecurityPodSecurityRule
}

// IsSuccess returns true when this get kubernetes pod security rule o k response has a 2xx status code
func (o *GetKubernetesPodSecurityRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes pod security rule o k response has a 3xx status code
func (o *GetKubernetesPodSecurityRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes pod security rule o k response has a 4xx status code
func (o *GetKubernetesPodSecurityRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes pod security rule o k response has a 5xx status code
func (o *GetKubernetesPodSecurityRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes pod security rule o k response a status code equal to that given
func (o *GetKubernetesPodSecurityRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKubernetesPodSecurityRuleOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{environmentId}/opa][%d] getKubernetesPodSecurityRuleOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesPodSecurityRuleOK) String() string {
	return fmt.Sprintf("[GET /kubernetes/{environmentId}/opa][%d] getKubernetesPodSecurityRuleOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesPodSecurityRuleOK) GetPayload() *models.PodsecurityPodSecurityRule {
	return o.Payload
}

func (o *GetKubernetesPodSecurityRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodsecurityPodSecurityRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesPodSecurityRuleBadRequest creates a GetKubernetesPodSecurityRuleBadRequest with default headers values
func NewGetKubernetesPodSecurityRuleBadRequest() *GetKubernetesPodSecurityRuleBadRequest {
	return &GetKubernetesPodSecurityRuleBadRequest{}
}

/*
GetKubernetesPodSecurityRuleBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type GetKubernetesPodSecurityRuleBadRequest struct {
}

// IsSuccess returns true when this get kubernetes pod security rule bad request response has a 2xx status code
func (o *GetKubernetesPodSecurityRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes pod security rule bad request response has a 3xx status code
func (o *GetKubernetesPodSecurityRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes pod security rule bad request response has a 4xx status code
func (o *GetKubernetesPodSecurityRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes pod security rule bad request response has a 5xx status code
func (o *GetKubernetesPodSecurityRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes pod security rule bad request response a status code equal to that given
func (o *GetKubernetesPodSecurityRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKubernetesPodSecurityRuleBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes/{environmentId}/opa][%d] getKubernetesPodSecurityRuleBadRequest ", 400)
}

func (o *GetKubernetesPodSecurityRuleBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes/{environmentId}/opa][%d] getKubernetesPodSecurityRuleBadRequest ", 400)
}

func (o *GetKubernetesPodSecurityRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
