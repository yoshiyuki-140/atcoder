package main

import (
	"fmt"
	"math"
	"sort"
)

type cards []int

func (a cards) Len() int           { return len(a) }
func (a cards) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a cards) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	// init
	n := 0
	Alice := make([]int, 0)
	Bob := make([]int, 0)

	// input
	fmt.Scan(&n)

	var a cards
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = 0
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}

	// sort
	sort.Sort(a)

	// handle
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			Alice = append(Alice, a[i])
		} else {
			Bob = append(Bob, a[i])
		}
	}

	// output
	resultAlice, resultBob := 0, 0

	resultAlice = sum(Alice)
	resultBob = sum(Bob)
	finalResult := math.Abs(float64(resultAlice - resultBob))
	fmt.Println(finalResult)
}

func sum(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}
