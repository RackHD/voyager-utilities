package models

// Inventory-Service AMQP Constants
const (
	InventoryServiceExchange     string = "voyager-inventory-service"
	InventoryServiceExchangeType string = "topic"
	InventoryServiceBindingKey   string = "requests"
	InventoryServiceConsumerTag  string = "consumer-tag"
)

// Constants for voyager-inventory-service DB node status
const (
	StatusAvailableLearning string = "Available-Learning"
	StatusAdded             string = "Added"
	StatusDiscovered        string = "Discovered"
)

// Constants for voyager-inventory-service DB node types
const (
	NodeTypeSwitch  string = "switch"
	NodeTypeCompute string = "compute"
)

// NodeMessage is a message about a node
type NodeMessage struct {
	ObjectType string `json:"type"`
	Action     string `json:"action"`
	NodeID     string `json:"nodeId"`
	NodeType   string `json:"nodeType"`
}

// NodeEntity is a node
type NodeEntity struct {
	ID     string `gorm:"primary_key"`
	Type   string `sql:"index"`
	Status string `sql:"index"`
}

// NodeEntityJSON is a NodeEntity, but for JSON structure
type NodeEntityJSON struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

// CmdMessage is a struct for an AMQP message for a command from the CLI
type CmdMessage struct {
	Command string      `json:"command"`
	Args    interface{} `json:"args"`
}

// CmdArg is a struct for a command argument
type CmdArg struct {
	Arg string `json:"arg"`
}

// IPEntity is a row in the IP Address table
type IPEntity struct {
	IPID      string `gorm:"primary_key" json:"ipID"`
	IPAddress string `sql:"index" json:"ipAddress"`
	NodeID    string `sql:"index" json:"nodeID"`
	PoolID    string `sql:"index" json:"poolID"`
	SubnetID  string `sql:"index" json:"subnetID"`
}
