package splittingDataset

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func SplittingDataset(csvPath string) {
	// Splitting data into 80:20 ratio for training and testing

	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	DFfile := dataframe.ReadCSV(f)

	// Define column types
	trainingNum := (4 * DFfile.Nrow()) / 5
	testNum := DFfile.Nrow() / 5
	if trainingNum+testNum < DFfile.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// Calc num of elements in each set
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Create subset indeces
	trainingDF := DFfile.Subset(trainingIdx)
	testDF := DFfile.Subset(testIdx)

	// Create map for writing the data to files
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create files by saving filtered data
	for idx, setName := range []string{"./linearRegression/training.csv", "./linearRegression/test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// Buffered writer
		w := bufio.NewWriter(f)

		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
