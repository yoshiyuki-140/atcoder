package main

import "fmt"

func main() {
	var N, n int
	fmt.Scan(&N)
	mochi := make(map[int]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		mochi[n]++
	}

	// main logic
	fmt.Println(len(mochi))
}
