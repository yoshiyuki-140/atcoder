package main

import (
	"fmt"
)

// 長さ5でそこまで長くないため、挿入ソート
// スライスを渡すことを想定
func insertSort(array []int) {
	var x, j int
	for i := 1; i < len(array); i++ {
		x = array[i]
		j = i
		for j > 0 && array[j-1] > x {
			array[j] = array[j-1]
			j--
		}
		array[j] = x
	}
}

// スライスの比較を行う関数
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// スライスのサイズ
	const arraySize int = 5
	// 入力のスライス
	input := make([]int, arraySize)

	// データ格納
	for i := 0; i < arraySize; i++ {
		fmt.Scan(&input[i])
	}

	// 処理
	// ソート結果を作成
	sortedInput := make([]int, arraySize) // コピーを作成 1
	copy(sortedInput, input)              // コピーを作成 2
	insertSort(sortedInput)               // 挿入ソート
	// 一つずつ交換していき、ソートしたものとどう違うのかを見る
	tmpInput := make([]int, arraySize) // ソート用にコピーを作成
	var x int
	for i := 1; i < arraySize; i++ {
		copy(tmpInput, input) // スライスのコピーを作成
		// 隣同士の交換
		x = tmpInput[i]
		tmpInput[i] = tmpInput[i-1]
		tmpInput[i-1] = x

		// ソート済みの配列と等しいかチェック
		if equalSlices(tmpInput, sortedInput) {
			// 等しい場合結果を出力して終了
			fmt.Println("Yes")
			return
		}
	}
	// 等しくない場合の結果出力
	fmt.Println("No")
}
