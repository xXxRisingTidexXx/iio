package test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// A shortcut for ginkgo/gomega DSL to execute specs in a panic-safe
// manner. Personally this function guarantees that the specified
// goroutine won't pass a panic to the higher level of ginkgo.
// Potential panic would be caught and asserted at the defer block.
// Obviously, panic's presence would cause ginkgo's tear down.
func With(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).To(gomega.BeNil())
		}()
		body()
	})
}

// A shortcut for ginkgo/gomega DSL to execute specs in a panic-safe
// manner. Personally this function expects a panic to occur. It's
// useful for specs specified to reproduce erroneous situations.
func Spare(text string, body func()) {
	ginkgo.It(text, func() {
		defer func() {
			gomega.Expect(recover()).NotTo(gomega.BeNil())
		}()
		body()
	})
}

// Universal equality assertion expecting arguments to equal each
// other. It differs from the standard gomega's function
// https://pkg.go.dev/github.com/onsi/gomega?tab=doc#Equal `cause it
// leverages https://golang.org/pkg/reflect/#DeepEqual but this one
// https://pkg.go.dev/github.com/google/go-cmp/cmp?tab=doc#Equal
// is more preferred `cause of a higher safety and more flexible
// customization.
func Equate(actual interface{}, expected interface{}) {
	gomega.Expect(cmp.Equal(actual, expected)).To(gomega.BeTrue())
}

// Universal inequality assertion based on widely-spread in our app
// https://pkg.go.dev/github.com/google/go-cmp/cmp?tab=doc#Equal
// equality function.
func Discern(actual interface{}, expected interface{}) {
	gomega.Expect(cmp.Equal(actual, expected)).To(gomega.BeFalse())
}
