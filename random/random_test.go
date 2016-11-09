package random_test

import (
	. "github.com/RackHD/voyager-utilities/random"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rand", func() {

	Context("When RandQueue is called multiple times", func() {
		It("should return a different string each time", func() {
			string1 := RandQueue()
			string2 := RandQueue()
			string3 := RandQueue()

			Expect(string1).ToNot(Equal(string2))
			Expect(string1).ToNot(Equal(string3))
			Expect(string2).ToNot(Equal(string3))
		})

	})

})
