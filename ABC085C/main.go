package main

import (
	"fmt"
)

func main() {
	const (
		tenThousand  = 10000
		fiveThousand = 5000
		oneThousand  = 1000
	)
	// input
	N, Y := 0, 0
	fmt.Scanf("%d %d", &N, &Y)

	// main logic
	// 単純に思いつくのは、全探索
	for x := 0; x <= N; x++ {
		for y := 0; y <= N-x; y++ {
			z := N - x - y
			if tenThousand*x+fiveThousand*y+oneThousand*z == Y {
				fmt.Printf("%d %d %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
