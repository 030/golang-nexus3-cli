// Code generated by go-swagger; DO NOT EDIT.

package security_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetUserSourcesReader is a Reader for the GetUserSources structure.
type GetUserSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSourcesOK creates a GetUserSourcesOK with default headers values
func NewGetUserSourcesOK() *GetUserSourcesOK {
	return &GetUserSourcesOK{}
}

/*
GetUserSourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserSourcesOK struct {
	Payload []*models.APIUserSource
}

// IsSuccess returns true when this get user sources o k response has a 2xx status code
func (o *GetUserSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user sources o k response has a 3xx status code
func (o *GetUserSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user sources o k response has a 4xx status code
func (o *GetUserSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user sources o k response has a 5xx status code
func (o *GetUserSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user sources o k response a status code equal to that given
func (o *GetUserSourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user sources o k response
func (o *GetUserSourcesOK) Code() int {
	return 200
}

func (o *GetUserSourcesOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/user-sources][%d] getUserSourcesOK  %+v", 200, o.Payload)
}

func (o *GetUserSourcesOK) String() string {
	return fmt.Sprintf("[GET /v1/security/user-sources][%d] getUserSourcesOK  %+v", 200, o.Payload)
}

func (o *GetUserSourcesOK) GetPayload() []*models.APIUserSource {
	return o.Payload
}

func (o *GetUserSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSourcesForbidden creates a GetUserSourcesForbidden with default headers values
func NewGetUserSourcesForbidden() *GetUserSourcesForbidden {
	return &GetUserSourcesForbidden{}
}

/*
GetUserSourcesForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type GetUserSourcesForbidden struct {
}

// IsSuccess returns true when this get user sources forbidden response has a 2xx status code
func (o *GetUserSourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user sources forbidden response has a 3xx status code
func (o *GetUserSourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user sources forbidden response has a 4xx status code
func (o *GetUserSourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user sources forbidden response has a 5xx status code
func (o *GetUserSourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user sources forbidden response a status code equal to that given
func (o *GetUserSourcesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user sources forbidden response
func (o *GetUserSourcesForbidden) Code() int {
	return 403
}

func (o *GetUserSourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/user-sources][%d] getUserSourcesForbidden ", 403)
}

func (o *GetUserSourcesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/security/user-sources][%d] getUserSourcesForbidden ", 403)
}

func (o *GetUserSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
