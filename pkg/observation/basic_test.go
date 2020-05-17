package observation_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/observation"
	"iio/pkg/test"
)

var _ = Describe("basic", func() {
	observe := func(observer observation.Observer, cost float64) {

	}
	test.With("should correctly collect all costs and produce a series", func() {
		observer := observe.NewB()
	})
})
