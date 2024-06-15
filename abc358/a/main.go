package main

import "fmt"

func main() {
	S, T := "", ""
	fmt.Scan(&S, &T)
	if S == "AtCoder" && T == "Land" {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
