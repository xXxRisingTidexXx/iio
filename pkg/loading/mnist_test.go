package loading_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"iio/pkg/loading"
	"iio/pkg/test"
)

var _ = Describe("mnist", func() {
	test.With("should correctly iterate over both training and test datasets", func() {
		trainingLoader, testLoader := loading.NewMNISTLoaders()
		trainingLoader.Shuffle()
		for trainingLoader.Next() {
			Expect(trainingLoader.Batch(50)).To(HaveLen(50))
		}
		testLoader.Shuffle()
		for testLoader.Next() {
			Expect(testLoader.Batch(50)).To(HaveLen(50))
		}
	})
})
