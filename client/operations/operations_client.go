// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetRecommendations(params *GetRecommendationsParams, opts ...ClientOption) (*GetRecommendationsOK, error)

	GetRecommenders(params *GetRecommendersParams, opts ...ClientOption) (*GetRecommendersOK, error)

	GetZoneRecommendations(params *GetZoneRecommendationsParams, opts ...ClientOption) (*GetZoneRecommendationsOK, error)

	SendAddToCart(params *SendAddToCartParams, opts ...ClientOption) (*SendAddToCartOK, error)

	SendClickReco(params *SendClickRecoParams, opts ...ClientOption) (*SendClickRecoOK, error)

	SendViewProduct(params *SendViewProductParams, opts ...ClientOption) (*SendViewProductOK, error)

	SendViewReco(params *SendViewRecoParams, opts ...ClientOption) (*SendViewRecoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetRecommendations Get a set of recommendations
*/
func (a *Client) GetRecommendations(params *GetRecommendationsParams, opts ...ClientOption) (*GetRecommendationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecommendationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRecommendations",
		Method:             "POST",
		PathPattern:        "/personalization/recs/{siteId}/{recommenderName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRecommendationsReader{formats: a.formats},
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
	success, ok := result.(*GetRecommendationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRecommendations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRecommenders Get a list of recommenders that can be used in recommendation requests.
*/
func (a *Client) GetRecommenders(params *GetRecommendersParams, opts ...ClientOption) (*GetRecommendersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecommendersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRecommenders",
		Method:             "GET",
		PathPattern:        "/personalization/recommenders/{siteId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRecommendersReader{formats: a.formats},
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
	success, ok := result.(*GetRecommendersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRecommenders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZoneRecommendations Get a set of recommendations for a zone
*/
func (a *Client) GetZoneRecommendations(params *GetZoneRecommendationsParams, opts ...ClientOption) (*GetZoneRecommendationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneRecommendationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getZoneRecommendations",
		Method:             "POST",
		PathPattern:        "/personalization/{siteId}/zones/{zoneName}/recs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZoneRecommendationsReader{formats: a.formats},
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
	success, ok := result.(*GetZoneRecommendationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getZoneRecommendations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendAddToCart Tells the Einstein engine when a user adds an item to their cart.
*/
func (a *Client) SendAddToCart(params *SendAddToCartParams, opts ...ClientOption) (*SendAddToCartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendAddToCartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendAddToCart",
		Method:             "POST",
		PathPattern:        "/activities/{siteId}/addToCart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendAddToCartReader{formats: a.formats},
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
	success, ok := result.(*SendAddToCartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendAddToCart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendClickReco Tells the Einstein engine when a user clicks on a recommendation
*/
func (a *Client) SendClickReco(params *SendClickRecoParams, opts ...ClientOption) (*SendClickRecoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendClickRecoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendClickReco",
		Method:             "POST",
		PathPattern:        "/activities/{siteId}/clickReco",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendClickRecoReader{formats: a.formats},
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
	success, ok := result.(*SendClickRecoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendClickReco: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendViewProduct Tells the Einstein engine when a user views a product.
*/
func (a *Client) SendViewProduct(params *SendViewProductParams, opts ...ClientOption) (*SendViewProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendViewProductParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendViewProduct",
		Method:             "POST",
		PathPattern:        "/activities/{siteId}/viewProduct",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendViewProductReader{formats: a.formats},
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
	success, ok := result.(*SendViewProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendViewProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendViewReco Tells the Einstein engine when a user views a set of recommendations
*/
func (a *Client) SendViewReco(params *SendViewRecoParams, opts ...ClientOption) (*SendViewRecoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendViewRecoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendViewReco",
		Method:             "POST",
		PathPattern:        "/activities/{siteId}/viewReco",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendViewRecoReader{formats: a.formats},
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
	success, ok := result.(*SendViewRecoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendViewReco: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}