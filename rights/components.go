// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

import "github.com/TheThingsNetwork/ttn/core/types"

const (
	// ComponentSettings is the types.Right to read and write access to the settings and access key of a network component
	ComponentSettings types.Right = "component:settings"

	// ComponentDelete is the types.Right to delete the network component
	ComponentDelete types.Right = "component:delete"

	// ComponentCollaborators is the types.Right to view and edit component collaborators
	ComponentCollaborators = "component:collaborators"
)

// AllComponentUserRights are all the rights a user can have to a component
var AllComponentUserRights = []types.Right{
	ComponentSettings,
	ComponentDelete,
	ComponentCollaborators,
}
