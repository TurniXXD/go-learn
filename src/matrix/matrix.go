package matrix

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func Matrix() {
	// init vector
	var vector []float64

	// Add components
	vector = append(vector, 12.5)
	vector = append(vector, 8.9)

	fmt.Printf("Vector: \n%v\n", vector)

	// Define a vector type
	newVector := mat64.NewVector(2, []float64{12.8, 4.3})

	fmt.Println(newVector)

	dataForMatrix := []float64{1.2, -4.1, -2.4, 5.6}

	//Form matrix
	a := mat64.NewDense(2, 2, dataForMatrix)

	// Format matrix in human readible way
	fA := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("\nMatrix: \nA = %v\n\n", fA)

	// You can get / set values from specific row / column / single element
	matRow := mat64.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are %v\n\n", matRow)

	// Calculate determinant of matrix
	// 1.2*5.6-(-4.1)*(-2.4) = -3.12
	detA := mat64.Det(a)
	fmt.Printf("det(a)= %.2f\n\n", detA)
}
