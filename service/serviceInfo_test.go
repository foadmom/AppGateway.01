package service

import (
	"AppGateway/types"
	"reflect"
	"testing"
	"time"
)

var testTime time.Time

// var JSONMessage_1 string = `{"Name":"Service.A","Host":"localhost","Port":"443","Staus":0,"Log":"nothing to log","TimeStamp":"2022-12-13T16:47:24.928715693Z"}`
// var StruMessage_1 ServiceInfo = ServiceInfo{"Service.A", "localhost", "443", types.ONLINE, "nothing to log", testTime}
var JSONMessage_1 string
var StruMessage_1 ServiceInfo

func Init() {
	testTime, _ = time.Parse(time.RFC3339, "2022-12-13T16:47:24.928715693Z")
	JSONMessage_1 = `{"Name":"Service.A","Host":"localhost","Port":"443","Staus":0,"Log":"nothing to log","TimeStamp":"2022-12-13T16:47:24.928715693Z"}`
	StruMessage_1 = ServiceInfo{"Service.A", "localhost", "443", types.ONLINE, "nothing to log", testTime}
}

func TestServiceInfo_Marshall(t *testing.T) {
	Init()
	_got, _err := StruMessage_1.MarshallServiceInfo()
	if _err != nil {
		t.Errorf("ServiceInfo.MarshallServiceInfo() error = %v\n", _err)
	} else if !reflect.DeepEqual(_got, []byte(JSONMessage_1)) {
		t.Errorf("\nexpected %s\nbut got  %s\n", string(JSONMessage_1), string(_got))
	}

}

func TestServiceInfo_UnMarshall(t *testing.T) {
	Init()
	var _serviceInfo ServiceInfo
	_err := _serviceInfo.UnMarshalServiceInfo([]byte(JSONMessage_1))
	if _err != nil {
		t.Errorf("ServiceInfo.UnMarshallServiceInfo() error = %v\n", _err)
	} else if !reflect.DeepEqual(_serviceInfo, StruMessage_1) {
		t.Errorf("\nexpected %v\nbut got  %v\n", StruMessage_1, _serviceInfo)
	}

}
