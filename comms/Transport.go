package comms

import (
	"AppGateway/types"
)

type Transport struct {
}

var transport Transport = Transport{}

// ===========================================================================
//
// ===========================================================================
func init() {
	// read the config and setup the comms params and properties
	transport = Transport{}
}

// ===========================================================================
//
// ===========================================================================
func Instance() types.CommsInterface {
	return transport
}
