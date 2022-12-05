package register

import (
	s "ServiceTools/service"
	"ServiceTools/types"
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

var ServiceArray []s.ServiceInfo = []s.ServiceInfo{
	{Name: "AccountService-A", Host: "host-11", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-A", Host: "host-12", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-A", Host: "host-13", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-A", Host: "host-14", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-B", Host: "host-21", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-B", Host: "host-22", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-B", Host: "host-23", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-B", Host: "host-24", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-C", Host: "host-31", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-C", Host: "host-32", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-C", Host: "host-33", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-D", Host: "host-41", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
	{Name: "AccountService-D", Host: "host-42", Port: "8080", Status: types.ONLINE, TimeStamp: time.Now()},
}

func loadTestData() {
	for _index := 0; _index < len(ServiceArray); _index++ {
		UpdateStatus(ServiceArray[_index])
	}
}

func printRgistry(t *testing.T, _registry *Registry) {
	for _key, _value := range _registry.Register {
		t.Logf("key=%s  value=%v\n", _key, _value)
	}
}

func printService(t *testing.T, service *s.ServiceInfo) {
	t.Logf("service  %s    host=%s\n", service.Name, service.Host)
}

func findService(t *testing.T, serviceName string) s.ServiceInfo {
	var _service s.ServiceInfo
	for _index := 0; _index < 6; _index++ {
		_service = FindService("AccountService-B")
		printService(t, &_service)
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
	_service.Status = types.OFFLINE
	UpdateStatus(_service)
	_service = findService(t, "AccountService-B")
	_service = findService(t, "AccountService-B")
	t.Logf("\n............... %s\n", _service.Host)
	time.Sleep(7000 * time.Millisecond)
	_service = findService(t, "AccountService-B")
	UpdateStatus(_service)
	findService(t, "AccountService-B")
}
