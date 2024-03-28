package main

import "fmt"

func main() {
	// input
	var N, Y int
	fmt.Scan(&N, &Y)

	for i := N; i >= 0; i-- {
		for j := N - i; j >= 0; j-- {
			k := N - i - j
			if 10000*i+5000*j+1000*k == Y {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
