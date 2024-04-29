package main

import (
	"fmt"
	"math"
)

func main() {
	// init
	N := 0
	fmt.Scan(&N)
	ballSizes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ballSizes[i])
		ballSizes[i] = int(math.Pow(2, float64(ballSizes[i])))
	}

	// main logic
	column := make([]int, 0)
	for _, ballSize := range ballSizes {
		column = append(column, ballSize)
		for {
			if len(column) <= 1 {
				break
			}
			last := column[len(column)-1]
			penultimate := column[len(column)-2]
			if last == penultimate {
				column = column[:len(column)-2]
				column = append(column, last+penultimate)
			} else {
				break
			}
		}
	}
	fmt.Println(len(column))
}
