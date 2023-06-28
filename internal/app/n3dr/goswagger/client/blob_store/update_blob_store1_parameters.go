// Code generated by go-swagger; DO NOT EDIT.

package blob_store

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

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// NewUpdateBlobStore1Params creates a new UpdateBlobStore1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBlobStore1Params() *UpdateBlobStore1Params {
	return &UpdateBlobStore1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBlobStore1ParamsWithTimeout creates a new UpdateBlobStore1Params object
// with the ability to set a timeout on a request.
func NewUpdateBlobStore1ParamsWithTimeout(timeout time.Duration) *UpdateBlobStore1Params {
	return &UpdateBlobStore1Params{
		timeout: timeout,
	}
}

// NewUpdateBlobStore1ParamsWithContext creates a new UpdateBlobStore1Params object
// with the ability to set a context for a request.
func NewUpdateBlobStore1ParamsWithContext(ctx context.Context) *UpdateBlobStore1Params {
	return &UpdateBlobStore1Params{
		Context: ctx,
	}
}

// NewUpdateBlobStore1ParamsWithHTTPClient creates a new UpdateBlobStore1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBlobStore1ParamsWithHTTPClient(client *http.Client) *UpdateBlobStore1Params {
	return &UpdateBlobStore1Params{
		HTTPClient: client,
	}
}

/*
UpdateBlobStore1Params contains all the parameters to send to the API endpoint

	for the update blob store 1 operation.

	Typically these are written to a http.Request.
*/
type UpdateBlobStore1Params struct {

	// Body.
	Body *models.AzureBlobStoreAPIModel

	/* Name.

	   Name of the blob store to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update blob store 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBlobStore1Params) WithDefaults() *UpdateBlobStore1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update blob store 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBlobStore1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update blob store 1 params
func (o *UpdateBlobStore1Params) WithTimeout(timeout time.Duration) *UpdateBlobStore1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update blob store 1 params
func (o *UpdateBlobStore1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update blob store 1 params
func (o *UpdateBlobStore1Params) WithContext(ctx context.Context) *UpdateBlobStore1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update blob store 1 params
func (o *UpdateBlobStore1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update blob store 1 params
func (o *UpdateBlobStore1Params) WithHTTPClient(client *http.Client) *UpdateBlobStore1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update blob store 1 params
func (o *UpdateBlobStore1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update blob store 1 params
func (o *UpdateBlobStore1Params) WithBody(body *models.AzureBlobStoreAPIModel) *UpdateBlobStore1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update blob store 1 params
func (o *UpdateBlobStore1Params) SetBody(body *models.AzureBlobStoreAPIModel) {
	o.Body = body
}

// WithName adds the name to the update blob store 1 params
func (o *UpdateBlobStore1Params) WithName(name string) *UpdateBlobStore1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the update blob store 1 params
func (o *UpdateBlobStore1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBlobStore1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
