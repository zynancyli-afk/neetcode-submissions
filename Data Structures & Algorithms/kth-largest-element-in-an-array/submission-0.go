func findKthLargest(nums []int, k int) int {
	klargest := []int{}
	min := nums[0]
	minIndex := 0

	for i,v := range nums {
		if len(klargest) < k {
			klargest = append(klargest,v)
			if v < min {
				min = v
				minIndex = i
			}
		}else if v > min {
			klargest[minIndex] = v
			min = v
			//find new min,minIndex
			for j,w := range klargest {
				if w < min {
					min = w
					minIndex = j
				}
			}
		}
	}
	return min
}
