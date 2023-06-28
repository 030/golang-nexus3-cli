// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*
GetTasksOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTasksOK struct {
	Payload *models.PageTaskXO
}

// IsSuccess returns true when this get tasks o k response has a 2xx status code
func (o *GetTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tasks o k response has a 3xx status code
func (o *GetTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks o k response has a 4xx status code
func (o *GetTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks o k response has a 5xx status code
func (o *GetTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks o k response a status code equal to that given
func (o *GetTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tasks o k response
func (o *GetTasksOK) Code() int {
	return 200
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) String() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) GetPayload() *models.PageTaskXO {
	return o.Payload
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageTaskXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
