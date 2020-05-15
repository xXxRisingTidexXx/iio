package estimate

type Report struct {
	Classes  []*Record
	MacroAvg *Record
	Accuracy float64
}

func (report *Report) String() string {
	return ""
}
