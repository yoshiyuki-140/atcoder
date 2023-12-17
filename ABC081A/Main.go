package main

import "fmt"

func main() {
	s := ""
	cnt := 0

	fmt.Scanf("%s", &s)
	for _, c := range s {
		if c == '1' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
