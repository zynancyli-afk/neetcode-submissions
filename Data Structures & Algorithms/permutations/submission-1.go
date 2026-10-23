func permute(nums []int) [][]int {
	perms := [][]int{}
	curperm := []int{}
	seen := make(map[int]bool)
	
	var permutation func()
	permutation = func(){
		if len(curperm) == len(nums) {
			perms = append(perms, curperm)
		}
		for _,v := range nums {
			if yes,ok := seen[v];ok && yes {
				continue
			}
			seen[v]=true
			curperm = append(curperm, v)
			permutation()
			curperm = curperm[:len(curperm)-1]
			seen[v]=false
		}
	}
	permutation()
	return perms
}