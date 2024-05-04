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
	totalShoulderHeight := 0 // 肩の高さのみの合計値を先に計算する
	for _, g := range Gs {
		totalShoulderHeight += g.A
	}

	maxHeight, height := 0, 0
	for _, g := range Gs {
		height = totalShoulderHeight - g.A + g.B // 頂上に選ばれた巨人の肩の高さを引いて、頭の高さを足す
		if height > maxHeight {
			maxHeight = height
		}
	}

	fmt.Println(maxHeight)
}
