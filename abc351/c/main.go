package main

// このコードは解きなおしで書いたやつ

import (
	"fmt"
)

func main() {
	var n, l int
	var a [20000]int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[l])
		l++
		for l > 1 {
			if a[l-2] == a[l-1] {
				a[l-2]++ // 値が同じときは2^x+2^x = 2^(x+1)だから
				l--
			} else {
				break
			}
		}
	}
	fmt.Println(l)
}
