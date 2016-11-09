package amqp_test

import (
	"os"
	"time"

	"github.com/RackHD/voyager-utilities/amqp"
	"github.com/RackHD/voyager-utilities/random"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Amqp", func() {
	var rabbitMQURL string
	var client *amqp.Client
	var testExchange string
	var testExchangeType string
	var testQueueName string
	var testBindingKey string
	var testConsumerTag string
	var testReceiveMessage string
	var testCorrelationID string

	BeforeEach(func() {
		rabbitMQURL = os.Getenv("RABBITMQ_URL")
		testExchange = "amqp-service"
		testExchangeType = "topic"
		testQueueName = random.RandQueue()
		testBindingKey = "test-key"
		testConsumerTag = "testTag"
		testReceiveMessage = "Ginkgo Receive Test Message"
		testCorrelationID = "test ID"
	})

	AfterEach(func() {
		client.Close()
	})

	Describe("NewClient", func() {
		Context("when the RabbitMQ url is valid", func() {
			It("INTEGRATION should return non nil client", func() {
				client = amqp.NewClient(rabbitMQURL)
				Expect(client).ToNot(BeNil())
			})
		})

		Context("when the RabbitMQ url is not valid", func() {
			It("should retry 5 time before failure", func() {
				start := time.Now()
				client = amqp.NewClient("amqp://guest:guest@i-am-fake-server:5672/")
				end := time.Now()
				Expect(client).To(BeNil())
				Expect(end.Sub(start).Seconds()).To(BeNumerically(">=", 25))
			})
		})
	})

	Describe("AmqpListen", func() {
		Context("when the RabbitMQ url is valid", func() {
			It("INTEGRATION should be able to send and receive text message", func() {
				client = amqp.NewClient(rabbitMQURL)
				Expect(client).ToNot(BeNil())

				_, deliveries, err := client.Listen(testExchange, testExchangeType, testQueueName, testBindingKey, testConsumerTag)
				Expect(err).ToNot(HaveOccurred())

				err = client.Send(testExchange, testExchangeType, testBindingKey, testReceiveMessage, testCorrelationID, "")
				Expect(err).ToNot(HaveOccurred())

				d := <-deliveries
				d.Ack(false)
				Expect(string(d.Body)).To(Equal(testReceiveMessage))
				Expect(d.CorrelationId).To(Equal(testCorrelationID))
			})
		})
	})
})
