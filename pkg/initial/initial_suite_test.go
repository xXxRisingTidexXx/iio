package initial_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "initial suite")
}
