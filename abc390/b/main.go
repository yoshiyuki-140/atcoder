package main

import "fmt"

func main() {
	N := 0
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// 等比数列か否かを判定するので、隣り合う数の比が一定であるかどうかを判定する必要がある

	// 各々チェック
	for i := 1; i < N-1; i++ {
		// 浮動小数点数を扱うときのコンピュータの誤差を打ち消すための工夫が必要
		if A[i]*A[i] != A[i-1]*A[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
