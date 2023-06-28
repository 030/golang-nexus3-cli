// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository34Reader is a Reader for the CreateRepository34 structure.
type CreateRepository34Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository34Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository34Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository34Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository34Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository34MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository34Created creates a CreateRepository34Created with default headers values
func NewCreateRepository34Created() *CreateRepository34Created {
	return &CreateRepository34Created{}
}

/*
CreateRepository34Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository34Created struct {
}

// IsSuccess returns true when this create repository34 created response has a 2xx status code
func (o *CreateRepository34Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository34 created response has a 3xx status code
func (o *CreateRepository34Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository34 created response has a 4xx status code
func (o *CreateRepository34Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository34 created response has a 5xx status code
func (o *CreateRepository34Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository34 created response a status code equal to that given
func (o *CreateRepository34Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository34 created response
func (o *CreateRepository34Created) Code() int {
	return 201
}

func (o *CreateRepository34Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Created ", 201)
}

func (o *CreateRepository34Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Created ", 201)
}

func (o *CreateRepository34Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository34Unauthorized creates a CreateRepository34Unauthorized with default headers values
func NewCreateRepository34Unauthorized() *CreateRepository34Unauthorized {
	return &CreateRepository34Unauthorized{}
}

/*
CreateRepository34Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository34Unauthorized struct {
}

// IsSuccess returns true when this create repository34 unauthorized response has a 2xx status code
func (o *CreateRepository34Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository34 unauthorized response has a 3xx status code
func (o *CreateRepository34Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository34 unauthorized response has a 4xx status code
func (o *CreateRepository34Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository34 unauthorized response has a 5xx status code
func (o *CreateRepository34Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository34 unauthorized response a status code equal to that given
func (o *CreateRepository34Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository34 unauthorized response
func (o *CreateRepository34Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository34Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Unauthorized ", 401)
}

func (o *CreateRepository34Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Unauthorized ", 401)
}

func (o *CreateRepository34Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository34Forbidden creates a CreateRepository34Forbidden with default headers values
func NewCreateRepository34Forbidden() *CreateRepository34Forbidden {
	return &CreateRepository34Forbidden{}
}

/*
CreateRepository34Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository34Forbidden struct {
}

// IsSuccess returns true when this create repository34 forbidden response has a 2xx status code
func (o *CreateRepository34Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository34 forbidden response has a 3xx status code
func (o *CreateRepository34Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository34 forbidden response has a 4xx status code
func (o *CreateRepository34Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository34 forbidden response has a 5xx status code
func (o *CreateRepository34Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository34 forbidden response a status code equal to that given
func (o *CreateRepository34Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository34 forbidden response
func (o *CreateRepository34Forbidden) Code() int {
	return 403
}

func (o *CreateRepository34Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Forbidden ", 403)
}

func (o *CreateRepository34Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34Forbidden ", 403)
}

func (o *CreateRepository34Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository34MethodNotAllowed creates a CreateRepository34MethodNotAllowed with default headers values
func NewCreateRepository34MethodNotAllowed() *CreateRepository34MethodNotAllowed {
	return &CreateRepository34MethodNotAllowed{}
}

/*
CreateRepository34MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository34MethodNotAllowed struct {
}

// IsSuccess returns true when this create repository34 method not allowed response has a 2xx status code
func (o *CreateRepository34MethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository34 method not allowed response has a 3xx status code
func (o *CreateRepository34MethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository34 method not allowed response has a 4xx status code
func (o *CreateRepository34MethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository34 method not allowed response has a 5xx status code
func (o *CreateRepository34MethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository34 method not allowed response a status code equal to that given
func (o *CreateRepository34MethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the create repository34 method not allowed response
func (o *CreateRepository34MethodNotAllowed) Code() int {
	return 405
}

func (o *CreateRepository34MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34MethodNotAllowed ", 405)
}

func (o *CreateRepository34MethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/repositories/cocoapods/proxy][%d] createRepository34MethodNotAllowed ", 405)
}

func (o *CreateRepository34MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
