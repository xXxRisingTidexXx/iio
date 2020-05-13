package costs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCosts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "costs suite")
}
