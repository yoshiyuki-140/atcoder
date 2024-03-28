package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	letters := strings.Split(s, "")
	counter := 0
	for _, letter := range letters {
		if letter == "1" {
			counter++
		}
	}
	fmt.Println(counter)
}
