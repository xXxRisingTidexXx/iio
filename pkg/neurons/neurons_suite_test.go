package neurons_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNeurons(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "neurons suite")
}
