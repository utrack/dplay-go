package dplay

// ServerOptions sets options, common to DirectPlay.
type ServerOptions struct {
	// TODO: Allows point values? Use int otherwise.
	// Version is a server's version integer encoded as a string.
	Version int

	// TODO make sep struct with checkers.
	// This server's ID as a string.
	// Format: 00000000-00000000-00000000-00000000
	ServerID string
	// Description is this own server's description.
	Description string
	// Name is this server's visible name.
	Name string

	// MaxPlayers imposes max upper limit on player count.
	MaxPlayers uint

	// IsPassworded is true if passowrd is required to connect
	// to the server.
	IsPassworded bool

	// ApplicationInstanceGUID is the instance GUID that
	// identifies the game session. Refer to MS-DPDX 2.2.5.
	ApplicationInstanceGUID []byte
	// ApplicationGUID is the current application's GUID.
	ApplicationGUID []byte
}
