package evaluation

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/gonum/stat"
	"gonum.org/v1/gonum/integrate"
)

func ContinuousMetric() {
	// Open the continuos observations and predictions
	f, err := os.Open("datasets/boston_house_prices.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file
	reader := csv.NewReader(f)

	// Observed and  predicted will hold the parsed observed and predicted values
	// form the continuous data file.
	var observed []float64
	var predicted []float64

	// Line will track row numbers for logging
	line := 1

	// Read in the records looking for unexpected types in the columns.
	for {
		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed and predicted values.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calculate the mean absolute error and mean squared error.
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// Output the MAE and MSE value to standard out.
	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n\n", mSE)
}

func CategoricalMetric() {
	// Open the categorical observations and predictions
	f, err := os.Open("datasets/threeClass.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []int
	var predicted []int

	line := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if line == 1 {
			line++
			continue
		}

		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	var truePosNeg int

	// Accumulate the true, negative and positive count
	for i, oVal := range observed {
		if oVal == predicted[i] {
			truePosNeg++
		}
	}
	// Calculate accuracy
	accuracy := float64(truePosNeg) / float64(len(observed))

	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)

	classes := []int{0, 1, 2}

	for _, class := range classes {
		var truePos int
		var falsePos int
		var falseNeg int

		for i, oVal := range observed {
			switch oVal {
			case class:
				if predicted[i] == class {
					truePos++
					continue
				}
				falseNeg++
			default:
				if predicted[i] == class {
					falsePos++
				}
			}
		}

		precision := float64(truePos) / float64(truePos+falsePos)

		recall := float64(truePos) / float64(truePos+falseNeg)

		fmt.Printf("\nPrecision (class %d) = %0.2f", class, precision)
		fmt.Printf("\nRecall (class %d) = %0.2f\n\n", class, recall)
	}

	scores := []float64{0.1, 0.35, 0.4, 0.8}

	classes1 := []bool{true, false, true, false}

	// Calculate true positive rates, recalls and false positive rates
	tpr, fpr := stat.ROC(0, scores, classes1, nil)

	// Calculate area under curve
	auc := integrate.Trapezoidal(fpr, tpr)

	fmt.Printf("true positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}
