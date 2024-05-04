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
	P := make([]int, N)
	for i := range P {
		P[i] = i
	}

	// main logic
	maxHeight := 0
	for NextPermutation(P) {
		height := getHeight(&Gs, P)
		if height > maxHeight {
			maxHeight = height
		}
	}
	fmt.Println(maxHeight)
}

func getHeight(gArray *[]Giant, P []int) int {
	// その順列の時の積み上げた巨人達の高さ
	H := 0
	for i, p := range P {
		if i == len(P)-1 {
			H += (*gArray)[p].B
			break
		}
		H += (*gArray)[p].A
	}
	return H
}

func NextPermutation(a []int) bool {
	// 与えられた順列の次の順列をみつける関数

	// ピボットを見つける
	last := len(a) - 2
	for last >= 0 && a[last] >= a[last+1] {
		last--
	}
	if last < 0 {
		return false // これ以上新しい順列は存在しない
	}

	// ピボットより大きい最小の要素を見つける
	next := len(a) - 1
	for a[next] <= a[last] {
		next--
	}

	// 要素を交換
	a[last], a[next] = a[next], a[last]

	// ピボットの直後から末尾までを反転
	reverse(a, last+1, len(a)-1)

	return true
}

// reverse はスライスの指定された範囲を反転させます。
func reverse(a []int, start, end int) {
	// スライスオブジェクトの範囲反転
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}
