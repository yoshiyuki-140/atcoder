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
	// 前から順にする場合はこれは早いと思う
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if 500*i+100*j+50*k == X {
					cnt++
				}
				if X < 50*k {
					break
				}
			}
			if X < 100*j {
				break
			}
		}
		if X < 500*i {
			break
		}
	}

	// 出力
	fmt.Println(cnt)

}
