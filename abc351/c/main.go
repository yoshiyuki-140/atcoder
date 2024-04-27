package main

import (
	"fmt"
	"math"
)

func main() {
	// init
	N := 0
	fmt.Scan(&N)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}
	column := make([]int, 0)

	// main logic
	for i := 0; i < N; i++ {
		column = append(column, int(math.Pow(2, float64(A[i]))))

		if len(column) <= 1 {
			continue
		}

		for len(column) > 1 {
			penultimateSize := column[len(column)-2]
			lastSize := column[len(column)-1]
			// fmt.Println(penultimateSize, lastSize)
			if penultimateSize == lastSize {
				column = column[:len(column)-2]
				column = append(column, penultimateSize+lastSize)
			} else {
				break
			}
		}
	}
	fmt.Println(len(column))
}
