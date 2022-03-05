// Code generated by go-swagger; DO NOT EDIT.

package edge_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// EdgeTemplateListReader is a Reader for the EdgeTemplateList structure.
type EdgeTemplateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeTemplateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeTemplateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeTemplateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeTemplateListOK creates a EdgeTemplateListOK with default headers values
func NewEdgeTemplateListOK() *EdgeTemplateListOK {
	return &EdgeTemplateListOK{}
}

/* EdgeTemplateListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeTemplateListOK struct {
	Payload []*models.PortainerTemplate
}

func (o *EdgeTemplateListOK) Error() string {
	return fmt.Sprintf("[GET /edge_templates][%d] edgeTemplateListOK  %+v", 200, o.Payload)
}
func (o *EdgeTemplateListOK) GetPayload() []*models.PortainerTemplate {
	return o.Payload
}

func (o *EdgeTemplateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeTemplateListInternalServerError creates a EdgeTemplateListInternalServerError with default headers values
func NewEdgeTemplateListInternalServerError() *EdgeTemplateListInternalServerError {
	return &EdgeTemplateListInternalServerError{}
}

/* EdgeTemplateListInternalServerError describes a response with status code 500, with default header values.

EdgeTemplateListInternalServerError edge template list internal server error
*/
type EdgeTemplateListInternalServerError struct {
}

func (o *EdgeTemplateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_templates][%d] edgeTemplateListInternalServerError ", 500)
}

func (o *EdgeTemplateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
