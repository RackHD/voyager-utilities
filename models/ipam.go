package models

// Action types
const (
	CreateAction string = "create"
	ShowAction   string = "show"
	UpdateAction string = "update"
	DeleteAction string = "delete"
)

// Exchange types
const (
	PoolType        string = "pool"
	SubnetType      string = "subnet"
	ReservationType string = "reservation"
	LeaseType       string = "lease"
)

// Name types
const (
	DefaultKey string = "DEFAULT-MGN"
)

// range
// 172.31.128.1 is RackHD. 172.31.128.2 is the laptop
const (
	DefaultRange string = "172.31.128.0/24"
	DefaultStart string = "172.31.128.3"
	DefaultEnd   string = "172.31.128.255"
)

// IPAM RabbitMQ
const (
	IpamExchange     string = "voyager-ipam-service"
	IpamExchangeType string = "topic"
	IpamSendQueue    string = "ipam-send-queue"
	IpamReceiveQueue string = "ipam-receive-queue"
	IpamConsumerTag  string = "ipam-consumer-tag"
)

// Action types
const (
	RequestIPAction string = "request-ip"
	ReleaseIPAction string = "release-ip"
)

// IpamLeaseReq is used to send ip request to IPAM
type IpamLeaseReq struct {
	Action   string `json:"action"`
	SubnetID string `json:"subnetid"`
}

// AmqpResp is a general response for amqp response
type AmqpResp struct {
	Failed bool   `json:"failed"`
	Error  string `json:"error"`
}

// IpamLeaseResp is used to receive ip response from IPAM
type IpamLeaseResp struct {
	AmqpResp
	IP          string `json:"ip"`
	Reservation string `json:"reservation"`
}

// AMQPMsg is an AMQP message to be processed as an IPAM resource
type AMQPMsg struct {
	ObjectType string `json:"type"`
	Action     string `json:"action"`
}

// IpamReleaseReq is used to release ip
type IpamReleaseReq struct {
	Action      string `json:"action"`
	Reservation string `json:"reservation"`
}

// IPAMPoolMsg is a struct wrapping PoolV1 struct with AMQP objects inserted
type IPAMPoolMsg struct {
	ObjectType string      `json:"type"`
	Action     string      `json:"action"`
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Tags       []string    `json:"tags"`
	Metadata   interface{} `json:"metadata"`
}

// IPAMSubnetMsg is a struct wrapping SubnetV1 struct with AMQP objects inserted
type IPAMSubnetMsg struct {
	ObjectType string      `json:"type"`
	Action     string      `json:"action"`
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Tags       []string    `json:"tags"`
	Metadata   interface{} `json:"metadata"`
	Pool       string      `json:"pool"`
	Start      string      `json:"start"`
	End        string      `json:"end"`
}

// IPAMReservationMsg is a struct wrapping ReservationV1 struct with AMQP objects inserted
type IPAMReservationMsg struct {
	ObjectType string      `json:"type"`
	Action     string      `json:"action"`
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Tags       []string    `json:"tags"`
	Metadata   interface{} `json:"metadata"`
	Subnet     string      `json:"subnet"`
}

// IPAMLeaseMsg is a struct wrapping LeaseV1 struct with AMQP objects inserted
type IPAMLeaseMsg struct {
	ObjectType  string      `json:"type"`
	Action      string      `json:"action"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Tags        []string    `json:"tags"`
	Metadata    interface{} `json:"metadata"`
	Subnet      string      `json:"subnet"`
	Reservation string      `json:"reservation"`
	Address     string      `json:"address"`
}
