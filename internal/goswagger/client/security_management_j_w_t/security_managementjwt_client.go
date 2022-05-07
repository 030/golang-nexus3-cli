// Code generated by go-swagger; DO NOT EDIT.

package security_management_j_w_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security management j w t API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security management j w t API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSecret(params *GetSecretParams, opts ...ClientOption) error

	UpdateSecret(params *UpdateSecretParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSecret gets j w t secret
*/
func (a *Client) GetSecret(params *GetSecretParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecretParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSecret",
		Method:             "GET",
		PathPattern:        "/v1/security/jwt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecretReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdateSecret sets j w t secret note that session will be expired for the all logged in users
*/
func (a *Client) UpdateSecret(params *UpdateSecretParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSecretParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSecret",
		Method:             "PUT",
		PathPattern:        "/v1/security/jwt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSecretReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}