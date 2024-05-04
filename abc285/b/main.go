package main

// これは解けてないコード

import "fmt"

func main() {
	// init
	N, S := 0, ""
	fmt.Scan(&N, &S)
	runes := []rune(S)

	// main logic
Label1:
	for i := 1; i <= N-1; i++ {
		result := 0
		for l := 2; l <= N-i; l++ {
			if condition1(l, i, N) {
				fmt.Println(l)
				continue Label1
			}
			if condition2(i, l, &runes) {
				fmt.Println(l)
				continue Label1
			}
			result = l
		}
		fmt.Println(result)
	}

}

func condition1(l, i, N int) bool {
	return l <= N-i
}

func condition2(i, l int, runes *[]rune) bool {
	for k := 1; k <= l; k++ {
		if (*runes)[k-1] == (*runes)[k-1+i] {
			return false
		}
	}
	return true
}
