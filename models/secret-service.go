package models

// AMQP constants for voyager-secret-service
const (
  SecretExchange     string = "voyager-secret-service"
  SecretExchangeType string = "topic"
  SecretQueueName    string = "voyager-secret-service-queue"
  SecretBindingKey   string = "requests"
  SecretConsumerTag  string = "consumer-tag"
)

// Credentials is the about the credential to return
type Credentials struct {
  Username string `json:"username"`
  Password string `json:"password"`
}
