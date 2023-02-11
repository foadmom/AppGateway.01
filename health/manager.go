package health

import (
	"AppGateway/comms"
	"fmt"
)

type Config struct {
}

var myCongig Config = Config{}

// ===========================================================================
//
// ===========================================================================
func init() {
	// read config and initialise module
}

func Init(channel chan string) {
	go mainLoop(channel)
}

// ===========================================================================
// listen for a heartbeat message. convert json to ServiceInfo struct.
// return the ServiceInfo through the channel
// ===========================================================================
func mainLoop(channel chan string) {
	var _HBChannel chan string = make(chan string)

	comms.Instance().Start(_HBChannel)
	for true {
		var _message string
		select {
		case _message = <-_HBChannel:
			fmt.Println(_message)
		}
	}

}
