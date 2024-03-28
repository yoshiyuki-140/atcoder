package main

import "fmt"

func main() {
	N := 0
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		fmt.Printf("%d", N)
	}
	fmt.Println()
}
