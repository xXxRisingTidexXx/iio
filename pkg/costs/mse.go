package costs

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func NewMSECostFunction() CostFunction {
	return &mseCostFunction{}
}

type mseCostFunction struct{}

func (costFunction *mseCostFunction) Evaluate(actual mat.Vector) int {
	if actual == nil {
		panic("costs: mse cost function got nil actual")
	}
	index, max := 0, actual.AtVec(0)
	for i := 1; i < actual.Len(); i++ {
		if value := actual.AtVec(i); value > max {
			index = i
			max = value
		}
	}
	return index
}

func (costFunction *mseCostFunction) Cost(actual mat.Vector, label int) float64 {
	costFunction.checkInput(actual, label)
	cost := 0.0
	for i := 0; i < actual.Len(); i++ {
		value := actual.AtVec(i)
		if i == label {
			value = 1.0 - value
		}
		cost += value * value
	}
	return cost
}

func (costFunction *mseCostFunction) checkInput(actual mat.Vector, label int) {
	if actual == nil {
		panic("costs: mse cost function got nil actual")
	}
	if label < 0 || label >= actual.Len() {
		panic(fmt.Sprintf("costs: mse cost function got invalid label, %d", label))
	}
}

func (costFunction *mseCostFunction) Differentiate(actual mat.Vector, label int) mat.Vector {
	costFunction.checkInput(actual, label)
	diffs := mat.NewVecDense(actual.Len(), nil)
	diffs.SetVec(label, 1)
	diffs.SubVec(actual, diffs)
	return diffs
}
