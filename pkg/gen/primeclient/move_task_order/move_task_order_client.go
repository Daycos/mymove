// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new move task order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for move task order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FetchMTOUpdates gets all move orders

Gets all move orders
*/
func (a *Client) FetchMTOUpdates(params *FetchMTOUpdatesParams) (*FetchMTOUpdatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchMTOUpdatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fetchMTOUpdates",
		Method:             "GET",
		PathPattern:        "/move-task-orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FetchMTOUpdatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FetchMTOUpdatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchMTOUpdates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMoveTaskOrderCustomer gets a the customer associated with a move task order ID

Gets a the customer associated with a move task order ID
*/
func (a *Client) GetMoveTaskOrderCustomer(params *GetMoveTaskOrderCustomerParams) (*GetMoveTaskOrderCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMoveTaskOrderCustomerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMoveTaskOrderCustomer",
		Method:             "GET",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}/customer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMoveTaskOrderCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMoveTaskOrderCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMoveTaskOrderCustomer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateMoveTaskOrderPostCounselingInformation updates move task order s post counseling information

Updates move task order's post counseling information
*/
func (a *Client) UpdateMoveTaskOrderPostCounselingInformation(params *UpdateMoveTaskOrderPostCounselingInformationParams) (*UpdateMoveTaskOrderPostCounselingInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMoveTaskOrderPostCounselingInformationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMoveTaskOrderPostCounselingInformation",
		Method:             "PATCH",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}/post-counseling-info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateMoveTaskOrderPostCounselingInformationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMoveTaskOrderPostCounselingInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMoveTaskOrderPostCounselingInformation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
