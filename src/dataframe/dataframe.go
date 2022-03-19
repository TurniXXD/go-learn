package dataframe

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func Dataframe() {
	sumFile, err := os.Open("./datasets/boston_house_prices.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer sumFile.Close()

	sumDF := dataframe.ReadCSV(sumFile)

	// Calc summary statistics for all the columns of dataset
	sumResult := sumDF.Describe()

	fmt.Println(sumResult)
}
