package main

import "fmt"

type City struct {
	A int
	B int
}

func main() {

	// init
	N := 0
	fmt.Scan(&N)
	city := make([]City, N)
	for i := range city {
		fmt.Scan(&city[i].A, &city[i].B)
	}

	// main logic
}
