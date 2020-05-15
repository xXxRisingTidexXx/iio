package estimate

type Estimator interface {
	Track(int, int)
	Estimate() *Report
}
