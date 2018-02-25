package dplay

// ServerOptions sets options, common to DirectPlay.
type ServerOptions struct {
	// EnumResponseData returns application-specific data package
	// for enum requests.
	EnumResponseData func() []byte

	// Name is this server's visible name.
	Name string

	// MaxPlayers imposes max upper limit on player count.
	MaxPlayers uint

	// IsPassworded is true if passowrd is required to connect
	// to the server.
	// IsPassworded bool

	// ApplicationInstanceGUID is the instance GUID that
	// identifies the game session. Refer to MS-DPDX 2.2.5.
	ApplicationInstanceGUID []byte
	// ApplicationGUID is the current application's GUID.
	ApplicationGUID []byte
}
