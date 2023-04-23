// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepositoryReader is a Reader for the CreateRepository structure.
type CreateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepositoryCreated creates a CreateRepositoryCreated with default headers values
func NewCreateRepositoryCreated() *CreateRepositoryCreated {
	return &CreateRepositoryCreated{}
}

/*
CreateRepositoryCreated describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepositoryCreated struct {
}

// IsSuccess returns true when this create repository created response has a 2xx status code
func (o *CreateRepositoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository created response has a 3xx status code
func (o *CreateRepositoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository created response has a 4xx status code
func (o *CreateRepositoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository created response has a 5xx status code
func (o *CreateRepositoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository created response a status code equal to that given
func (o *CreateRepositoryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository created response
func (o *CreateRepositoryCreated) Code() int {
	return 201
}

func (o *CreateRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryCreated ", 201)
}

func (o *CreateRepositoryCreated) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryCreated ", 201)
}

func (o *CreateRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepositoryUnauthorized creates a CreateRepositoryUnauthorized with default headers values
func NewCreateRepositoryUnauthorized() *CreateRepositoryUnauthorized {
	return &CreateRepositoryUnauthorized{}
}

/*
CreateRepositoryUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepositoryUnauthorized struct {
}

// IsSuccess returns true when this create repository unauthorized response has a 2xx status code
func (o *CreateRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository unauthorized response has a 3xx status code
func (o *CreateRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository unauthorized response has a 4xx status code
func (o *CreateRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository unauthorized response has a 5xx status code
func (o *CreateRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository unauthorized response a status code equal to that given
func (o *CreateRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository unauthorized response
func (o *CreateRepositoryUnauthorized) Code() int {
	return 401
}

func (o *CreateRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryUnauthorized ", 401)
}

func (o *CreateRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryUnauthorized ", 401)
}

func (o *CreateRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepositoryForbidden creates a CreateRepositoryForbidden with default headers values
func NewCreateRepositoryForbidden() *CreateRepositoryForbidden {
	return &CreateRepositoryForbidden{}
}

/*
CreateRepositoryForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepositoryForbidden struct {
}

// IsSuccess returns true when this create repository forbidden response has a 2xx status code
func (o *CreateRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository forbidden response has a 3xx status code
func (o *CreateRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository forbidden response has a 4xx status code
func (o *CreateRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository forbidden response has a 5xx status code
func (o *CreateRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository forbidden response a status code equal to that given
func (o *CreateRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository forbidden response
func (o *CreateRepositoryForbidden) Code() int {
	return 403
}

func (o *CreateRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryForbidden ", 403)
}

func (o *CreateRepositoryForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/maven/group][%d] createRepositoryForbidden ", 403)
}

func (o *CreateRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
