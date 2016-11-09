ORGANIZATION = RackHD
PROJECT = utils
export RABBITMQ_URL = amqp://guest:guest@localhost:5672

default: deps test

deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
	go get ./...

test:
	ginkgo -r -race -trace -cover -randomizeAllSpecs -v
