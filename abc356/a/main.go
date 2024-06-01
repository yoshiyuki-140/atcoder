package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	N, L, R := 0, 0, 0
	fmt.Scan(&N, &L, &R)
	A := make([]int, N)

	if L == R {
		for i := range A {
			A[i] = i + 1
		}
	} else {
		for i := 0; i < N; i++ {
			if L-1 <= i && i <= R-1 {
				A[i] = N - L - i + 1
				continue
			}
			A[i] = i + 1
		}
	}

	// convert int to str
	tmp := make([]string, N)
	for i, v := range A {
		tmp[i] = strconv.Itoa(v)
	}
	result := strings.Join(tmp, " ")
	fmt.Println(result)
}
