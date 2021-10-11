// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/go-swagger/models"
)

// NewCreateParams creates a new CreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateParams() *CreateParams {
	return &CreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateParamsWithTimeout creates a new CreateParams object
// with the ability to set a timeout on a request.
func NewCreateParamsWithTimeout(timeout time.Duration) *CreateParams {
	return &CreateParams{
		timeout: timeout,
	}
}

// NewCreateParamsWithContext creates a new CreateParams object
// with the ability to set a context for a request.
func NewCreateParamsWithContext(ctx context.Context) *CreateParams {
	return &CreateParams{
		Context: ctx,
	}
}

// NewCreateParamsWithHTTPClient creates a new CreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateParamsWithHTTPClient(client *http.Client) *CreateParams {
	return &CreateParams{
		HTTPClient: client,
	}
}

/* CreateParams contains all the parameters to send to the API endpoint
   for the create operation.

   Typically these are written to a http.Request.
*/
type CreateParams struct {

	/* Body.

	   A role configuration
	*/
	Body *models.RoleXORequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateParams) WithDefaults() *CreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create params
func (o *CreateParams) WithTimeout(timeout time.Duration) *CreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create params
func (o *CreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create params
func (o *CreateParams) WithContext(ctx context.Context) *CreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create params
func (o *CreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) WithHTTPClient(client *http.Client) *CreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create params
func (o *CreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create params
func (o *CreateParams) WithBody(body *models.RoleXORequest) *CreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create params
func (o *CreateParams) SetBody(body *models.RoleXORequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
