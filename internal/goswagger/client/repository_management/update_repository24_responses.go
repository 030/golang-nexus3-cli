// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository24Reader is a Reader for the UpdateRepository24 structure.
type UpdateRepository24Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository24Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository24NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository24Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository24Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository24NoContent creates a UpdateRepository24NoContent with default headers values
func NewUpdateRepository24NoContent() *UpdateRepository24NoContent {
	return &UpdateRepository24NoContent{}
}

/* UpdateRepository24NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository24NoContent struct {
}

func (o *UpdateRepository24NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/proxy/{repositoryName}][%d] updateRepository24NoContent ", 204)
}

func (o *UpdateRepository24NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository24Unauthorized creates a UpdateRepository24Unauthorized with default headers values
func NewUpdateRepository24Unauthorized() *UpdateRepository24Unauthorized {
	return &UpdateRepository24Unauthorized{}
}

/* UpdateRepository24Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository24Unauthorized struct {
}

func (o *UpdateRepository24Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/proxy/{repositoryName}][%d] updateRepository24Unauthorized ", 401)
}

func (o *UpdateRepository24Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository24Forbidden creates a UpdateRepository24Forbidden with default headers values
func NewUpdateRepository24Forbidden() *UpdateRepository24Forbidden {
	return &UpdateRepository24Forbidden{}
}

/* UpdateRepository24Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository24Forbidden struct {
}

func (o *UpdateRepository24Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/helm/proxy/{repositoryName}][%d] updateRepository24Forbidden ", 403)
}

func (o *UpdateRepository24Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
