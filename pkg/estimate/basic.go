package estimate

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"sync"
)

func NewBasicEstimator(classNumber int) Estimator {
	if classNumber <= 1 {
		panic(fmt.Sprintf("reports: got invalid class number, %d", classNumber))
	}
	return &basicEstimator{
		classNumber:     classNumber,
		confusionMatrix: mat.NewDense(classNumber, classNumber, nil),
	}
}

type basicEstimator struct {
	sync.Mutex
	classNumber     int
	confusionMatrix *mat.Dense
}

func (estimator *basicEstimator) Track(actual, ideal int) {
	if actual < 0 || actual >= estimator.classNumber {
		panic(fmt.Sprintf("reports: invalid actual label, %d", actual))
	}
	if ideal < 0 || ideal >= estimator.classNumber {
		panic(fmt.Sprintf("reports: invalid ideal label, %d", ideal))
	}
	estimator.Lock()
	estimator.confusionMatrix.Set(actual, ideal, estimator.confusionMatrix.At(actual, ideal)+1)
	estimator.Unlock()
}

func (estimator *basicEstimator) Estimate() *Report {
	classes := make([]*Record, estimator.classNumber)
	totalAccuracy, totalSupport, totalPrecision, totalRecall, totalF1Score := 0.0, 0.0, 0.0, 0.0, 0.0
	for i := 0; i < estimator.classNumber; i++ {
		accuracy := estimator.confusionMatrix.At(i, i)
		precision := 0.0
		if predicted := mat.Sum(estimator.confusionMatrix.RowView(i)); predicted != 0 {
			precision = accuracy / predicted
		}
		recall, support := 0.0, mat.Sum(estimator.confusionMatrix.ColView(i))
		if support != 0 {
			recall = accuracy / support
		}
		f1Score := 0.0
		if precision != 0 && recall != 0 {
			f1Score = 2 * precision * recall / (precision + recall)
		}
		totalAccuracy += accuracy
		totalSupport += support
		totalPrecision += precision
		totalRecall += recall
		totalF1Score += f1Score
		classes[i] = NewRecord(int(support), precision, recall, f1Score)
	}
	if totalSupport != 0 {
		totalAccuracy /= totalSupport
	}
	return NewReport(
		classes,
		NewRecord(
			int(totalSupport),
			totalPrecision/float64(estimator.classNumber),
			totalRecall/float64(estimator.classNumber),
			totalF1Score/float64(estimator.classNumber),
		),
		totalAccuracy,
	)
}
