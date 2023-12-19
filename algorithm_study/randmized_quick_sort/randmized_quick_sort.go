// 未完成
package main

import (
	"fmt"
	"math/rand"
)

func std_quickSort(slice []int) []int {
}

func randmized_quickSort(slice []int, pivot int) []int {

	if len(slice) < 2 {
		return slice
	}

	left, right := 0, len(slice)-1

	pivotIndex := rand.Int() % len(slice)

	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]

	for i := 0; i < len(slice); i++ {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	randmized_quickSort(slice[:left], rand.Intn(left)+len(slice))
	randmized_quickSort(slice[left+1:], rand.Intn(right-left-1))

	return slice

}

func main() {
	min := 1
	max := 15

	slice_size := 1000

	slice := []int{}
	for i := 0; i < slice_size; i++ {
	}

	randValue := rand.Intn(max) + min

	result := randmized_quickSort(slice, randValue)

	fmt.Println(result)
}
