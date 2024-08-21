package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	if (a*b)%2 == 0 {
		fmt.Println("Odd")
		return
	}
	fmt.Println("Even")
}
