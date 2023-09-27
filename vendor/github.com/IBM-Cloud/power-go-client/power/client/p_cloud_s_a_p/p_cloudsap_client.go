// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud s a p API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud s a p API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudSapGet(params *PcloudSapGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapGetOK, error)

	PcloudSapGetall(params *PcloudSapGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapGetallOK, error)

	PcloudSapPost(params *PcloudSapPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapPostOK, *PcloudSapPostCreated, *PcloudSapPostAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PcloudSapGet gets the information on an s a p profile
*/
func (a *Client) PcloudSapGet(params *PcloudSapGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudSapGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.sap.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudSapGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudSapGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.sap.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudSapGetall gets list of s a p profiles
*/
func (a *Client) PcloudSapGetall(params *PcloudSapGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudSapGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.sap.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/sap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudSapGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudSapGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.sap.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudSapPost creates a new s a p p VM instance
*/
func (a *Client) PcloudSapPost(params *PcloudSapPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudSapPostOK, *PcloudSapPostCreated, *PcloudSapPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudSapPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.sap.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/sap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudSapPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudSapPostOK:
		return value, nil, nil, nil
	case *PcloudSapPostCreated:
		return nil, value, nil, nil
	case *PcloudSapPostAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_s_a_p: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
