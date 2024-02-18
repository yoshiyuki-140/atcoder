package main

import (
	"fmt"
	"math"
)

func main() {
	var A, M, L, R int
	fmt.Scanf("%d %d %d %d", &A, &M, &L, &R)

	// 高橋君と青木君の位置にある最初のクリスマスツリーを見つけるため、各々の最小値 (minTree) と最大値 (maxTree) を計算する
	minTree := int(math.Ceil(float64(L-A) / float64(M)))
	maxTree := int(math.Floor(float64(R-A) / float64(M)))

	// minTree から maxTree までの範囲でクリスマスツリーがある本数を計算する
	treeCount := maxTree - minTree + 1
	if treeCount < 0 {
		// もし本数が負になった場合は0にセット
		treeCount = 0
	}

	fmt.Println(treeCount)
}
