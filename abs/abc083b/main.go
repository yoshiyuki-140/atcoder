package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	sum := 0
	for i := 1; i <= N; i++ {
		if isSumOfEachDigitBetweenAandB(i, A, B) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isSumOfEachDigitBetweenAandB(n, A, B int) bool {
	// 各桁の総和がA以上B以下であるか否かを求める関数
	nstr := strconv.Itoa(n)      // 文字列へ変換
	D := strings.Split(nstr, "") // 一文字づつ分解してスライスに格納

	sum := 0
	for _, d := range D {
		n, err := strconv.Atoi(d)
		if err != nil {
			fmt.Println("文字列から数値に変換できてない:", err)
		}
		sum += n
	}
	if A <= sum && sum <= B {
		return true
	}
	return false
}
