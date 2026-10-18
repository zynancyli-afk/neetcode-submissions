func maxSubArray(nums []int) int {
    sum := 0
	maxSum := nums[0]
	for _,v := range nums {
		if sum <= 0 {
			sum = v
		}else {
			sum+=v
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
