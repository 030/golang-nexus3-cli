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
)

// NewDeleteBlobStoreParams creates a new DeleteBlobStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlobStoreParams() *DeleteBlobStoreParams {
	return &DeleteBlobStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlobStoreParamsWithTimeout creates a new DeleteBlobStoreParams object
// with the ability to set a timeout on a request.
func NewDeleteBlobStoreParamsWithTimeout(timeout time.Duration) *DeleteBlobStoreParams {
	return &DeleteBlobStoreParams{
		timeout: timeout,
	}
}

// NewDeleteBlobStoreParamsWithContext creates a new DeleteBlobStoreParams object
// with the ability to set a context for a request.
func NewDeleteBlobStoreParamsWithContext(ctx context.Context) *DeleteBlobStoreParams {
	return &DeleteBlobStoreParams{
		Context: ctx,
	}
}

// NewDeleteBlobStoreParamsWithHTTPClient creates a new DeleteBlobStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlobStoreParamsWithHTTPClient(client *http.Client) *DeleteBlobStoreParams {
	return &DeleteBlobStoreParams{
		HTTPClient: client,
	}
}

/* DeleteBlobStoreParams contains all the parameters to send to the API endpoint
   for the delete blob store operation.

   Typically these are written to a http.Request.
*/
type DeleteBlobStoreParams struct {

	/* Name.

	   The name of the blob store to delete
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlobStoreParams) WithDefaults() *DeleteBlobStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete blob store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlobStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete blob store params
func (o *DeleteBlobStoreParams) WithTimeout(timeout time.Duration) *DeleteBlobStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete blob store params
func (o *DeleteBlobStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete blob store params
func (o *DeleteBlobStoreParams) WithContext(ctx context.Context) *DeleteBlobStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete blob store params
func (o *DeleteBlobStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete blob store params
func (o *DeleteBlobStoreParams) WithHTTPClient(client *http.Client) *DeleteBlobStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete blob store params
func (o *DeleteBlobStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete blob store params
func (o *DeleteBlobStoreParams) WithName(name string) *DeleteBlobStoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete blob store params
func (o *DeleteBlobStoreParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlobStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}