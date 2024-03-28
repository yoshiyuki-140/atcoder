package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	// main logic
	sum := a * b
	if sum%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
