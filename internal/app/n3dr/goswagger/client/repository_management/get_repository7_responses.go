// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetRepository7Reader is a Reader for the GetRepository7 structure.
type GetRepository7Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository7Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository7OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository7OK creates a GetRepository7OK with default headers values
func NewGetRepository7OK() *GetRepository7OK {
	return &GetRepository7OK{}
}

/*
GetRepository7OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository7OK struct {
	Payload *models.SimpleAPIHostedRepository
}

// IsSuccess returns true when this get repository7 o k response has a 2xx status code
func (o *GetRepository7OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository7 o k response has a 3xx status code
func (o *GetRepository7OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository7 o k response has a 4xx status code
func (o *GetRepository7OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository7 o k response has a 5xx status code
func (o *GetRepository7OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository7 o k response a status code equal to that given
func (o *GetRepository7OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository7 o k response
func (o *GetRepository7OK) Code() int {
	return 200
}

func (o *GetRepository7OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/raw/hosted/{repositoryName}][%d] getRepository7OK  %+v", 200, o.Payload)
}

func (o *GetRepository7OK) String() string {
	return fmt.Sprintf("[GET /v1/repositories/raw/hosted/{repositoryName}][%d] getRepository7OK  %+v", 200, o.Payload)
}

func (o *GetRepository7OK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepository7OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
