// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository2Reader is a Reader for the UpdateRepository2 structure.
type UpdateRepository2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository2NoContent creates a UpdateRepository2NoContent with default headers values
func NewUpdateRepository2NoContent() *UpdateRepository2NoContent {
	return &UpdateRepository2NoContent{}
}

/* UpdateRepository2NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository2NoContent struct {
}

func (o *UpdateRepository2NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] updateRepository2NoContent ", 204)
}

func (o *UpdateRepository2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository2Unauthorized creates a UpdateRepository2Unauthorized with default headers values
func NewUpdateRepository2Unauthorized() *UpdateRepository2Unauthorized {
	return &UpdateRepository2Unauthorized{}
}

/* UpdateRepository2Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository2Unauthorized struct {
}

func (o *UpdateRepository2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] updateRepository2Unauthorized ", 401)
}

func (o *UpdateRepository2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository2Forbidden creates a UpdateRepository2Forbidden with default headers values
func NewUpdateRepository2Forbidden() *UpdateRepository2Forbidden {
	return &UpdateRepository2Forbidden{}
}

/* UpdateRepository2Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository2Forbidden struct {
}

func (o *UpdateRepository2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/maven/proxy/{repositoryName}][%d] updateRepository2Forbidden ", 403)
}

func (o *UpdateRepository2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
