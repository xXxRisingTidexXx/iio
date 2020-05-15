package loading_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestLoading(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "loading suite")
}
