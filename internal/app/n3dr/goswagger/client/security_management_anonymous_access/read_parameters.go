// Code generated by go-swagger; DO NOT EDIT.

package security_management_anonymous_access

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

// NewReadParams creates a new ReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadParams() *ReadParams {
	return &ReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadParamsWithTimeout creates a new ReadParams object
// with the ability to set a timeout on a request.
func NewReadParamsWithTimeout(timeout time.Duration) *ReadParams {
	return &ReadParams{
		timeout: timeout,
	}
}

// NewReadParamsWithContext creates a new ReadParams object
// with the ability to set a context for a request.
func NewReadParamsWithContext(ctx context.Context) *ReadParams {
	return &ReadParams{
		Context: ctx,
	}
}

// NewReadParamsWithHTTPClient creates a new ReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadParamsWithHTTPClient(client *http.Client) *ReadParams {
	return &ReadParams{
		HTTPClient: client,
	}
}

/*
ReadParams contains all the parameters to send to the API endpoint

	for the read operation.

	Typically these are written to a http.Request.
*/
type ReadParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadParams) WithDefaults() *ReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read params
func (o *ReadParams) WithTimeout(timeout time.Duration) *ReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read params
func (o *ReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read params
func (o *ReadParams) WithContext(ctx context.Context) *ReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read params
func (o *ReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read params
func (o *ReadParams) WithHTTPClient(client *http.Client) *ReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read params
func (o *ReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
