package main

import "fmt"

func main() {
	H, A := 0, 0
	fmt.Scan(&H, &A)

	// 割り切れる場合はそのまま計算して出力する
	if H%A == 0 {
		fmt.Println((int)(H / A))
		return
	}

	// 割り切れない場合は、上へ丸める
	fmt.Println((int)(H/A) + 1)
}
