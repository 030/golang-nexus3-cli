// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository20Reader is a Reader for the CreateRepository20 structure.
type CreateRepository20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository20Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository20Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository20Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository20Created creates a CreateRepository20Created with default headers values
func NewCreateRepository20Created() *CreateRepository20Created {
	return &CreateRepository20Created{}
}

/* CreateRepository20Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository20Created struct {
}

func (o *CreateRepository20Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/group][%d] createRepository20Created ", 201)
}

func (o *CreateRepository20Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository20Unauthorized creates a CreateRepository20Unauthorized with default headers values
func NewCreateRepository20Unauthorized() *CreateRepository20Unauthorized {
	return &CreateRepository20Unauthorized{}
}

/* CreateRepository20Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository20Unauthorized struct {
}

func (o *CreateRepository20Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/group][%d] createRepository20Unauthorized ", 401)
}

func (o *CreateRepository20Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository20Forbidden creates a CreateRepository20Forbidden with default headers values
func NewCreateRepository20Forbidden() *CreateRepository20Forbidden {
	return &CreateRepository20Forbidden{}
}

/* CreateRepository20Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository20Forbidden struct {
}

func (o *CreateRepository20Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/group][%d] createRepository20Forbidden ", 403)
}

func (o *CreateRepository20Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}