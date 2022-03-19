package main

import (
	dt "github.com/TurniXXD/go-learn/dataframe"
	eval "github.com/TurniXXD/go-learn/evaluation"
	knn "github.com/TurniXXD/go-learn/goLearnKnn"

	linR "github.com/TurniXXD/go-learn/linearRegression"
	mat "github.com/TurniXXD/go-learn/matrix"
)

func main() {
	knn.GoLearnKnn()
	mat.Matrix()
	eval.ContinuousMetric()
	eval.CategoricalMetric()
	dt.Dataframe()
	// CSV path, dependant var, records count
	linR.LinearRegression("./datasets/exams.csv", "FINAL", 4)
}
