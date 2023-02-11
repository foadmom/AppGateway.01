package types

import (
	"time"
)

type ServiceStatus int

const DROP_DEAD_TIMEOUT = 5000

// ===================================
const ONLINE ServiceStatus = 0
const DRAINING ServiceStatus = 1
const OFFLINE ServiceStatus = -1

// ===================================
//
// ===================================
type ServiceData interface {
	Name() string
	Host() string
	Port() string
	Status() ServiceStatus
	Log() string
	TimeStamp() time.Time
}

// ===================================
//
// ===================================
type CommsInterface interface {
	ListenForServiceInfo(HBChannel chan string)
	ListenForHeartbeat(HBChannel chan string)
	Start(HBChannel chan string)

	SetHeartbeatTestData(testData []string)
}

// ===================================
// standard json message
// ===================================
type JSON string

type MessageHeader struct {
	UUID          string    `json:"UUID"`       // one transaction/Meesage could have multiple messages with multiple requests
	SequenceID    string    `json:"SequenceID"` // Optional: this could be id or a sequence no. in one UUID/transaction
	HostName      string    `json:"HostName"`
	ApplicationID string    `json:"ApplicationID"`
	Timestamp     time.Time `json:"Timestamp"`
	UserId        string    `json:"UserId"`   // optional. used for some services
	Password      string    `json:"Password"` // optional. used for some services
}

type MessagePayload struct {
	RequestType string `json:"RequestType"` // Request ID or Code identifying what is being requested
	Data        JSON   `json:"Data"`        // This is the actual payload data which is different for each RequstType
}

type Message struct {
	Header MessageHeader `json:"Header"`
}
