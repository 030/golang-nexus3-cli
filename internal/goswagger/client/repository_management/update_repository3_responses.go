// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository3Reader is a Reader for the UpdateRepository3 structure.
type UpdateRepository3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepository3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository3NoContent creates a UpdateRepository3NoContent with default headers values
func NewUpdateRepository3NoContent() *UpdateRepository3NoContent {
	return &UpdateRepository3NoContent{}
}

/* UpdateRepository3NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository3NoContent struct {
}

func (o *UpdateRepository3NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/apt/hosted/{repositoryName}][%d] updateRepository3NoContent ", 204)
}

func (o *UpdateRepository3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository3Unauthorized creates a UpdateRepository3Unauthorized with default headers values
func NewUpdateRepository3Unauthorized() *UpdateRepository3Unauthorized {
	return &UpdateRepository3Unauthorized{}
}

/* UpdateRepository3Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository3Unauthorized struct {
}

func (o *UpdateRepository3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/apt/hosted/{repositoryName}][%d] updateRepository3Unauthorized ", 401)
}

func (o *UpdateRepository3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository3Forbidden creates a UpdateRepository3Forbidden with default headers values
func NewUpdateRepository3Forbidden() *UpdateRepository3Forbidden {
	return &UpdateRepository3Forbidden{}
}

/* UpdateRepository3Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository3Forbidden struct {
}

func (o *UpdateRepository3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/apt/hosted/{repositoryName}][%d] updateRepository3Forbidden ", 403)
}

func (o *UpdateRepository3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository3NotFound creates a UpdateRepository3NotFound with default headers values
func NewUpdateRepository3NotFound() *UpdateRepository3NotFound {
	return &UpdateRepository3NotFound{}
}

/* UpdateRepository3NotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRepository3NotFound struct {
}

func (o *UpdateRepository3NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/apt/hosted/{repositoryName}][%d] updateRepository3NotFound ", 404)
}

func (o *UpdateRepository3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
