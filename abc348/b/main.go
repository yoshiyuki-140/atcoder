package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	X, Y int
}

type Point struct {
	n, near_n int
	point     coordinate
	near      coordinate
}

func main() {
	// input
	var N int
	fmt.Scan(&N)

	points := make([]Point, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&points[i].point.X, &points[i].point.Y)
		points[i].n = i + 1
	}

	// main logic

	// 点との距離をすべて試す
	for i := 0; i < N; i++ {
		// nearを保存する
		points[i].near, points[i].near_n = findNearPoint(points[i].point, &points)
	}

	for _, p := range points {
		fmt.Println(p.near_n)
	}
}

func findNearPoint(point coordinate, points *[]Point) (coordinate, int) {
	var n int // point番号
	var d float64
	var result coordinate
	d = 0

	// for i := 0; i < len(*points); i++ {
	for i := len(*points) - 1; i > 0; i-- {
		if point.X == (*points)[i].point.X && point.Y == (*points)[i].point.Y {
			continue
		}
		d_tmp := distance(point, (*points)[i].point)
		fmt.Println(d_tmp)
		if d_tmp >= d {
			result = (*points)[i].point
			n = i + 1
			d = d_tmp
		}
	}
	fmt.Println("_____")
	// for i, p := range *points {
	// 	if point.X == p.point.X && point.Y == p.point.Y {
	// 		continue
	// 	}
	// 	d_tmp := distance(point, p.point)
	// 	if d_tmp < d {
	// 		result = p.point
	// 		n = i + 1
	// 	}
	// }

	// store
	return result, n
}

func distance(A, B coordinate) float64 {
	// 距離を計算する関数
	return math.Sqrt(float64(Abs((A.X-B.X)*(A.X-B.X) - (A.Y-B.Y)*(A.Y-B.Y))))
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
