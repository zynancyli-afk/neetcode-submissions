func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}else if len(nums) == 2 {
		return max(nums[0],nums[1])
	}
	house := make([]int,len(nums))
	house[0]=nums[0]
	house[1]= max(nums[0],nums[1])
	for i:=2;i < len(nums);i++{
		house[i]=max(house[i-2]+nums[i],house[i-1])
	}
	return house[len(nums)-1]
}