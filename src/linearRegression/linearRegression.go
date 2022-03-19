package linearRegression

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	hist "github.com/TurniXXD/go-learn/histogram"
	sp "github.com/TurniXXD/go-learn/scatterPlot"
	sd "github.com/TurniXXD/go-learn/splittingDataset"
	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func LinearRegression(csvPath, depValCol string, records int) {
	hist.Histogram(csvPath)
	sp.ScatterPlot(csvPath, depValCol)
	sd.SplittingDataset(csvPath)

	fTrain, err := os.Open("./linearRegression/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fTrain.Close()

	readerTrain := csv.NewReader(fTrain)

	// Read all CSV records
	readerTrain.FieldsPerRecord = records
	trainingData, err := readerTrain.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Init regression
	var r regression.Regression
	r.SetObserved("Dependant Variable")

	readerCMD := bufio.NewReader(os.Stdin)

	fmt.Println("\n Enter Independent Variable based on plots: ")
	fmt.Print("> ")

	indepValName, _ := readerCMD.ReadString('\n')

	indepValName = strings.TrimSuffix(indepValName, "\n")

	r.SetVar(0, indepValName)

	// Loop trough CSV records
	for i, record := range trainingData {
		if i == 0 {
			continue
		}

		indepVal, err := strconv.ParseFloat(record[records-1], 64)
		if err != nil {
			log.Fatal(err)
		}

		depVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(indepVal, []float64{depVal}))
	}

	r.Run()
	fmt.Printf("\n Regression Formula: \n%v\n\n", r.Formula)

	fTest, err := os.Open("./linearRegression/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fTest.Close()

	readerTest := csv.NewReader(fTest)

	// Read all CSV records
	readerTest.FieldsPerRecord = records
	testData, err := readerTrain.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop trough data and evaluate prediction based on Mean Absolute Error
	var mAE float64
	for i, record := range testData {
		if i == 0 {
			continue
		}

		indepVal, err := strconv.ParseFloat(record[records-1], 64)
		if err != nil {
			log.Fatal(err)
		}

		depVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict `y` with trained model
		yPredicted, err := r.Predict([]float64{depVal})
		if err != nil {
			log.Fatal(err)
		}

		mAE += math.Abs(indepVal-yPredicted) / float64(len(testData))
	}

	fmt.Printf("MAE = %0.2f\n\n", mAE)

	fDataset, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fDataset.Close()

	DFdataset := dataframe.ReadCSV(fDataset)

	yVals := DFdataset.Col(depValCol).Float()

	// Vals for plotting
	pts := make(plotter.XYs, DFdataset.Nrow())

	// Predicted vals for plotting
	ptsPred := make(plotter.XYs, DFdataset.Nrow())

	// Fill the points with data
	for i, floatVal := range DFdataset.Col(indepValName).Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// Create the plot
	p := plot.New()
	p.X.Label.Text = indepValName
	p.Y.Label.Text = depValCol
	p.Add(plotter.NewGrid())

	// Add scatter plot for the observations
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Line plot for the predictions
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Save plot
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./linearRegression/regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

// From caluclated Regession Formula, just haven't figured out how to pass it in here
func predict(n float64) float64 {
	return 15.5313 + n*1.8436
}
