package register

import "time"

// ===================================
// ===================================
type ServiceList struct {
	Services     []*ServiceInfo
	CurrentIndex int
}

// ===================================
// ===================================
func createServiceList() *ServiceList {
	var _services []*ServiceInfo = make([]*ServiceInfo, 0)
	var _serviceList ServiceList = ServiceList{_services, 0}

	return &_serviceList
}

func (b *ServiceList) addService(service *ServiceInfo) {
	b.Services = append(b.Services, service)
}

func (b *ServiceList) checkForDeadService() {
	for _index := 0; _index < len(b.Services); _index++ {
		if b.Services[_index].Status != ONLINE {
			var _now uint = uint(time.Now().UnixMilli())
			var _lastUpdated uint = uint(b.Services[_index].TimeStamp.UnixMilli())

			var _diff int = int(_now - _lastUpdated)
			if _diff < 0 {
				return
			}
			if _diff > DROP_DEAD_TIMEOUT {
				b.removeService(_index)
			}

		}
	}
}

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

func (b *ServiceList) getService() *ServiceInfo {
	var _len int = len(b.Services)

	b.adjustIndex()
	if b.CurrentIndex == -1 {
		return nil
	}

	for _loopIndex := 0; _loopIndex < _len; _loopIndex++ {
		if b.Services[b.CurrentIndex].Status == ONLINE {
			return b.Services[b.CurrentIndex]
		} else {
			b.adjustIndex()
		}
	}
	return nil
}

func (b *ServiceList) findServiceInfoRecord(thisServiceInfo *ServiceInfo) *ServiceInfo {
	var _len int = len(b.Services)

	for _loopIndex := 0; _loopIndex < _len; _loopIndex++ {
		var _nextService *ServiceInfo = b.Services[_loopIndex]
		if (_nextService.Name == thisServiceInfo.Name) &&
			(_nextService.Host == thisServiceInfo.Host) &&
			(_nextService.Port == thisServiceInfo.Port) {
			return _nextService
		}
	}
	return nil
}
