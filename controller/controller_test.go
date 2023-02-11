package controller

import (
	"AppGateway/comms"
	"AppGateway/service"
	"AppGateway/types"
	"encoding/json"
	"testing"
	"time"
)

var testTime time.Time
var heartBeat_1 service.ServiceInfo

var myComms types.CommsInterface = comms.Instance()
var healthChannel chan string = make(chan string, 10)

func Init() {
	testTime, _ = time.Parse(time.RFC3339, "2022-12-13T16:47:24.928715693Z")
	heartBeat_1 = service.ServiceInfo{"Service.A", "localhost", "443", types.ONLINE, "nothing to log", testTime}

	myComms.Start(healthChannel)
}

func setUpTestData() []string {
	var _services int = 6
	var _loopCount int = 20
	var _dataSize int = _services * _loopCount
	var _data []string = make([]string, _dataSize)
	var _heartBeat [6]service.ServiceInfo
	_heartBeat[0] = service.ServiceInfo{"Service.A", "Server_A", "443", types.ONLINE, "nothing to log", testTime}
	_heartBeat[1] = service.ServiceInfo{"Service.A", "Server_B", "443", types.ONLINE, "nothing to log", testTime}
	_heartBeat[2] = service.ServiceInfo{"Service.A", "Server_C", "443", types.ONLINE, "nothing to log", testTime}
	_heartBeat[3] = service.ServiceInfo{"Service.B", "Server_A", "443", types.ONLINE, "nothing to log", testTime}
	_heartBeat[4] = service.ServiceInfo{"Service.C", "Server_B", "443", types.ONLINE, "nothing to log", testTime}
	_heartBeat[5] = service.ServiceInfo{"Service.C", "Server_C", "443", types.ONLINE, "nothing to log", testTime}

	for _i := 0; _i < _loopCount; _i++ {
		for _j := 0; _j < _services; _j++ {
			var _message []byte
			var _index int = (_i * 6) + _j
			var _timeVariance int = (_dataSize - _index) * -1
			_heartBeat[_j].TimeStamp = time.Now().Add(time.Second * time.Duration(_timeVariance))
			if _index%5 == 0 {
				_heartBeat[_j].Status = types.OFFLINE
			} else {
				_heartBeat[_j].Status = types.ONLINE
			}
			_message, _ = json.Marshal(_heartBeat[_j])
			_data[_index] = string(_message)
			//			time.Sleep(1 * time.Second)

		}

	}
	return _data
}

// ===========================================================================
//
// ===========================================================================
func TestSimulateMessageReceived(t *testing.T) {

	_testData := setUpTestData()

	myComms.SetHeartbeatTestData(_testData)
	MainLoop()

}
