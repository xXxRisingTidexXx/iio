package sampling_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSampling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sampling Suite")
}
