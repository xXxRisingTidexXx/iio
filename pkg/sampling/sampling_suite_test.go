package sampling_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"math/rand"
	"testing"
)

func TestSampling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "sampling suite")
}

var _ = BeforeSuite(func() {
	rand.Seed(42)
})

type Benchmarker ginkgo.Benchmarker

var RunSpecs = ginkgo.RunSpecs
var Fail = ginkgo.Fail
var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
var Measure = ginkgo.Measure
var BeforeSuite = ginkgo.BeforeSuite
var RegisterFailHandler = gomega.RegisterFailHandler
var Expect = gomega.Expect
var Equal = gomega.Equal
var BeNil = gomega.BeNil
var BeTrue = gomega.BeTrue
var BeFalse = gomega.BeFalse

func With(text string, body func()) {
	It(text, func() {
		defer func() {
			Expect(recover()).To(BeNil())
		}()
		body()
	})
}

func Spare(text string, body func()) {
	It(text, func() {
		defer func() {
			Expect(recover()).NotTo(BeNil())
		}()
		body()
	})
}
