package service

import (
	t "AppGateway/types"
	"encoding/json"
	"fmt"
	"time"
)

// ===================================
type ServiceInfo struct {
	Name      string          `json:"Name"`
	Host      string          `json:"Host"`
	Port      string          `json:"Port"`
	Status    t.ServiceStatus `json:"Staus"`
	Log       string          `json:"Log"` // in case the status has additional comments or errors or info
	TimeStamp time.Time       `json:"TimeStamp"`
}

// ===========================================================================
//
// ===========================================================================
func (s *ServiceInfo) Create(name, alias, host, port string, status t.ServiceStatus, log string, timestamp time.Time) *ServiceInfo {
	var _serviceInfo ServiceInfo = ServiceInfo{name, host, port, status, log, timestamp}

	return &_serviceInfo
}

// ===========================================================================
// Marshal a serviceInfo into a json
// ===========================================================================
func (s *ServiceInfo) MarshallServiceInfo() ([]byte, error) {
	var _json []byte
	var _err error

	_json, _err = json.Marshal(s)
	return _json, _err
}

// ===========================================================================
//
// ===========================================================================
func (s *ServiceInfo) UnMarshalServiceInfo(jsonService []byte) error {
	var _err error

	_err = json.Unmarshal(jsonService, s)
	return _err
}

// ===========================================================================
// Although it looks like a useless function that should be inlined, it is
// useful for simulating and testing events without comms or queues
// ===========================================================================
// func (s ServiceInfo) ReturnMessage(channel chan string) {
// 	channel <- s
// }

// ===========================================================================
//
// ===========================================================================
func (s *ServiceInfo) ToString() string {
	var _serviceSt string = fmt.Sprintf("%s - %s - %s - %s - %d - %s - %v", s.Name, s.Host, s.Port, s.Status, s.Log, s.TimeStamp)

	return _serviceSt
}
