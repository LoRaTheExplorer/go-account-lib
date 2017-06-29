// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

import "github.com/TheThingsNetwork/ttn/core/types"

const (
	// AppSettings is the types.Right to read and write access to the settings, devices and access keys of the application
	AppSettings types.Right = "settings"

	// AppCollaborators is the types.Right to edit and modify collaborators of the application
	AppCollaborators types.Right = "collaborators"

	// AppDelete is the types.Right to delete the application
	AppDelete types.Right = "delete"

	// Devices is the types.Right to list, edit and remove devices for the application on a handler
	Devices types.Right = "devices"

	// ReadUplink is the types.Right to view messages sent by devices of the application
	ReadUplink types.Right = "messages:up:r"

	// WriteUplink is the types.Right to send messages to the application
	WriteUplink types.Right = "messages:up:w"

	// WriteDownlink is the types.Right to send messages to devices of the application
	WriteDownlink types.Right = "messages:down:w"
)

// AllAppUserRights are all the rights a user can have on an application
var AllAppUserRights = []types.Right{
	AppSettings,
	AppCollaborators,
	AppDelete,
	Devices,
}
