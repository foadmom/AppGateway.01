package comms

import "time"

type testComms struct {
}

type commsConfig struct {
}

var myConfig commsConfig = commsConfig{}
var myHeartbeatTestData []string

// ===========================================================================
// this is a testing CommsInterface
// ====================================
// SetHeartbeatTestData (test Heartbeat array)  // for unit test comm
// InitHeartbeat (channel)
//     set the heartbeatChannel
//     go heartBeatFlow
//         loop
//             delay
//             get next message from test data array
//             heartbeatChannel <- message

// SetConfigTestData (test Config array)  // for unit test comm
// InitConfig (channel)
//     set the Configchannel
//     go configFlow
//         loop
//             delay
//             get next config test Config array
//             Configchannel <- config
// ===========================================================================

// ===========================================================================
//
// ===========================================================================
func init() {
	// read the config and setup the comms params and properties
}

// ===========================================================================
// set the test data for the loop
// ===========================================================================
func (t Transport) SetHeartbeatTestData(testData []string) {
	myHeartbeatTestData = testData
}

// ===========================================================================
//
// ===========================================================================
func (t Transport) Start(HBChannel chan string) {
	go t.ListenForHeartbeat(HBChannel)
}

// ===========================================================================
// for unit testing only
// ===========================================================================
func (t Transport) ListenForHeartbeat(channel chan string) {

	for _index := 0; _index < len(myHeartbeatTestData); _index++ {
		time.Sleep(1 * time.Second)
		channel <- myHeartbeatTestData[_index]
	}
}

// ===========================================================================
// for unit testing only
// ===========================================================================
func (t Transport) ListenForServiceInfo(channel chan string) {

	for _index := 0; _index < len(myHeartbeatTestData); _index++ {
		time.Sleep(1 * time.Second)
		channel <- myHeartbeatTestData[_index]
	}
}
