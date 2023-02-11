package register

import (
	s "AppGateway/service"
	t "AppGateway/types"
	"time"
)

// ===================================
// ===================================
type ServiceList struct {
	Services     []*s.ServiceInfo
	CurrentIndex int
}

// ==================================================================
//
// ==================================================================
func createServiceList() *ServiceList {
	var _services []*s.ServiceInfo = make([]*s.ServiceInfo, 0)
	var _serviceList ServiceList = ServiceList{_services, 0}

	return &_serviceList
}

// ==================================================================
//
// ==================================================================
func (b *ServiceList) addService(service *s.ServiceInfo) {
	b.Services = append(b.Services, service)
}

// ==================================================================
//
// ==================================================================
func (b *ServiceList) checkForDeadService() {
	for _index := 0; _index < len(b.Services); _index++ {
		if b.Services[_index].Status != t.ONLINE {
			var _now uint = uint(time.Now().UnixMilli())
			var _lastUpdated uint = uint(b.Services[_index].TimeStamp.UnixMilli())

			var _diff int = int(_now - _lastUpdated)
			if _diff < 0 {
				return
			}
			if _diff > t.DROP_DEAD_TIMEOUT {
				b.removeService(_index)
			}

		}
	}
}

// ==================================================================
//
// ==================================================================
func (b *ServiceList) removeService(index int) {
	var _len int = len(b.Services)
	b.Services[index] = b.Services[_len-1]
	b.Services = b.Services[:_len-1]
}

// ==================================================================
// to be added
// this function should use semaphore to lock access to CurrentIndex,
// or use atomic operations
// this is only needed if you are running multiple threads/goroutines
// within this one process
// ==================================================================
func (b *ServiceList) adjustIndex() {
	var _len int = len(b.Services)
	var _index = b.CurrentIndex

	if _len == 0 {
		_index = -1
	} else if _len == 1 {
		_index = 0
	} else if _len > 1 {
		_index++
		if _index >= _len {
			_index = 0 // if the index has passed the last one reset it
		}
	}
	// this operation may have to be atomised
	b.CurrentIndex = _index
}

// ==================================================================
//
// ==================================================================
func (b *ServiceList) getService() *s.ServiceInfo {
	var _len int = len(b.Services)

	b.adjustIndex()
	if b.CurrentIndex == -1 {
		return nil
	}

	for _loopIndex := 0; _loopIndex < _len; _loopIndex++ {
		if b.Services[b.CurrentIndex].Status == t.ONLINE {
			return b.Services[b.CurrentIndex]
		} else {
			b.adjustIndex()
		}
	}
	return nil
}

// ==================================================================
//
// ==================================================================
func (b *ServiceList) findServiceInfoRecord(thisServiceInfo *s.ServiceInfo) *s.ServiceInfo {
	var _len int = len(b.Services)

	for _loopIndex := 0; _loopIndex < _len; _loopIndex++ {
		var _nextService *s.ServiceInfo = b.Services[_loopIndex]
		if (_nextService.Name == thisServiceInfo.Name) &&
			(_nextService.Host == thisServiceInfo.Host) &&
			(_nextService.Port == thisServiceInfo.Port) {
			return _nextService
		}
	}
	return nil
}

// ==================================================================
//
// ==================================================================
func (sl *ServiceList) ToString() string {
	var _list string

	for _, _item := range sl.Services {
		_list = _list + _item.ToString() + "\n"
	}

	return _list
}
