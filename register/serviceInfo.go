package register

import "time"

// ===================================
type ServiceInfo struct {
	Name      string
	Host      string
	Port      string
	Status    ServiceStatus
	TimeStamp time.Time
}

// ===================================
// ===================================
func (s *ServiceInfo) Create(name, alias, host, port string, status ServiceStatus, timestamp time.Time) *ServiceInfo {
	var _serviceInfo ServiceInfo = ServiceInfo{name, host, port, status, timestamp}

	return &_serviceInfo
}
