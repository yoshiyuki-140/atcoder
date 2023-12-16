package main

import "fmt"

/*
	0≤A,B,C≤50
	A+B+C≥1
	50≤X≤20,000
	A,B,C は整数である
	X は
	50 の倍数である
*/

func main() {
	var A, B, C, X, cnt int // Xは50の倍数

	// 入力
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)

	// 処理
	// 全探索アルゴリズム -> ブルートフォース
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if 500*i+100*j+50*k == X {
					cnt++
				}
			}
		}
	}

	// 出力
	fmt.Println(cnt)

}
