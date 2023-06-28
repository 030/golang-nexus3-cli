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

// NewCreateBlobStoreParams creates a new CreateBlobStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlobStoreParams() *CreateBlobStoreParams {
	return &CreateBlobStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlobStoreParamsWithTimeout creates a new CreateBlobStoreParams object
// with the ability to set a timeout on a request.
func NewCreateBlobStoreParamsWithTimeout(timeout time.Duration) *CreateBlobStoreParams {
	return &CreateBlobStoreParams{
		timeout: timeout,
	}
}

// NewCreateBlobStoreParamsWithContext creates a new CreateBlobStoreParams object
// with the ability to set a context for a request.
func NewCreateBlobStoreParamsWithContext(ctx context.Context) *CreateBlobStoreParams {
	return &CreateBlobStoreParams{
		Context: ctx,
	}
}

// NewCreateBlobStoreParamsWithHTTPClient creates a new CreateBlobStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlobStoreParamsWithHTTPClient(client *http.Client) *CreateBlobStoreParams {
	return &CreateBlobStoreParams{
		HTTPClient: client,
	}
}

/*
CreateBlobStoreParams contains all the parameters to send to the API endpoint

	for the create blob store operation.

	Typically these are written to a http.Request.
*/
type CreateBlobStoreParams struct {

	// Body.
	Body *models.S3BlobStoreAPIModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlobStoreParams) WithDefaults() *CreateBlobStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlobStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blob store params
func (o *CreateBlobStoreParams) WithTimeout(timeout time.Duration) *CreateBlobStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blob store params
func (o *CreateBlobStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blob store params
func (o *CreateBlobStoreParams) WithContext(ctx context.Context) *CreateBlobStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blob store params
func (o *CreateBlobStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blob store params
func (o *CreateBlobStoreParams) WithHTTPClient(client *http.Client) *CreateBlobStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blob store params
func (o *CreateBlobStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blob store params
func (o *CreateBlobStoreParams) WithBody(body *models.S3BlobStoreAPIModel) *CreateBlobStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blob store params
func (o *CreateBlobStoreParams) SetBody(body *models.S3BlobStoreAPIModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlobStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
