package main

// 解説見た

import (
	"fmt"
	"math"
)

func main() {
	N := 0
	fmt.Scan(&N)

	// 2**n回ループ
	bitsArray := make([]int, N)
	count := int(math.Pow(2, float64(N)))
	score := 0

Label1:
	for i := count; i > 0; i-- {
		// 条件判定
		score = 0
		bitsArray = intToBitsArray(i, N)
		for j, b := range bitsArray {
			if b == 1 {
				score += 1
			} else {
				score -= 1
			}
			if score < 0 {
				// scoreが0以下ならばそとがわのループをスキップ
				continue Label1
			}

			// loopの最後に使う処理
			if j == len(bitsArray)-1 && score != 0 {
				continue Label1
			}
		}
		PrintParentheses(&bitsArray)
	}
}

func PrintParentheses(bitsArray *[]int) {
	// bitを受け取って、対応する丸かっこを標準出力に出力する関数。
	for _, b := range *bitsArray {
		if b == 1 {
			fmt.Printf("%s", "(")
		} else {
			fmt.Printf("%s", ")")
		}
	}
	fmt.Printf("\n") // 改行
}

func intToBitsArray(num, N int) []int {
	// N[bit]の配列

	// 結果を格納するスライスを初期化
	result := make([]int, N)

	// 各ビットを検査し、対応する配列の位置に格納
	for i := 0; i < N; i++ {
		// numを右にiビットシフトし、1とAND演算を行って最下位ビットを取得
		bit := (num >> i) & 1
		// 配列の先頭に近い位置に低いビット（LSB）を格納
		result[N-i-1] = bit
	}

	return result
}
