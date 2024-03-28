package main

import (
	"fmt"
)

type Plan struct {
	t, x, y int
}

func isTravelPlanFeasible(N int, schedules []Plan) string {
	if len(schedules) == 1 {
		if schedules[0].t < schedules[0].x+schedules[0].y || (schedules[0].t-schedules[0].x-schedules[0].y)%2 != 0 {
			return "No"
		}
		return "Yes"
	}
	for i := 0; i < N-1; i++ {
		dt := schedules[i+1].t - schedules[i].t      // 時間の差分
		dx := abs(schedules[i+1].x - schedules[i].x) // x座標の差分
		dy := abs(schedules[i+1].y - schedules[i].y) // y座標の差分

		if dt < dx+dy || (dt-dx-dy)%2 != 0 { // 時間の差分が座標の差分より小さい、または時間と座標の偶奇が一致しない場合
			return "No"
		}
	}
	return "Yes"
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	var N int
	fmt.Scan(&N)

	schedules := make([]Plan, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&schedules[i].t, &schedules[i].x, &schedules[i].y)
	}

	result := isTravelPlanFeasible(N, schedules)
	fmt.Println(result)
}
