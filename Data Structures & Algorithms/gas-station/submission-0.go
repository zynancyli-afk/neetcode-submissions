func canCompleteCircuit(gas []int, cost []int) int {
	total,totalFromStart, start := 0,0,0
    for i := 0; i < len(gas); i++ {
		total += gas[i]-cost[i]
		totalFromStart += gas[i]-cost[i]
		if totalFromStart < 0 {
			start = i+1
			totalFromStart = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}





