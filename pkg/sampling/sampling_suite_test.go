package sampling_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"math/rand"
	"testing"
)

func TestSampling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sampling Suite")
}

var _ = BeforeSuite(func() {
	rand.Seed(42)
})

// Declarations for Ginkgo DSL
type Benchmarker ginkgo.Benchmarker

var RunSpecs = ginkgo.RunSpecs
var Fail = ginkgo.Fail
var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
var Measure = ginkgo.Measure
var BeforeSuite = ginkgo.BeforeSuite

// Declarations for Gomega DSL
var RegisterFailHandler = gomega.RegisterFailHandler
var Expect = gomega.Expect

// Declarations for Gomega Matchers
var Equal = gomega.Equal
var BeNil = gomega.BeNil
var BeTrue = gomega.BeTrue
var BeFalse = gomega.BeFalse
var HaveOccurred = gomega.HaveOccurred

// Custom checkers
func ExpectNoPanic() {
	Expect(recover()).To(BeNil())
}

func ExpectPanic() {
	Expect(recover()).NotTo(BeNil())
}
