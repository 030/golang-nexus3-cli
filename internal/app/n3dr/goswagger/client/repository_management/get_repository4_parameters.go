// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

// NewGetRepository4Params creates a new GetRepository4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepository4Params() *GetRepository4Params {
	return &GetRepository4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepository4ParamsWithTimeout creates a new GetRepository4Params object
// with the ability to set a timeout on a request.
func NewGetRepository4ParamsWithTimeout(timeout time.Duration) *GetRepository4Params {
	return &GetRepository4Params{
		timeout: timeout,
	}
}

// NewGetRepository4ParamsWithContext creates a new GetRepository4Params object
// with the ability to set a context for a request.
func NewGetRepository4ParamsWithContext(ctx context.Context) *GetRepository4Params {
	return &GetRepository4Params{
		Context: ctx,
	}
}

// NewGetRepository4ParamsWithHTTPClient creates a new GetRepository4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepository4ParamsWithHTTPClient(client *http.Client) *GetRepository4Params {
	return &GetRepository4Params{
		HTTPClient: client,
	}
}

/* GetRepository4Params contains all the parameters to send to the API endpoint
   for the get repository 4 operation.

   Typically these are written to a http.Request.
*/
type GetRepository4Params struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository4Params) WithDefaults() *GetRepository4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 4 params
func (o *GetRepository4Params) WithTimeout(timeout time.Duration) *GetRepository4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 4 params
func (o *GetRepository4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 4 params
func (o *GetRepository4Params) WithContext(ctx context.Context) *GetRepository4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 4 params
func (o *GetRepository4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 4 params
func (o *GetRepository4Params) WithHTTPClient(client *http.Client) *GetRepository4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 4 params
func (o *GetRepository4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 4 params
func (o *GetRepository4Params) WithRepositoryName(repositoryName string) *GetRepository4Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 4 params
func (o *GetRepository4Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepository4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}