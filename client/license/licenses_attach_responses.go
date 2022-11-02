// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// LicensesAttachReader is a Reader for the LicensesAttach structure.
type LicensesAttachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LicensesAttachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLicensesAttachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLicensesAttachOK creates a LicensesAttachOK with default headers values
func NewLicensesAttachOK() *LicensesAttachOK {
	return &LicensesAttachOK{}
}

/*
LicensesAttachOK describes a response with status code 200, with default header values.

Success license data will be in `body.Licenses`, Failures will be in `body.FailedKeys[key] = error`
*/
type LicensesAttachOK struct {
	Payload *models.LicensesAttachResponse
}

// IsSuccess returns true when this licenses attach o k response has a 2xx status code
func (o *LicensesAttachOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this licenses attach o k response has a 3xx status code
func (o *LicensesAttachOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this licenses attach o k response has a 4xx status code
func (o *LicensesAttachOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this licenses attach o k response has a 5xx status code
func (o *LicensesAttachOK) IsServerError() bool {
	return false
}

// IsCode returns true when this licenses attach o k response a status code equal to that given
func (o *LicensesAttachOK) IsCode(code int) bool {
	return code == 200
}

func (o *LicensesAttachOK) Error() string {
	return fmt.Sprintf("[POST /licenses][%d] licensesAttachOK  %+v", 200, o.Payload)
}

func (o *LicensesAttachOK) String() string {
	return fmt.Sprintf("[POST /licenses][%d] licensesAttachOK  %+v", 200, o.Payload)
}

func (o *LicensesAttachOK) GetPayload() *models.LicensesAttachResponse {
	return o.Payload
}

func (o *LicensesAttachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicensesAttachResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
