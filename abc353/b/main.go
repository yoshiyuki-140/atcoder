package main

import "fmt"

func main() {
	N, K := 0, 0
	fmt.Scan(&N, &K)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}

	// main logic
	mark, counter, empty := 0, 0, K
	for {
		if mark == N-1 {
			break
		}
		if A[mark] > empty {
			counter++
			empty = K
		} else {
			empty -= A[mark]
			mark++
		}
	}
	fmt.Println(counter)
}
