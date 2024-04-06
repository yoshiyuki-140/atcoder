package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	X, Y int
}

func main() {
	var N int
	fmt.Scan(&N)

	points := make([]coordinate, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&points[i].X, &points[i].Y)
	}

	// main logic
	nears := make([]int, N)

	for pi, p := range points {
		var d float64
		for i := 0; i < N; i++ {
			if p.X == points[i].X && p.Y == points[i].Y {
				continue
			}
			d_tmp := distance(p, points[i])
			if d < d_tmp {
				nears[pi] = i + 1
				d = d_tmp
			}
		}
	}

	for _, p := range nears {
		fmt.Println(p)
	}
}

func distance(A, B coordinate) float64 {
	return math.Sqrt(float64(square(A.X-B.X) + square(A.Y-B.Y)))
}

func square(n int) int {
	return n * n
}
