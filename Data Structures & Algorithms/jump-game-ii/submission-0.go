func jump(nums []int) int {
	jumps := make(map[int]int)
	var minjump func(index int) int
	minjump = func(index int) int {
		if v, ok := jumps[index]; ok {
			return v
		}
		if nums[index]+index >= len(nums)-1 {
			jumps[index] = 1
			return 1
		}
		min := len(nums)
		for i := index + 1; i < len(nums) && i <= index+nums[index]; i++ {
			j := minjump(i) + 1
			if j < min {
				min = j
			}
		}
		jumps[index] = min
		return min
	}
	if len(nums) <= 1 {
		return 0
	}
	return minjump(0)
}