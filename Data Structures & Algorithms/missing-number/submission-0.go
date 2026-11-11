func missingNumber(nums []int) int {
	num := len(nums)
	for i:=0;i<len(nums);i++{
		num=num^i^nums[i]
	}
	return num
}
