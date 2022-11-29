package register

import (
	"testing"
	"time"
)

// type ServiceInfo struct {
// 	Name      string
// 	Host      string
// 	Port      string
// 	Status    ServiceStatus
// 	TimeStamp time.Time
// }

var ServiceArray []ServiceInfo = []ServiceInfo{
	{"AccountService-A", "host-11", "8080", ONLINE, time.Now()},
	{"AccountService-A", "host-12", "8080", ONLINE, time.Now()},
	{"AccountService-A", "host-13", "8080", ONLINE, time.Now()},
	{"AccountService-A", "host-14", "8080", ONLINE, time.Now()},
	{"AccountService-B", "host-21", "8080", ONLINE, time.Now()},
	{"AccountService-B", "host-22", "8080", ONLINE, time.Now()},
	{"AccountService-B", "host-23", "8080", ONLINE, time.Now()},
	{"AccountService-B", "host-24", "8080", ONLINE, time.Now()},
	{"AccountService-C", "host-31", "8080", ONLINE, time.Now()},
	{"AccountService-C", "host-32", "8080", ONLINE, time.Now()},
	{"AccountService-C", "host-33", "8080", ONLINE, time.Now()},
	{"AccountService-D", "host-41", "8080", ONLINE, time.Now()},
	{"AccountService-D", "host-42", "8080", ONLINE, time.Now()},
}

func loadTestData() {
	for _index := 0; _index < len(ServiceArray); _index++ {
		UpdateStatus(&ServiceArray[_index])
	}
}

func printRgistry(t *testing.T, _registry *Registry) {
	for _key, _value := range _registry.Register {
		t.Logf("key=%s  value=%v\n", _key, _value)
	}
}

func printService(t *testing.T, service *ServiceInfo) {
	t.Logf("service  %s    host=%s\n", service.Name, service.Host)
}

func findService(t *testing.T, serviceName string) *ServiceInfo {
	var _service *ServiceInfo
	for _index := 0; _index < 6; _index++ {
		_service = FindService("AccountService-B")
		printService(t, _service)
	}
	t.Log("===============================\n")
	return _service
}

func TestFindService(t *testing.T) {
	loadTestData()
	printRgistry(t, getRegistry())
	findService(t, "AccountService-B")
	_service := FindService("AccountService-B")
	t.Logf("\n............... %s\n", _service.Host)
	_service.Status = OFFLINE
	UpdateStatus(_service)
	_service = findService(t, "AccountService-B")
	_service = findService(t, "AccountService-B")
	t.Logf("\n............... %s\n", _service.Host)
	time.Sleep(7000 * time.Millisecond)
	_service = findService(t, "AccountService-B")
	UpdateStatus(_service)
	findService(t, "AccountService-B")
}
