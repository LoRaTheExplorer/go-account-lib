// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package account

import (
	"errors"
	"fmt"

	"github.com/TheThingsNetwork/go-account-lib/auth"
	"github.com/TheThingsNetwork/go-account-lib/scope"
	"github.com/TheThingsNetwork/ttn/core/types"
)

// ListGateways list all gateways
func (a *Account) ListGateways() (gateways []Gateway, err error) {
	err = a.get(a.auth, "/api/v2/gateways", &gateways)
	return gateways, err
}

// FindGateway returns the information about a specific gateay
func (a *Account) FindGateway(gatewayID string) (gateway Gateway, err error) {
	err = a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), &gateway)
	return gateway, err
}

type registerGatewayReq struct {
	// ID is the ID of the new gateway (required)
	ID string `json:"id"`

	// Country is the country code of the new gateway (required)
	FrequencyPlan string `json:"frequency_plan"`

	// Location is the location of the new gateway
	Location *Location `json:"location,omitempty"`
}

// RegisterGateway registers a new gateway on the account server
func (a *Account) RegisterGateway(gatewayID string, frequencyPlan string, location *Location) (gateway Gateway, err error) {
	if gatewayID == "" {
		return gateway, errors.New("Cannot create gateway: no ID given")
	}

	if frequencyPlan == "" {
		return gateway, errors.New("Cannot create gateway: no FrequencyPlan given")
	}

	req := registerGatewayReq{
		ID:            gatewayID,
		FrequencyPlan: frequencyPlan,
		Location:      location,
	}

	err = a.post(a.auth, "/api/v2/gateways", req, &gateway)
	return gateway, err
}

// FindGateway returns the information about a specific gateay
func (a *Account) GetGatewayToken(gatewayID string) (*Token, error) {
	var gateway Gateway
	err := a.get(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), &gateway)
	if err != nil {
		return nil, err
	}

	// already have the token!
	if gateway.Token != nil {
		return gateway.Token, nil
	}

	// didn't get a token, but we can use the Key to get one
	if gateway.Key != "" {
		var token Token
		err = a.get(auth.AccessKey(gateway.Key), fmt.Sprintf("/api/v2/gateways/%s/token", gatewayID), &token)
		if err != nil {
			return nil, err
		}

		return &token, nil
	}

	return nil, errors.New("Cannot get token using this authentication method")
}

// DeleteGateway removes a gateway from the account server
func (a *Account) DeleteGateway(gatewayID string) error {
	return a.del(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID))
}

// GrantGatewayRights grants rights to a collaborator of the gateway
func (a *Account) GrantGatewayRights(gatewayID string, username string, rights []types.Right) error {
	req := grantReq{
		Rights: rights,
	}
	return a.put(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/collaborators/%s", gatewayID, username), req, nil)
}

// RetractGatewayRights removes rights from a collaborator of the gateway
func (a *Account) RetractGatewayRights(gatewayID string, username string) error {
	return a.del(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s/collaborators/%s", gatewayID, username))
}

// GatewayEdits contains editable fields of gateways
type GatewayEdits struct {
	Owner         string        `json:"owner,omitempty"`
	PublicRights  []types.Right `json:"public_rights,omitempty"`
	FrequencyPlan string        `json:"frequency_plan,omitempty"`
	Location      *Location     `json:"location,omitempty"`
}

// EditGateway edits the fields of a gateway
func (a *Account) EditGateway(gatewayID string, edits GatewayEdits) error {
	return a.patch(a.auth.WithScope(scope.Gateway(gatewayID)), fmt.Sprintf("/api/v2/gateways/%s", gatewayID), edits, nil)
}

// TransferOwnership transfers the owenership of the gateway to another user
func (a *Account) TransferOwnership(gatewayID, username string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Owner: username,
	})
}

// SetPublicRights changes the publicily visible rights of the gateway
func (a *Account) SetPublicRights(gatewayID string, rights []types.Right) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		PublicRights: rights,
	})
}

// ChangeFrequencyPlan changes the requency plan of a gateway
func (a *Account) ChangeFrequencyPlan(gatewayID, plan string) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		FrequencyPlan: plan,
	})
}

// ChangeLocation changes the location of the gateway
func (a *Account) ChangeLocation(gatewayID string, latitude, longitude float64) error {
	return a.EditGateway(gatewayID, GatewayEdits{
		Location: &Location{
			Longitude: longitude,
			Latitude:  latitude,
		},
	})
}
