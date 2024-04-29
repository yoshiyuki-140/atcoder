package main

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
				a[l-2]++
				l--
			} else {
				break
			}
		}
	}
	fmt.Println(l)
}
