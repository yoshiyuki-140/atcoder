package main

import (
	"fmt"
	"sort"
)

func main() {
	// init
	N, M := 0, 0
	fmt.Scan(&N, &M)

	A, B := make([]int, N), make([]int, M)
	for i := range A {
		fmt.Scan(&A[i])
	}
	for i := range B {
		fmt.Scan(&B[i])
	}

	// main logic
	// A + Bの作成
	C := append(A, B...)
	// A+Bをソート
	sort.Ints(C)
	// Aをソート
	sort.Ints(A)

	// Cから2個ずつとって来る。
	for i := 0; i < len(C)-1; i++ {
		// s:START,e:END
		s, e := C[i], C[i+1]
		if isIncludeNumInA(s, &A) && isIncludeNumInA(e, &A) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

// Aに引数に与えた数値が含まれているのかを調べる関数
// 戻り値はbool型
func isIncludeNumInA(n int, A *[]int) bool {
	for _, a := range *A {
		if n == a {
			return true
		}
	}
	return false
}
