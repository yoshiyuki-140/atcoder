package main

import "fmt"

func main() {
	B, G, result := 0, 0, ""

	fmt.Scanf("%d %d", &B, &G)
	if B > G {
		result = "Bat"
	} else {
		result = "Glove"
	}
	fmt.Println(result)
}
