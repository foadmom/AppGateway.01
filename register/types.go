package register

type ServiceStatus int

const DROP_DEAD_TIMEOUT = 5000

// ===================================
const ONLINE ServiceStatus = 0
const DRAINING ServiceStatus = 1
const OFFLINE ServiceStatus = -1
