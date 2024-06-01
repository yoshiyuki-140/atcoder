package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// init
	N, L, R := 0, 0, 0
	fmt.Scan(&N, &L, &R)
	A := make([]int, N)

	// main logic
	for i := range A {
		A[i] = i + 1
	}

	target := A[L-1 : R]
	sort.Slice(target, func(i, j int) bool {
		return target[i] > target[j]
	})

	// convert int to str
	tmp := make([]string, N)
	for i, v := range A {
		tmp[i] = strconv.Itoa(v)
	}
	result := strings.Join(tmp, " ")
	fmt.Println(result)
}
