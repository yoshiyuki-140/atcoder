package main

import "fmt"

type Plan struct {
	t, x, y int
}

func main() {
	var N int
	fmt.Scan(&N)
	plans := make([]Plan, N)
	for _, plan := range plans {
		fmt.Scan(&(plan.t), &(plan.x), &(plan.y))
	}
	// main logic
	for i := 0; i < N-1; i++ {
		if !canExist(plans[i], plans[i+1]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func canExist(planN, planN1 Plan) bool {
	if planN.x+planN.t == planN1.x && planN.y == planN1.y || planN.x-planN.t == planN1.x && planN.y == planN1.y || planN.y+planN.t == planN1.y && planN.x == planN1.x || planN.y-planN.t == planN1.y && planN.x == planN1.x {
		return true
	}
	return false
}
