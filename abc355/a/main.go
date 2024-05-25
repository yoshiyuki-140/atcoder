package main

import "fmt"

func main() {
	// init
	A, B := 0, 0
	fmt.Scan(&A, &B)

	// main logic

	// 回答が同じ場合には、判定不可能として処理を終了する。
	if A == B {
		fmt.Println("-1")
		return
	}
	ans := 6 - (A + B)
	fmt.Println(ans)
}
