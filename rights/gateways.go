// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

import "github.com/TheThingsNetwork/ttn/core/types"

const (
	// GatewaySettings is the types.Right to read and write access to the gateway settings
	GatewaySettings types.Right = "gateway:settings"

	// GatewayCollaborators is the types.Right to edit the gateway collaborators
	GatewayCollaborators types.Right = "gateway:collaborators"

	// GatewayDelete is the types.Right to delete a gatweay
	GatewayDelete types.Right = "gateway:delete"

	// GatewayLocation is the types.Right to view the exact location of the gateway, otherwise only approximate location will be shown
	GatewayLocation types.Right = "gateway:location"

	// GatewayStatus is the types.Right to view the gateway status and metrics about the gateway
	GatewayStatus types.Right = "gateway:status"

	// GatewayOwner is the types.Right that states that a collaborator is an owner
	GatewayOwner types.Right = "gateway:owner"

	// GatewayMessages is the types.Right to view the messages of a gateway
	GatewayMessages types.Right = "gateway:messages"
)

// AllGatewayUserRights are all the rights a user can have on a gateway
var AllGatewayUserRights = []types.Right{
	GatewaySettings,
	GatewayCollaborators,
	GatewayDelete,
	GatewayLocation,
	GatewayStatus,
	GatewayOwner,
	GatewayMessages,
}
