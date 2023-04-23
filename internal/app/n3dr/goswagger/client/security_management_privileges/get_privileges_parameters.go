// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

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

// NewGetPrivilegesParams creates a new GetPrivilegesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPrivilegesParams() *GetPrivilegesParams {
	return &GetPrivilegesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivilegesParamsWithTimeout creates a new GetPrivilegesParams object
// with the ability to set a timeout on a request.
func NewGetPrivilegesParamsWithTimeout(timeout time.Duration) *GetPrivilegesParams {
	return &GetPrivilegesParams{
		timeout: timeout,
	}
}

// NewGetPrivilegesParamsWithContext creates a new GetPrivilegesParams object
// with the ability to set a context for a request.
func NewGetPrivilegesParamsWithContext(ctx context.Context) *GetPrivilegesParams {
	return &GetPrivilegesParams{
		Context: ctx,
	}
}

// NewGetPrivilegesParamsWithHTTPClient creates a new GetPrivilegesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPrivilegesParamsWithHTTPClient(client *http.Client) *GetPrivilegesParams {
	return &GetPrivilegesParams{
		HTTPClient: client,
	}
}

/*
GetPrivilegesParams contains all the parameters to send to the API endpoint

	for the get privileges operation.

	Typically these are written to a http.Request.
*/
type GetPrivilegesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get privileges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivilegesParams) WithDefaults() *GetPrivilegesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get privileges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivilegesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get privileges params
func (o *GetPrivilegesParams) WithTimeout(timeout time.Duration) *GetPrivilegesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get privileges params
func (o *GetPrivilegesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get privileges params
func (o *GetPrivilegesParams) WithContext(ctx context.Context) *GetPrivilegesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get privileges params
func (o *GetPrivilegesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get privileges params
func (o *GetPrivilegesParams) WithHTTPClient(client *http.Client) *GetPrivilegesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get privileges params
func (o *GetPrivilegesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivilegesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
