package main

import "fmt"

func main() {
	N := 0
	fmt.Scan(&N)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}

	// main logic
	result, sum := 0, 0
	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			sum = (A[i-1] + A[j-1])
			if sum >= 100000000 {
				sum %= 100000000
			}
			result += sum
		}
	}
	fmt.Println(result)
}
