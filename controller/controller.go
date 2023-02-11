package controller

import (
	"AppGateway/comms"
	"AppGateway/register"
	"AppGateway/service"
	"fmt"
)

// ===========================================================================
//
// ===========================================================================
func init() {

}

// ===========================================================================
//
// ===========================================================================
func MainLoop() {
	var _healthChannel chan string = make(chan string, 10)
	//    var _configChannel chan string

	_comms := comms.Instance()
	_comms.Start(_healthChannel)
	for true {
		var _message string
		select {
		case _message = <-_healthChannel:
			processHeartBeat(_message)
		}
	}

}

// ===========================================================================
//
// ===========================================================================
func getConfig() {
}

// ===========================================================================
//
// ===========================================================================
func processHeartBeat(message string) {
	fmt.Printf("received %v\n", message)
	var _serviceInfo service.ServiceInfo = service.ServiceInfo{}
	var _err error = _serviceInfo.UnMarshalServiceInfo([]byte(message))
	if _err == nil {
		_err = register.UpdateStatus(_serviceInfo)
		if _err == nil {
			fmt.Println("==============================================================")
			fmt.Println(register.GetRegistry().ToString())
			// register.PrintRegistry(register.GetRegistry())
		}
	}

}

// ===========================================================================
//
// ===========================================================================
func processRegistration(message string) {
	fmt.Println("received", message)

}

// ===========================================================================
//
// ===========================================================================
func processConfigUpdate(message string) {

}

// ===========================================================================
//
// ===========================================================================
func processCertificateUpdate() {

}
