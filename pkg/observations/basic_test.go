package observations_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/observations"
	"iio/pkg/test"
)

var _ = Describe("basic", func() {
	observe := func(observer observations.Observer, cost float64) {

	}
	test.With("should correctly collect all costs and produce a series", func() {
		observer := observe.NewB()
	})
})
