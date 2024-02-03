// 1 以上 N 以下の整数のうち、10 進法での各桁の和が A 以上 B 以下であるものの総和を求めてください。
package main

import "fmt"

func main() {
	result, N, A, B := 0, 0, 0, 0
	// 読み込み
	fmt.Scanf("%d %d %d", &N, &A, &B)

	for i := 1; i <= N; i++ {
		sumOfDigit := sumOfEachDigit(fmt.Sprintf("%d", i))
		if A <= sumOfDigit && sumOfDigit <= B {
			result += i
		}
	}
	fmt.Println(result)
}

func sumOfEachDigit(s string) int {
	result := 0
	for _, e := range s {
		// eを数値に変換
		result += int(e - '0')
	}
	return result
}
