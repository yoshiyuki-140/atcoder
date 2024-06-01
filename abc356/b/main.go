package main

import "fmt"

func main() {

	// init
	N, M := 0, 0
	fmt.Scan(&N, &M)
	A := make([]int, M)
	for i := range A {
		fmt.Scan(&A[i])
	}
	X := make([][]int, N)
	for i := range X {
		X[i] = make([]int, M)
		for j := range X[i] {
			fmt.Scan(&X[i][j])
		}
	}

	// main logic
	// 栄養素ごとの合計摂取量
	nutrition_i := make([]int, M)
	for i := range X {
		for j := range X[i] {
			nutrition_i[j] += X[i][j]
		}
	}
	// 判定
	for i := range nutrition_i {
		if nutrition_i[i] < A[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
