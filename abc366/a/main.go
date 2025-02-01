package main

import "fmt"

func main() {
	N, T, A, ans := 0, 0, 0, ""
	fmt.Scan(&N, &T, &A)
	if T > N/2 || A > N/2 {
		ans = "Yes"
	} else {
		ans = "No"
	}
	fmt.Println(ans)
}
