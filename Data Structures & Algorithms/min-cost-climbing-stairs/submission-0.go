func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
    minCost := make([]int,len(cost)+1)
	minCost[0]=0
	minCost[1]=0
	for i:=2;i<=len(cost);i++ {
		minCost[i]=min(minCost[i-1]+cost[i-1],minCost[i-2]+cost[i-2])
	}
	return minCost[len(cost)]
}


/*
0 - 0
1 - 0
2 - []



*/
