package main

import "fmt"

func main() {
	N, A, t := 0, 0, 0
	fmt.Scan(&N, &A)
	T := make([]int, N)
	for i := range T {
		fmt.Scan(&T[i])
	}

	// main logic
	count := 0 // 来た人の数
	Q := make([]int, N)
	a := 0
	for i := 0; ; i++ {
		if t == T[a] {
			Q = append(Q, t)
			if t < T[a+1] && T[a+1] < t+A {
			}
			a++
		}

		t++ // 時間進める
	}
}
