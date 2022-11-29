package register

type Registry struct {
	Register map[string]*ServiceList
}

var ServiceRegister Registry = Registry{}
var ServiceRegisterPtr *Registry = &ServiceRegister

func init() {
	ServiceRegister.Register = make(map[string]*ServiceList)
}

// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================
// ==================================================================

// ==================================================================
// Check if there is a serviceName list for the service name
// If list for serviceName exists
//
//	    Check the list for service name, host, port
//	    if exists:
//	        Update the status and timestamp
//	    else
//	        Create a new entry for this service
//	        Add to the serviceName list
//	else
//	    Create a serviceName list
//	    Create a new entry for this service
//	    Add entry to the serviceName list
//	    Add serviceName list to registry map
//
// ==================================================================
func UpdateStatus(receivedServiceInfo *ServiceInfo) error {
	var _err error

	var _matchedService *ServiceInfo
	var _serviceList *ServiceList = ServiceRegisterPtr.Register[receivedServiceInfo.Name]
	if _serviceList == nil {
		_serviceList = ServiceRegisterPtr.createAndAddServiceList(receivedServiceInfo)
	} else {
		_matchedService = _serviceList.findServiceInfoRecord(receivedServiceInfo)
	}
	if _matchedService == nil {
		_serviceList.addService(receivedServiceInfo)
	} else {
		_matchedService.Status = receivedServiceInfo.Status
		_matchedService.TimeStamp = receivedServiceInfo.TimeStamp
	}
	_serviceList.checkForDeadService()

	return _err
}

func FindService(serviceName string) *ServiceInfo {
	return ServiceRegisterPtr.findService(serviceName)
}

func (r *Registry) createAndAddServiceList(receivedServiceInfo *ServiceInfo) *ServiceList {
	var _serviceList *ServiceList = createServiceList()
	r.Register[receivedServiceInfo.Name] = _serviceList
	return _serviceList
}

// ==================================================================
// Get list from registry map for the serviceName
// If does not exist return nil
// else
//
//	Scan the list, starting from the currentIndex for an entry with status = ONLINE
//	If found return
//	Increment and adjust currentIndex
//	Return service info
//
// else return nil
// Do we need a config backup in case heartbeat
// ==================================================================
func (r *Registry) findService(serviceName string) *ServiceInfo {
	var _service *ServiceInfo = nil
	var _serviceList *ServiceList = r.Register[serviceName]
	if _serviceList != nil {
		_service = _serviceList.getService()
	}

	return _service
}

func getRegistry() *Registry {
	return ServiceRegisterPtr
}