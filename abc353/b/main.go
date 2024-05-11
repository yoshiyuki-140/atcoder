package main

import (
	"fmt"
)

func main() {
	N, K := 0, 0
	fmt.Scan(&N, &K)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}

	// main logic
	counter, mark := 0, 0
	for mark < N {
		empty := K
		for mark < N {
			if A[mark] <= empty {
				empty -= A[mark]
			} else {
				break
			}
			mark++
		}
		counter++
	}
	fmt.Println(counter)
}
