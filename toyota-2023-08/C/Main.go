// できてない
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := ""
	fmt.Scanf("%s", &s)

	reps := []int{1, 11, 111}
	var sumReps []int

	// 全探索
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				sumReps = append(sumReps, reps[i]+reps[j]+reps[k])
			}
		}
	}
	sort.Ints(sumReps)

	prev := 0
	result := []int{}

	for _, v := range sumReps {
		if prev == v {
			continue
		}
		prev = v
		result = append(result, v)
	}
}
