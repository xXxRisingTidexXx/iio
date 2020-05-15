package reports

type Reporter interface {
	Track(int, int)
	Report() *Report
}
