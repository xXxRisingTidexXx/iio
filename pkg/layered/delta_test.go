package layered_test

import (
	. "github.com/onsi/ginkgo"
	"iio/pkg/layered"
	"iio/pkg/test"
)

var _ = Describe("delta", func() {
	test.With("should correctly calculate layer delta", func() {
		delta := layered.NewDelta(
			test.Vector(0.41794311, 0.61151192),
			test.Vector(0.69353618, 0.59955494, 0.4991476, 0.41237347),
			-0.25,
		)
		test.Equate(
			delta.Weights,
			test.Matrix(
				2,
				4,
				-0.07246466699167994,
				-0.06264496405986585,
				-0.052153825073259004,
				-0.04308716263332293,
				-0.1060264102553164,
				-0.0916587481262212,
				-0.076308676809848,
				-0.0630428230991906,
			),
		)
		test.Equate(delta.Biases, test.Vector(-0.1044857775, -0.15287798))
	})
})
