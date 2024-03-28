package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)

	for i := range A {
		fmt.Scan(&A[i])
	}

	// main logic
	counter := 0
	for isDivisionAllNum(A) {
		for i := range A {
			A[i] = A[i] / 2
		}
		counter++
	}
	fmt.Println(counter)
}

func isDivisionAllNum(A []int) bool {
	// すべての数値に対して2で除算できるか否かを判定するロジック
	for _, a := range A {
		if a%2 != 0 {
			return false
		}
	}
	return true
}
