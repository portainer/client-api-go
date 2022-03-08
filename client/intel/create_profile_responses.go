// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateProfileReader is a Reader for the CreateProfile structure.
type CreateProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProfileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProfileOK creates a CreateProfileOK with default headers values
func NewCreateProfileOK() *CreateProfileOK {
	return &CreateProfileOK{}
}

/* CreateProfileOK describes a response with status code 200, with default header values.

Success
*/
type CreateProfileOK struct {
}

func (o *CreateProfileOK) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles][%d] createProfileOK ", 200)
}

func (o *CreateProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProfileBadRequest creates a CreateProfileBadRequest with default headers values
func NewCreateProfileBadRequest() *CreateProfileBadRequest {
	return &CreateProfileBadRequest{}
}

/* CreateProfileBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CreateProfileBadRequest struct {
}

func (o *CreateProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles][%d] createProfileBadRequest ", 400)
}

func (o *CreateProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProfileConflict creates a CreateProfileConflict with default headers values
func NewCreateProfileConflict() *CreateProfileConflict {
	return &CreateProfileConflict{}
}

/* CreateProfileConflict describes a response with status code 409, with default header values.

Profile name already exists
*/
type CreateProfileConflict struct {
}

func (o *CreateProfileConflict) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles][%d] createProfileConflict ", 409)
}

func (o *CreateProfileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProfileInternalServerError creates a CreateProfileInternalServerError with default headers values
func NewCreateProfileInternalServerError() *CreateProfileInternalServerError {
	return &CreateProfileInternalServerError{}
}

/* CreateProfileInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateProfileInternalServerError struct {
}

func (o *CreateProfileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fdo/profiles][%d] createProfileInternalServerError ", 500)
}

func (o *CreateProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
