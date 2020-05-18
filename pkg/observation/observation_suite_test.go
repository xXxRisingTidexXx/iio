package observation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestObserve(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "observation suite")
}
