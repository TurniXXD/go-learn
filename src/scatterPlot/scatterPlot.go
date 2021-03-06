package scatterPlot

import (
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func ScatterPlot(csvPath, depValCol string) {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	DFfile := dataframe.ReadCSV(f)

	yVals := DFfile.Col(depValCol).Float()
	for _, colName := range DFfile.Names() {
		pts := make(plotter.XYs, DFfile.Nrow())

		for i, floatVal := range DFfile.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		p := plot.New()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./linearRegression/"+colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
