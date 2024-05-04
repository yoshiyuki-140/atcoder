package main

import (
	"fmt"
)

type Giant struct {
	A, B int
}

func main() {
	// init
	N := 0
	fmt.Scan(&N)
	Gs := make([]Giant, N)
	for i := range Gs {
		fmt.Scan(&Gs[i].A, &Gs[i].B)
	}

	// main logic
	maxHeight, height := 0, 0
	for i := range Gs {
		height = getHeight(&Gs, i)
		if height > maxHeight {
			maxHeight = height
		}
	}
	fmt.Println(maxHeight)
}

func getHeight(gArray *[]Giant, top int) int {
	// その順列の時の積み上げた巨人達の高さ
	H := 0
	for i, g := range *gArray {
		if i == top {
			H += g.B
		} else {
			H += g.A
		}
	}
	return H
}
