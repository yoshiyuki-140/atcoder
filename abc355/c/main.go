package main

import "fmt"

func main() {
	// init
	N, T := 0, 0
	fmt.Scan(&N, &T)

	A := make([]int, T)
	for i := range A {
		fmt.Scan(&A[i])
	}

	// main logic
	// make NxN matrix
	// var M [N][N]int
	M := make([][]int, N)
	for i := 0; i < N; i++ {
		M[i] = make([]int, N)
	}

}

// ビンゴか否かを判定する関数
// 戻り値はbool
func isBingo(M *[][]int, matrixSize int) bool {
	// 列のチェック
	for i := 0; i < matrixSize; i++ {
		sum := 0
		// 合計値の計算
		for j := 0; j < matrixSize; j++ {
			sum += (*M)[i][j]
		}
		if matrixSize == sum {
			return true
		}
	}
	// 行のチェック
	for i := 0; i < matrixSize; i++ {
		sum := 0
		for j := 0; j < matrixSize; j++ {
			sum += (*M)[i][j]
		}
		if matrixSize == sum {
			return true
		}
	}
	// 斜めのチェック
	sum := 0
	for i := 0; i < matrixSize; i++ {
		sum += (*M)[i][i]
	}

	if sum == matrixSize {
		return true
	}
	return false

}

// スライスの合計値を計算する関数
func Sum(S []int) int {
	sum := 0
	for i := range S {
		sum += S[i]
	}
	return sum
}
