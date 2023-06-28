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

// NewGetRepository5Params creates a new GetRepository5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepository5Params() *GetRepository5Params {
	return &GetRepository5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepository5ParamsWithTimeout creates a new GetRepository5Params object
// with the ability to set a timeout on a request.
func NewGetRepository5ParamsWithTimeout(timeout time.Duration) *GetRepository5Params {
	return &GetRepository5Params{
		timeout: timeout,
	}
}

// NewGetRepository5ParamsWithContext creates a new GetRepository5Params object
// with the ability to set a context for a request.
func NewGetRepository5ParamsWithContext(ctx context.Context) *GetRepository5Params {
	return &GetRepository5Params{
		Context: ctx,
	}
}

// NewGetRepository5ParamsWithHTTPClient creates a new GetRepository5Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepository5ParamsWithHTTPClient(client *http.Client) *GetRepository5Params {
	return &GetRepository5Params{
		HTTPClient: client,
	}
}

/*
GetRepository5Params contains all the parameters to send to the API endpoint

	for the get repository 5 operation.

	Typically these are written to a http.Request.
*/
type GetRepository5Params struct {

	// RepositoryName.
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository5Params) WithDefaults() *GetRepository5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repository 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepository5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repository 5 params
func (o *GetRepository5Params) WithTimeout(timeout time.Duration) *GetRepository5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository 5 params
func (o *GetRepository5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository 5 params
func (o *GetRepository5Params) WithContext(ctx context.Context) *GetRepository5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository 5 params
func (o *GetRepository5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository 5 params
func (o *GetRepository5Params) WithHTTPClient(client *http.Client) *GetRepository5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository 5 params
func (o *GetRepository5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get repository 5 params
func (o *GetRepository5Params) WithRepositoryName(repositoryName string) *GetRepository5Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get repository 5 params
func (o *GetRepository5Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepository5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
