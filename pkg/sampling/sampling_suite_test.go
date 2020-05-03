package sampling_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestSampling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "sampling suite")
}

type Benchmarker ginkgo.Benchmarker

var RunSpecs = ginkgo.RunSpecs
var Fail = ginkgo.Fail
var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
var Measure = ginkgo.Measure
var RegisterFailHandler = gomega.RegisterFailHandler
var Expect = gomega.Expect
var And = gomega.And
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
