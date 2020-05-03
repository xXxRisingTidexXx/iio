package test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func With(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).To(gomega.BeNil())
		}()
		body()
	})
}

func Spare(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).NotTo(gomega.BeNil())
		}()
		body()
	})
}

func Equate(actual interface{}, expected interface{}) {
	gomega.Expect(actual).To(gomega.Equal(expected))
}

func Comply(actual interface{}, expected interface{}) {
	gomega.Expect(cmp.Equal(actual, expected)).To(gomega.BeTrue())
}

func Discern(actual interface{}, expected interface{}) {
	gomega.Expect(cmp.Equal(actual, expected)).To(gomega.BeFalse())
}
