package main

import (
	// eval "github.com/TurniXXD/go-learn/evaluation"
	// knn "github.com/TurniXXD/go-learn/goLearnKnn"
	// mat "github.com/TurniXXD/go-learn/matrix"
	eval "github.com/TurniXXD/go-learn/evaluation"
	knn "github.com/TurniXXD/go-learn/goLearnKnn"
	mat "github.com/TurniXXD/go-learn/matrix"
)

func main() {
	knn.GoLearnKnn()
	mat.Matrix()
	eval.ContinuousMetric()
	eval.CategoricalMetric()
}
