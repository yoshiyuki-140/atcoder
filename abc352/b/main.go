package main

import "fmt"

func main() {
	S, T := "", ""
	fmt.Scan(&S, &T)
	runeS, runeT := []rune(S), []rune(T)

	counter := 0
	for i, t := range runeT {
		if runeS[counter] == t {
			counter++
			fmt.Printf("%d ", i+1)
		}
	}
}
