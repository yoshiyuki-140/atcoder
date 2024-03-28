package main

import (
	"fmt"
)

func isTravelPlanFeasible(N int, schedule [][]int) string {
	for i := 0; i < N-1; i++ {
		dt := schedule[i+1][0] - schedule[i][0]      // 時間の差分
		dx := abs(schedule[i+1][1] - schedule[i][1]) // x座標の差分
		dy := abs(schedule[i+1][2] - schedule[i][2]) // y座標の差分

		if dt < dx+dy || (dt-dx-dy)%2 != 0 { // 時間の差分が座標の差分より小さい、または時間と座標の偶奇が一致しない場合
			return "No"
		}
	}
	return "Yes"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var N int
	fmt.Scan(&N)

	schedule := make([][]int, N)
	for i := 0; i < N; i++ {
		schedule[i] = make([]int, 3)
		fmt.Scan(&schedule[i][0], &schedule[i][1], &schedule[i][2])
	}

	result := isTravelPlanFeasible(N, schedule)
	fmt.Println(result)
}
