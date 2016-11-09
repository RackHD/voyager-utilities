package random_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rand Suite")
}

var _ = BeforeSuite(func() {
	rand.Seed(time.Now().UTC().UnixNano())

})
