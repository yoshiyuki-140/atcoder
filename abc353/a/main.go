package main

import "fmt"

func main() {
	// init
	N := 0
	fmt.Scan(&N)
	H := make([]int, N)
	for i := range H {
		fmt.Scan(&H[i])
	}

	// main logic
	maxHeightBuilding := 1
	for i := range H {
		if H[i] > H[maxHeightBuilding-1] {
			maxHeightBuilding = i + 1
			break
		}
	}
	if maxHeightBuilding == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(maxHeightBuilding)
	}
}
