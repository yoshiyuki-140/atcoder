package main

import "fmt"

func main() {
	// init
	N := 0
	fmt.Scan(&N)

	A, B := make([][]rune, N), make([][]rune, N)
	tmp := ""
	for i := 0; i < N; i++ {
		A[i] = make([]rune, N)
		fmt.Scan(&tmp)
		// I didn't write error handling
		A[i] = []rune(tmp)
	}
	for i := 0; i < N; i++ {
		B[i] = make([]rune, N)
		fmt.Scan(&tmp)
		// I didn't write error handling
		B[i] = []rune(tmp)
	}

	// main logic
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i][j] != B[i][j] {
				fmt.Printf("%d %d\n", i+1, j+1)
				break
			}
		}
	}
}
