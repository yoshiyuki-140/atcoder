package main

import (
	"fmt"
	"strings"
)

func canTransform(S string) string {
	keywords := []string{"dream", "dreamer", "erase", "eraser"}

	for len(S) > 0 {
		found := false
		for _, kw := range keywords {
			if strings.HasSuffix(S, kw) {
				found = true
				S = strings.TrimSuffix(S, kw)
				break
			}
		}
		if !found {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	var S string
	fmt.Scan(&S)
	result := canTransform(S)
	fmt.Println(result)
}
