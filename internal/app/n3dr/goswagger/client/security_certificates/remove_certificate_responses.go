// Code generated by go-swagger; DO NOT EDIT.

package security_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveCertificateReader is a Reader for the RemoveCertificate structure.
type RemoveCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewRemoveCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveCertificateForbidden creates a RemoveCertificateForbidden with default headers values
func NewRemoveCertificateForbidden() *RemoveCertificateForbidden {
	return &RemoveCertificateForbidden{}
}

/*
RemoveCertificateForbidden describes a response with status code 403, with default header values.

Insufficient permissions to remove certificate from the trust store
*/
type RemoveCertificateForbidden struct {
}

// IsSuccess returns true when this remove certificate forbidden response has a 2xx status code
func (o *RemoveCertificateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove certificate forbidden response has a 3xx status code
func (o *RemoveCertificateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove certificate forbidden response has a 4xx status code
func (o *RemoveCertificateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove certificate forbidden response has a 5xx status code
func (o *RemoveCertificateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove certificate forbidden response a status code equal to that given
func (o *RemoveCertificateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove certificate forbidden response
func (o *RemoveCertificateForbidden) Code() int {
	return 403
}

func (o *RemoveCertificateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/security/ssl/truststore/{id}][%d] removeCertificateForbidden ", 403)
}

func (o *RemoveCertificateForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/security/ssl/truststore/{id}][%d] removeCertificateForbidden ", 403)
}

func (o *RemoveCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
