package histogram

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Histogram(csvPath string) {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	DFfile := dataframe.ReadCSV(f)

	// Create histogram for each col
	for _, colName := range DFfile.Names() {
		plotVals := make(plotter.Values, DFfile.Nrow())
		for i, floatVal := range DFfile.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		p := plot.New()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Histogram for values drawn
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		h.Normalize(1)

		// Add histogram to plot
		p.Add(h)

		// Save histogram into PNG
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./linearRegression/"+colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
