package main

import (
	"fmt"
)

func main() {
	A := make([]int, 9)
	B := make([]int, 8)
	for i := 0; i < len(A); i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < len(B); i++ {
		fmt.Scan(&B[i])
	}
	fmt.Println(sum(A) - sum(B) + 1)
}

func sum(scores []int) int {
	result := 0
	for _, score := range scores {
		result += score
	}
	return result
}
