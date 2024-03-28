package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, A, B int

	fmt.Scan(&N)

	C := make([]int, N)

	for i := range C {
		fmt.Scan(&C[i])
	}

	// main logic

	// ソート
	sort.Ints(C)
	// for i, c := range C {
	for i, j := len(C)-1, 0; i >= 0; i, j = i-1, j+1 {
		if j%2 == 0 {
			A += C[i]
		} else {
			B += C[i]
		}
	}
	fmt.Println(A - B)
}
