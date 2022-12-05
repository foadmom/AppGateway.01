package service

import (
	t "ServiceTools/types"
	"time"
)

// ===================================
type ServiceInfo struct {
	Name      string
	Host      string
	Port      string
	Status    t.ServiceStatus
	TimeStamp time.Time
}

// ===================================
// ===================================
func (s *ServiceInfo) Create(name, alias, host, port string, status t.ServiceStatus, timestamp time.Time) *ServiceInfo {
	var _serviceInfo ServiceInfo = ServiceInfo{name, host, port, status, timestamp}

	return &_serviceInfo
}
