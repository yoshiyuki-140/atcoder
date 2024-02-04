package main

import "fmt"

func main() {
	// Code for ABC085B
	// input
	n := 0
	fmt.Scan(&n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}

	// main logic
	// sliceから重複を削除する
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := m[d[i]]; !ok {
			// m[d[i]]が存在しない場合のみtrueを代入
			// つまり、mの長さは重複を削除したdの長さとなる
			m[d[i]] = true
		}
	}

	// output
	fmt.Println(len(m))
}
