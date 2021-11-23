// Code generated by go-swagger; DO NOT EDIT.

package manage_i_q_server_configuration

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

// NewEnableIqParams creates a new EnableIqParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableIqParams() *EnableIqParams {
	return &EnableIqParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableIqParamsWithTimeout creates a new EnableIqParams object
// with the ability to set a timeout on a request.
func NewEnableIqParamsWithTimeout(timeout time.Duration) *EnableIqParams {
	return &EnableIqParams{
		timeout: timeout,
	}
}

// NewEnableIqParamsWithContext creates a new EnableIqParams object
// with the ability to set a context for a request.
func NewEnableIqParamsWithContext(ctx context.Context) *EnableIqParams {
	return &EnableIqParams{
		Context: ctx,
	}
}

// NewEnableIqParamsWithHTTPClient creates a new EnableIqParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableIqParamsWithHTTPClient(client *http.Client) *EnableIqParams {
	return &EnableIqParams{
		HTTPClient: client,
	}
}

/* EnableIqParams contains all the parameters to send to the API endpoint
   for the enable iq operation.

   Typically these are written to a http.Request.
*/
type EnableIqParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable iq params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableIqParams) WithDefaults() *EnableIqParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable iq params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableIqParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable iq params
func (o *EnableIqParams) WithTimeout(timeout time.Duration) *EnableIqParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable iq params
func (o *EnableIqParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable iq params
func (o *EnableIqParams) WithContext(ctx context.Context) *EnableIqParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable iq params
func (o *EnableIqParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable iq params
func (o *EnableIqParams) WithHTTPClient(client *http.Client) *EnableIqParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable iq params
func (o *EnableIqParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EnableIqParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
