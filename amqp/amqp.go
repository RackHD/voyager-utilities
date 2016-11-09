package amqp

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	samqp "github.com/streadway/amqp"
)

// Client is a struct containing an AMQP connection
type Client struct {
	conn *samqp.Connection
	done chan error
}

// NewClient creates a new amqp client
func NewClient(url string) *Client {
	client := Client{}
	var err error

	for i := 1; i <= 5; i++ {
		client.conn, err = samqp.Dial(url)
		if err == nil {
			return &client
		}
		log.Errorf("Could not connect due to: %s, retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	log.Errorf("Failed to create new client for %s\n", url)
	return nil
}

// Close closes the  AMQP server
func (a *Client) Close() (err error) {
	if a != nil && a.conn != nil {
		return a.conn.Close()
	}
	return nil
}

// Send sends a message to the specified exchange
func (a *Client) Send(exchange, exchangeType, routingKey, body, correlationID, replyTo string) error {
	channel, err := a.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to create channel: %s", err)
	}

	log.Debugf("declaring %q Exchange (%q)", exchangeType, exchange)
	err = channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exhange: %s", err)
	}

	log.Debugf("publishing %dB body (%q)", len(body), body)
	err = channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		samqp.Publishing{
			Headers:         samqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			CorrelationId:   correlationID,
			ReplyTo:         replyTo,
			Body:            []byte(body),
			DeliveryMode:    samqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,               // 0-9
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish: %s", err)
	}
	return nil
}

// Listen connects to an AMQP exchange and returns a channel that will contain messages
func (a *Client) Listen(exchange, exchangeType, queueName, bindingKey, consumerTag string) (*samqp.Channel, <-chan samqp.Delivery, error) {
	channel, err := a.conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create channel: %s", err)
	}

	log.Debugf("declaring exchange (%q)", exchange)
	err = channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare exchange: %s", err)
	}

	log.Debugf("declaring queue %q", queueName)
	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		true,      // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare queue: %s", err)
	}

	log.Debugf("declared queue (%q %d messages, %d consumers), binding to exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, bindingKey)

	err = channel.QueueBind(
		queue.Name, // name of the queue
		bindingKey, // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to bind queue: %s", err)
	}

	deliveries, err := channel.Consume(
		queue.Name,  // name
		consumerTag, // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to consume queue: %s", err)
	}

	log.Infof("Listening on exchange: %s\n", exchange)
	return channel, deliveries, nil
}
