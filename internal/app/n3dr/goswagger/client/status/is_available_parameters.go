// Code generated by go-swagger; DO NOT EDIT.

package status

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
)

// NewIsAvailableParams creates a new IsAvailableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIsAvailableParams() *IsAvailableParams {
	return &IsAvailableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIsAvailableParamsWithTimeout creates a new IsAvailableParams object
// with the ability to set a timeout on a request.
func NewIsAvailableParamsWithTimeout(timeout time.Duration) *IsAvailableParams {
	return &IsAvailableParams{
		timeout: timeout,
	}
}

// NewIsAvailableParamsWithContext creates a new IsAvailableParams object
// with the ability to set a context for a request.
func NewIsAvailableParamsWithContext(ctx context.Context) *IsAvailableParams {
	return &IsAvailableParams{
		Context: ctx,
	}
}

// NewIsAvailableParamsWithHTTPClient creates a new IsAvailableParams object
// with the ability to set a custom HTTPClient for a request.
func NewIsAvailableParamsWithHTTPClient(client *http.Client) *IsAvailableParams {
	return &IsAvailableParams{
		HTTPClient: client,
	}
}

/* IsAvailableParams contains all the parameters to send to the API endpoint
   for the is available operation.

   Typically these are written to a http.Request.
*/
type IsAvailableParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the is available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsAvailableParams) WithDefaults() *IsAvailableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the is available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IsAvailableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the is available params
func (o *IsAvailableParams) WithTimeout(timeout time.Duration) *IsAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is available params
func (o *IsAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is available params
func (o *IsAvailableParams) WithContext(ctx context.Context) *IsAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is available params
func (o *IsAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is available params
func (o *IsAvailableParams) WithHTTPClient(client *http.Client) *IsAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is available params
func (o *IsAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}