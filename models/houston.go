package models

// Houston RabbitMQ
const (
	HoustonExchange     string = "voyager-houston"
	HoustonExchangeType string = "topic"
	HoustonSendQueue    string = "voyager-houston-send-queue"
	HoustonReceiveQueue string = "voyager-houston-receive-queue"
	HoustonConsumerTag  string = "voyager-houston-consumer-tag"
	HoustonRepliesQueue string = "replies"
)

// Constants
const (
	IPAMServiceName string = "ipam"
)

// ServiceMessage is ...
type ServiceMessage struct {
	ServiceName string `json:"service_name"`
	State       string `json:"state"`
}

// PoolEntity is a
type PoolEntity struct {
	ID      string         `gorm:"primary_key"`
	Name    string         `sql:"index"`
	Subnets []SubnetEntity `gorm:"one2many:pool_subnets;"` // One-To-Many relationship (has many)
}

// SubnetEntity is a ...
type SubnetEntity struct {
	ID     string `gorm:"primary_key"`
	Name   string `sql:"index"`
	PoolID string // Foreign key for Pool (belongs to)
}
