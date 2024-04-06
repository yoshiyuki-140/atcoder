package main

import "fmt"

func main() {
	var N, L, K int
	fmt.Scan(&N, &L, &K)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// main logic

	// n個からk個選ぶ nCr
}

func nCr(n, r int) int {
	return factorial(n-r) / (factorial(n) * factorial(r))
}

func factorial(n int) int {
	// 階乗計算する関数
	result := 1
	for i := 1; i <= n; i++ {
		result *= n
	}
	return result
}
