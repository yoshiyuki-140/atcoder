package main

import "fmt"

func main() {
	cnt := 0
	N := 0

	fmt.Scanf("%d", &N)

	As := make([]int, N)

	// input
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &As[i])
	}

	// すべての要素が偶数かどうか
outter:
	for {
		// すべての要素が偶数かどうかを判定
		for _, a := range As {
			if a%2 != 0 {
				break outter
			}
		}
		// 全て偶数だったらcntの値を+1する
		cnt++
		// すべての要素を2で割る
		for i, a := range As {
			As[i] = a / 2
		}
	}
	// print ans
	fmt.Println(cnt)
}
