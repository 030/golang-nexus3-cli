// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository11Reader is a Reader for the UpdateRepository11 structure.
type UpdateRepository11Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository11Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository11NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository11Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository11Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository11NoContent creates a UpdateRepository11NoContent with default headers values
func NewUpdateRepository11NoContent() *UpdateRepository11NoContent {
	return &UpdateRepository11NoContent{}
}

/* UpdateRepository11NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository11NoContent struct {
}

func (o *UpdateRepository11NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] updateRepository11NoContent ", 204)
}

func (o *UpdateRepository11NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository11Unauthorized creates a UpdateRepository11Unauthorized with default headers values
func NewUpdateRepository11Unauthorized() *UpdateRepository11Unauthorized {
	return &UpdateRepository11Unauthorized{}
}

/* UpdateRepository11Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository11Unauthorized struct {
}

func (o *UpdateRepository11Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] updateRepository11Unauthorized ", 401)
}

func (o *UpdateRepository11Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository11Forbidden creates a UpdateRepository11Forbidden with default headers values
func NewUpdateRepository11Forbidden() *UpdateRepository11Forbidden {
	return &UpdateRepository11Forbidden{}
}

/* UpdateRepository11Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository11Forbidden struct {
}

func (o *UpdateRepository11Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/group/{repositoryName}][%d] updateRepository11Forbidden ", 403)
}

func (o *UpdateRepository11Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
