package dplay_test

import (
	"fmt"
	"log"
	"testing"

	dplay "github.com/utrack/dplay-go"
)

func TestRunServerFreelancer(t *testing.T) {

	ServerID := "00000000-00000000-00000000-00000000"
	Description := "desc"
	Version := "16777227" // grep server_version of FOS
	adata := dplay.NewDataPackage()
	ispvp := 1
	adata.AddString(fmt.Sprintf(`%v:1:%v:-1910309061:%v:`, ispvp, Version, ServerID))
	adata.AddStringUnicodeTerm(Description)

	opts := dplay.ServerOptions{
		EnumResponseData: func() []byte { return adata.Bytes() },
		Name:             "OpenFL Server",
		PlayerCount:      func() uint { return 0 },
		MaxPlayers:       5,
		ApplicationGUID: []byte{
			0x26, 0xf0, 0x90, 0xa6, 0xf0, 0x26, 0x57, 0x4e, 0xac, 0xa0, 0xec, 0xf8,
			0x68, 0xe4, 0x8d, 0x21,
		},
		ApplicationInstanceGUID: []byte{
			0xa8, 0xc6, 0x27, 0x1d, 0x41, 0x66, 0xd8, 0x49, 0x89, 0xeb, 0x1e,
			0xbc, 0x42, 0x21, 0xca, 0xe9,
		},
	}
	s, err := dplay.NewServer(":2302", opts)
	if err != nil {
		log.Fatal(err)
	}
	s.Listen()
}
