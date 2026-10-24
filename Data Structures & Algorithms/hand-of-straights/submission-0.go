func isNStraightHand(hand []int, groupSize int) bool {
    counts := make([]int,1001)
	for _,v := range hand {
		counts[v]++
	}
	for i,v := range counts {
		if v > 0 {
			for m := i;m<i+groupSize;m++{
				if m > len(counts) - 1 {
					return false
				}
				counts[m]-=v
			}
		}else if v < 0 {
			return false
		}
	}
	return true
}
