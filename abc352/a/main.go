package main

import "fmt"

func main() {
	N, X, Y, Z := 0, 0, 0, 0
	fmt.Scan(&N, &X, &Y, &Z)

	result := ""
	if X < Y {
		if X < Z && Z < Y {
			result = "Yes"
		} else {
			result = "No"
		}
	} else {
		if Y < Z && Z < X {
			result = "Yes"
		} else {
			result = "No"
		}
	}
	fmt.Println(result)
}
