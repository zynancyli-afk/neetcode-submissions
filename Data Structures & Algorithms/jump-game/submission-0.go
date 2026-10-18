func canJump(nums []int) bool {
	trackJump := make(map[int]bool)
	var canJumpWithStart func(index int)bool
	canJumpWithStart = func(index int)bool{
		if index == len(nums)-1 {
			return true
		}
		if yes,ok := trackJump[index];ok{
			return yes
		}
		for i := index+1; i <= index + nums[index];i++{
			yes, ok := trackJump[i]
			if !ok {
				yes = canJumpWithStart(i)
				trackJump[i] = yes
			}
			if yes {
				return yes
			}
		}
		return false
	}
	return canJumpWithStart(0)
}
