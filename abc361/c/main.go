package main

import (
	"fmt"
	"sort"
)

// 戦略
// ソートしてから大きいほうと小さいほう、交互にとっていく

func main() {
	N, K := 0, 0
	fmt.Scan(&N, &K)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}
	sort.SliceStable(A, func(i, j int) bool { return A[i] < A[j] })
	A_cp := A
	// 小さいほうからとる
	flag := true
	for i := 0; i < K; i++ {
		if flag {
			A = A[1:]
		} else {
			A = A[:len(A)-1]
		}
		flag = !flag
	}
	result1 := objective(A)
	// 大きいほうからとる
	A = A_cp
	flag = false
	for i := 0; i < K; i++ {
		if flag {
			A = A[1:]
		} else {
			A = A[:len(A)-1]
		}
		flag = !flag
	}
	result2 := objective(A)
	if result1 > result2 {
		fmt.Println(result2)
	} else {
		fmt.Println(result1)
	}
}

// 目的関数(ソート済みのスライスを受け取るものとする)
func objective(A []int) int {
	result := A[0] - A[len(A)-1]
	if result < 0 {
		result = -result
	}
	return result
}
