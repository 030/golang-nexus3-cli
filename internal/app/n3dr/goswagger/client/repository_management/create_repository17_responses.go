// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository17Reader is a Reader for the CreateRepository17 structure.
type CreateRepository17Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository17Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository17Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository17Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository17Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository17Created creates a CreateRepository17Created with default headers values
func NewCreateRepository17Created() *CreateRepository17Created {
	return &CreateRepository17Created{}
}

/*
CreateRepository17Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository17Created struct {
}

// IsSuccess returns true when this create repository17 created response has a 2xx status code
func (o *CreateRepository17Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository17 created response has a 3xx status code
func (o *CreateRepository17Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository17 created response has a 4xx status code
func (o *CreateRepository17Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository17 created response has a 5xx status code
func (o *CreateRepository17Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository17 created response a status code equal to that given
func (o *CreateRepository17Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository17 created response
func (o *CreateRepository17Created) Code() int {
	return 201
}

func (o *CreateRepository17Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Created ", 201)
}

func (o *CreateRepository17Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Created ", 201)
}

func (o *CreateRepository17Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository17Unauthorized creates a CreateRepository17Unauthorized with default headers values
func NewCreateRepository17Unauthorized() *CreateRepository17Unauthorized {
	return &CreateRepository17Unauthorized{}
}

/*
CreateRepository17Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository17Unauthorized struct {
}

// IsSuccess returns true when this create repository17 unauthorized response has a 2xx status code
func (o *CreateRepository17Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository17 unauthorized response has a 3xx status code
func (o *CreateRepository17Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository17 unauthorized response has a 4xx status code
func (o *CreateRepository17Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository17 unauthorized response has a 5xx status code
func (o *CreateRepository17Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository17 unauthorized response a status code equal to that given
func (o *CreateRepository17Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository17 unauthorized response
func (o *CreateRepository17Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository17Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Unauthorized ", 401)
}

func (o *CreateRepository17Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Unauthorized ", 401)
}

func (o *CreateRepository17Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository17Forbidden creates a CreateRepository17Forbidden with default headers values
func NewCreateRepository17Forbidden() *CreateRepository17Forbidden {
	return &CreateRepository17Forbidden{}
}

/*
CreateRepository17Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository17Forbidden struct {
}

// IsSuccess returns true when this create repository17 forbidden response has a 2xx status code
func (o *CreateRepository17Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository17 forbidden response has a 3xx status code
func (o *CreateRepository17Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository17 forbidden response has a 4xx status code
func (o *CreateRepository17Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository17 forbidden response has a 5xx status code
func (o *CreateRepository17Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository17 forbidden response a status code equal to that given
func (o *CreateRepository17Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository17 forbidden response
func (o *CreateRepository17Forbidden) Code() int {
	return 403
}

func (o *CreateRepository17Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Forbidden ", 403)
}

func (o *CreateRepository17Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/group][%d] createRepository17Forbidden ", 403)
}

func (o *CreateRepository17Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
