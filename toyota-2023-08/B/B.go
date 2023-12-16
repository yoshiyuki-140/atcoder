package main

import (
	"fmt"
	"math"
)

func main() {
	//文字を一文字ずつ読み込む
	m := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4}

	s1, s2 := "", ""
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)
	runes := []rune(s1 + s2)

	S1, S2, T1, T2 := string(runes[0]), string(runes[1]), string(runes[2]), string(runes[3])

	sd, td := math.Abs(float64((m[S1]-m[S2])%3)), math.Abs(float64((m[T1]-m[T2])%3))

	if sd == td {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
